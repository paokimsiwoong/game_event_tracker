package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq" // _ "github.com/lib/pq" 는 postgres driver를 사용한다고 알리는 것. main.go 내부에서 직접 코드 작성할 때 쓰이지는 않음
	"github.com/paokimsiwoong/game_event_tracker/internal/commands"
	"github.com/paokimsiwoong/game_event_tracker/internal/config"
	"github.com/paokimsiwoong/game_event_tracker/internal/database"
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

	// sqlc가 생성한 database 패키지 사용
	dbQueries := database.New(db)

	// state, commands 구조체들 초기화
	stateInstance := commands.State{
		PtrCfg: &cfg,
		PtrDB:  dbQueries,
	}
	cmds := commands.Commands{
		CommandMap: make(map[string]func(*commands.State, commands.Command) error),
	}

	//  command 등록
	cmds.Register("addsite", commands.HandlerAddSite)
	cmds.Register("sites", commands.HandlerSites)

	// 유저 명령어 입력 확인
	if len(os.Args) < 2 {
		// os.Args의 첫번째 arg는 무조건 프로그램 이름이므로 명령어가 포함되어 있으려면 길이가 2 이상이어야 한다
		log.Fatal("error checking arguments : not enough arguments were provided")
	}
	// 유저 명령어 command 구조체에 저장
	cmd := commands.Command{
		Name: os.Args[1], // 0은 프로그램 이름, 1은 명령어 이름
		Args: os.Args[2:],
	}

	// 명령어 실행
	if err := cmds.Run(&stateInstance, cmd); err != nil {
		// log.Fatalf("%v", err)
		// @@@ 해답 log.Fatal 함수 사용 반영
		log.Fatal(err)
	}
}
