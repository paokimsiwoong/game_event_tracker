package commands

import (
	"context"
	"errors"
	"fmt"
	"slices"
)

// DB에 site 추가하는 "deletesite" command 입력 시 실행되는 함수
func HandlerDeleteSite(s *State, cmd Command) error {
	// deletesite 뒤의 추가 명령어들은 cmd.args에 저장되어 있음
	// args의 길이가 2이 아니면 deletesite <n, name, u, url> <n이면 이름, u면 url> 형태가 아니므로 에러
	if len(cmd.Args) != 2 {
		return errors.New("the deletesite handler expects two arguments, keyword type(one of n, name, u, url) and keyword value")
	}

	// cmd.Args[0] 확인
	allowed := []string{"n", "name", "u", "url"}
	if !slices.Contains(allowed, cmd.Args[0]) {
		return errors.New("the deletesite handler expects two arguments, keyword type(one of n, name, u, url) and keyword value")
	}

	if cmd.Args[0][0] == 'n' {
		site, err := s.PtrDB.GetSiteByName(context.Background(), cmd.Args[1])
		if err != nil {
			return fmt.Errorf("error deleting site: site with the provided name not found: %w", err)
		}

		if err := s.PtrDB.DeleteSiteByName(context.Background(), cmd.Args[1]); err != nil {
			return fmt.Errorf("error deleting site: %w", err)
		}

		fmt.Printf("site %s(url: %s) has been deleted\n", site.Name, site.Url)

		return nil
	}

	site, err := s.PtrDB.GetSiteByURL(context.Background(), cmd.Args[1])

	if err != nil {
		return fmt.Errorf("error deleting site: site with the provided url not found: %w", err)
	}

	if err := s.PtrDB.DeleteSiteByURL(context.Background(), cmd.Args[1]); err != nil {
		return fmt.Errorf("error deleting site: %w", err)
	}

	fmt.Printf("site %s(url: %s) has been deleted\n", site.Name, site.Url)

	return nil
}
