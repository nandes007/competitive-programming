package src

import (
	"fmt"
	"sync"
)

// Soal race codition sebagai berikut. Dan solusinya ada pada RaceCondtion() function.
// var counter = 0
// func main() {
// 	for i := 0; i < 100; i++ {
// 		go func() {
// 			counter++
//                                    fmt.Println(counter)
// 		}()
// 	}
// 	time.Sleep(time.Second)

// }

func RaceCondition() {
	var (
		counter = 0
		lock    sync.Mutex
		wg      sync.WaitGroup
	)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			lock.Lock()
			counter++
			fmt.Println(counter)
			lock.Unlock()
		}()
	}

	wg.Wait()
}
