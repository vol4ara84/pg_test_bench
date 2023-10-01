package config

type Postgres struct {
	DSN         string `required:"true"`
	MaxIdleConn int    `envconfig:"default=2"`
	MaxOpenConn int    `envconfig:"default=5"`
}
