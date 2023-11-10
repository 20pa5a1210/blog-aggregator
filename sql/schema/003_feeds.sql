
-- +goose Up

CREATE TABLE feeds(
    id UUID PRIMARY KEY,
    created_at timestamp not null,
    updated_at timestamp not null,
    name text not null,
    url text not null unique,
    user_id UUID not null REFERENCES users(id) on delete cascade
);

-- +goose Down
DROP TABLE feeds;
