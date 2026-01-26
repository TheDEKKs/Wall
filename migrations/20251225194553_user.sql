-- +goose Up
CREATE TABLE users(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_name VARCHAR UNIQUE NOT NULL,
    password_hash VARCHAR NOT NULL,
    -- telegram_id VARCHAR UNIQUE,
    -- telegram_user_name VARCHAR,

    registration_at TIMESTAMPTZ DEFAULT now()

);

-- +goose Down
DROP TABLE users;
