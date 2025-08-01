-- +goose Up
CREATE TABLE events (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    tag INTEGER NOT NULL,
    tag_text TEXT NOT NULL,
    starts_at TIMESTAMP WITH TIME ZONE,
    ends_at TIMESTAMP WITH TIME ZONE,
    names TEXT[] NOT NULL,
    posted_ats TIMESTAMP WITH TIME ZONE[] NOT NULL,
    post_urls TEXT[] NOT NULL,
    post_ids UUID[] NOT NULL,
    site_id UUID NOT NULL,
    CONSTRAINT fk_sites
    FOREIGN KEY (site_id)
    REFERENCES sites(id)
    ON DELETE CASCADE,
    UNIQUE(tag, starts_at, ends_at) 
);


-- +goose Down
DROP TABLE events;