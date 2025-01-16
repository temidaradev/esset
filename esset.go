package esset

import (
	"bytes"
	"embed"
	"image"
	"image/color"
	_ "image/png"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// Assets
func GetAsset(efs embed.FS, path string) *ebiten.Image {
	file, err := efs.Open(path)
	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func GetMultiAssets(efs embed.FS, path string) []*ebiten.Image {
	matches, err := fs.Glob(efs, path)
	if err != nil {
		panic(err)
	}

	images := make([]*ebiten.Image, len(matches))
	for i, match := range matches {
		images[i] = GetAsset(efs, match)
	}

	return images
}

// Fonts
func GetFont(data []byte, fontSize int) (text.Face, error) {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	return &text.GoTextFace{
		Source: s,
		Size:   float64(fontSize),
	}, nil
}

func DrawText(screen *ebiten.Image, str string, fontSize int, posX, posY float64, textFace text.Face, color color.Color) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(posX, posY)
	op.ColorScale.ScaleWithColor(color)
	op.LayoutOptions.LineSpacing = float64(fontSize)

	text.Draw(screen, str, textFace, op)
}
