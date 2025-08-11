package commands

import (
	"context"
	"errors"
	"fmt"

	gocal "google.golang.org/api/calendar/v3"
	// @@@ google calendar와 내 internal calendar가 이름이 겹침 => 별칭(alias) 사용하기

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/paokimsiwoong/game_event_tracker/internal/calendar"
	"github.com/paokimsiwoong/game_event_tracker/internal/database"
)

// db에 저장된 event들을 구글 캘린더에 일정으로 등록하는 핸들러
func HandlerCalendar(s *State, cmd Command) error {

	switch len(cmd.Args) {
	case 0:
		return addToCalendar(s, "", "wr")
	case 1:
		switch cmd.Args[0] {
		case "ongoing":
			return addToCalendar(s, "ongoing", "wr")
		case "upcoming":
			return addToCalendar(s, "upcoming", "wr")
		case "nr":
			return addToCalendar(s, "", "nr")
		case "wr":
			return addToCalendar(s, "", "wr")
		case "or":
			return addToCalendar(s, "", "or")
		}
	case 2:
		if cmd.Args[0] == "nr" || cmd.Args[0] == "wr" || cmd.Args[0] == "or" {
			return addToCalendar(s, cmd.Args[1], cmd.Args[0])
		}
		return addToCalendar(s, cmd.Args[0], cmd.Args[1])
	default:
		return errors.New("error invalid command")
	}
	// @@@ TODO: 기간 지정하고 그 기간 내 이벤트 등록하는 부분

	return errors.New("error invalid command")
}

// 주어진 옵션들에 맞게 DB에서 데이터를 불러와 구글 캘린더에 저장하는 함수
func addToCalendar(s *State, opt1, opt2 string) error {
	Adder := map[string]func(*gocal.Service, string, *calendar.Event) error{
		"nr": calendar.AddEvent,
		"wr": calendar.AddEventWithReminds,
		"or": calendar.AddOnlyReminds,
	}

	events, err := getEvents(s, opt1)
	if err != nil {
		return fmt.Errorf("error registering calendar events: error getting event data from db: %w", err)
	}

	count := 0

	for _, event := range events {

		// 이미 구글 캘린더에 등록된 경우면 EventCalID 값이 존재
		// => event.EventCalID.Valid이 true
		if len(event.EventCalIds) != 0 {
			fmt.Printf("Event %v has already been added to Google Calendar\n", event.ID)
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

		addF, ok := Adder[opt2]
		if !ok {
			return errors.New("error registering calendar events: error invalid opt2")
		}

		// err = calendar.AddEvent(s.PtrCalSrv, s.PtrCfg.CalendarID, data)
		err = addF(s.PtrCalSrv, s.PtrCfg.CalendarID, data)
		if err != nil {
			return fmt.Errorf("error registering calendar events: error calling calendar.AddEvent: %w", err)
		}
		// calendar.AddEvent를 실행하고 나면 data의 EventCalID 필드 값이 차있다
		// // 이 필드 값을 db에 저장해야 한다

		if len(data.EventCalIDs) == 0 {
			return errors.New("len(data.EventCalIDs) == 0")
		}

		// @@@ db에 구글 캘린더에 등록된 event 아이디(data의 EventCalID 필드)를 입력
		err = s.PtrDB.SetEventCalIDsByID(context.Background(), database.SetEventCalIDsByIDParams{
			Column1: data.EventCalIDs,
			ID:      event.ID,
		})
		if err != nil {
			return fmt.Errorf("error registering calendar events: error updating events table: %w", err)
		}

		for _, eventCalID := range data.EventCalIDs {
			fmt.Printf("Event %s with cal id %s registered\n", event.ID, eventCalID)
		}

		count += len(data.EventCalIDs)
	}

	fmt.Printf("%v events registered in Google Calendar\n", count)

	return nil
}

type GetEvent struct {
	ID          pgtype.UUID
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
	Tag         int32
	TagText     string
	StartsAt    pgtype.Timestamptz
	EndsAt      pgtype.Timestamptz
	EventCalIds []string
	Names       []string
	PostedAts   []pgtype.Timestamptz
	PostUrls    []string
	PostIds     []pgtype.UUID
	SiteID      pgtype.UUID
	SiteName    string
	SiteUrl     string
}

// DB에서 반환한 구조는 같지만 이름이 다른 구조체들을 하나의 구조체로 변환해서 반환하는 함수
func getEvents(s *State, opt1 string) ([]GetEvent, error) {
	var result []GetEvent

	switch opt1 {
	case "":
		rows, err := s.PtrDB.GetEventsAndSite(context.Background())
		if err != nil {
			return nil, err
		}
		for _, row := range rows {
			result = append(result, GetEvent{
				ID:          row.ID,
				CreatedAt:   row.CreatedAt,
				UpdatedAt:   row.UpdatedAt,
				Tag:         row.Tag,
				TagText:     row.TagText,
				StartsAt:    row.StartsAt,
				EndsAt:      row.EndsAt,
				EventCalIds: row.EventCalIds,
				Names:       row.Names,
				PostedAts:   row.PostedAts,
				PostUrls:    row.PostUrls,
				PostIds:     row.PostIds,
				SiteID:      row.SiteID,
				SiteName:    row.SiteName,
				SiteUrl:     row.SiteUrl,
			})
		}
		return result, nil
	case "ongoing":
		rows, err := s.PtrDB.GetEventsOnGoing(context.Background())
		if err != nil {
			return nil, err
		}
		for _, row := range rows {
			result = append(result, GetEvent{
				ID:          row.ID,
				CreatedAt:   row.CreatedAt,
				UpdatedAt:   row.UpdatedAt,
				Tag:         row.Tag,
				TagText:     row.TagText,
				StartsAt:    row.StartsAt,
				EndsAt:      row.EndsAt,
				EventCalIds: row.EventCalIds,
				Names:       row.Names,
				PostedAts:   row.PostedAts,
				PostUrls:    row.PostUrls,
				PostIds:     row.PostIds,
				SiteID:      row.SiteID,
				SiteName:    row.SiteName,
				SiteUrl:     row.SiteUrl,
			})
		}
		return result, nil
	case "upcoming":
		rows, err := s.PtrDB.GetEventsOnGoingAndUpcoming(context.Background())
		if err != nil {
			return nil, err
		}
		for _, row := range rows {
			result = append(result, GetEvent{
				ID:          row.ID,
				CreatedAt:   row.CreatedAt,
				UpdatedAt:   row.UpdatedAt,
				Tag:         row.Tag,
				TagText:     row.TagText,
				StartsAt:    row.StartsAt,
				EndsAt:      row.EndsAt,
				EventCalIds: row.EventCalIds,
				Names:       row.Names,
				PostedAts:   row.PostedAts,
				PostUrls:    row.PostUrls,
				PostIds:     row.PostIds,
				SiteID:      row.SiteID,
				SiteName:    row.SiteName,
				SiteUrl:     row.SiteUrl,
			})
		}
		return result, nil
	default:
		return nil, errors.New("error invalid opt1")
	}
}
