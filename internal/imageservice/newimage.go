package imageservice

import (
	"golang.org/x/image/bmp"
	"image"
	"image/color"
	"internshipApplicationTemplate/internal/makeuuid"
	"log"
	"os"
	"strconv"
)

func NewImage(width string, height string) (string, error) {
	uuid := makeuuid.GenerateUUID(width, height)

	w, err := strconv.Atoi(width)
	if err != nil {
		log.Println("in NewImage: convert string to int error")
		return "", err
	}

	h, err := strconv.Atoi(height)
	if err != nil {
		log.Println("in NewImage: convert string to int error")
	}

	img := image.NewRGBA(image.Rect(0, 0, w, h))

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			img.Set(x, y, color.White)
		}
	}

	fileName := uuid + ".bmp"

	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println("create file error")
		return "", err
	}

	defer func() {
		if cerr := f.Close(); err != nil {
			log.Println("file close error", err)
			err = cerr
		}
	}()

	if err := bmp.Encode(f, img); err != nil {
		f.Close()
		log.Println("bmp encoding error")
		return "", err
	}

	return uuid, nil
}
