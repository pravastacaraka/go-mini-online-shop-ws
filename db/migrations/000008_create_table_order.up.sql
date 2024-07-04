CREATE TABLE IF NOT EXISTS "order" (
    id bigserial PRIMARY KEY,
    payment_id bigint NOT NULL,
    user_id bigint NOT NULL,
    address_id bigint NOT NULL,
    invoice varchar(50) NOT NULL,
    total_price int NOT NULL,
    total_weight double precision NOT NULL,
    shipping_price int NOT NULL,
    "status" int NOT NULL DEFAULT 100,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    FOREIGN KEY (payment_id) REFERENCES payment (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES "user" (id),
    FOREIGN KEY (address_id) REFERENCES "address" (id)
);