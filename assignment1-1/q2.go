package cos418_hw1_1

import (
	"bufio"
	"io"
	"strconv"
	"os"
	"log"
)

// Sum numbers from channel `nums` and output sum to `out`.
// You should only output to `out` once.
// Do NOT modify function signature.
func sumWorker(nums chan int, out chan int) {
	// TODO: implement me
	// HINT: use for loop over `nums`
	sum := 0
	// for v := range nums {
	// 	sum += v
	// }
	for i  := 0; i < len(nums); i++ {
		sum += <- nums
	}
	//close(nums)
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

	reader, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	arr, err := readInts(reader)
	if err != nil {
		log.Fatal(err)
	}
	length := len(arr)
	seg := length / num
	nums := make(chan int, seg)
	out := make(chan int)

	for i := 0; i < num; i++ {
		for  j := range arr[i*seg:(i+1)*seg] {
			nums <- j
		}
		
		go sumWorker(nums, out)
	}
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
