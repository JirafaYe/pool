package main

import (
	"fmt"
	"pool/pool"
)

var nums [10000]int

func main() {
	p := pool.InitPool(100)
	p.Run()
	for i := 0; i < 10000; i++ {
		num := i
		go p.Submit(func() {
			fmt.Println("task", num)
			nums[num] = 1
		})
	}
	p.Wait()
	fmt.Println(nums[len(nums)-1])
	var flag = true
	for _, num := range nums {
		if num == 0 {
			flag = false
			break
		}
	}
	fmt.Println(flag)
}
