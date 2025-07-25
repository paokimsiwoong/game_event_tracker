package commands

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"slices"
	"strconv"

	"github.com/paokimsiwoong/game_event_tracker/internal/crawler"
	"github.com/paokimsiwoong/game_event_tracker/internal/database"
	"github.com/paokimsiwoong/game_event_tracker/internal/parser"
)

// @@@ TODO: 동일 이벤트가 여러번 공지되는 경우 events table에 여러번 등록되는 문제
// // @@@ ==> calendar 등록시 걸러내는 방식 vs 여기서 걸러내는 방식?
func HandlerCrawl(s *State, cmd Command) error {
	// args의 길이가 2이 아니면 crawl <n, name, u, url> <n이면 이름, u면 url> 형태가 아니므로 에러
	if len(cmd.Args) != 3 {
		return errors.New("the crawl handler expects three arguments, keyword type(one of n, name, u, url), keyword value, and crawl duration in days")
	}
	// cmd.Args[0] 확인
	allowed := []string{"n", "name", "u", "url"}
	if !slices.Contains(allowed, cmd.Args[0]) {
		return errors.New("the crawl handler expects three arguments, keyword type(one of n, name, u, url), keyword value, and crawl duration in days")
	}

	var site database.Site
	var err error

	if cmd.Args[0][0] == 'n' {
		site, err = s.PtrDB.GetSiteByName(context.Background(), cmd.Args[1])
		if err != nil {
			return fmt.Errorf("error crawling site: site with the provided name not found: %w", err)
		}
	} else {
		site, err = s.PtrDB.GetSiteByURL(context.Background(), cmd.Args[1])
		if err != nil {
			return fmt.Errorf("error crawling site: site with the provided url not found: %w", err)
		}
	}

	// @@@ 현재는 포켓몬 스/바 이벤트 일정 crawl, parse 함수 밖에 없음
	duration, err := strconv.Atoi(cmd.Args[2])
	if err != nil {
		return fmt.Errorf("error crawling site: invalid crawl duration: %w", err)
	}
	crawled, err := crawler.PokeCrawl(site.Url, duration)
	if err != nil {
		return fmt.Errorf("error crawling site: failed to crawl: %w", err)
	}
	parsed, err := parser.PokeParse(crawled)
	if err != nil {
		return fmt.Errorf("error crawling site: failed to parse: %w", err)
	}

	var count int

	for _, p := range parsed {
		if es, err := s.PtrDB.GetEventsByNameAndPostedAtAndSiteID(context.Background(), database.GetEventsByNameAndPostedAtAndSiteIDParams{
			Name: p.Title,
			PostedAt: sql.NullTime{
				Time:  p.PostedAt,
				Valid: true,
			},
			SiteID: site.ID,
		}); err == nil && len(es) != 0 { // 기존에 등록된 이벤트일 경우
			log.Printf("Event %s posted at %v has already been registered\n", p.Title, p.PostedAt)
			continue
		}

		// @@@ p.StartsAt이 여러개 들어있는 경우 처리해야함
		for i := 0; i < len(p.StartsAt); i++ {
			count += 1

			if i < len(p.EndsAt) {
				s.PtrDB.CreateEvent(context.Background(), database.CreateEventParams{
					Name:    p.Title,
					Tag:     int32(p.Kind),
					TagText: p.KindTxt,
					PostedAt: sql.NullTime{
						Time:  p.PostedAt,
						Valid: true,
					},
					StartsAt: sql.NullTime{
						Time:  p.StartsAt[i],
						Valid: true,
					},
					EndsAt: sql.NullTime{
						Time:  p.EndsAt[i],
						Valid: true,
					},
					Body:     p.Body,
					EventUrl: p.Url,
					SiteID:   site.ID,
				})
			} else { // @@@ 종료시점 없는 경우 처리해야함
				s.PtrDB.CreateEvent(context.Background(), database.CreateEventParams{
					Name:    p.Title,
					Tag:     int32(p.Kind),
					TagText: p.KindTxt,
					PostedAt: sql.NullTime{
						Time:  p.PostedAt,
						Valid: true,
					},
					StartsAt: sql.NullTime{
						Time:  p.StartsAt[i],
						Valid: true,
					},
					EndsAt: sql.NullTime{
						Valid: false,
					},
					Body:     p.Body,
					EventUrl: p.Url,
					SiteID:   site.ID,
				})
			}

		}

	}

	log.Printf("Total %d events have been registered", count)

	s.PtrDB.MarkSiteFetched(context.Background(), site.ID)

	return nil
}
