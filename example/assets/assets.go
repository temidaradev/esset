package assets

import (
	"embed"

	"github.com/temidaradev/esset"
)

//go:embed *
var assets embed.FS

var Idle = esset.GetSingleImage("mainchar.png")
