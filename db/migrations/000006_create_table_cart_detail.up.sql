CREATE TABLE IF NOT EXISTS cart_detail (
    id bigserial PRIMARY KEY,
    cart_id bigint NOT NULL,
    product_id bigint NOT NULL,
    quantity int NOT NULL,
    FOREIGN KEY (cart_id) REFERENCES cart (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES product (id)
);