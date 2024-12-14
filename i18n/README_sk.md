[![Go Reference](https://pkg.go.dev/badge/github.com/temidaradev/esset.svg)](https://pkg.go.dev/github.com/temidaradev/esset)

# Esset

Esset je základný implementátor materiálov pre ebitengine.

# Použitie

Najskôr spustite `go get github.com/temidaradev/esset` a vytvorte priečinok pre materiály ak ju ešte nemáte. Po vytvorení priečinka umiestnite .png materiály do priečinka a vytvorte `assets.go`. Potom pridajte

```
//go:embed *
var assets embed.FS
```

tento embed príkaz hneď po import časti. Teraz už môžete používať esset ako váš implementátor materiálov. Vyžaduje to 2 parametre z vášho embed príkazu a potom váš `"asset.png"`.

## GetAsset

Tu je príklad: `var Idle = esset.GetAsset(assets, "path/to/your/asset.png")`

## GetMultipleAssets

Dôležité je vytvoriť priečinok a umiestniť každú jednotlivú dlaždicu (.png) takto

<img src="resources/image.png" height="400">

a potom môžete jednoducho použiť túto funkciu takto: `var Tile = esset.GetMultipleAssets(assets, "path/to/your/*.png")` Kedže vyberáte viac ako 1 obraz, náš `*ebiten.Image` je pole z ktorého môžete vyberať podľa indexu takto: `TileComponent := assets.Tile[0]`, alebo ak vám treba získať náhodný materiál z toho priečinka tak môžete takto: `TileRandom := assets.Tile[rand.Intn(len(assets.Tile))]`

## UseFont

Pre písma ich musíte vložiť oddelene takto:

```
//go:embed path/to/your/font.ttf
var MyFont []byte
```

Potom musíte vytvoriť `&text.DrawOptions{}` vo vašej `Draw()` funkcií takto:

```
opF := &text.DrawOptions{}
opF.GeoM.Translate(x, y)
opF.ColorScale.ScaleWithColor(color.White)
```

Potom môžete použiť `esset.UseFont` funkciu takto: `esset.UseFont(screen, assets.MyFont, "wassup", 24, opF)`

Veľmi vďaka pre [@m110](https://github.com/m110) za podporu so zdrojom <3
