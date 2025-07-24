package commands

import (
	"context"
	"errors"
	"fmt"
	"slices"
)

// "delete" command 입력 시 실행되는 함수
func HandlerDelete(s *State, cmd Command) error {
	// delete 뒤의 추가 명령어들은 cmd.args에 저장되어 있음

	// 첫번째 추가 명령어가 site이면
	if cmd.Args[0] == "site" {
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
	} else if cmd.Args[0] == "event" {
		if len(cmd.Args) == 2 {
			// DeleteOldEvents
			if cmd.Args[1] == "old" {
				err := s.PtrDB.DeleteOldEvents(context.Background())
				if err != nil {
					return fmt.Errorf("error deleting old events: %w", err)
				}
				fmt.Println("Old events have been deleted")

				return nil
			} else if cmd.Args[1] == "all" {
				// ResetEvents
				err := s.PtrDB.ResetEvents(context.Background())
				if err != nil {
					return fmt.Errorf("error deleting all events: %w", err)
				}
				fmt.Println("All events have been deleted")

				return nil
			}
			return errors.New("the delete event command expects one or two additional arguments in the form of <one of old, all, name, n, url, u> <(if name, n) name_value, (if url, u) url_value>")

		} else if len(cmd.Args) == 3 {
			// cmd.Args[1] 확인
			allowed := []string{"n", "name", "u", "url"}
			if !slices.Contains(allowed, cmd.Args[1]) {
				return errors.New("the delete event command expects one or two additional arguments in the form of <one of old, all, name, n, url, u> <(if name, n) name_value, (if url, u) url_value>")
			}
			// DeleteEventsBySiteName
			if cmd.Args[1][0] == 'n' {
				site, err := s.PtrDB.GetSiteByName(context.Background(), cmd.Args[2])
				if err != nil {
					return fmt.Errorf("error deleting event: site with the provided name not found: %w", err)
				}

				if err := s.PtrDB.DeleteEventsBySiteName(context.Background(), cmd.Args[2]); err != nil {
					return fmt.Errorf("error deleting event by site name: %w", err)
				}
				fmt.Printf("events by site %s(url: %s) have been deleted\n", site.Name, site.Url)

				return nil
			}
			// DeleteEventsBySiteUrl
			site, err := s.PtrDB.GetSiteByURL(context.Background(), cmd.Args[2])

			if err != nil {
				return fmt.Errorf("error deleting event: site with the provided url not found: %w", err)
			}

			if err := s.PtrDB.DeleteEventsBySiteUrl(context.Background(), cmd.Args[2]); err != nil {
				return fmt.Errorf("error deleting event by site url: %w", err)
			}

			fmt.Printf("events by site %s(url: %s) has been deleted\n", site.Name, site.Url)

			return nil
		}

		return errors.New("the delete event command expects one or two additional arguments in the form of <one of old, all, name, n, url, u> <(if name, n) name_value, (if url, u) url_value>")
	} else {
		return errors.New("the delete handler expects its first argument to be either site or events")
	}
}

// // DB에 site 추가하는 "deletesite" command 입력 시 실행되는 함수
// func HandlerDeleteSite(s *State, cmd Command) error {
// 	// deletesite 뒤의 추가 명령어들은 cmd.args에 저장되어 있음
// 	// args의 길이가 2이 아니면 deletesite <n, name, u, url> <n이면 이름, u면 url> 형태가 아니므로 에러
// 	if len(cmd.Args) != 2 {
// 		return errors.New("the deletesite handler expects two arguments, keyword type(one of n, name, u, url) and keyword value")
// 	}

// 	// cmd.Args[0] 확인
// 	allowed := []string{"n", "name", "u", "url"}
// 	if !slices.Contains(allowed, cmd.Args[0]) {
// 		return errors.New("the deletesite handler expects two arguments, keyword type(one of n, name, u, url) and keyword value")
// 	}

// 	if cmd.Args[0][0] == 'n' {
// 		site, err := s.PtrDB.GetSiteByName(context.Background(), cmd.Args[1])
// 		if err != nil {
// 			return fmt.Errorf("error deleting site: site with the provided name not found: %w", err)
// 		}

// 		if err := s.PtrDB.DeleteSiteByName(context.Background(), cmd.Args[1]); err != nil {
// 			return fmt.Errorf("error deleting site: %w", err)
// 		}

// 		fmt.Printf("site %s(url: %s) has been deleted\n", site.Name, site.Url)

// 		return nil
// 	}

// 	site, err := s.PtrDB.GetSiteByURL(context.Background(), cmd.Args[1])

// 	if err != nil {
// 		return fmt.Errorf("error deleting site: site with the provided url not found: %w", err)
// 	}

// 	if err := s.PtrDB.DeleteSiteByURL(context.Background(), cmd.Args[1]); err != nil {
// 		return fmt.Errorf("error deleting site: %w", err)
// 	}

// 	fmt.Printf("site %s(url: %s) has been deleted\n", site.Name, site.Url)

// 	return nil
// }
