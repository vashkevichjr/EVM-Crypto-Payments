-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS invoices (
    id UUID PRIMARY KEY,
    merchant_id UUID NOT NULL REFERENCES merchants(id),
    order_id TEXT NOT NULL,
    amount NUMERIC(20, 6) NOT NULL,
    currency TEXT NOT NULL,
    deposit_address TEXT NOT NULL,
    private_key_encrypted TEXT NOT NULL,
    status TEXT NOT NULL CHECK (status IN ('PENDING', 'PAID', 'EXPIRED')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMPTZ NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS invoices_merchant_order_id_uindex
    ON invoices (merchant_id, order_id);

CREATE INDEX IF NOT EXISTS invoices_status_address_index
    ON invoices (status, deposit_address);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS invoices_status_address_index;
DROP INDEX IF EXISTS invoices_merchant_order_id_uindex;
DROP TABLE IF EXISTS invoices;
-- +goose StatementEnd

