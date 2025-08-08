package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/paokimsiwoong/game_event_tracker/internal/calendar"
	"github.com/paokimsiwoong/game_event_tracker/internal/database"
)

// db에 저장된 event들을 구글 캘린더에 일정으로 등록하는 핸들러
func HandlerCalendar(s *State, cmd Command) error {
	// 추가 명령어 없으면 db에 있는 event 전부 등록
	if len(cmd.Args) == 0 {
		events, err := s.PtrDB.GetEventsAndSite(context.Background())
		if err != nil {
			return fmt.Errorf("error registering calendar events: error getting event data from db: %w", err)
		}

		count := 0

		for _, event := range events {
			// 이미 구글 캘린더에 등록된 경우면 EventCalID 값이 존재
			// => event.EventCalID.Valid이 true
			if event.EventCalID.Valid {
				check, err := calendar.CheckEvent(s.PtrCalSrv, s.PtrCfg.CalendarID, event.EventCalID.String)
				if err != nil {
					return fmt.Errorf("error registering calendar events: error calling calendar.CheckEvent: %w", err)
				}
				if check {
					fmt.Printf("Event %v with cal id %s has already been added to Google Calendar\n", event.ID, event.EventCalID.String)
					continue
				}

				fmt.Printf("Can't find the event %v with cal id %s in Google Calendar\nWill register the event again with a new cal id\n", event.ID, event.EventCalID.String)
			}

			data := &calendar.Event{
				Tag:       event.Tag,
				TagText:   event.TagText,
				StartsAt:  event.StartsAt,
				EndsAt:    event.EndsAt,
				PostNames: event.Names,
				EventUrls: event.PostUrls,
				SiteName:  event.SiteName,
				SiteUrl:   event.SiteUrl,
			}

			err = calendar.AddEvent(s.PtrCalSrv, s.PtrCfg.CalendarID, data)
			if err != nil {
				return fmt.Errorf("error registering calendar events: error calling calendar.AddEvent: %w", err)
			}
			// calendar.AddEvent를 실행하고 나면 data의 EventCalID 필드 값이 차있다
			// // 이 필드 값을 db에 저장해야 한다

			// @@@ db에 구글 캘린더에 등록된 event 아이디(data의 EventCalID 필드)를 입력
			err = s.PtrDB.SetEventCalIDByID(context.Background(), database.SetEventCalIDByIDParams{
				EventCalID: pgtype.Text{
					String: data.EventCalID,
					Valid:  true,
				},
				ID: event.ID,
			})
			if err != nil {
				return fmt.Errorf("error registering calendar events: error updating events table: %w", err)
			}

			fmt.Printf("Event with cal id %s registered\n", data.EventCalID)

			count++
		}

		fmt.Printf("%v events registered in Google Calendar\n", count)

		return nil
	}

	if cmd.Args[0] == "ongoing" {
		events, err := s.PtrDB.GetEventsOnGoing(context.Background())
		if err != nil {
			return fmt.Errorf("error registering calendar events: error getting event data from db: %w", err)
		}

		count := 0

		for _, event := range events {
			// 이미 구글 캘린더에 등록된 경우면 EventCalID 값이 존재
			// => event.EventCalID.Valid이 true
			if event.EventCalID.Valid {
				continue
			}

			data := &calendar.Event{
				Tag:       event.Tag,
				TagText:   event.TagText,
				StartsAt:  event.StartsAt,
				EndsAt:    event.EndsAt,
				PostNames: event.Names,
				EventUrls: event.PostUrls,
				SiteName:  event.SiteName,
				SiteUrl:   event.SiteUrl,
			}

			err = calendar.AddEvent(s.PtrCalSrv, s.PtrCfg.CalendarID, data)
			if err != nil {
				return fmt.Errorf("error registering calendar events: error calling calendar.AddEvent: %w", err)
			}
			// calendar.AddEvent를 실행하고 나면 data의 EventCalID 필드 값이 차있다
			// // 이 필드 값을 db에 저장해야 한다

			// @@@ db에 구글 캘린더에 등록된 event 아이디(data의 EventCalID 필드)를 입력
			err = s.PtrDB.SetEventCalIDByID(context.Background(), database.SetEventCalIDByIDParams{
				EventCalID: pgtype.Text{
					String: data.EventCalID,
					Valid:  true,
				},
				ID: event.ID,
			})
			if err != nil {
				return fmt.Errorf("error registering calendar events: error updating events table: %w", err)
			}

			fmt.Printf("Event with cal id %s registered\n", data.EventCalID)

			count++
		}

		fmt.Printf("%v ongoing events registered in Google Calendar\n", count)

		return nil
	}

	// @@@ TODO: 기간 지정하고 그 기간 내 이벤트 등록하는 부분

	return errors.New("calcal")
}
