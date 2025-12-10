package main

// This file contains implementation of or channels and how they work
// When you have to listen to multiple channels, or channel pattern combines them together and return one channel which you can listen to
// When ANY channel receives data, it closes all other channels

func Or(channels ...<-chan any) <-chan any {
	// base condition
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[1]
	}

	orDone := make(chan any)

	go func() {
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-Or(append(channels[3:], orDone)...):
			}
		}
	}()

	return orDone
}
