package carbon

import (
    "image"
    
    "github.com/ukhobotov/gui"
    "github.com/ukhobotov/gui/win"
)

type Window struct {
    env    gui.Env
    mux    *gui.Mux
    Bounds image.Rectangle
}

func NewWindow(title string, resizable bool, size image.Point) *Window {
    opts := []win.Option{win.Title(title)}
    if resizable {
        opts = append(opts, win.Resizable())
    }
    if size != (image.Point{}) {
        opts = append(opts, win.Size(size.X, size.Y))
    } else {
        opts = append(opts, win.Maximized())
    }
    w, err := win.New(opts...)
    if err != nil {
        panic(err)
    }
    mux, env := gui.NewMux(w)
    return &Window{env: env, mux: mux}
}

func (w *Window) Handle(f func(event gui.Event)) {
    defer close(w.env.Draw())
    for event := range w.env.Events() {
        switch e := event.(type) {
        case win.WiClose:
            return
        case gui.Resize:
            w.Bounds = e.Bounds()
        }
        if f != nil {
            f(event)
        }
    }
}

func (w *Window) CenterRect(size image.Point) image.Rectangle {
    return image.Rectangle{
        Min: w.Bounds.Min.Add(w.Bounds.Max).Sub(size).Div(2),
        Max: w.Bounds.Min.Add(w.Bounds.Max).Add(size).Div(2),
    }
}

func (w *Window) MakeEnv() gui.Env {
    return w.mux.MakeEnv()
}
