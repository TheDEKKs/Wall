-- +goose Up
CREATE TABLE telegram (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	telegram_id BIGINT NOT NULL,
	first_name VARCHAR NOT NULL,
	last_name VARCHAR NOT NULL,
	username VARCHAR NOT NULL,

    registration_at TIMESTAMPTZ DEFAULT now()

);

CREATE TABLE users(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_name VARCHAR UNIQUE NOT NULL,
    password_hash VARCHAR NOT NULL,
    telegram_id UUID UNIQUE NOT NULL REFERENCES telegram(id),

    registration_at TIMESTAMPTZ DEFAULT now()

);

-- +goose Down
DROP TABLE users;
DROP TABLE telegram;

