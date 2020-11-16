-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS order_items (
    id                  BIGSERIAL NOT NULL PRIMARY KEY,
    quantity            INT NOT NULL,
    product_id          BIGINT NOT NULL REFERENCES products(id),
    order_id            BIGINT NOT NULL REFERENCES orders(id),
    organization_id     BIGINT NOT NULL REFERENCES organizations(id),
    created_at          TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at          TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS order_items;
