package commands

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/paokimsiwoong/game_event_tracker/internal/calendar"
)

// "delete" command 입력 시 실행되는 함수
func HandlerDelete(s *State, cmd Command) error {
	// delete 뒤의 추가 명령어들은 cmd.args에 저장되어 있음

	switch cmd.Args[0] {
	// 첫번째 추가 명령어가 site이면
	case "site":
		// args의 길이가 3이 아니면 delete site <n, name, u, url> <n이면 이름, u면 url> 형태가 아니므로 에러
		if len(cmd.Args) != 3 {
			return errors.New("the delete site command expects two additional arguments, keyword type(one of n, name, u, url) and keyword value")
		}

		// cmd.Args[1] 확인
		allowed := []string{"n", "name", "u", "url"}
		if !slices.Contains(allowed, cmd.Args[1]) {
			return errors.New("the delete site command expects two additional arguments, keyword type(one of n, name, u, url) and keyword value")
		}

		if cmd.Args[1][0] == 'n' {
			site, err := s.PtrDB.GetSiteByName(context.Background(), cmd.Args[2])
			if err != nil {
				return fmt.Errorf("error deleting site: site with the provided name not found: %w", err)
			}

			if err := s.PtrDB.DeleteSiteByName(context.Background(), cmd.Args[2]); err != nil {
				return fmt.Errorf("error deleting site by name: %w", err)
			}

			fmt.Printf("site %s(url: %s) has been deleted\n", site.Name, site.Url)

			return nil
		}

		site, err := s.PtrDB.GetSiteByURL(context.Background(), cmd.Args[2])

		if err != nil {
			return fmt.Errorf("error deleting site: site with the provided url not found: %w", err)
		}

		if err := s.PtrDB.DeleteSiteByURL(context.Background(), cmd.Args[2]); err != nil {
			return fmt.Errorf("error deleting site by url: %w", err)
		}

		fmt.Printf("site %s(url: %s) has been deleted\n", site.Name, site.Url)

		return nil

	// 첫번째 추가 명령어가 post이면
	case "post":
		if len(cmd.Args) == 2 {
			switch cmd.Args[1] {
			// DeleteOldPosts
			case "old":
				err := s.PtrDB.DeleteOldPosts(context.Background())
				if err != nil {
					return fmt.Errorf("error deleting old posts: %w", err)
				}
				fmt.Println("Old posts have been deleted")

				return nil

			// ResetPosts
			case "all":
				err := s.PtrDB.ResetPosts(context.Background())
				if err != nil {
					return fmt.Errorf("error deleting all posts: %w", err)
				}
				fmt.Println("All posts have been deleted")

				return nil
			}

			return errors.New("the delete post command expects one or two additional arguments in the form of <one of old, all, name, n, url, u> <(if name, n) name_value, (if url, u) url_value>")

		} else if len(cmd.Args) == 3 {
			// cmd.Args[1] 확인
			allowed := []string{"n", "name", "u", "url"}
			if !slices.Contains(allowed, cmd.Args[1]) {
				return errors.New("the delete post command expects one or two additional arguments in the form of <one of old, all, name, n, url, u> <(if name, n) name_value, (if url, u) url_value>")
			}
			// DeletePostsBySiteName
			if cmd.Args[1][0] == 'n' {
				site, err := s.PtrDB.GetSiteByName(context.Background(), cmd.Args[2])
				if err != nil {
					return fmt.Errorf("error deleting post: site with the provided name not found: %w", err)
				}

				if err := s.PtrDB.DeletePostsBySiteName(context.Background(), cmd.Args[2]); err != nil {
					return fmt.Errorf("error deleting post by site name: %w", err)
				}
				fmt.Printf("posts by site %s(url: %s) have been deleted\n", site.Name, site.Url)

				return nil
			}
			// DeletePostsBySiteUrl
			site, err := s.PtrDB.GetSiteByURL(context.Background(), cmd.Args[2])

			if err != nil {
				return fmt.Errorf("error deleting post: site with the provided url not found: %w", err)
			}

			if err := s.PtrDB.DeletePostsBySiteUrl(context.Background(), cmd.Args[2]); err != nil {
				return fmt.Errorf("error deleting post by site url: %w", err)
			}

			fmt.Printf("posts by site %s(url: %s) has been deleted\n", site.Name, site.Url)

			return nil
		}

		return errors.New("the delete post command expects one or two additional arguments in the form of <one of old, all, name, n, url, u> <(if name, n) name_value, (if url, u) url_value>")
	case "event":
		// db에서 이벤트들을 삭제하기전에 불러와서 event_cal_id 칼럼의 데이터를 이용해 구글 캘린더에서 삭제부터 하기
		events, err := s.PtrDB.GetEventsAndSite(context.Background())
		if err != nil {
			return fmt.Errorf("error deleting calendar events: error getting event data from db: %w", err)
		}

		for _, event := range events {
			data := &calendar.Event{
				Tag:        event.Tag,
				TagText:    event.TagText,
				StartsAt:   event.StartsAt,
				EndsAt:     event.EndsAt,
				PostNames:  event.Names,
				EventUrls:  event.PostUrls,
				SiteName:   event.SiteName,
				SiteUrl:    event.SiteUrl,
				EventCalID: event.EventCalID.String,
			}

			check, err := calendar.CheckEvent(s.PtrCalSrv, s.PtrCfg.CalendarID, data)
			if err != nil {
				fmt.Printf("Failed to check if an event %v is in Google Calendar: %v\n", data.EventCalID, err)
				continue
			}

			if check {
				err = calendar.DeleteEvent(s.PtrCalSrv, s.PtrCfg.CalendarID, data)
				if err != nil {
					fmt.Printf("Failed to delete the event %v in Google Calendar: %v\n", data.EventCalID, err)
					continue
				}
				fmt.Printf("The event %v deleted in Google Calendar\n", data.EventCalID)
			} else {
				fmt.Printf("The event %v is not in Google Calendar\n", data.EventCalID)
			}
		}

		// db events 테이블 초기화
		err = s.PtrDB.ResetEvents(context.Background())
		if err != nil {
			return fmt.Errorf("error reseting events table in db: %w", err)
		}

		return nil
	default:
		return errors.New("the delete handler expects its first argument to be either site or events")
	}

}
