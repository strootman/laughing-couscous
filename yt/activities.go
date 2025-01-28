package yt

import (
	"context"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func Activities() *youtube.ActivityListResponse {
	ctx := context.Background()

	service, err1 := youtube.NewService(ctx, option.WithAPIKey("AIzaSyCV2JStskbQz-mw1AosCEFk5kcjEHlOdiI"))
	if err1 != nil {
		log.Fatalf("Unable to create Youtube service %v", err1)
	}
	parts := []string{"snippet"}
	call := service.Activities.List(parts)
	call.Mine(true)
	response, err2 := call.Do()

	if err2 != nil {
		log.Fatalf("Unable to fetch activities %v", err2)
	}
	return response
}
