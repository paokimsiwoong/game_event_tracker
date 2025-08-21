package commands

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/paokimsiwoong/game_event_tracker/internal/calendar"
	"google.golang.org/api/googleapi"
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
		if len(cmd.Args) == 1 {
			// all 없이 post로 끝나도 전체 삭제
			err := s.PtrDB.ResetPosts(context.Background())
			if err != nil {
				return fmt.Errorf("error deleting all posts: %w", err)
			}
			fmt.Println("All posts have been deleted")

			return nil
		}
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

			return errors.New("the delete post command expects one or two additional arguments in the form of <one of old, all, name, n, url, u, id, ID> <(if name, n) name_value, (if url, u) url_value, (if id, ID) id_value>")

		} else if len(cmd.Args) == 3 {
			// cmd.Args[1] 확인
			allowed := []string{"n", "name", "u", "url", "ID", "id"}
			if !slices.Contains(allowed, cmd.Args[1]) {
				return errors.New("the delete post command expects one or two additional arguments in the form of <one of old, all, name, n, url, u, id, ID> <(if name, n) name_value, (if url, u) url_value, (if id, ID) id_value>")
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
			} else if cmd.Args[1][0] == 'u' {
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

			// DeletePostByID
			id, err := uuid.Parse(cmd.Args[2])
			if err != nil {
				return fmt.Errorf("error deleting post: error parsing string uuid: %w", err)
			}

			pgId := pgtype.UUID{
				Bytes: id,
				Valid: true,
			}

			// _, err = s.PtrDB.GetPostByID(context.Background(), pgId)
			// if err != nil {
			// 	return fmt.Errorf("error deleting post: post with the provided id not found: %w", err)
			// }

			if err := s.PtrDB.DeletePostByID(context.Background(), pgId); err != nil {
				return fmt.Errorf("error deleting post by id: %w", err)
			}

			fmt.Printf("post with id %s has been deleted\n", id)

			return nil
		}

		return errors.New("the delete post command expects one or two additional arguments in the form of <one of old, all, name, n, url, u, id, ID> <(if name, n) name_value, (if url, u) url_value, (if id, ID) id_value>")

	// 첫번째 추가 명령어가 event이면
	case "event":
		// id 제공하면 해당 id 가진 이벤트의 구글 일정 삭제
		if len(cmd.Args) == 3 {
			switch cmd.Args[1] {
			case "ID":
				fallthrough
			case "id":
				id, err := uuid.Parse(cmd.Args[2])
				if err != nil {
					return fmt.Errorf("error deleting event: error parsing string uuid: %w", err)
				}

				pgId := pgtype.UUID{
					Bytes: id,
					Valid: true,
				}

				// db에서 이벤트를 삭제하기전에 불러와서 event_cal_id 칼럼의 데이터를 이용해 구글 캘린더에서 삭제부터 하기
				event, err := s.PtrDB.GetEventByID(context.Background(), pgId)
				if err != nil {
					return fmt.Errorf("error deleting calendar events: error getting event data from db: %w", err)
				}

				// @@@@@@@@ event 삭제하면서 posts 테이블의 registered 칼럼 값 다시 false로 변경
				for _, postId := range event.PostIds {
					err = s.PtrDB.SetPostRegisteredFalse(context.Background(), postId)
					if err != nil {
						return fmt.Errorf("error deleting calendar events: error updating a post: %w", err)
					}
				}

				if len(event.EventCalIds) == 0 {
					fmt.Printf("The event %v has not been added to Goole Calendar\n", event.ID)
				} else {
					for _, eventCalID := range event.EventCalIds {

						check, err := calendar.CheckEvent(s.PtrCalSrv, s.PtrCfg.CalendarID, eventCalID)
						if err != nil {
							fmt.Printf("Failed to check if an event %v is in Google Calendar: %v\n", eventCalID, err)
							continue
						}
						if check {
							err = calendar.DeleteEvent(s.PtrCalSrv, s.PtrCfg.CalendarID, eventCalID)
							if err != nil {
								if googleAPIErr, ok := err.(*googleapi.Error); ok && googleAPIErr.Code == 410 {
									// @@@ 이미 지워진 일정이어도 CheckEvent는 통과하는 상황이라 예외처리 필요
									fmt.Printf("The event %v has been deleted: %v\n", eventCalID, err)
									continue
								}
								fmt.Printf("Failed to delete the event %v in Google Calendar: %v\n", eventCalID, err)
								continue
							}

							fmt.Printf("The event %v with cal id %v deleted in Google Calendar\n", event.ID, eventCalID)
						} else {
							fmt.Printf("The event %v with cal id %v is not in Google Calendar\n", event.ID, eventCalID)
						}
					}
				}

				// db events 에서 이벤트 삭제
				err = s.PtrDB.DeleteEventByID(context.Background(), pgId)
				if err != nil {
					return fmt.Errorf("error deleting event: %w", err)
				}

				fmt.Printf("event with id %s has been deleted\n", id)

				return nil
			default:
				return errors.New("invalid command: it must be either delete event or delete event <id or ID> <id value>")
			}
		}

		// 추가 명령어 없으면 전체 삭제

		// db에서 이벤트들을 삭제하기전에 불러와서 event_cal_id 칼럼의 데이터를 이용해 구글 캘린더에서 삭제부터 하기
		events, err := s.PtrDB.GetEventsAndSite(context.Background())
		if err != nil {
			return fmt.Errorf("error deleting calendar events: error getting event data from db: %w", err)
		}

		count := 0

		for _, event := range events {

			// @@@@@@@@ events 테이블 초기화 하면서 posts 테이블의 registered 칼럼 값 다시 false로 변경
			for _, postId := range event.PostIds {
				err = s.PtrDB.SetPostRegisteredFalse(context.Background(), postId)
				if err != nil {
					return fmt.Errorf("error deleting calendar events: error updating a post: %w", err)
				}
			}

			if len(event.EventCalIds) == 0 {
				fmt.Printf("The event %v has not been added to Goole Calendar\n", event.ID)
				continue
			}

			for _, eventCalID := range event.EventCalIds {

				check, err := calendar.CheckEvent(s.PtrCalSrv, s.PtrCfg.CalendarID, eventCalID)
				if err != nil {
					fmt.Printf("Failed to check if an event %v is in Google Calendar: %v\n", eventCalID, err)
					continue
				}
				if check {
					err = calendar.DeleteEvent(s.PtrCalSrv, s.PtrCfg.CalendarID, eventCalID)
					if err != nil {
						if googleAPIErr, ok := err.(*googleapi.Error); ok && googleAPIErr.Code == 410 {
							// @@@ 이미 지워진 일정이어도 CheckEvent는 통과하는 상황이라 예외처리 필요
							fmt.Printf("The event %v has been deleted: %v\n", eventCalID, err)
							continue
						}
						fmt.Printf("Failed to delete the event %v in Google Calendar: %v\n", eventCalID, err)
						continue
					}

					fmt.Printf("The event %v with cal id %v deleted in Google Calendar\n", event.ID, eventCalID)
					count++
				} else {
					fmt.Printf("The event %v with cal id %v is not in Google Calendar\n", event.ID, eventCalID)
				}
			}

		}

		// db events 테이블 초기화
		err = s.PtrDB.ResetEvents(context.Background())
		if err != nil {
			return fmt.Errorf("error reseting events table in db: %w", err)
		}

		fmt.Printf("%v events in db deleted\n %v events in Google Calendar deleted\n", len(events), count)

		return nil
	default:
		return errors.New("the delete handler expects its first argument to be one of site, post, or event")
	}

}
