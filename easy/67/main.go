package main

import "fmt"

/*Given two binary strings a and b, return their sum as a binary string.



Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"

Constraints:

1 <= a.length, b.length <= 104
a and b consist only of '0' or '1' characters.
Each string does not contain leading zeros except for the zero itself.*/

func addBinary(a string, b string) string {
	c := []byte(a)

	if len(a) < len(b) {
		c = []byte(b)
	}
	curr := 0
	i, j := len(a)-1, len(b)-1
	k := max(i, j)
	for i >= 0 && j >= 0 || curr > 0 {
		tmp := 0

		if i < 0 {
			tmp = int(b[j]-48) + curr
		} else if j < 0 {
			tmp = int(a[i]-48) + curr
		} else {
			tmp = int(a[i]-48) + int(b[j]-48) + curr
		}
		if tmp > 1 {
			curr = 1
		} else {
			curr = 0
		}
		c[k] = byte(tmp%2) + 48
		i--
		j--
		k--
		fmt.Println(i, j)
		if i < 0 && j < 0 && curr == 1 {
			c = append([]byte{49}, c...)
			curr = 0
			break
		}

	}
	return string(c)
}

func main() {
	a := "11"
	b := "1"
	fmt.Println("Output: \"100\"")
	fmt.Println(addBinary(a, b))
	a = "1010"
	b = "1011"
	fmt.Println("Output: \"10101\"")
	fmt.Println(addBinary(a, b))
}
