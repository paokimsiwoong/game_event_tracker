-- +goose Up
CREATE TABLE events (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    tag TEXT NOT NULL,
    starts_at TIMESTAMP,
    ends_at TIMESTAMP,
    body TEXT NOT NULL,
    site_id UUID NOT NULL,
    CONSTRAINT fk_sites
    FOREIGN KEY (site_id)
    REFERENCES sites(id)
    ON DELETE CASCADE 
);


-- +goose Down
DROP TABLE events;