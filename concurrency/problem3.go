/*
Problem 4: Sum Using Channel

Create a function that takes a slice of ints
Calculate sum in a goroutine
Send result via channel

Example:
Input: [1,2,3,4]
Output: 10
*/

package concurrency

func ChannelSum(nums []int, ch chan int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	ch <- sum
}

/*
func main() {
	ch := make(chan int)
	nums := []int{1, 2, 3, 4}
	go concurrency.ChannelSum(nums, ch)

	msg := <-ch
	fmt.Println("Result: ", msg)
}
*/
