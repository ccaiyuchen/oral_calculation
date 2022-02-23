package main

import (
"fmt"
"math/rand"
"time"
)

func main() {

	for i := 0; i < 5; i++ {
		Create(5)
	}
	println()
}

func Create(maxNum int32) {
	symbol := make(map[int32]string, 0)
	symbol[0] = "+"
	symbol[1] = "-"
	symbol[2] = "+"
	symbol[3] = "-"
	fmt.Printf("打卡时间：%v \t用时：%v\t正确率：%v\n", "______", "______", "______")

	for i := 0; i < 10; i++ {
		if i != 0 && i%2 == 0 {
			fmt.Println()
		}
		symbolIndex := RandInt(0, 4)
		num1 := RandInt(1, maxNum)
		num2 := RandInt(1, maxNum)

		if symbol[symbolIndex] == "-" && num1 <= num2 {
			tmpNum2 := num1
			num1 = num2
			num2 = tmpNum2
		}
		fmt.Printf("%v %v %v = \t", num1, symbol[symbolIndex], num2)

	}

	fmt.Println()
}

// RandInt ...
func RandInt(min, max int32) int32 {
	/*if min == 0 && max == 0 || min >= max {
		fmt.Errorf("err")
		return max
	}*/
	rand.Seed(time.Now().UnixNano())
	return rand.Int31n(max-min) + min
}
