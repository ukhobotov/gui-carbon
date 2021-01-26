package carbon

import (
    "image"
    "image/color"
    "image/draw"
)

type Drawer = func(d draw.Image) image.Rectangle

func DrawRect(dst draw.Image, r image.Rectangle, clr color.Color) {
    DrawImage(dst, r, image.NewUniform(clr))
}

func DrawHLine(dst draw.Image, x1, x2, y int, clr color.Color) {
    DrawRect(dst, image.Rect(x1, y, x2, y+1), clr)
}

func DrawVLine(dst draw.Image, x, y1, y2 int, clr color.Color) {
    DrawRect(dst, image.Rect(x, y1, x+1, y2), clr)
}

type Border struct {
    Color color.Color
    Width int
}

func DrawBorder(dst draw.Image, r image.Rectangle, border Border) {
    if border.Width == 0 {
        border.Width = 1
    }
    x1, y1, x2, y2 := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y
    DrawRect(dst, image.Rect(x1, y1, x1+border.Width, y2), border.Color)
    DrawRect(dst, image.Rect(x1, y2-border.Width, x2, y2), border.Color)
    DrawRect(dst, image.Rect(x2-border.Width, y1, x2, y2), border.Color)
    DrawRect(dst, image.Rect(x1, y1, x2, y1+border.Width), border.Color)
}

func DrawImage(dst draw.Image, r image.Rectangle, img image.Image) {
    draw.Draw(dst, r, img, image.Point{}, draw.Over)
}
