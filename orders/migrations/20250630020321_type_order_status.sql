-- +goose Up
-- +goose StatementBegin
CREATE TYPE "order".status AS ENUM ('created', 'payed', 'success', 'cancelled');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE "order".status;
-- +goose StatementEnd
