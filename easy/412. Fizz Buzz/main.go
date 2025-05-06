package main

import (
	"fmt"
	"strconv"
)

/*
 Given an integer n, return a string array answer (1-indexed) where:

answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
answer[i] == "Fizz" if i is divisible by 3.
answer[i] == "Buzz" if i is divisible by 5.
answer[i] == i (as a string) if none of the above conditions are true.


Example 1:

Input: n = 3
Output: ["1","2","Fizz"]
Example 2:

Input: n = 5
Output: ["1","2","Fizz","4","Buzz"]
Example 3:

Input: n = 15
Output: ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]


Constraints:*/

func main() {
	n := 3
	fmt.Println("Output: [\"1\",\"2\",\"Fizz\"]")
	fmt.Println(fizzBuzz(n))
	n = 5
	fmt.Println("Output: [\"1\",\"2\",\"Fizz\",\"4\",\"Buzz\"]")
	fmt.Println(fizzBuzz(n))
	n = 15
	fmt.Println("Output: [\"1\",\"2\",\"Fizz\",\"4\",\"Buzz\",\"Fizz\"," +
		"\"7\",\"8\",\"Fizz\",\"Buzz\",\"11\",\"Fizz\",\"13\",\"14\",\"FizzBuzz\"]")
	fmt.Println(fizzBuzz(n))
}

func fizzBuzz(n int) []string {
	if n == 0 {
		return nil
	}
	str := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		s := strconv.Itoa(i)
		if i%3 == 0 && i%5 == 0 {
			s = "FizzBuzz"
		} else if i%3 == 0 {
			s = "Fizz"
		} else if i%5 == 0 {
			s = "Buzz"
		}
		str = append(str, s)
	}
	return str
}
