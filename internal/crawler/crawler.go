package crawler

import (
	"encoding/json"
	"fmt"
	"io"
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

	for _, data := range pokeSVJSON.Data {
		stAtUnix, err := strconv.ParseInt(data.StAt, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("error calling strconv.ParseInt: %w", err)
		}
		// 지정한 시점보다 과거의 이벤트는 무시
		if stAtUnix < ts {
			continue
		}

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
