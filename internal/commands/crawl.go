package commands

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/paokimsiwoong/game_event_tracker/internal/crawler"
	"github.com/paokimsiwoong/game_event_tracker/internal/database"
	"github.com/paokimsiwoong/game_event_tracker/internal/parser"
)

// @@@ TODO: 동일 이벤트가 여러번 공지되는 경우 events table에 여러번 등록되는 문제
// // @@@ ==> calendar 등록시 걸러내는 방식 vs 여기서 걸러내는 방식?
func HandlerCrawl(s *State, cmd Command) error {
	// args의 길이가 2이 아니면 crawl <site name> <duration> 형태가 아니므로 에러
	if len(cmd.Args) != 2 {
		return errors.New("the crawl handler expects two arguments, crawl site name and crawl duration in days")
	}

	var count int

	if cmd.Args[0] == "all" {
		// all 일 경우 모든 사이트 크롤링 진행
		sites, err := s.PtrDB.GetSites(context.Background())
		if err != nil {
			return fmt.Errorf("error crawling site: %w", err)
		}
		duration, err := strconv.Atoi(cmd.Args[1])
		if err != nil {
			return fmt.Errorf("error crawling site: invalid crawl duration: %w", err)
		}

		for _, site := range sites {
			c, err := crawlAndParse(s, cmd, site, duration)
			if err != nil {
				return err
			}
			count += c
		}

	} else {
		site, err := s.PtrDB.GetSiteByName(context.Background(), cmd.Args[0])
		if err != nil {
			return fmt.Errorf("error crawling site: site with the provided name not found: %w", err)
		}

		// @@@ 현재는 포켓몬 스/바 이벤트 일정 crawl, parse 함수 밖에 없음
		duration, err := strconv.Atoi(cmd.Args[1])
		if err != nil {
			return fmt.Errorf("error crawling site: invalid crawl duration: %w", err)
		}

		c, err := crawlAndParse(s, cmd, site, duration)
		if err != nil {
			return err
		}
		count += c
	}

	log.Printf("Total %d posts have been registered", count)

	return nil
}

func crawlAndParse(s *State, cmd Command, site database.Site, duration int) (int, error) {
	var count int

	crawled, err := crawler.Crawl(site.Name, site.Url, duration)
	if err != nil {
		return 0, fmt.Errorf("error crawling site: failed to crawl: %w", err)
	}
	parsed, err := parser.Parse(site.Name, crawled)
	if err != nil {
		return 0, fmt.Errorf("error crawling site: failed to parse: %w", err)
	}

	for _, p := range parsed {
		if posts, err := s.PtrDB.GetPostsByNameAndPostedAtAndSiteID(context.Background(), database.GetPostsByNameAndPostedAtAndSiteIDParams{
			Name: p.Title,
			PostedAt: pgtype.Timestamptz{
				Time:  p.PostedAt,
				Valid: true,
			},
			SiteID: site.ID,
		}); err == nil && len(posts) != 0 { // 기존에 등록된 이벤트일 경우
			log.Printf("Post %s posted at %v has already been registered\n", p.Title, p.PostedAt)
			continue
		}

		// @@@ p.StartsAt이 여러개 들어있는 경우 처리해야함
		for i := 0; i < len(p.StartsAt); i++ {
			count += 1

			if i < len(p.EndsAt) {
				_, err = s.PtrDB.CreatePost(context.Background(), database.CreatePostParams{
					Name:    p.Title,
					Tag:     int32(p.Kind),
					TagText: p.KindTxt,
					PostedAt: pgtype.Timestamptz{
						Time:  p.PostedAt,
						Valid: true,
					},
					StartsAt: pgtype.Timestamptz{
						Time:  p.StartsAt[i],
						Valid: true,
					},
					EndsAt: pgtype.Timestamptz{
						Time:  p.EndsAt[i],
						Valid: true,
					},
					Body:    p.Body,
					PostUrl: p.Url,
					SiteID:  site.ID,
				})
				if err != nil {
					return 0, fmt.Errorf("error creating a post: %w", err)
				}
			} else { // @@@ 종료시점 없는 경우 처리해야함
				_, err = s.PtrDB.CreatePost(context.Background(), database.CreatePostParams{
					Name:    p.Title,
					Tag:     int32(p.Kind),
					TagText: p.KindTxt,
					PostedAt: pgtype.Timestamptz{
						Time:  p.PostedAt,
						Valid: true,
					},
					StartsAt: pgtype.Timestamptz{
						Time:  p.StartsAt[i],
						Valid: true,
					},
					EndsAt: pgtype.Timestamptz{
						Valid: false,
					},
					Body:    p.Body,
					PostUrl: p.Url,
					SiteID:  site.ID,
				})
				if err != nil {
					return 0, fmt.Errorf("error creating a post: %w", err)
				}
			}

		}

	}
	s.PtrDB.MarkSiteFetched(context.Background(), site.ID)

	return count, nil
}
