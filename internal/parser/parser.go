package parser

import (
	"time"

	"github.com/paokimsiwoong/game_event_tracker/internal/crawler"
)

type PokeSVParsedResult struct {
	Title    string
	Kind     string
	KindTxt  string
	Body     string
	PostedAt int64
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

	return nil, nil
}
