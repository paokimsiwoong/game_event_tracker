package commands

import (
	"context"
	"errors"
	"fmt"
)

// DB에 저장된 post 목록 보여주는 핸들러
func HandlerPosts(s *State, cmd Command) error {
	// 추가 명령어 없으면 db에 있는 post 전부 출력
	if len(cmd.Args) == 1 && cmd.Args[0] == "all" {
		posts, err := s.PtrDB.GetPostsAndSites(context.Background())
		if err != nil {
			return fmt.Errorf("error getting posts table : %w", err)
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
			return fmt.Errorf("error getting posts table : %w", err)
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
	} else {
		return errors.New("the posts handler expects one argument. If all is given, print all posts, if ongoing is given, print ongoing posts)")
	}
}
