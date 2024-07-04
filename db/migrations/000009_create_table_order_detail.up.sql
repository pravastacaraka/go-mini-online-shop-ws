CREATE TABLE IF NOT EXISTS order_detail (
    id bigserial PRIMARY KEY,
    order_id bigint NOT NULL,
    product_id bigint NOT NULL,
    product_name varchar(150) NOT NULL,
    quantity int NOT NULL,
    price int NOT NULL,
    subtotal_price int NOT NULL,
    "weight" double precision NOT NULL,
    subtotal_weight double precision NOT NULL,
    category_id int NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    FOREIGN KEY (order_id) REFERENCES "order" (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES product (id)
);