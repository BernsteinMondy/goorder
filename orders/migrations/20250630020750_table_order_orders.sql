-- +goose Up
-- +goose StatementBegin
CREATE TABLE "order".orders
(
    row_id         uuid           NOT NULL,
    row_created_at timestamptz    NOT NULL,
    id             uuid           NOT NULL,
    created_at     timestamptz    NOT NULL,
    user_id        uuid           NOT NULL,
    status         "order".status NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "order".orders;
-- +goose StatementEnd
