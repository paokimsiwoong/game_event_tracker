package commands

import (
	"fmt"

	"github.com/paokimsiwoong/game_event_tracker/internal/config"
	"github.com/paokimsiwoong/game_event_tracker/internal/database"
	"google.golang.org/api/calendar/v3"
)

// command들이 사용할 config.Config, database.Queries 포인터를 저장하는 구조체
type State struct {
	PtrCfg    *config.Config
	PtrDB     *database.Queries
	PtrCalSrv *calendar.Service
}

// 명령어 한개의 정보를 저장하는 구조체
type Command struct {
	Name string
	Args []string
}

// 사용가능한 명령어들을 저장한 구조체
type Commands struct {
	// map: 명령어 이름 => 해당 명령어 handler 함수
	CommandMap map[string]func(*State, Command) error
}

// 새 명령어 handler를 commands 구조체에 저장하는 메소드
func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.CommandMap[name] = f
}

// commands 구조체에서 주어진 cmd를 찾아 실행하는 메소드
func (c *Commands) Run(s *State, cmd Command) error {
	f, ok := c.CommandMap[cmd.Name]
	if !ok {
		return fmt.Errorf("error no such command: %s", cmd.Name)
	}

	if err := f(s, cmd); err != nil {
		return fmt.Errorf("error running command %s: %w", cmd.Name, err)
	}
	return nil
}
