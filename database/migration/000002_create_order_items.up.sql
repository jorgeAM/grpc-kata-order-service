BEGIN;

CREATE TABLE IF NOT EXISTS order_schema.order_items
(
    id uuid PRIMARY KEY,
    order_id uuid NOT NULL,
    product_code VARCHAR(50) NOT NULL,
    quantity INT NOT NULL,
    unit_price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

COMMIT;
