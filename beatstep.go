package beatstep

import (
	"errors"

	"github.com/rakyll/portmidi"
)

// BeatStep represents an instance of BeatStep device.
type BeatStep struct {
	input  *portmidi.Stream
	output *portmidi.Stream
}

const deviceName = "Arturia BeatStep MIDI 1"

func discover() (input, output portmidi.DeviceID, err error) {
	inputDevice, outputDevice := -1, -1
	for device := 0; device < portmidi.CountDevices(); device++ {
		info := portmidi.Info(portmidi.DeviceID(device))
		if info.Name == deviceName {
			if info.IsInputAvailable {
				inputDevice = device
			}
			if info.IsOutputAvailable {
				outputDevice = device
			}
		}
	}
	if inputDevice == -1 || outputDevice == -1 {
		err = errors.New("beatstep: no beatstep connected")
		return
	}
	return portmidi.DeviceID(inputDevice), portmidi.DeviceID(outputDevice), nil
}

// Open initializes new BeatStep instance.
func Open() (*BeatStep, error) {
	inputDevice, outputDevice, err := discover()
	if err != nil {
		return nil, err
	}
	inputStream, err := portmidi.NewInputStream(inputDevice, 1024)
	if err != nil {
		return nil, err
	}
	outputStream, err := portmidi.NewOutputStream(outputDevice, 1024, 0)
	if err != nil {
		inputStream.Close()
		return nil, err
	}
	bs := &BeatStep{
		input:  inputStream,
		output: outputStream,
	}
	return bs, nil
}

func (bs *BeatStep) read() (states []State, err error) {
	events, err := bs.input.Read(64)
	if err != nil {
		return nil, err
	}
	for _, event := range events {
		state := decode(event)
		states = append(states, state)
	}
	return states, nil
}

// Listen reads all the States from the device and writes
// them in to the channel it returns.
// Listen returns immidiately.
func (bs *BeatStep) Listen() <-chan State {
	ch := make(chan State)
	go func(ch chan<- State) {
		for {
			states, err := bs.read()
			if err != nil {
				close(ch)
				return
			}
			for _, state := range states {
				ch <- state
			}
		}
	}(ch)
	return ch
}

// Close closes internal resourses of BeatStep.
func (bs *BeatStep) Close() error {
	bs.input.Close()
	bs.output.Close()
	return nil
}
