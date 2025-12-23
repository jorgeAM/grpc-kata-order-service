BEGIN;

CREATE TABLE IF NOT EXISTS grpc_kata.orders
(
    id uuid PRIMARY KEY,
    customer_id VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

COMMIT;
