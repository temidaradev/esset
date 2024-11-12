package assets

import (
	"embed"

	"github.com/temidaradev/esset"
)

//go:embed *
var assets embed.FS

var Idle = esset.GetSingleImage(assets, "mainchar.png")
var Left = esset.GetSingleImage(assets, "left.png")
var Right = esset.GetSingleImage(assets, "right.png")
