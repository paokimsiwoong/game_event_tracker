package commands

import (
	"context"
	"fmt"
)

// DB에 저장된 event 목록 보여주는 핸들러
func HandlerEvents(s *State, cmd Command) error {

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
				"Name: %s\nCreated at: %v\nPosted at: %v\nTag: %v\nTag text: %v\nStarts at: %v\nEnds at: %v\nSite name: %v\nSite url: %v\n",
				event.Name,
				event.CreatedAt,
				event.PostedAt.Time,
				event.Tag,
				event.TagText,
				event.StartsAt.Time,
				event.EndsAt.Time,
				event.SiteName,
				event.SiteUrl,
			)
		} else {
			fmt.Printf(
				"Name: %s\nCreated at: %v\nPosted at: %v\nTag: %v\nTag text: %v\nStarts at: %v\nEnds at: permanent\nSite name: %v\nSite url: %v\n",
				event.Name,
				event.CreatedAt,
				event.PostedAt.Time,
				event.Tag,
				event.TagText,
				event.StartsAt.Time,
				event.SiteName,
				event.SiteUrl,
			)
		}
		fmt.Println("--------------------------------------------------")
	}
	fmt.Println("--------------------------------------------------")

	return nil
}
