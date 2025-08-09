-- +goose Up
CREATE TABLE events (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    tag INTEGER NOT NULL,
    tag_text TEXT NOT NULL,
    starts_at TIMESTAMP WITH TIME ZONE,
    ends_at TIMESTAMP WITH TIME ZONE,
    event_cal_ids TEXT[],
    names TEXT[] NOT NULL,
    posted_ats TIMESTAMP WITH TIME ZONE[] NOT NULL,
    post_urls TEXT[] NOT NULL,
    post_ids UUID[] NOT NULL,
    site_id UUID NOT NULL,
    CONSTRAINT fk_sites
    FOREIGN KEY (site_id)
    REFERENCES sites(id),
    -- ON DELETE CASCADE, -- 캘린더에 등록된 많은 일정을 지우기 전에 site가 지워져서 events 테이블에 연결된 데이터들이 지워지면 큰 문제
    UNIQUE(tag, starts_at, ends_at) 
);


-- +goose Down
DROP TABLE events;