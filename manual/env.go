package manual

import (
    "image"
    "image/draw"
    
    "github.com/ukhobotov/gui"
)

type Env struct {
    events chan gui.Event
    draw   chan<- func(d draw.Image) image.Rectangle
    Hidden bool
}

var EventsCap = 4

func NewEnv(draw chan<- func(d draw.Image) image.Rectangle) *Env {
    return &Env{
        events: make(chan gui.Event, EventsCap),
        draw:   draw,
    }
}

func (e *Env) Events() <-chan gui.Event {
    return e.events
}

func (e *Env) Draw() chan<- func(draw.Image) image.Rectangle {
    return e.draw
}

func (e *Env) Entry() chan<- gui.Event {
    return e.events
}

func SendVisible(event gui.Event, envs ...*Env) {
    for _, env := range envs {
        if env != nil && !env.Hidden {
            env.events <- event
        }
    }
}
