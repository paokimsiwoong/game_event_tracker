package commands

import (
	"fmt"

	markdown "github.com/MichaelMure/go-term-markdown"
)

func HandlerHelp(s *State, cmd Command) error {
	md := `
# 명령어 리스트
## help
* 프로그램 사용법 출력

## sites
* sites 테이블에 저장된 크롤 가능한 사이트 리스트. Name에 표시된 값을 crawl 명령어에 사용

## crawl
* 주어진 기간 내에 게시된 이벤트 공지 글을 받아 posts 테이블에 데이터를 저장
* crawl <siteName> <duration> 과 같은 형태로 크롤링할 사이트 이름과 기간을 같이 입력
	* <siteName>: 현재 pokesv 또는 epic 가능
	* <duration>: 정수로 크롤링 일수 입력

## posts
* 저장된 이벤트 게시글 전부를 리스트로 출력
* 추가 옵션
	* ongoing
		* 진행중인 이벤트의 게시글만 리스트로 출력
	* upcoming
		* 진행중이거나 진행 예정인 이벤트의 게시글만 리스트로 출력
	* period <duration>
		* 주어진 기간 내에 게시된 게시글만 리스트로 출력

## events
* 이벤트의 종류, 진행 기간 등을 담은 events 테이블의 데이터들을 리스트로 출력
	* posts는 동일한 이벤트에 대한 공지를 여러번 게시한 경우 그 중복 공지들이 전부 표시되지만, events는 중복 게시된 이벤트여도 한번만 표시
* 추가 옵션
	* r 
		* posts 테이블의 데이터들을 events 테이블에 입력한 후 리스트 출력
		* -r, register, -register 도 동일 기능

## calendar
* events 테이블에 저장된 이벤트들을 구글 캘린더에 입력
* 추가 옵션
	* ongoing
		* 진행 중인 이벤트만 입력
	* upcoming
		* 진행 중이거나 진행 예정인 이벤트만 입력
	* wr
		* 시작, 중간, 종료 리마인드를 추가 (기본값)
	* nr
		* 시작, 중간, 종료 리마인드를 미추가
	* or
		* 시작, 중간, 종료 리마인드만 구글 캘린더에 입력

## delete
* db의 데이터를 지우는 명령어
* 필수 옵션
	* site
		* sites 테이블에 저장된 데이터를 삭제. post나 event와 다르게 전체 삭제 기능 없음
		* site 필수 옵션
			* name siteName 또는 n siteName
				* 이름으로 지정된 사이트 삭제
			* url siteURL 또는 u siteURL
				* url로 지정된 사이트 삭제
	* post
		* posts 테이블에 저장된 데이터를 삭제
		* post 추가 옵션
			* old
				* 종료된 이벤트에 관련한 게시글들만 삭제
			* name <siteName>, n <siteName>
				* 이름으로 지정한 사이트에서 크롤링한 게시글들만 삭제
			* url <siteURL>, u <siteURL>
				* url로 지정한 사이트에서 크롤링한 게시글들만 삭제
			* id <postUUID>, ID <postUUID>
				* 해당 UUID를 가지는 post를 posts 테이블에서 삭제
	* event
		* events 테이블에 저장된 데이터를 지우고, 해당 데이터가 구글 캘린더에 입력되어 있을 경우 그 구글 캘린더 일정도 삭제
		* event 추가 옵션
			* id <postUUID>, ID <postUUID>
				* 해당 UUID를 가지는 event를 events 테이블에서 삭제하고 구글 캘린더에서도 삭제

## addsite
* sites 테이블에 데이터를 추가하는 명령어
	* addsite로 추가한 뒤, internal/crawler/crawler.go와 internal/parser/parser.go에 해당 사이트 크롤링, 파싱 함수를 추가해야 crawl 명령어에서 추가한 사이트로 크롤링 가능
* addsite <siteName> <siteURL>과 같은 형태로 사이트 이름과 사이트 url 입력
`

	// markdown.Render 함수는 raw markdown string을 렌더링이 된 마크다운 글로 변경
	rendered := markdown.Render(md, 80, 3)
	// 80은 라인 너비
	// 6은 왼쪽 들여쓰기 값

	fmt.Println(string(rendered))

	return nil
}
