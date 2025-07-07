-- +goose Up
CREATE TABLE sites (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    last_fetched_at TIMESTAMP,
    name TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL -- 동일 url 중복을 막기
);


-- +goose Down
DROP TABLE sites;