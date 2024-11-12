[![Go Reference](https://pkg.go.dev/badge/github.com/temidaradev/esset.svg)](https://pkg.go.dev/github.com/temidaradev/esset)
# Esset

Esset is a basic asset implementer for ebitengine.

## Usage

First `go get github.com/temidaradev/esset` and create an assets folder if you don't have. After creating folder put .png assets into that folder and create `assets.go`. After this add

```
//go:embed *
var assets embed.FS
```

this embed statement after import part. Now you can use esset as you asset implementer, It wants 2 parameters from you firstly embed statement and then your `"asset.png"`. Here is an example: `var Idle = esset.GetSingleAsset(assets, "path/to/your/asset.png")`
