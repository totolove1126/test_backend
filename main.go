package main

import (
	"fmt"
)

func main() {
	var choice string
	fmt.Print("Input Char : ")
	fmt.Scan(&choice)
	GenarateNumberBycode(choice)
}

func GenarateNumberBycode(str string) (num int) {
	total := 0
	for i := 0; i < len(str); i++ {

		// current คือตัวอักษร
		current := str[i]
		// ส่งสตริงไปดึงค่าจากฟังก์ชัน GetcharValue
		currentValue := GetcharValue(current)

		//Check Special Case + เช็คตัวแปรต่อไปว่ามีจริงไหม
		if i+1 < len(str) && IsSpecialCase(current, str[i+1]) {
			nextValue := GetcharValue(str[i+1])
			currentValue = nextValue - currentValue
			i++
		}
		total += currentValue
	}

	fmt.Println(total)
	return

}

func GetcharValue(str byte) (num int) {
	switch str {
	case 'A':
		return 1
	case 'B':
		return 5
	case 'Z':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'R':
		return 1000
	default:
		return 0
	}

}

func IsSpecialCase(current byte, next byte) (suc bool) {
	return (current == 'A' && (next == 'B' || next == 'Z')) || (current == 'Z' && (next == 'L' || next == 'C')) || (current == 'C' && (next == 'D' || next == 'R'))
}
