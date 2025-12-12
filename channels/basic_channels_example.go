package channels

import (
	"fmt"
	"sync"
)

// This example contains implementation of print odd and even numbers in a proper sequecen but using two go routines
func InitBasicChannelsExample() {
	var wg sync.WaitGroup
	wg.Add(2)
	oddChannel := make(chan struct{})
	evenChannel := make(chan struct{})

	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			<-oddChannel              // wait for channel before go routine should start printing it's data
			fmt.Println(i)            // print the data
			evenChannel <- struct{}{} // tell the even channel to print it's data now and wait for it finish, since this is blocking in nature
		}
	}()

	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			<-evenChannel
			fmt.Println(i)
			if i < 10 { // without this condition, evenChannel go routine will write to odd channel which would have been already closed by then causing deadlock
				oddChannel <- struct{}{}
			}
		}
	}()

	oddChannel <- struct{}{} // start the odd go routine
	wg.Wait()
	close(oddChannel)
	close(evenChannel)
}
