-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA "order";
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA "order";
-- +goose StatementEnd
