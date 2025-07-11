package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/paokimsiwoong/game_event_tracker/internal/database"
)

// DB에 site 추가하는 "addsite" command 입력 시 실행되는 함수
func HandlerAddSite(s *State, cmd Command) error {
	// addsite 뒤의 추가 명령어들은 cmd.args에 저장되어 있음
	// args의 길이가 2이 아니면 addsite <siteName> <siteURL> 형태가 아니므로 에러
	if len(cmd.Args) != 2 {
		return errors.New("the addsite handler expects two arguments, the site name and the site url")
	}

	// 이미 있는 사이트를 또 등록하려 하는 경우
	if _, err := s.PtrDB.GetSiteByURL(context.Background(), cmd.Args[1]); err == nil { // 기존에 있는 사이트라 get하는데 문제없어서 err == nil 이면
		return errors.New("error adding site: can not add existing url")
	}

	site, err := s.PtrDB.CreateSite(
		context.Background(),
		database.CreateSiteParams{
			Name: cmd.Args[0],
			Url:  cmd.Args[1],
		},
	)
	if err != nil {
		return fmt.Errorf("error adding site: %w", err)
	}

	fmt.Printf("site %s was added: %+v\n", cmd.Args[0], site)
	return nil
}
