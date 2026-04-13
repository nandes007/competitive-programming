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

// Concurrent bug
// func multiply(n int) int {
//     result := 0
//     go func() {
//         if n%2 == 0 {
//             result = n * 2
//         } else {
//             result = n * 3
//         }
//     }()
//     return result
// }

func multiply(n int) int {
	result := 0
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		if n%2 == 0 {
			result = n * 2
		} else {
			result = n * 3
		}
	}()

	wg.Wait()
	return result
}

func largestNumber(a, b []int) int {
	mergedArr := append(a, b...)
	largestNumber := 0
	for i := 0; i < len(mergedArr); i++ {
		if mergedArr[i] > largestNumber {
			largestNumber = mergedArr[i]
		}
	}

	return largestNumber
}
