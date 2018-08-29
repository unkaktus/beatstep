package beatstep

import (
	"errors"
)

// ToggleLight toggles light of pad to state on.
// If pad is 0, ToggleLight toggles light of all the pads.
func (bs *BeatStep) ToggleLight(pad int, on bool) error {
	if pad > 16 {
		return errors.New("beatstep: no such pad")
	}
	// If pad is 0 toggle all the lights.
	if pad == 0 {
		for p := 1; p <= 16; p++ {
			if err := bs.ToggleLight(p, on); err != nil {
				return err
			}
		}
		return nil
	}
	var note int64
	for k, v := range padMap {
		if v == int64(pad) {
			note = k
		}
	}
	velocity := int64(0)
	if on {
		velocity = int64(127)
	}
	return bs.output.WriteShort(0x90, note, velocity)
}
