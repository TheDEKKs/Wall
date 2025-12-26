-- +goose Up
CREATE TABLE users(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_name VARCHAR UNIQUE NOT NULL,
    password_hash VARCHAR NOT NULL,

    registration_at TIMESTAMPTZ DEFAULT now()

);

-- +goose Down
DROP TABLE users;
