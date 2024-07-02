CREATE TABLE IF NOT EXISTS "user" (
    id bigserial PRIMARY KEY,
    email varchar(300) NOT NULL UNIQUE,
    "password" text NOT NULL,
    "name" varchar(100) NOT NULL,
    token text,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now()
);