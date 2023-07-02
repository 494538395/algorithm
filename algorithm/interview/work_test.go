package interview

import (
	"fmt"
	"sync"
	"testing"
)

// goroutine 交替打印
// 一个打印奇数、一个打印偶数
// 从小到大打印
func TestWork1(t *testing.T) {
	info := make(chan struct{}) // 无缓冲 channel 用于控制顺序

	printOddNumber := func() {
		for i := 1; i <= 100; i++ {
			info <- struct{}{}
			if i%2 != 0 {
				fmt.Println("printOddNumber  -->", i)
			}
		}
	}

	printEvenNumber := func() {
		for i := 1; i <= 100; i++ {
			<-info
			if i%2 == 0 {

				fmt.Println("printEvenNumber  -->", i)
			}
		}
	}
	go printEvenNumber()
	go printOddNumber()

	for {
	}
}

// 交替打印第二种解法
// 偶数和奇数处理不一样，
// 如果是偶数，是奇数 goroutine 最后一个负责传递
// 如果是奇数，是偶数 goroutine 最后一个负责传递
func TestWork2(t *testing.T) {
	evenNumber := make(chan int, 100) // 偶数
	oddNumber := make(chan int, 100)  // 奇数

	wg := sync.WaitGroup{}
	wg.Add(2)

	oddNumber <- 1
	go func() {
		defer wg.Done()
		// 1  3  5
		// ->2  4  6
		for num := range oddNumber {
			fmt.Println("oddNumber -->", num)
			num++
			evenNumber <- num

			//if num == 100 {
			//	close(oddNumber)
			//	break
			//}
		}

	}()

	go func() {
		defer wg.Done()
		//2 4
		// ->3  5
		for num := range evenNumber {
			//if num == 100 {
			//	fmt.Println("evenNumber -->", num)
			//	close(evenNumber)
			//	break
			//}
			fmt.Println("evenNumber -->", num)
			num++
			oddNumber <- num

		}
	}()

	wg.Wait()
}

func TestWrok3(t *testing.T) {

	var wg sync.WaitGroup
	evenNumber := make(chan int, 100) // 偶数
	oddNumber := make(chan int, 100)  // 奇数

	wg.Add(1)
	oddNumber <- 1
	go func() {
		defer wg.Done()
		for num := range oddNumber {
			if num == 100 {
				close(oddNumber)
				return
			}
			fmt.Println("oddNumber -->", num)
			num++
			evenNumber <- num
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range evenNumber {
			if num == 101 { // 这里需要使用 101 而不是 100
				close(evenNumber)
				return
			}
			fmt.Println("evenNumber -->", num)
			num++
			oddNumber <- num
		}
	}()

	wg.Wait()
}
