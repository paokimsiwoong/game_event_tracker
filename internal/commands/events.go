package commands

import (
	"context"
	"errors"
	"fmt"
)

// DB에 저장된 event 목록 보여주는 핸들러
func HandlerEvents(s *State, cmd Command) error {
	// 추가 명령어 없으면 db에 있는 이벤트 전부 출력
	if len(cmd.Args) == 1 && cmd.Args[0] == "all" {
		events, err := s.PtrDB.GetEventsAndSites(context.Background())
		if err != nil {
			return fmt.Errorf("error getting events table : %w", err)
		}

		fmt.Println("--------------------------------------------------")
		fmt.Println("--------------------------------------------------")
		fmt.Printf("%d events in the table\n", len(events))
		fmt.Println("--------------------------------------------------")
		fmt.Println("--------------------------------------------------")

		if len(events) == 0 {
			return nil
		}

		for _, event := range events {
			if event.EndsAt.Valid {
				fmt.Printf(
					"Name: %s\nCreated at: %v\nPosted at: %v\nTag: %v\nTag text: %v\nStarts at: %v\nEnds at: %v\nEvent url: %v\nSite name: %v\nSite url: %v\n",
					event.Name,
					event.CreatedAt,
					event.PostedAt.Time,
					event.Tag,
					event.TagText,
					event.StartsAt.Time,
					event.EndsAt.Time,
					event.EventUrl,
					event.SiteName,
					event.SiteUrl,
				)
			} else {
				fmt.Printf(
					"Name: %s\nCreated at: %v\nPosted at: %v\nTag: %v\nTag text: %v\nStarts at: %v\nEnds at: permanent\nEvent url: %v\nSite name: %v\nSite url: %v\n",
					event.Name,
					event.CreatedAt,
					event.PostedAt.Time,
					event.Tag,
					event.TagText,
					event.StartsAt.Time,
					event.EventUrl,
					event.SiteName,
					event.SiteUrl,
				)
			}
			fmt.Println("--------------------------------------------------")
		}
		fmt.Println("--------------------------------------------------")

		return nil

	} else if len(cmd.Args) == 1 && cmd.Args[0] == "ongoing" {
		events, err := s.PtrDB.GetEventsOnGoingAndSites(context.Background())
		if err != nil {
			return fmt.Errorf("error getting events table : %w", err)
		}

		fmt.Println("--------------------------------------------------")
		fmt.Println("--------------------------------------------------")
		fmt.Printf("%d ongoing events in the table\n", len(events))
		fmt.Println("--------------------------------------------------")
		fmt.Println("--------------------------------------------------")

		if len(events) == 0 {
			return nil
		}

		for _, event := range events {
			if event.EndsAt.Valid {
				fmt.Printf(
					"Name: %s\nCreated at: %v\nPosted at: %v\nTag: %v\nTag text: %v\nStarts at: %v\nEnds at: %v\nEvent url: %v\nSite name: %v\nSite url: %v\n",
					event.Name,
					event.CreatedAt,
					event.PostedAt.Time,
					event.Tag,
					event.TagText,
					event.StartsAt.Time,
					event.EndsAt.Time,
					event.EventUrl,
					event.SiteName,
					event.SiteUrl,
				)
			} else {
				fmt.Printf(
					"Name: %s\nCreated at: %v\nPosted at: %v\nTag: %v\nTag text: %v\nStarts at: %v\nEnds at: permanent\nEvent url: %v\nSite name: %v\nSite url: %v\n",
					event.Name,
					event.CreatedAt,
					event.PostedAt.Time,
					event.Tag,
					event.TagText,
					event.StartsAt.Time,
					event.EventUrl,
					event.SiteName,
					event.SiteUrl,
				)
			}
			fmt.Println("--------------------------------------------------")
		}
		fmt.Println("--------------------------------------------------")

		return nil
	} else {
		return errors.New("the events handler expects one argument. If all is given, print all events, if ongoing is given, print ongoing events)")
	}
}
