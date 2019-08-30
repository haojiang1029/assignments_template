package cos418_hw1_1

import (
	"bufio"
	"io"
	"strconv"
	"os"
	"log"
	"fmt"
)

//var wg sync.WaitGroup
// Sum numbers from channel `nums` and output sum to `out`.
// You should only output to `out` once.
// Do NOT modify function signature.
func sumWorker(nums chan int, out chan int) {
	// TODO: implement me
	// HINT: use for loop over `nums`
	//defer wg.Done()
	fmt.Println("bbb")
	sum := 0
	for v := range nums {
		sum += v
		fmt.Println("sum: ", sum)
	}
	
	out <- sum 
}

// Read integers from the file `fileName` and return sum of all values.
// This function must launch `num` go routines running
// `sumWorker` to find the sum of the values concurrently.
// You should use `checkError` to handle potential errors.
// Do NOT modify function signature.
func sum(num int, fileName string) int {
	// TODO: implement me
	// HINT: use `readInts` and `sumWorkers`
	// HINT: used buffered channels for splitting numbers between workers
	//wg.Add(num)
	
	reader, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	arr, err := readInts(reader)
	if err != nil {
		log.Fatal(err)
	}
	
	nums := make(chan int, num)
	out := make(chan int)
	fmt.Println("num: , len: ", num, len(arr))

	for i := 0; i < num; i++ {
		fmt.Println("i: ", i)
		go sumWorker(nums, out)

	}
	
	for v := range arr {
		fmt.Println("v: ", v)
		nums <- v

	}

	fmt.Println("aaa")
	close(nums)
	//wg.Wait()
	
	
	var sum int
	for i:= 0; i < num; i++ {
		sum += <-out
	}
	return sum
	
	
}

// Read a list of integers separated by whitespace from `r`.
// Return the integers successfully read with no error, or
// an empty slice of integers and the error that occurred.
// Do NOT modify this function.
func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var elems []int
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return elems, err
		}
		elems = append(elems, val)
	}
	return elems, nil
}
