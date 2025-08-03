package calendar

import (
	"context"
	"database/sql"
	"testing"

	// _ "github.com/lib/pq"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/paokimsiwoong/game_event_tracker/internal/database"
	"github.com/stretchr/testify/require"
)

func TestNewCalendar(t *testing.T) {
	// @@@ html 내용이 테스트 시점에 따라 바뀌므로 다른 테스트 방식 고려 필요 @@@
	// Test: create calendar api client
	srv, err := NewCalendar()
	require.NoError(t, err)
	// Test: add a event to Google calendar
	err = TestCalendar(srv)
	require.NoError(t, err)
	// require.Error(t, err)
	// assert.Equal(t, ErrTemp, err)

	// Test: add a event to Google calendar with db data

	// sql.Open의 첫번째 인자로 사용하는 sql 드라이버를 지정(_ "github.com/lib/pq"이 postgres)
	// 두번째 인자로는 connection string(postgres://username:password@localhost:5432/dbname?sslmode=disable 형태)로 database 연결
	db, err := sql.Open("pgx", "postgres://postgres:20151223@localhost:5432/getker?sslmode=disable")
	require.NoError(t, err)
	// db는 *sql.DB 타입
	defer db.Close()

	// sqlc가 생성한 database 패키지 사용
	dbQueries := database.New(db)
	posts, err := dbQueries.GetPostsOnGoingAndSites(context.Background())
	require.NoError(t, err)
	e := &Event{
		Name:     posts[1].Name,
		Tag:      posts[1].Tag,
		TagText:  posts[1].TagText,
		StartsAt: posts[1].StartsAt,
		EndsAt:   posts[1].EndsAt,
		EventUrl: posts[1].PostUrl,
		SiteName: posts[1].SiteName,
		SiteUrl:  posts[1].SiteUrl,
	}
	calendarID := "primary"
	err = AddEvent(srv, calendarID, e)
	require.NoError(t, err)
}
