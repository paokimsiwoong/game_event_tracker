package commands

import (
	"context"
	"fmt"

	"github.com/paokimsiwoong/game_event_tracker/internal/database"
)

// posts 테이블의 post 데이터들을 정리해 events 테이블에 저장하는 핸들러
func HandlerEvents(s *State, cmd Command) error {
	// @@@ events 테이블에 데이터 등록하는 부분 제외하는 옵션 추가
	if len(cmd.Args) != 0 {
		switch cmd.Args[0] {
		case "-r":
			fallthrough
		case "-register":
			fallthrough
		case "register":
			fallthrough
		case "r":
			count := 0

			posts, err := s.PtrDB.GetPostsAndSites(context.Background())
			if err != nil {
				return fmt.Errorf("error getting posts table: %w", err)
			}

			fmt.Println("--------------------------------------------------")
			fmt.Println("--------------------------------------------------")
			fmt.Printf("%d posts in the table\n", len(posts))
			fmt.Println("--------------------------------------------------")
			fmt.Println("--------------------------------------------------")

			if len(posts) == 0 {
				fmt.Println("0 event registered")
				fmt.Println("--------------------------------------------------")
				fmt.Println("--------------------------------------------------")
				return nil
			}

			for _, post := range posts {
				// 이미 등록된 post면 continue
				if post.Registered {
					fmt.Printf("The post %s(url:%s) has already been registered to events table in db\n", post.Name, post.PostUrl)
					continue
				}

				event, err := s.PtrDB.CreateEvent(context.Background(), database.CreateEventParams{
					Tag:      post.Tag,
					TagText:  post.TagText,
					StartsAt: post.StartsAt,
					EndsAt:   post.EndsAt,
					Column5:  post.Name,
					Column6:  post.PostedAt,
					Column7:  post.PostUrl,
					Column8:  post.ID,
					SiteID:   post.SiteID,
				})
				if err != nil {
					return fmt.Errorf("error creating an event: %w", err)
				}

				if len(event.Names) == 1 {
					count++
				}

				// 등록 완료한 post는 registered 칼럼 값 true로 변경
				err = s.PtrDB.SetPostRegisteredTrue(context.Background(), post.ID)
				if err != nil {
					return fmt.Errorf("error updating a post: %w", err)
				}
			}

			fmt.Printf("%d events registered\n", count)

		default:
			fmt.Println("Invalid argument. Will list events without registering")
		}
	}

	fmt.Println("--------------------------------------------------")
	fmt.Println("--------------------------------------------------")

	events, err := s.PtrDB.GetEvents(context.Background())
	if err != nil {
		return fmt.Errorf("error getting events table: %w", err)
	}

	for _, event := range events {
		fmt.Printf(
			"ID: %s\nCreated at: %v\nUpdated at: %v\nTag: %v\nTag text: %v\nStarts at: %v\n",
			event.ID,
			event.CreatedAt,
			event.UpdatedAt,
			event.Tag,
			event.TagText,
			event.StartsAt.Time,
		)
		if event.EndsAt.Valid {
			fmt.Printf("Ends at: %v\n", event.EndsAt.Time)
		} else {
			fmt.Println("Ends at: permanent")
		}
		fmt.Println("Urls:")

		for _, url := range event.PostUrls {
			fmt.Printf("%v\n", url)
		}

		fmt.Println("--------------------------------------------------")
	}
	fmt.Println("--------------------------------------------------")
	fmt.Printf("%d events in events table\n", len(events))
	fmt.Println("--------------------------------------------------")
	fmt.Println("--------------------------------------------------")

	return nil
}
