package channels

func Tee(done chan struct{}, inChannel <-chan any) (<-chan any, <-chan any) {
	outChan1 := make(chan any)
	outChan2 := make(chan any)

	go func() {
		defer close(outChan1)
		defer close(outChan2)

		for val := range inChannel {
			o1, o2 := outChan1, outChan2
			for range 2 {
				select {
				case <-done:
					return
				case o1 <- val:
					o1 = nil
				case o2 <- val:
					o2 = nil
				}
			}
		}
	}()

	return outChan1, outChan2
}
