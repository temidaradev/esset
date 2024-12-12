[![Go Reference](https://pkg.go.dev/badge/github.com/temidaradev/esset.svg)](https://pkg.go.dev/github.com/temidaradev/esset)

# Esset

Esset は、ebitengine の基本的なアセット実装者です。

# 使用法

まず `github.com/temidaradev/esset`にアクセスし、アセットフォルダを作成します（まだ作成していない場合）。フォルダを作成したら`.png`アセットをそのフォルダに入れて`assets.go`を作成します。その後、以下を追加します。

```
//go:embed *
var assets embed.FS
```

この埋め込みステートメントは、インポート部分の後にあります。これで、esset をアセット実装者として使用できます。最初に埋め込みステートメント、次に「asset.png」という 2 つのパラメータが必要です。

## GetAsset

例: `var Idle = esset.GetAsset(assets, "path/to/your/asset.png")`

## GetMultipleAssets

重要なのは、フォルダを作成し、すべてのタイルアイテム「.png」を次のように配置することです。

<img src="resources/image.png" height="400">

そして、その関数は次のように簡単に使用できます: `var Tile = esset.GetMultipleAssets(assets, "path/to/your/*.png")` 1 つ以上の画像を選択しているため、`*ebiten.Image` は次のようにインデックスで選択できるスライスです: `TileComponent := asset.Tile[0]` または、そのフォルダーからランダムなアセットを取得する必要がある場合は、次のように実行できます: `TileRandom := asset.Tile[rand.Intn(len(assets.Tile))]`

## UseFont

フォントの場合は、次のようにフォントを個別に埋め込む必要があります。

```
//go:embed path/to/your/font.ttf
var MyFont []byte
```

その後、次のように `Draw()` 関数に `&text.DrawOptions{}` を作成する必要があります。

```
opF := &text.DrawOptions{}
opF.GeoM.Translate(x, y)
opF.ColorScale.ScaleWithColor(color.White)
```

その後、`esset.UseFont` 関数を次のように使用できます: `esset.UseFont(screen, asset.MyFont, "wassup", 24, opF)`

ソースサポートをしてくれた [@m110](https://github.com/m110) に感謝します <3
