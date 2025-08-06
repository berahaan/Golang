package services

import "GOLANG/entity"

type VideoService interface {
	Greet() string
	SaveVideo(video entity.Video)  // this just saaves the video to local storage woth specified paths
	FindAllVideos() []entity.Video // this is return slices of videos for storage or from databse
}
type VideoServices struct {
	Videos []entity.Video
}
