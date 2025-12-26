-- +goose Up
-- +goose StatementBegin
CREATE TABLE walls (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID UNIQUE REFERENCES users(id),
    created_at TIMESTAMPTZ DEFAULT now()

);

CREATE TABLE comments (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID UNIQUE NOT NULL REFERENCES users(id),
    wall_id UUID UNIQUE NOT NULL REFERENCES walls(id),

    text VARCHAR NOT NULL,

    created_at TIMESTAMPTZ DEFAULT now()


);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE comments;
DROP TABLE walls;
-- +goose StatementEnd
