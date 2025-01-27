package yt

import (
	"context"
	"log"

	"google.golang.org/api/youtube/v3"
)

func LikedVideos() *youtube.VideoListResponse {
	ctx := context.Background()

	service, err1 := youtube.NewService(ctx)
	if err1 != nil {
		log.Fatalf("Unable to create YouTube service %v", err1)
	}

	parts := []string{"snippet", "contentDetails", "statistics"}
	call := service.Videos.List(parts)
	call.MyRating("like")
	response, err2 := call.Do()

	if err2 != nil {
		log.Fatalf("Unable to fetch like videos %v", err2)
	}
	return response
}
