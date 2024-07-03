CREATE TABLE IF NOT EXISTS cart (
    id bigserial PRIMARY KEY,
    user_id bigint NOT NULL,
    product_id bigint NOT NULL,
    quantity int NOT NULL,
    FOREIGN KEY (user_id) REFERENCES "user" (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES product (id)
);