package carbon

type State byte

const (
    Default  State = 0
    Hover    State = 0b1
    Active   State = 0b10
    Focused  State = 0b100
    Disabled State = 0b1000
)
