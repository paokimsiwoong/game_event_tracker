package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/paokimsiwoong/game_event_tracker/internal/crawler"
)

func Parse(opt string, input []crawler.PokeSVResult) ([]PokeSVParsedResult, error) {
	switch opt {
	case "pokesv":
		return PokeParse(input)
	// case "epic":
	default:
		return nil, errors.New("parse function for provided opt not yet implemented")
	}
}

type PokeSVParsedResult struct {
	Title string
	// Kind    string // @@@ int로 변환
	Kind    int
	KindTxt string
	Body    string
	Url     string
	// PostedAt int64 // @@@ 모두 time.Time으로 통일
	PostedAt time.Time
	StartsAt []time.Time // @@@ 이벤트가 두 기간에 걸쳐있는 경우도 있으므로 [] 슬라이스
	EndsAt   []time.Time
}

// 포켓몬 스/바 이벤트 내용에서 기간을 얻어내는 파싱함수
// func PokeParse(body string) ([]string, error) {
func PokeParse(input []crawler.PokeSVResult) ([]PokeSVParsedResult, error) {
	// <h1>기간</h1>과 다음 <h1> 사이의 내용 추출 필요
	// @@@ 기간 대신 엔트리 기간, 개최 기간 같은 변형(ex: 25.04.03)도 있고 기간이 없는 공지(ex: 24.12.23)도 있음
	// PokeSVResult의 Kind, KindTxt 필드 활용 필요
	// | kind 값 | kindTxt 값      |
	// |:--------|:---------------|
	// | 1       | 테라 레이드배틀 |
	// | 2       | 랭크배틀        | // 필요없음
	// | 3       | 레귤레이션      | // 필요없음
	// | 4       | 인터넷 대회     | // 필요없음
	// | 5       | 이상한 소포     |
	// | 6       | 기타 소식       | // @@@ 기타 소식에 【예고】검은 레쿠쟈 강림! 특별한 이벤트 테라 레이드배틀과 이벤트 대량발생도 개최! 와 같이 여러 이벤트 공지가 합쳐진 경우도 있음
	// | 8       | 대량발생        |

	// 결과 담을 슬라이스 선언
	var output []PokeSVParsedResult

	// 문자열 매칭에 사용할 정규표현식 패턴 컴파일
	re, err := regexp.Compile("<h1>.*기간.*?<h1>")
	if err != nil {
		return nil, fmt.Errorf("error compiling regexp: %w", err)
	}
	// ex: 2025년 7월 11일(금) 9:00~7월 14일(월) 8:59 과 같은 문자열 찾기
	rere, err := regexp.Compile(`\d{4}년\s*\d{1,2}월\s*\d{1,2}일.*?\d{1,2}:\d{2}~.*?<br />`)
	if err != nil {
		return nil, fmt.Errorf("error compiling regexp: %w", err)
	}

	// 2025년 7월 11일(금) 9:00 과 같은 형태의 문자열에서 숫자 추출에 사용
	rerere, err := regexp.Compile(`(\d{4})년\s*(\d{1,2})월\s*(\d{1,2})일.*?(\d{1,2}):(\d{2})`)
	if err != nil {
		return nil, fmt.Errorf("error compiling regexp: %w", err)
	}
	// ~ 뒤의 7월 14일(월) 8:59 과 같이 년도 표시가 없는 문자열에서 숫자 추출에 사용
	rererere, err := regexp.Compile(`(\d{1,2})월\s*(\d{1,2})일.*?(\d{1,2}):(\d{2})`)
	if err != nil {
		return nil, fmt.Errorf("error compiling regexp: %w", err)
	}

	// time.Date에 쓰일 KST 시간대 설정
	loc := time.FixedZone("KST", 9*60*60)

	for _, r := range input {
		if r.Kind == "2" || r.Kind == "3" || r.Kind == "4" {
			continue
		}

		match := re.FindString(r.Body)
		// match 못찾았을 경우 예외 처리
		if match == "" {
			continue
		}

		// log.Printf("re match: %s", match)

		// int64 유닉스 타임을 time.Time으로 변환
		stAt := time.Unix(r.StAt, 0)
		postedAt := stAt.In(loc)

		// r.Kind의 string을 int로 변환
		kind, _ := strconv.Atoi(r.Kind)

		o := PokeSVParsedResult{
			Title:    r.Title,
			Kind:     kind,
			KindTxt:  r.KindTxt,
			Body:     r.Body,
			Url:      r.Url,
			PostedAt: postedAt,
		}

		// <h1>.*기간.*<\h1> 뒤의 시간 찾기
		// split := strings.Split(match, "</h1>")
		// @@@ rere로 시간 추출하므로 굳이 split 필요 없음

		// log.Printf("split 0: %s", split[0])

		// ex: 2025년 7월 11일(금) 9:00~7월 14일(월) 8:59 과 같은 형태의 문자열 찾기
		timeMatches := rere.FindAllString(match, -1)

		for _, m := range timeMatches {
			split := strings.Split(m, "~")

			// 2025년 7월 11일(금) 9:00 과 같은 형태의 문자열에서 숫자 추출
			start := rerere.FindStringSubmatch(split[0])
			if len(start) < 6 {
				// FindStringSubmatch의 결과의 0번 인덱스는 매칭되는 문자열 전체가 저장되어 있고
				// 1번 인덱스부터 subgroup들이 저장되어 있다
				// // `(\d{4})년\s*(\d{1,2})월\s*(\d{1,2})일.*?(\d{1,2}):(\d{2})`는
				// // 5개의 subgroup이 있으므로 매칭결과의 길이는 6(매칭된 전체문자열, subgroup 5개)이어야 한다
				fmt.Println("시간 포맷 불일치")
				continue
			}
			year, _ := strconv.Atoi(start[1])
			month, _ := strconv.Atoi(start[2])
			day, _ := strconv.Atoi(start[3])
			hour, _ := strconv.Atoi(start[4])
			min, _ := strconv.Atoi(start[5])

			t := time.Date(year, time.Month(month), day, hour, min, 0, 0, loc)

			o.StartsAt = append(o.StartsAt, t)

			end := rerere.FindStringSubmatch(split[1])
			if len(end) < 6 {
				// 연말에서 연초에 걸치는 이벤트가 아닌 경우
				// 이벤트 종료일시에는 년도가 생략되므로
				// 대부분의 이벤트는 len(end) < 6 조건이 참

				// 7월 14일(월) 8:59 과 같이 년도 표시가 없는 문자열에서 숫자 추출
				end = rererere.FindStringSubmatch(split[1])
				if len(end) < 5 {
					// 이벤트 시작 후 상시 진행 이벤트라서 종료 일시가 없는 경우
					continue
				}

				month, _ = strconv.Atoi(end[1])
				day, _ = strconv.Atoi(end[2])
				hour, _ = strconv.Atoi(end[3])
				min, _ = strconv.Atoi(end[4])
			} else {
				year, _ = strconv.Atoi(end[1])
				month, _ = strconv.Atoi(end[2])
				day, _ = strconv.Atoi(end[3])
				hour, _ = strconv.Atoi(end[4])
				min, _ = strconv.Atoi(end[5])
			}

			t = time.Date(year, time.Month(month), day, hour, min, 0, 0, loc)

			o.EndsAt = append(o.EndsAt, t)
		}

		output = append(output, o)
	}

	return output, nil
}
