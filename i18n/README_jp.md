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

<img src="../resources/image.png" height="400">

そして、その関数は次のように簡単に使用できます: `var Tile = esset.GetMultipleAssets(assets, "path/to/your/*.png")` 1 つ以上の画像を選択しているため、`*ebiten.Image` は次のようにインデックスで選択できるスライスです: `TileComponent := asset.Tile[0]` または、そのフォルダーからランダムなアセットを取得する必要がある場合は、次のように実行できます: `TileRandom := asset.Tile[rand.Intn(len(assets.Tile))]`

## DrawText

フォントの場合は、フォントを個別に埋め込む必要があり、次のように `text.Face` 変数を追加する必要があります:

```
//go:embed path/to/your/font.ttf
var MyFont []byte
var FontFaceS text.Face
var FontFaceM text.Face
var FontFaceM text.Face
```

次に、フォントを毎回ロードしないようにするには、次のように、メイン ファイル (`Game{}` を含む) にある `init()` 関数に `GetFont()` 関数を配置する必要があります:

```
func init() {
    assets.FontFaceS, _ = esset.GetFont(assets.MyFont, 48)
}
```

`GetFont()` 関数を設定したら、`DrawText()` 関数を使用する準備が整います

このために特別な DrawOptions を作成する必要はありません。 (screen、text、fontSize、posX、posY、text.Face、color) だけです

`esset.DrawText` 関数は次のように使用できます: `esset.DrawText(screen、"こんにちわ", 16, 100, 50, assets.FontFaceS, color.White)`

ソースサポートをしてくれた [@m110](https://github.com/m110) に感謝します <3
