-- +goose Up
-- +goose StatementBegin
CREATE TABLE files
(
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    mask bit varying
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE  files;
-- +goose StatementEnd
