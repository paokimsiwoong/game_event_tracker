package crawler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type PokeSVJSON struct {
	Hash string `json:"hash"`
	Data []struct {
		ID          string `json:"id"`
		Reg         string `json:"reg"`
		Title       string `json:"title"`
		Kind        string `json:"kind"`
		KindTxt     string `json:"kindTxt"`
		Banner      string `json:"banner"`
		IsImportant string `json:"isImportant"`
		StAt        string `json:"stAt"`
		NewAt       string `json:"newAt"`
		Link        string `json:"link"`
	} `json:"data"`
}

type PokeSVResult struct {
	Title   string
	Kind    string
	KindTxt string
	Body    string
	StAt    int64
}

const pokeURL = "https://sv-news.pokemon.co.jp/ko/"

// const pokeURL = "https://sv-news.pokemon.co.jp/ko/json/list.json"

func PokeCrawl(url string, duration int) ([]PokeSVResult, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error calling http.Get: %w", err)
	}
	defer resp.Body.Close()

	// resp body의 json을 담을 구조체 선언
	var pokeSVJSON PokeSVJSON

	// 디코딩
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&pokeSVJSON); err != nil {
		return nil, fmt.Errorf("error decoding resp body: %w", err)
	}

	// json data 항목안에 들어있는 이벤트 게시글들 html 데이터들 모아서 담을 구조체 선언
	var result []PokeSVResult

	// data의 stAt과 비교할 Unix 타임스탬프 생성하기(Unix 타임스탬프는 UTC 기준)
	now := time.Now()
	checkPoint := now.AddDate(0, 0, -duration)
	ts := checkPoint.Unix()

	// 시간순 정렬 확실하게 하기위해 미리 정렬 실행
	// sort.Slice(pokeSVJSON.Data, func(i, j int) bool {
	// 	// return pokeSVJSON.Data[i].StAt < pokeSVJSON.Data[j].StAt
	// 	ith, _ := strconv.ParseInt(pokeSVJSON.Data[i].StAt, 10, 64)
	// 	jth, _ := strconv.ParseInt(pokeSVJSON.Data[j].StAt, 10, 64)
	//  @@@ 변환 두번 할 필요 없도록 stAt의 타입이 int 64인 새 struct를 만들어
	//  @@@ stAt의 타입 변환 후 저장한 것을 가지고 정렬?
	// 	return ith < jth
	// })

	for _, data := range pokeSVJSON.Data {

		stAtUnix, err := strconv.ParseInt(data.StAt, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("error calling strconv.ParseInt: %w", err)
		}

		// 지정한 시점보다 과거의 이벤트는 무시
		if stAtUnix < ts {
			// continue
			// 일단 정렬이 되어있으므로 break로 변경?
			// 확실하게 정렬 추가?
			break
		}

		// 임시로 data 타이틀과 날짜 확인
		t := time.Unix(stAtUnix, 0)
		// kst := t.In(time.FixedZone("KST", 9*60*60))
		// time.FixedZone은 지정한 이름과 UTC 기준 오프셋을 입력하여 만드는 커스텀 시간대
		kst := t.Local()
		log.Printf("data title: %v\nposted at: %v\n", data.Title, kst)

		dataResp, err := http.Get(pokeURL + data.Link)
		if err != nil {
			return nil, fmt.Errorf("error calling http.Get data: %w", err)
		}

		rawBytes, err := io.ReadAll(dataResp.Body)
		if err != nil {
			return nil, fmt.Errorf("error calling io.ReadAll: %w", err)
		}

		result = append(result, PokeSVResult{
			Title:   data.Title,
			Kind:    data.Kind,
			KindTxt: data.KindTxt,
			Body:    string(rawBytes),
			StAt:    stAtUnix,
		})

		dataResp.Body.Close()
	}

	return result, nil
}
