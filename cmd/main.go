package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // _ "github.com/lib/pq" 는 postgres driver를 사용한다고 알리는 것. main.go 내부에서 직접 코드 작성할 때 쓰이지는 않음
	"github.com/paokimsiwoong/game_event_tracker/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Fatal: %v", err)
	}

	// sql.Open의 첫번째 인자로 사용하는 sql 드라이버를 지정(_ "github.com/lib/pq"이 postgres)
	// 두번째 인자로는 cfg.DBURL에 저장된 connection string(postgres://username:password@localhost:5432/dbname?sslmode=disable 형태)로 database 연결
	db, errr := sql.Open("postgres", cfg.DBURL)
	// db는 *sql.DB 타입
	if errr != nil {
		log.Fatalf("error connecting to db : %v", err)
	}
	defer db.Close()
}
