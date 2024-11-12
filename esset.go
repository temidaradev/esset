package esset

import (
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func GetSingleImage(name string) *ebiten.Image {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(img)
}
