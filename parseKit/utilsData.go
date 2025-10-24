package parsekit

type Room struct {
	Name    string
	X, Y    int
	IsStart bool
	IsEnd   bool
	Link    []*Room
}

var (
	AntNum    int
	Err       error
	StartRoom string
	EndRoom   string
	Rooms     = make(map[string]*Room)
)
