# ebiten-tutorial

## 1. 空っぽの画面を描く

* screen_blank/blank.go

## 2. スクリーンを色で塗りつぶす

* screen_fill/screen_fill.go

## 3. 画像をレンダリングする

* render_image/render_image.go

## 4. 位置を指定して画像をレンダリングする

* set_image_position/set_position.go

## 5. フィルタを設定する

* ebiten.Imageの拡大縮小や回転時の保管方法を指定する
* フィルターの方式は2種
  * FilterNearest：デフォルトのフィルター
  * FilterLinear：少しぼやけたような感じになる

## 6. 画像を回転する

* ebiten.DrawImageOptionで回転を指定する

## 画像の中心を原点となるように回転する

* image_rotate_center.go
* 画像の中心座標が原点となるように座標を移動させる。そして画像を回転させて、最後に座標を元に戻すと画像の中心を軸として回転しているようにみえる

### 6.1 単純な回転

### 6.2 回転中心

## 7． オフスクリーン