package calendar

// 새이벤트는 캘린더에 넣고, 지나간 이벤트는 지울지 남길지 결정?
// 테라레이드 외의 공지는?

// 캘린더에 넣을 일정의 길이, 시간 설정 기능?

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
)

var ErrTemp = errors.New("temp")

// const clientSecretFilePath = "../../client_secret.json"
// const tokFilePath = "../../token.json"

// 사용자 인증 토큰 가져오기
func getToken(config *oauth2.Config, tokFilePath string) (*oauth2.Token, error) {
	tok, err := tokenFromFile(tokFilePath)
	if err != nil {
		// 사용자 인증 토큰이 없을 경우
		tok, err = getTokenFromWeb(config)
		saveToken(tokFilePath, tok)
	}
	// @@@ refresh 토큰 정보가 있으므로 만료된 토큰을 oauth2.TokenSource를 사용해 자동 갱신 가능
	// } else if time.Now().After(tok.Expiry) {
	// 	// tok.Expiry를 확인해 만료된 토큰인지 확인하는 부분
	// 	tok, err = getTokenFromWeb(config)
	// 	saveToken(tokFile, tok)
	// }
	return tok, err
}

// 로컬에 저장된 json 파일에서 토큰정보를 불러와 디코딩하고 그 정보를 저장한 구조체를 반환하는 함수
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	// oauth2.Token 구조체는 불러온 json 파일의 정보를 저장
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// 로컬에 토큰 json이 없거나 만료된 경우 웹 브라우저 인증을 통해 새 토큰을 생성하는 함수
func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	// AuthCodeURL 메서드는 OAuth 2.0 인증 과정에서
	// 사용자로 하여금 권한 승인을 받기 위해 방문해야 할 URL을 생성하는 역할
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	// state 인자:
	// // CSRF(사이트 간 요청 위조) 공격을 방지하기 위해 클라이언트가 임의로 생성해서 넘기는 임의의 문자열
	// // 인증 성공 후 이 값이 그대로 돌아오므로, 클라이언트는 이를 검증해야 한다
	// opts ...AuthCodeOption:
	// // 추가 인증 URL 파라미터(옵션)
	// // // oauth2.AccessTypeOffline (오프라인 액세스, 갱신 토큰 요청)
	// // // oauth2.AccessTypeOnline (기본, 온라인 액세스)
	// // // oauth2.ApprovalForce (사용자에게 동의 화면 강제 표시) 등 필요에 따라 여러 옵션들을 전달 가능
	fmt.Printf("브라우저에서 URL을 열고 인증코드를 입력하세요:\n%v\n", authURL)
	// 사용자가 권한 승인을 마치고 나면, OAuth2 서버는 redirect_uri로 인증 코드를 돌려보낸다.
	// redirect_uri는 client_secret.json의 redirect_uris 필드 값 (현재: "redirect_uris":["http://localhost"])
	var authCode string
	// 사용자가 입력한 인증 코드를 authCode에 저장
	if _, err := fmt.Scan(&authCode); err != nil {
		// log.Fatalf("인증 코드 읽기 실패: %v", err)
		return nil, fmt.Errorf("인증 코드 읽기 실패: %v", err)
	}
	// config.Exchange 메소드는 OAuth 2.0 인증 플로우에서
	// Authorization Code(여기서는 authCode)를 사용해 액세스 토큰을 실제로 발급받는 핵심 함수
	tok, err := config.Exchange(context.TODO(), authCode)
	// 권한 승인 과정에서 사용자에게 부여된 Authorization Code(인증 코드)를
	// Google이나 기타 OAuth2 제공자에게 보내어,
	// 실제 Access Token과 Refresh Token 등의 인증 토큰을 받아오는 역할
	// // 내부적으로 OAuth2 서버의 토큰 엔드포인트에 POST 요청을 보내 인가 코드를 토큰으로 교환
	// // // oauth2.Config 구조체의 Endpoint 필드에 oauth2.Endpoint 구조체로 저장되어 있음
	// // // // AuthURL 필드는 client_secret.json의 auth_uri값, TokenURL 필드는 client_secret.json의 token_uri값
	if err != nil {
		// log.Fatalf("토큰 교환 실패: %v", err)
		return nil, fmt.Errorf("토큰 교환 실패: %v", err)
	}
	return tok, nil
}

// 생성된 token을 로컬에 파일로 저장하는 함수
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("토큰을 저장: %s\n", path)
	f, _ := os.Create(path)
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// refresh 토큰으로 갱신되는 access token을 다시 token.json에 저장하기 위해 필요한
// oauth2.TokenSource를 구현하는 custom 구조체
type savingTokenSource struct {
	src  oauth2.TokenSource
	path string
}

// Token 메소드는 oauth2.TokenSource 인터페이스 구현 조건
// 감싸여진 oauth2.TokenSource의 Token 메소드를 그대로 실행하면서
// 추가로 *oauth2.Token 정보를 로컬에 파일로 저장
func (s *savingTokenSource) Token() (*oauth2.Token, error) {
	t, err := s.src.Token()
	if err == nil {
		f, _ := os.Create(s.path)
		defer f.Close()
		json.NewEncoder(f).Encode(t)
	}
	return t, err
}

// Google Calender API client를 생성하는 함수
func NewCalendar(clientSecretFilePath, tokFilePath string) (*calendar.Service, error) {
	// 1. 인증 정보 로드
	b, err := os.ReadFile(clientSecretFilePath)
	// client_secret.json는 OAuth 클라이언트 인증 정보를 담고 있음
	if err != nil {
		// log.Fatalf("client_secret.json 읽기 실패: %v", err)
		return nil, fmt.Errorf("client_secret.json 읽기 실패: %v", err)
	}
	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	// ConfigFromJSON:
	// // Google API의 OAuth 2.0 인증을 위한 설정 정보를
	// // JSON 형태(보통 Google Cloud Console에서 받은 credentials.json)로부터 파싱하여
	// // oauth2.Config 구조체를 생성하는 역할
	// oauth2.Config:
	// // OAuth 인증과 토큰 발급에 필요한 클라이언트 ID, 클라이언트 시크릿, 인증 URI, 토큰 URI, 리디렉션 URI 등 정보를 포함
	// // 이 정보들은 oauth2 라이브러리에서 인증 URL 생성, 액세스 토큰 발급 등 OAuth 2.0 플로우를 수행할 때 사용하는 핵심 설정값
	// // // ==> config를 통해 config.AuthCodeURL(), config.Exchange() 등 OAuth2 관련 메서드 호출
	// // 추가적으로 Scopes 필드에 ConfigFromJSON 두번째 인자로 입력된 scope 저장. 여기서는 calendar.CalendarScope (Scopes specifies optional requested permissions.)
	// calendar.CalendarScope:
	// // calendar.CalendarScope는 구글 캘린더 API 접근을 위한 권한 범위를 지정하는 상수
	if err != nil {
		// log.Fatalf("OAuth2 config 생성 실패: %v", err)
		return nil, fmt.Errorf("OAuth2 config 생성 실패: %v", err)
	}

	// getToken함수로 OAuth2 access token 생성
	token, err := getToken(config, tokFilePath)
	if err != nil {
		return nil, fmt.Errorf("OAuth2 access token 생성 실패: %v", err)
	}

	// access token이 만료됐을 때 코드에서 refresh token으로 자동 재발급 시도를 하게 하려면
	// oauth2.StaticTokenSource 대신 config.TokenSource로 token을 감싸야한다.
	baseTokenSource := config.TokenSource(context.Background(), token)
	// oauth2.NewClient에 TokenSource를 넘기면
	// 실제 HTTP 요청을 보낼 때마다 access token 만료 여부를 확인하고
	// 만료시 refresh token으로 갱신 후, 새로운 access token을 자동 사용

	// 갱신되는 토큰을 자동으로 로컬에 저장해주는 기능을 추가한 savingTokenSource 구조체로
	// baseTokenSource를 감싸기
	savingTokenSource := &savingTokenSource{src: baseTokenSource, path: tokFilePath}

	// oauth2.NewClient 함수는 context.Context와 TokenSource를 받아서
	// OAuth2 인증을 자동으로 처리하는 *http.Client를 생성하는 함수
	// // (OAuth2 인증 헤더(Authorization: Bearer <token>)를 자동으로 붙여주는 *http.Client)
	client := oauth2.NewClient(context.Background(), savingTokenSource)
	// ctx context.Context:
	// // 요청 취소, 타임아웃 관리 등을 위한 컨텍스트
	// src TokenSource:
	// // OAuth2 토큰을 제공하는 인터페이스(TokenSource)로, 액세스 토큰을 관리(갱신 등)하는 역할
	// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
	// 클라이언트가 HTTP 요청을 할 때마다, TokenSource에서 현재 유효한 액세스 토큰을 얻어 자동으로 Authorization 헤더에 입력
	// 토큰이 만료되면 내부에서 자동으로 토큰을 갱신하여 새 토큰으로 요청
	// 따라서 별도로 토큰 갱신 로직을 구현할 필요 없이, OAuth2 인증된 클라이언트를 사용 가능
	// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

	// client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))
	// @@@ oauth2.StaticTokenSource로 token을 감싸면 refresh 기능이 동작하지 않는다
	// // @@@ StaticTokenSource는 "항상 동일한 토큰만 반환"하며, 토큰이 만료되어도 기존 토큰만 고정적으로 반환

	// 2. Calendar API 서비스 생성
	// calendar.NewService는 Google Calendar API를 Go에서 실제로 사용하기 위한 "API 클라이언트"의 생성 함수
	// 여기에 인증 정보(HTTP 클라이언트, credentials 등)를 옵션으로 넘겨주어야 하며, 반환된 서비스 객체를 기반으로 구글 캘린더의 모든 기능을 호출 가능
	srv, err := calendar.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		// log.Fatalf("Calendar Service 생성 실패: %v", err)
		return nil, fmt.Errorf("calendar service 생성 실패: %v", err)
	}

	return srv, nil
}

// AddEvent에 입력되는 이벤트 데이터를 담는 구조체
// @@@ internal/database의 자동생성된 구조체를 쓰는 대신 여기서 구조체를 정의해서 편하게 수정 가능하도록 하기
type Event struct {
	Tag        int32
	TagText    string
	StartsAt   pgtype.Timestamptz
	EndsAt     pgtype.Timestamptz
	PostNames  []string
	EventUrls  []string
	SiteName   string
	SiteUrl    string
	EventCalID string // google 캘린더에 저장될 때 생성되는 id를 저장하는 필드
}

// 이벤트 데이터를 받아 구글 캘린더에 일정을 추가하는 함수
func AddEvent(srv *calendar.Service, calendarID string, data *Event) error {
	start := data.StartsAt.Time.Format(time.RFC3339)

	var end string
	// 이벤트 종료시점이 없는 경우와 있는 경우 구분
	if data.EndsAt.Valid {
		end = data.EndsAt.Time.Format(time.RFC3339)
	} else {
		end = data.StartsAt.Time.Add(time.Hour).Format(time.RFC3339)
	}

	name := "(" + data.SiteName + ") " + data.TagText

	desc := ""

	for i, name := range data.PostNames {
		desc += name + " (" + data.EventUrls[i] + ")\n"
	}

	event := &calendar.Event{
		Summary:     name,
		Location:    "서울",
		Description: desc,
		Start: &calendar.EventDateTime{
			DateTime: start,
			TimeZone: "Asia/Seoul",
		},
		End: &calendar.EventDateTime{
			DateTime: end,
			TimeZone: "Asia/Seoul",
		},
	}

	event, err := srv.Events.Insert(calendarID, event).Do()
	if err != nil {
		// log.Fatalf("일정 생성 실패: %v", err)
		return fmt.Errorf("일정 생성 실패: %v", err)
	}
	fmt.Printf("이벤트 생성됨: %s\n", event.HtmlLink)

	// srv.Events.Insert 호출 후 반환되는 event안에는 id 부분이 생성되어 있음
	data.EventCalID = event.Id

	return nil
}

// 해당 ID를 가진 이벤트가 구글 캘린더에 있는지 확인하는 함수
func CheckEvent(srv *calendar.Service, calendarID string, data *Event) (bool, error) {
	_, err := srv.Events.Get(calendarID, data.EventCalID).Do()
	if err != nil {
		// get 요청이 실패했을 때, 이벤트가 없어서 에러가 난건지, 그 외 이유인지 구분하기
		if googleAPIErr, ok := err.(*googleapi.Error); ok && googleAPIErr.Code == 404 {
			// @@@ err.(*googleapi.Error)는 type assertion
			// @@@ // Get.Do()가 반환한 err은
			// @@@ // Do 함수 내부에서 error 인터페이스를 구현하는 *googleapi.Error를 error로 반환한것
			// @@@ // // ==> type assertion을 통해 underlying type인 *googleapi.Error로 형변환 가능
			// 404 : Not Found
			return false, nil
		} else {
			// 404(Not Found)가 아닌 이유로 이벤트 조회 자체가 실패한 경우
			return false, err
		}
	}

	return true, nil
}

// 이전에 추가한 이벤트를 구글 캘린더에서 삭제하는 함수
func DeleteEvent(srv *calendar.Service, calendarID string, data *Event) error {
	return srv.Events.Delete(calendarID, data.EventCalID).Do()
}

func TestCalendar(srv *calendar.Service) error {
	var err error
	// 3. 일정(이벤트) 생성 예시
	t := time.Now()
	event := &calendar.Event{
		Summary:     "Go Calendar 연동 테스트",
		Location:    "서울",
		Description: "이것은 Go로 추가한 이벤트입니다.",
		Start: &calendar.EventDateTime{
			DateTime: t.Format(time.RFC3339),
			TimeZone: "Asia/Seoul",
		},
		End: &calendar.EventDateTime{
			DateTime: t.Add(time.Hour).Format(time.RFC3339),
			TimeZone: "Asia/Seoul",
		},
	}

	calendarId := "primary"
	event, err = srv.Events.Insert(calendarId, event).Do()
	if err != nil {
		// log.Fatalf("일정 생성 실패: %v", err)
		return fmt.Errorf("일정 생성 실패: %v", err)
	}
	fmt.Printf("이벤트 생성됨: %s\n", event.HtmlLink)

	// 4. 일정 조회 예시 (향후 10개)
	events, err := srv.Events.List(calendarId).ShowDeleted(false).
		SingleEvents(true).TimeMin("2025-07-01T00:00:00+09:00").MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		// log.Fatalf("일정 조회 실패: %v", err)
		return fmt.Errorf("일정 조회 실패: %v", err)
	}
	fmt.Println("Upcoming events:")
	for _, item := range events.Items {
		date := item.Start.DateTime
		if date == "" {
			date = item.Start.Date
		}
		fmt.Printf("%s (%s)\n", item.Summary, date)
	}

	return nil
	// return ErrTemp
}
