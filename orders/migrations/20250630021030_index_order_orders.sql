-- +goose Up
-- +goose StatementBegin
CREATE INDEX index_order_orders ON "order".orders (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX "order".index_order_orders;
-- +goose StatementEnd
