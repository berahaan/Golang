package services

import "GOLANG/entity"

type VideoService interface {
	SaveVideo(video entity.Video)  // this just saaves the video to local storage woth specified paths
	FindAllVideos() []entity.Video // this is return slices of videos
}
