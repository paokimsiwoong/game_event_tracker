package commands

import (
	"context"
	"fmt"
)

// DB에 저장된 site 목록 보여주는 핸들러
func HandlerSites(s *State, cmd Command) error {
	sites, err := s.PtrDB.GetSites(context.Background())
	if err != nil {
		return fmt.Errorf("error getting sites table : %w", err)
	}

	fmt.Println("--------------------------------------------------")
	fmt.Println("--------------------------------------------------")
	fmt.Printf("%d sites in the table\n", len(sites))
	fmt.Println("--------------------------------------------------")
	fmt.Println("--------------------------------------------------")

	for _, site := range sites {
		if site.LastFetchedAt.Valid {
			fmt.Printf(
				"Name: %s\nUrl: %s\nCreated at: %v\nLast fetched at: %v\n",
				site.Name,
				site.Url,
				site.CreatedAt,
				site.LastFetchedAt.Time,
			)
			fmt.Println("--------------------------------------------------")
			continue
		}
		fmt.Printf(
			"Name: %s\nUrl: %s\nCreated at: %v\nLast fetched at: not fetched\n",
			site.Name,
			site.Url,
			site.CreatedAt,
		)
		fmt.Println("--------------------------------------------------")
	}
	fmt.Println("--------------------------------------------------")

	return nil
}
