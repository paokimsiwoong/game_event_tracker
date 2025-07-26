# game_event_tracker

## 게임 이벤트 일정 캘린더 입력 프로젝트
---
### 유저가 게임 이벤트 공지를 제공하는 url (ex:https://sv-news.pokemon.co.jp/ko/list => https://sv-news.pokemon.co.jp/ko/json/list.json) 을 입력하면 그 내용들을 긁어와서 일정이 적혀 있는 이벤트들을 구글 캘린더에 저장합니다.

```
game-event-calendar/
├── cmd/
│   └── main.go            // 엔트리 포인트
├── internal/
│   ├── commands/          // CLI 명령어 기능
│   │   ├── commands.go
│   │   ├── addsite.go
│   │   ├── deletesite.go
│   │   └── sites.go
│   ├── crawler/           // 웹 크롤러(공지 긁어오기)
│   │   └── crawler.go
│   ├── parser/            // 일정 정보 파싱
│   │   └── parser.go
│   ├── calendar/          // 구글 캘린더 연동
│   │   └── calendar.go
│   └── config/            // 설정(토큰, URL 등)
│       └── config.go
├── go.mod
└── README.md
```

TODO: events에는 동일 이벤트가 중복되어 들어가지 않도록 변경하기(posts테이블과 events 테이블로 분리? / 시작, 종료 시점 pair를 UNIQUE(,)로 하면 동일 일정 다른 이벤트가 있음 => tag, 시작, 종료 시점 trio를 unique로?)