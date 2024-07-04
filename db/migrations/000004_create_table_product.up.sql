CREATE TABLE IF NOT EXISTS product (
    id bigserial PRIMARY KEY,
    category_id int NOT NULL,
    "name" varchar(150) NOT NULL,
    "desc" text NOT NULL,
    sku varchar(20),
    stock int NOT NULL,
    pictures text [] NOT NULL,
    price int NOT NULL,
    "weight" double precision NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    FOREIGN KEY (category_id) REFERENCES category (id)
);