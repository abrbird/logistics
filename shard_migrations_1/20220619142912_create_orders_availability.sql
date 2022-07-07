-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS postgres_fdw;
CREATE TABLE public.logistics_orders_availability_shard_1 (
                                               order_id bigint NOT NULL,
                                               issue_point_id bigint NOT NULL,
                                               status VARCHAR NOT NULL,
                                               updated_at TIMESTAMP NOT NULL default current_timestamp,
                                               PRIMARY KEY (order_id, issue_point_id)
);
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON public.logistics_orders_availability_shard_1
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.logistics_orders_availability_shard_1;
-- +goose StatementEnd
