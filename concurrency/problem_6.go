/*
Problem 6: Parallel Square Calculator

Input: [1,2,3,4,5]
Each number processed in a goroutine
Output squares
Expected:
1 4 9 16 25

*/

package concurrency

func Squares(num int, ch chan int) {

	ch <- num * num

}
