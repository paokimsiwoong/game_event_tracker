-- +goose Up
ALTER TABLE posts
ADD COLUMN registered BOOLEAN NOT NULL DEFAULT false;


-- +goose Down
ALTER TABLE posts
DROP COLUMN registered;



-- 아주 큰 테이블의 경우 여러 단계로 나눠서 진행
-- ALTER TABLE posts ADD COLUMN registered BOOLEAN;  -- NOT NULL 없이 추가
-- UPDATE posts SET registered = false WHERE registered IS NULL;  -- 기존 row 채우기
-- ALTER TABLE posts ALTER COLUMN registered SET NOT NULL;  -- NOT NULL 제약 추가
-- 추가로 default 설정도 하면
-- ALTER TABLE posts ALTER COLUMN registered SET DEFAULT false;

