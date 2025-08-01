-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    name TEXT NOT NULL,
    tag INTEGER NOT NULL,
    tag_text TEXT NOT NULL,
    posted_at TIMESTAMP WITH TIME ZONE NOT NULL,
    starts_at TIMESTAMP WITH TIME ZONE,
    -- starts_at TIMESTAMP WITH TIME ZONE[], ???
    ends_at TIMESTAMP WITH TIME ZONE,
    -- ends_at TIMESTAMP WITH TIME ZONE[], ???
    body TEXT NOT NULL,
    post_url TEXT NOT NULL,
    site_id UUID NOT NULL,
    CONSTRAINT fk_sites
    FOREIGN KEY (site_id)
    REFERENCES sites(id)
    ON DELETE CASCADE 
);


-- +goose Down
DROP TABLE posts;