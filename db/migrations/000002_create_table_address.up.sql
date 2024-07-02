CREATE TABLE IF NOT EXISTS "address" (
    id bigserial PRIMARY KEY,
    user_id bigint NOT NULL,
    "address" text NOT NULL,
    postal_code int NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    FOREIGN KEY (user_id) REFERENCES "user" (id) ON DELETE CASCADE
);