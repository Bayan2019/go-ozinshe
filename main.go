package main

import (
	"os"

	"github.com/Bayan2019/go-ozinshe/configuration"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	dirImages := os.Getenv("DIR_IMAGES")
	if dirImages == "" {
		dirImages = "/images"
	}
	dirVideos := os.Getenv("DIR_VIDEOS")
	if dirVideos == "" {
		dirVideos = "/videos"
	}

	configuration.ApiCfg = &configuration.ApiConfiguration{
		DirImages: dirImages,
		DirVideos: dirVideos,
	}
}
