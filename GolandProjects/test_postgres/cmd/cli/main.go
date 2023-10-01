package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"golang.org/x/sync/errgroup"
	"test_postgres/cmd/internal/config"
	"test_postgres/internal/generated/db"
	"test_postgres/internal/storage"

	"github.com/jmoiron/sqlx"
)

func postgresOpen(cfg *config.Postgres) (db *sqlx.DB, err error) {
	db, err = sqlx.Open("postgres", cfg.DSN)
	return
}

func Postgres(cfg *config.Postgres) (*sqlx.DB, error) {
	pgConn, err := postgresOpen(cfg)
	if err != nil {
		return nil, fmt.Errorf("init pg connection failed: %s", err)
	}
	pgConn.SetMaxOpenConns(cfg.MaxOpenConn)
	pgConn.SetMaxIdleConns(cfg.MaxIdleConn)

	return pgConn, nil
}

func runWithInterval(s Storage.Storage, ctx context.Context, id int64) error {
	timer := time.NewTimer(TIME_TO_DO_MIN * time.Minute)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-timer.C:
				done <- true
				return
			default:
				if err := s.UpdateFileMask(ctx, id); err != nil {
					ctx.Done()
					fmt.Println(err)
					break
				}
				time.Sleep(5 * time.Second)
			}
		}
	}()
	<-done
	return nil

}

const (
	TIME_TO_DO_MIN = 5
	STR_LENGTH     = 15000
	IDS            = 20
)

func main() {
	cfg := config.Postgres{
		DSN:         "host=localhost port=5432 dbname=tile_generator user=tile_generator password=tile_generator connect_timeout=5 statement_timeout=60000 sslmode=disable binary_parameters=yes",
		MaxIdleConn: 5,
		MaxOpenConn: 7,
	}
	pgConn, err := Postgres(&cfg)
	if err != nil {
		fmt.Println(err)
	}
	s := Storage.NewStorage(pgConn, db.New(pgConn))
	ctx := context.Background()
	basicString := make([]string, STR_LENGTH)

	filesIds := make([]int64, IDS)
	for i := 0; i < IDS; i++ {
		id, err := s.InsertFile(ctx, strings.Join(basicString, "1"))
		if err != nil {
			fmt.Println(err)
			break
		}
		filesIds[i] = id
	}
	g, gCtx := errgroup.WithContext(ctx)
	for i := range filesIds {
		i := i
		g.Go(func() error {
			return runWithInterval(*s, gCtx, filesIds[i])
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("DONE")
}
