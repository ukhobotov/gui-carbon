package manual

import (
    "image"
    "image/draw"
    
    "github.com/ukhobotov/gui"
)

func NewMux(draw chan<- func(d draw.Image) image.Rectangle) *Mux {
    eventCh := make(chan gui.Event)
    envCh := make(chan *Env)
    
    go mux(eventCh, envCh)
    
    return &Mux{
        childEvs: eventCh,
        envCh:    envCh,
        drawCh:   draw,
    }
}

func mux(eventCh <-chan gui.Event, envCh <-chan *Env) {
    var envs []*Env
    for {
        select {
        case event, ok := <-eventCh:
            if !ok {
                for _, env := range envs {
                    close(env.events)
                }
                return
            }
            for _, env := range envs {
                env.events <- event
            }
        case env := <-envCh:
            envs = append(envs, env)
        }
    }
}

type Mux struct {
    childEvs  chan<- gui.Event
    parentEvs <-chan gui.Event
    envCh     chan<- *Env
    drawCh    chan<- func(d draw.Image) image.Rectangle
}

func (mux *Mux) MakeEnv() gui.Env {
    env := NewEnv(mux.drawCh)
    mux.envCh <- env
    return env
}

func (mux *Mux) Entry() chan<- gui.Event {
    return mux.childEvs
}
