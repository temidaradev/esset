package assets

import (
	"embed"

	"github.com/temidaradev/esset"
)

//go:embed *
var assets embed.FS

//go:embed font/OpenSans-Medium.ttf
var MyFont []byte

var Idle = esset.GetAsset(assets, "mainchar.png")
var Left = esset.GetAsset(assets, "left.png")
var Right = esset.GetAsset(assets, "right.png")

var GopherTile = esset.GetMultiAssets(assets, "gopherTile/*.png")
