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

TODO: https://sv-news.pokemon.co.jp/ko/page/373.html, https://sv-news.pokemon.co.jp/ko/page/370.html 과 같이 한 게시글에 테라레이드 기간과 이후의 이상한 소포 선물 기간이 같이 있는 경우 이상한 소포 선물 기간의 tag와 tag text가 1, 테라 레이드배틀이 되는 문제 해결?