package commands

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// DB에 저장된 post 목록 보여주는 핸들러
func HandlerPosts(s *State, cmd Command) error {
	// 추가 명령어 없으면 db에 있는 post 전부 출력
	if len(cmd.Args) == 0 {
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
			return nil
		}

		for _, post := range posts {
			if post.EndsAt.Valid {
				fmt.Printf(
					"Name: %s\nCreated at: %v\nPosted at: %v\nTag: %v\nTag text: %v\nStarts at: %v\nEnds at: %v\nPost url: %v\nSite name: %v\nSite url: %v\n",
					post.Name,
					post.CreatedAt,
					post.PostedAt,
					post.Tag,
					post.TagText,
					post.StartsAt.Time,
					post.EndsAt.Time,
					post.PostUrl,
					post.SiteName,
					post.SiteUrl,
				)
			} else {
				fmt.Printf(
					"Name: %s\nCreated at: %v\nPosted at: %v\nTag: %v\nTag text: %v\nStarts at: %v\nEnds at: permanent\nPost url: %v\nSite name: %v\nSite url: %v\n",
					post.Name,
					post.CreatedAt,
					post.PostedAt,
					post.Tag,
					post.TagText,
					post.StartsAt.Time,
					post.PostUrl,
					post.SiteName,
					post.SiteUrl,
				)
			}
			fmt.Println("--------------------------------------------------")
		}
		fmt.Println("--------------------------------------------------")

		return nil

	} else if len(cmd.Args) == 1 && cmd.Args[0] == "ongoing" {
		posts, err := s.PtrDB.GetPostsOnGoingAndSites(context.Background())
		if err != nil {
			return fmt.Errorf("error getting posts table: %w", err)
		}

		fmt.Println("--------------------------------------------------")
		fmt.Println("--------------------------------------------------")
		fmt.Printf("%d ongoing posts in the table\n", len(posts))
		fmt.Println("--------------------------------------------------")
		fmt.Println("--------------------------------------------------")

		if len(posts) == 0 {
			return nil
		}

		for _, post := range posts {
			if post.EndsAt.Valid {
				fmt.Printf(
					"Name: %s\nCreated at: %v\nPosted at: %v\nTag: %v\nTag text: %v\nStarts at: %v\nEnds at: %v\nPost url: %v\nSite name: %v\nSite url: %v\n",
					post.Name,
					post.CreatedAt,
					post.PostedAt,
					post.Tag,
					post.TagText,
					post.StartsAt.Time,
					post.EndsAt.Time,
					post.PostUrl,
					post.SiteName,
					post.SiteUrl,
				)
			} else {
				fmt.Printf(
					"Name: %s\nCreated at: %v\nPosted at: %v\nTag: %v\nTag text: %v\nStarts at: %v\nEnds at: permanent\nPost url: %v\nSite name: %v\nSite url: %v\n",
					post.Name,
					post.CreatedAt,
					post.PostedAt,
					post.Tag,
					post.TagText,
					post.StartsAt.Time,
					post.PostUrl,
					post.SiteName,
					post.SiteUrl,
				)
			}
			fmt.Println("--------------------------------------------------")
		}
		fmt.Println("--------------------------------------------------")

		return nil
	} else if len(cmd.Args) == 2 && cmd.Args[0] == "period" {
		now := time.Now()
		p, err := strconv.Atoi(cmd.Args[1])
		if err != nil {
			return fmt.Errorf("error calling strconv.Atoi: %w", err)
		}
		t := now.AddDate(0, 0, -p)
		posts, err := s.PtrDB.GetPostsWithinGivenPeriod(context.Background(), pgtype.Timestamptz{
			Time:  t,
			Valid: true,
		})
		if err != nil {
			return fmt.Errorf("error getting posts table: %w", err)
		}

		fmt.Println("--------------------------------------------------")
		fmt.Println("--------------------------------------------------")
		fmt.Printf("There are %d posts in the table with an end time within the last %d days\n", len(posts), p)
		fmt.Println("--------------------------------------------------")
		fmt.Println("--------------------------------------------------")

		if len(posts) == 0 {
			return nil
		}

		for _, post := range posts {
			fmt.Printf(
				"Name: %s\nCreated at: %v\nPosted at: %v\nTag: %v\nTag text: %v\nStarts at: %v\n",
				post.Name,
				post.CreatedAt,
				post.PostedAt,
				post.Tag,
				post.TagText,
				post.StartsAt.Time,
			)
			if post.EndsAt.Valid {
				fmt.Printf(
					"Ends at: %v\n",
					post.EndsAt.Time,
				)
			} else {
				fmt.Println("Ends at: permanent")
			}
			fmt.Printf(
				"Post url: %v\nSite name: %v\nSite url: %v\n",
				post.PostUrl,
				post.SiteName,
				post.SiteUrl,
			)
			fmt.Println("--------------------------------------------------")
		}
		fmt.Println("--------------------------------------------------")

		return nil
	} else {
		return errors.New("the posts handler behavior:\nIf no additional argument is given, print all posts\nIf ongoing is given, print ongoing posts\nIf period and number are given, print posts with an end time later than {the given number} days ago")
	}
}
