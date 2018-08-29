package beatstep

import (
	"time"

	"github.com/rakyll/portmidi"
)

var encoderMap map[int64]int64 = map[int64]int64{
	7:   0,
	10:  1,
	74:  2,
	71:  3,
	76:  4,
	77:  5,
	93:  6,
	73:  7,
	75:  8,
	114: 9,
	18:  10,
	19:  11,
	16:  12,
	17:  13,
	91:  14,
	79:  15,
	72:  16,
}

var padMap map[int64]int64 = map[int64]int64{
	44: 1,
	45: 2,
	46: 3,
	47: 4,
	48: 5,
	49: 6,
	50: 7,
	51: 8,
	36: 9,
	37: 10,
	38: 11,
	39: 12,
	40: 13,
	41: 14,
	42: 15,
	43: 16,
}

func decode(e portmidi.Event) State {
	state := State{
		Type:      UnrecognizedState,
		Number:    int64(e.Data1),
		Value:     int64(e.Data2),
		Timestamp: time.Duration(int64(e.Timestamp)) * time.Millisecond,
	}
	if number, ok := encoderMap[e.Data1]; ok {
		state.Type = EncoderState
		state.Number = number
	}
	if number, ok := padMap[e.Data1]; ok {
		state.Type = PadState
		state.Number = number
	}
	return state
}
