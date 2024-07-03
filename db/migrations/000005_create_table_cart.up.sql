CREATE TABLE IF NOT EXISTS cart (
    id bigserial PRIMARY KEY,
    user_id bigint NOT NULL,
    FOREIGN KEY (user_id) REFERENCES "user" (id) ON DELETE CASCADE
);