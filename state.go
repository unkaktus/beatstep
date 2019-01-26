package beatstep

import (
	"time"
)

// ControlType designates type of BeatStep control.
type ControlType int

const (
	// UnrecognizedState reperesents unrecognized control.
	UnrecognizedState ControlType = iota
	// EncoderState represents encoder (knob) control.
	EncoderState
	// PadState represents pad control.
	PadState
)

// State represents a state of a control on the device.
type State struct {
	// Type is the type of control
	Type ControlType
	// Number, as labeled on the device
	Number int64
	// Value is the value of the control
	Value int64
	// Timestamp is portmidi's internal timestamp
	Timestamp time.Duration
}
