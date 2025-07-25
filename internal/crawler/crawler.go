package crawler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"sync"
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
	Url     string
	StAt    int64
	Success bool
}

type PokeSVIntermediateResult struct {
	Title   string
	Kind    string
	KindTxt string
	StAt    int64
}

const pokeURL = "https://sv-news.pokemon.co.jp/ko/"

// const pokeURL = "https://sv-news.pokemon.co.jp/ko/json/list.json"

// 포켓몬 스/바 이벤트 일정 crawl 함수
func PokeCrawl(url string, duration int) ([]PokeSVResult, error) {
	// 이 함수 안에서 여러번 get 요청을 하므로 http.Client를 생성해 하나만 사용하기
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	// @@@ 여러 go 루틴에서 사용될 때 동일 인스턴스가 사용되어야 하므로
	// @@@ 포인터로 생성
	// // @@@ 값으로 생성하고 &를 써도 되지만,
	// // @@@ 여러번 사용하다가 실수로 &를 빼먹어 client가 복사되는 것을 원천 방지하기 위해
	// // @@@ 처음부터 포인터로 생성한다

	resp, err := client.Get(url)
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

	// 여러개의 go 루틴들이 모두 끝날때까지 기다리도록 해주는 sync.WaitGroup 생성
	var wg sync.WaitGroup
	// go 루틴들의 결과를 안전하게 받기 위한 채널 생성
	ch := make(chan PokeSVResult, len(pokeSVJSON.Data))

	// 너무 많은 goroutine이 동시에 실행되면 시스템 리소스가 부족해지거나
	// 리퀘스트를 받는 서버(API 등)에 동시에 리퀘스트 가능한 수가 제한되어 있거나 할 수 있음.
	// 이럴 때는 세마포어 패턴(버퍼 채널로 동시 실행 개수 제한)을 활용
	sem := make(chan struct{}, 24) // 동시에 최대 24개만 실행

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
		// t := time.Unix(stAtUnix, 0)
		// kst := t.In(time.FixedZone("KST", 9*60*60))
		// time.FixedZone은 지정한 이름과 UTC 기준 오프셋을 입력하여 만드는 커스텀 시간대
		// kst := t.Local()
		// log.Printf("data title: %v\nposted at: %v\n", data.Title, kst)

		wg.Add(1)
		go fetchURL(client, pokeURL+data.Link, ch, &wg, sem, PokeSVIntermediateResult{
			Title:   data.Title,
			Kind:    data.Kind,
			KindTxt: data.KindTxt,
			StAt:    stAtUnix,
		})
	}

	// wg.Wait()은 wg.Add(1) 로 알고 있는 go 루틴 개수만큼 wg.Done()가 실행될때까지 block
	wg.Wait()
	// 데이터가 더 들어오지 않을 ch close
	close(ch)

	// 닫힌 채널로 for 루프를 돌면 그 채널 안에 들어있는 데이터들을 다 돌고나서 for 루프가 자동 종료된다
	for r := range ch {
		result = append(result, r)
	}

	// 결과 정렬

	// 시간순 정렬 확실하게 하기위해 정렬 실행
	sort.Slice(result, func(i, j int) bool {
		// return result[i].StAt < result[j].StAt
		return result[i].StAt > result[j].StAt
		// @@@ 최신이 제일 앞에 오도록 변경
	})

	return result, nil
}

// 주어진 url에 get 리퀘스트를 하고 그 결과를 채널에 입력하는 함수
// go 루틴으로 동시에 여러개가 실행되어도 문제 없도록 설계되어 있음
func fetchURL(client *http.Client, url string, ch chan<- PokeSVResult, wg *sync.WaitGroup, sem chan struct{}, intermediate PokeSVIntermediateResult) {
	// go 루틴으로 실행된 함수가 종료 되었음을 함수 외부에 알리기 위해
	// 함수 종료 시 wg.Done() 실행
	defer wg.Done()

	// 세마포어 획득
	sem <- struct{}{}
	// @@@ sem 채널에 입력 시도 => 채널 버퍼가 꽉차있으면 여기서 블락되고 버퍼가 비워질 때까지 기다린다

	// 세마포어 반환을 defer
	defer func() { <-sem }()
	// @@@ 함수가 종료되면서 sem 채널안의 버퍼 한칸을 비운다
	// @@@ => 블락되어 있던 다른 go 루틴에서 sem <- struct{}{} 실행 가능해져 블락이 풀린다

	dataResp, err := client.Get(url)
	if err != nil {
		ch <- PokeSVResult{
			Title:   intermediate.Title,
			Kind:    intermediate.Kind,
			KindTxt: intermediate.KindTxt,
			StAt:    intermediate.StAt,
			Body:    "Failed to make a GET request",
			Url:     url,
			Success: false,
		}
		return
	}
	defer dataResp.Body.Close()

	rawBytes, err := io.ReadAll(dataResp.Body)
	if err != nil {
		ch <- PokeSVResult{
			Title:   intermediate.Title,
			Kind:    intermediate.Kind,
			KindTxt: intermediate.KindTxt,
			StAt:    intermediate.StAt,
			Body:    "Failed to read a response",
			Url:     url,
			Success: false,
		}
		return
	}

	ch <- PokeSVResult{
		Title:   intermediate.Title,
		Kind:    intermediate.Kind,
		KindTxt: intermediate.KindTxt,
		StAt:    intermediate.StAt,
		Body:    string(rawBytes),
		Url:     url,
		Success: true,
	}
}
