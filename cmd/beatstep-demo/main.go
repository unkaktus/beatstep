package main

import (
	"log"
	"time"

	"github.com/nogoegst/beatstep"
)

func main() {
	bs, err := beatstep.Open()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for i := 0; i < 5; i++ {
			bs.ToggleLight(0, true)
			time.Sleep(300 * time.Millisecond)
			bs.ToggleLight(0, false)
			time.Sleep(300 * time.Millisecond)
		}

		snakePattern := []int{1, 9, 10, 2, 3, 11, 12, 4, 5, 13, 14, 6, 7, 15, 16, 8}
		for {
			for _, pad := range snakePattern {
				bs.ToggleLight(pad, true)
				time.Sleep(100 * time.Millisecond)
				bs.ToggleLight(pad, false)

			}
		}

	}()
	ch := bs.Listen()
	defer bs.Close()
	for state := range ch {
		log.Printf("state: %+v", state)
	}
}
