package main

import (
	"fmt"
	"image"
	_ "image/gif"

	/*
		func init() {
			image.RegisterFormat("gif", "GIF8?a", Decode, DecodeConfig)
		}
	*/
	_ "image/jpeg"
	/*
		func init() {
			image.RegisterFormat("jpeg", "\xff\xd8", Decode, DecodeConfig)
		}
	*/
	_ "image/png"
	/*
		func init() {
			image.RegisterFormat("png", pngHeader, Decode, DecodeConfig)
		}
	*/
	"os"
)

func main() {
	width, height, err := imageSize(os.Args[1])
	if err != nil {
		fmt.Println("get image size error:", err)
		return
	}

	fmt.Printf("image size: [%d, %d]\n", width, height)
}

func imageSize(imageFile string) (int, int, error) {
	f, _ := os.Open(imageFile)
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return 0, 0, err
	}

	b := img.Bounds()

	return b.Max.X, b.Max.Y, nil
}
