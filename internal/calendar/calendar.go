package calendar

// 새이벤트는 캘린더에 넣고, 지나간 이벤트는 지울지 남길지 결정?
// 테라레이드 외의 공지는?

// 캘린더에 넣을 일정의 길이, 시간 설정 기능?

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

var ErrTemp = errors.New("temp")

// 사용자 인증 토큰 가져오기
func getToken(config *oauth2.Config) *oauth2.Token {
	tokFile := "../../token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return tok
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("브라우저에서 URL을 열고 인증코드를 입력하세요:\n%v\n", authURL)
	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("인증 코드 읽기 실패: %v", err)
	}
	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("토큰 교환 실패: %v", err)
	}
	return tok
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("토큰을 저장: %s\n", path)
	f, _ := os.Create(path)
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func PokeCalendar() error {
	// 1. 인증 정보 로드
	b, err := os.ReadFile("../../client_secret.json")
	if err != nil {
		// log.Fatalf("credentials.json 읽기 실패: %v", err)
		return fmt.Errorf("credentials.json 읽기 실패: %v", err)
	}
	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		// log.Fatalf("OAuth2 config 생성 실패: %v", err)
		return fmt.Errorf("OAuth2 config 생성 실패: %v", err)
	}
	client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(getToken(config)))

	// 2. Calendar API 서비스 생성
	srv, err := calendar.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		// log.Fatalf("Calendar Service 생성 실패: %v", err)
		return fmt.Errorf("calendar service 생성 실패: %v", err)
	}

	// 3. 일정(이벤트) 생성 예시
	event := &calendar.Event{
		Summary:     "Go Calendar 연동 테스트",
		Location:    "서울",
		Description: "이것은 Go로 추가한 이벤트입니다.",
		Start: &calendar.EventDateTime{
			DateTime: "2025-07-27T10:00:00+09:00",
			TimeZone: "Asia/Seoul",
		},
		End: &calendar.EventDateTime{
			DateTime: "2025-07-27T11:00:00+09:00",
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

	// return nil
	return ErrTemp
}
