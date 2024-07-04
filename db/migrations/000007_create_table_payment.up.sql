CREATE TABLE IF NOT EXISTS payment (
    id bigserial PRIMARY KEY,
    cart_id bigint NOT NULL,
    total_payment int NOT NULL,
    gateway_name varchar(50),
    payment_code text NOT NULL,
    "status" int NOT NULL DEFAULT -1,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now()
);