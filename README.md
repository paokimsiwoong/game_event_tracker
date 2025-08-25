# game_event_tracker
## 게임 이벤트 일정 캘린더 입력 프로젝트
### 유저가 게임 이벤트 공지를 제공하는 url (ex:https://sv-news.pokemon.co.jp/ko/list => https://sv-news.pokemon.co.jp/ko/json/list.json) 을 입력하면 그 내용들을 긁어와서 일정이 적혀 있는 이벤트들을 구글 캘린더에 저장합니다.

<details>
<summary> <h2> 프로젝트 구조 </h2> </summary>
<div markdown="1">

```
game-event-calendar/
├── getker/
│   └── main.go            // 엔트리 포인트
├── internal/
│   ├── commands/          // CLI 명령어 기능
│   │   ├── commands.go
│   │   ├── help.go
│   │   ├── addsite.go
│   │   ├── sites.go
│   │   ├── crawl.go
│   │   ├── posts.go
│   │   ├── events.go
│   │   ├── cal.go
│   │   └── delete.go
│   ├── crawler/           // 웹 크롤러(공지 긁어오기)
│   │   └── crawler.go
│   ├── parser/            // 일정 정보 파싱
│   │   └── parser.go
│   ├── calendar/          // 구글 캘린더 연동
│   │   └── calendar.go
│   ├── database/          // sqlc 생성 코드
│   │   ├── db.go
│   │   ├── models.go
│   │   ├── sites.sql.go
│   │   ├── posts.sql.go
│   │   ├── events.sql.go
│   │   └── evetns_manual.sql.go
│   └── config/            // 설정(토큰, URL 등) 관리
│       └── config.go
├── sql/
│   ├── schema/            // sql 데이터베이스 마이그레이션 모음 (goose)
│   │   ├── 001_sites.sql
│   │   ├── 002_posts.sql
│   │   ├── 003_events.sql
│   │   ├── 004_post_registered.sql
│   │   └── 005_event_event_cal_ids.sql
│   └── queries/           // sql 쿼리 모음
│       ├── sites.sql
│       ├── posts.sql
│       └── events.sql
├── go.mod
├── go.sum
├── .gitignore
├── .getker_env                   // 필요 설정 값 저장
├── .getker_env_example
└── README.md
```

</div>
</details>


<details>
<summary> <h2> 프로젝트 설치 </h2> </summary>
<div markdown="1">

### 1. go v1.24 또는 그 이후 버전 설치
```bash
curl -sS https://webi.sh/golang | sh
```

<details>
<summary> <h3> 2. Postgres v15 또는 이후 버전 설치 및 설정 </h3> </summary>
<div markdown="1">

#### 2-1. Postgres 설치(리눅스 Ubuntu 기준)
```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
```
* #### 설치가 완료되면 자동으로 운영체제(리눅스) 레벨의 `postgres`라는 유저 계정이 생성

#### 2-2. postgres 계정 비밀 번호 설정하기
```bash
sudo passwd postgres
```
* #### 입력 시 비밀번호를 2번 입력하는 프롬프트가 생성되고, 입력한 비밀번호가 `postgres` 계정 로그인 비밀번호로 설정

#### 2-3. Postgres server 백그라운드 실행
```bash
sudo service postgresql start
```
#### 2-4. psql 쉘 사용하기
```bash
sudo -u postgres psql
```
* #### 명령어 입력하면 psql shell 이 새 prompt (`postgres=#`)를 표시

#### 2-5. 새 데이터베이스 생성
```bash
# psql shell(postgres=#)에 입력하기
CREATE DATABASE <db_name>;
# ex: CREATE DATABASE tracker;
```
#### 2-6. 데이터베이스 내 사용자 비밀번호 설정
```bash
# psql shell(postgres=#)에 입력하기
# 생성한 데이터베이스에 연결
\c <db_name>
# <db_name>=# 형태의 새 프롬프트 표시

# 데이터베이스에 연결된 상태에서(<db_name>=#)
# DB 내 사용자 postgres 비밀번호 설정
ALTER USER postgres PASSWORD '<your_password>';
``` 
* #### 여기서 설정한 비밀번호가 뒤에 나올 `connection string`에 들어가는 비밀번호
* #### `sudo passwd postgres`로 위에서 설정한 리눅스 OS상 postgres 유저의 비밀번호와 별개

</div>
</details>

### 3. 프로젝트 로컬 다운로드
```bash
git clone https://github.com/paokimsiwoong/game_event_tracker
```

### 4. goose 설치 및 up migration 실행
#### 4-1. goose 설치
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```
#### 4-2. up migration 실행
```bash
# 프로젝트의 sql/schema directory 경로에서 아래 명령어를 실행
goose postgres <connection_string> up
```
* #### `connection string`은 `"postgres://postgres:<database user's password>@localhost:5432/<database name>"`의 형태. 
    * #### 위에서 `ALTER USER postgres PASSWORD '<your_password>';`로 설정한 비밀번호 입력
    * #### (postgres 기본 포트는 `5432`)
* #### up migration을 실행하고 나면 프로젝트에 필요한 데이터 테이블들이 데이터베이스 내부에 생성

<details>
<summary> <h3> 5. Google Cloud에서 새 프로젝트 생성하기 </h3> </summary>
<div markdown="1">

#### 5-1. 웹 브라우저에서 [Google Cloud Console](https://console.cloud.google.com) 접속
* #### Google 계정 필요

#### 5-2. 프로젝트 선택 도구로 새 프로젝트 생성 페이지 들어가기
* #### Google Cloud Console페이지 상단 왼쪽의 Google Cloud 로고 오른쪽에 있는 프로젝트 선택 도구 클릭
* #### 새 프로젝트 버튼 클릭

#### 5-3. 프로젝트 정보 입력
* #### 프로젝트 이름, 위치는 자유롭게 입력 가능
* #### 입력 완료 후 만들기 버튼 클릭

#### 5-4. 해당 프로젝트 선택하기
* #### Google Cloud 로고 오른쪽의 프로젝트 선택 도구 부분에 생성한 프로젝트가 선택되어 있는지 확인하기
* #### 선택되어 있지 않으면 선택 도구를 클릭해 프로젝트를 찾고 선택하기

#### 5-5. 사용자 인증 정보 만들기
* #### Google Cloud 로고 왼쪽의 탐색 메뉴(가로줄 3개 모양)을 선택하고 제품 탭 밑의 API 및 서비스 페이지 클릭
* #### API 및 서비스 페이지 왼쪽에 보이는 하위 메뉴에서 사용자 인증 정보 클릭
* #### 표시된 페이지에서 + 사용자 인증 정보 만들기 버튼을 찾아 클릭하고 표시된 선택지 중 OAuth 클라이언트 ID 선택
* #### 애플리케이션 유형은 데스크톱 앱으로 설정하고 이름 설정 뒤 만들기 버튼 클릭
    * #### ***생성 완료 후 표시되는 정보(`client id, client 보안 비밀번호`)는 다시 볼 수 없으므로 정보들을 따로 안전한 곳에 메모해두고, 반드시 json 파일을 다운로드하기***
* #### 다운로드한 json 파일을 프로젝트 폴더 내부에 저장
    * #### json 파일의 이름은 `client_secret_<client id>.apps.googleusercontent.com.json`와 같은 형태로 되어 있고, 원하는 이름으로 변경해도 문제 없음

#### 5-6. Google 인증 플랫폼 테스트 사용자 설정
* #### [Google 인증 플랫폼](https://console.cloud.google.com/auth/) 페이지에서 대상 하위 메뉴 선택
* #### 표시된 페이지에서 테스트 사용자 섹션 밑의 + Add users 버튼을 클릭하고 구글 캘린더에 일정을 추가하려고 하는 구글 계정을 입력
    * ####  프로젝트 프로그램을 최초 실행할 때, 로컬에 액세스 토큰을 저장하는 과정에 프로그램이 출력한 주소에 접속해 프로그램에서 사용하는 구글 API 기능 권한을 승인하는 과정이 이루어지는 데, 그 때 테스트 사용자에 등록하지 않은 구글 계정은 권한 승인이 불가능

</div>
</details>

### 6. .getker_env 파일 설정
#### 6-1. 프로젝트 폴더 루트 경로(.getker_env_example이 존재하는 경로)에 .getker_env 파일 생성
#### 6-2. .getker_env_example 을 참고하며 .getker_env 내용 작성
```bash
# db connection string
DB_URL="postgres://<username>:<password>@localhost:5432/<dbname>?sslmode=disable"
# 일정을 업로드할 캘린더 id (기본값 primary를 쓰면 로그인한 사용자의 기본 캘린더에 일정이 업로드)
CALENDAR_ID = "primary"
# 5. 에서 생성한 Google Cloud Console 사용자 인증 정보 json 파일 위치
CLIENT_SECRET_FILE_PATH="OAuth 2.0 클라이언트 인증 정보 json 절대경로 또는 실행파일 기준 상대경로"
# OAuth 2.0 인증 과정에서 생성되고 사용될 액세스 토큰 저장 위치
TOKEN_FILE_PATH="로컬 OAuth 2.0 액세스 토큰 절대경로 또는 실행파일 기준 상대경로"
```

### 7. 로컬 액세스 토큰 생성
#### 7-1. 프로그램 최초 실행
```bash
# 프로젝트 루트 폴더에서 실행
go run ./getker
```
* #### 실행하면 `브라우저에서 URL을 열고 인증코드를 입력하세요:`과 `https://accounts.google.com/o/oauth2/auth?access_type=offline&client_id=.....` 형태의 url이 출력되고 사용자의 인증코드 입력을 기다린다
#### 7-2. 표시된 url에 접속
* #### 출력된 url에 웹 브라우저 등을 이용해 접속하면 구글 계정에 로그인하는 페이지가 표시된다. 여기서 5-6에서 등록한 테스트 사용자 구글 계정으로 로그인한다.
* #### 로그인하면 연결되는 새 페이지의 링크에 포함된 인증 코드를 찾아 복사해 사용자 입력을 기다리는 터미널에 입력한다.   
    * #### 링크는 `http://localhost/?state=state-token&code={인증코드}&scope=https://www.googleapis.com/auth/calendar`의 형태로 `code=` 뒤에 나타나는 인증코드를 복사한다.
        * #### 로그인 시 연결되는 `http://localhost`는 Google Cloud Console 사용자 인증 정보 json 파일의 installed 키 안에 redirect_uris 필드에 저장된 값을 따른다.


</div>
</details>


<details>
<summary> <h2> 프로젝트 사용법 </h2> </summary>
<div markdown="1">

```bash
# build 없이 사용할 경우
go run ./getker <commmand name> <argument1> <argument2> ...
```
```bash
# build
go build -o <app_name> ./getker
# 빌드 후 실행
./<app_name> <commmand name> <argument1> <argument2> ...
```
```bash
# install
go install ./getker

# 환경변수 파일 .getker_env 파일을 홈경로/.myprogram에 복사
mkdir -p "$HOME/.myprogram"
cp .getker_env "$HOME/.myprogram/"

# 실행
getker <commmand name> <argument1> <argument2> ...
```

### 명령어
#### `help`
* #### 프로그램 사용법 출력
#### `sites`
* #### `sites` 테이블에 저장된 크롤 가능한 사이트 리스트. `Name`에 표시된 값을 `crawl` 명령어에 사용
#### `crawl`
* #### 주어진 기간 내에 게시된 이벤트 공지 글을 받아 `posts` 테이블에 데이터를 저장
* #### `crawl <siteName> <duration>`과 같은 형태로 크롤링할 사이트 이름과 기간을 같이 입력
    * #### `<siteName>`: 현재 `pokesv`, `epic`, `all` 가능
    * #### `<duration>`: 정수로 크롤링 일수 입력
#### `posts`
* #### 저장된 이벤트 게시글 전부를 리스트로 출력
* #### 추가 옵션
    * #### `ongoing`
        * #### 진행중인 이벤트의 게시글만 리스트로 출력
    * #### `upcoming`
        * #### 진행중이거나 진행 예정인 이벤트의 게시글만 리스트로 출력
    * #### `period <duration>`
        * #### 주어진 기간 내에 게시된 게시글만 리스트로 출력
#### `events`
* #### 이벤트의 종류, 진행 기간 등을 담은 `events` 테이블의 데이터들을 리스트로 출력
    * #### `posts`는 동일한 이벤트에 대한 공지를 여러번 게시한 경우 그 중복 공지들이 전부 표시되지만, `events`는 중복 게시된 이벤트여도 한번만 표시
* #### 추가 옵션
    * #### `r`, `-r`, `register`, `-register`
        * #### `posts` 테이블의 데이터들을 `events` 테이블에 입력한 후 리스트 출력
#### `calendar`
* #### `events` 테이블에 저장된 이벤트들을 구글 캘린더에 입력
* #### 추가 옵션
    * #### `ongoing`
        * #### 진행 중인 이벤트만 입력
    * #### `upcoming`
        * #### 진행 중이거나 진행 예정인 이벤트만 입력
    * #### `wr`
        * #### 시작, 중간, 종료 리마인드를 추가 (기본값)
    * #### `nr`
        * #### 시작, 중간, 종료 리마인드를 미추가
    * #### `or`
        * #### 시작, 중간, 종료 리마인드만 구글 캘린더에 입력
#### `delete`
* #### db의 데이터를 지우는 명령어
* #### 필수 옵션
    * #### `site`
        * #### `sites` 테이블에 저장된 데이터를 삭제. `post`나 `event`와 다르게 전체 삭제 기능 없음
        * #### `site` 필수 옵션
            * #### `name siteName` 또는 `n siteName`
                * #### 이름으로 지정된 사이트 삭제
            * #### `url siteURL` 또는 `u siteURL`
                * #### url로 지정된 사이트 삭제
    * #### `post`
        * #### `posts` 테이블에 저장된 데이터를 삭제
        * #### `post` 추가 옵션
            * #### `old`
                * #### 종료된 이벤트에 관련한 게시글들만 삭제
            * #### `name <siteName>`, `n <siteName>`
                * #### 이름으로 지정한 사이트에서 크롤링한 게시글들만 삭제
            * #### `url <siteURL>`, `u <siteURL>`
                * #### url로 지정한 사이트에서 크롤링한 게시글들만 삭제
            * #### `id <postUUID>`, `ID <postUUID>`
                * #### 해당 `UUID`를 가지는 `post`를 `posts` 테이블에서 삭제
    * #### `event`
        * #### `events` 테이블에 저장된 데이터를 지우고, 해당 데이터가 구글 캘린더에 입력되어 있을 경우 그 구글 캘린더 일정도 삭제
        * #### `event` 추가 옵션
            * #### `id <postUUID>`, `ID <postUUID>`
                * #### 해당 `UUID`를 가지는 `event`를 `events` 테이블에서 삭제하고 구글 캘린더에서도 삭제
#### `addsite`
* #### `sites` 테이블에 데이터를 추가하는 명령어
    * #### `addsite`로 추가한 뒤, `internal/crawler/crawler.go`와 `internal/parser/parser.go`에 해당 사이트 크롤링, 파싱 함수를 추가해야 `crawl` 명령어에서 추가한 사이트로 크롤링 가능
* #### `addsite <siteName> <siteURL>`과 같은 형태로 사이트 이름과 사이트 url 입력


</div>
</details>


## TODO
- [ ] https://sv-news.pokemon.co.jp/ko/page/373.html, https://sv-news.pokemon.co.jp/ko/page/370.html 과 같이 한 게시글에 테라레이드 기간과 이후의 이상한 소포 선물 기간이 같이 있는 경우 이상한 소포 선물 기간의 tag와 tag text가 1, 테라 레이드배틀이 되는 문제 해결?
- [X] 에픽게임즈 스토어 무료게임 공지
- [X] HELP 명령어 추가
- [X] Remind 일정에는 본 일정 기간을 description에 추가하기
- [ ] 확인이 끝난 일정을 체크하고, 체크한 일정만 events 테이블과 calendar에서 삭제하는 기능 추가
    - [ ] event를 지울 때, 그 이벤트와 연관 있는 post들 같이 삭제하는 옵션 추가