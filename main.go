package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	var expr string
	fmt.Println("Введите выражение")
	fmt.Scan(&expr)

	if len(expr) < 3 {
		err := errors.New("Вывод ошибки, так как строка не является математической операцией.")
		fmt.Print(err)
	}

	var sign []string

	re1 := regexp.MustCompile("[+,--,*,/]+")
	sign = re1.FindAllString(expr, -1)
	var sign1 = sign[0]
	if len(sign) > 1 {
		err := errors.New("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		fmt.Print(err)
	}

	var i [3]byte = [3]byte{'I', 'V', 'X'}
	var rim_num []string

	if expr[0] == i[0] ||
		expr[0] == i[1] ||
		expr[0] == i[2] {
		re2 := regexp.MustCompile("[I,II,III,IV,V,VI,VII,VIII,IX,X]+")
		rim_num = re2.FindAllString(expr, -1)
		if len(rim_num) == 1 {
			err := errors.New("Вывод ошибки, так как используются одновременно разные системы счисления.")
			fmt.Println(err)
		}
		//if rim_num[0]

		n1 := decoder(rim_num[0], rim_num[1])
		//fmt.Println(n1)

		var number int
		var number1 int

		number, err := strconv.Atoi(n1[0])
		if err != nil {
			panic(err)
		}
		number1, err = strconv.Atoi(n1[1])
		if err != nil {
			panic(err)
		}

		if sign1 == "+" {
			fmt.Println(coder(number + number1))
		} else if sign1 == "-" {
			fmt.Println(coder(number - number1))
		} else if sign1 == "*" {
			fmt.Println(coder(number * number1))
		} else if sign1 == "/" {
			fmt.Println(coder(number / number1))
		}

		if number-number1 <= 0 {
			err := errors.New("Вывод ошибки, так как в римской системе нет отрицательных чисел и ноля.")
			fmt.Print(err)
		}

	} else {

		var nums []string

		re := regexp.MustCompile("[0-9]+")
		nums = re.FindAllString(expr, -1)
		//fmt.Println(nums)
		if len(nums) == 1 {
			err := errors.New("Вывод ошибки, так как используются одновременно разные системы счисления.")
			fmt.Println(err)
		}

		var num1 int
		var num2 int

		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		num2, err = strconv.Atoi(nums[1])
		if err != nil {
			panic(err)

		}

		if sign1 == "+" {
			fmt.Println(num1 + num2)
		} else if sign1 == "-" {
			fmt.Println(num1 - num2)
		} else if sign1 == "*" {
			fmt.Println(num1 * num2)
		} else if sign1 == "/" {
			fmt.Println(num1 / num2)
		}

	}
}

func decoder(x string, y string) []string {
	var rim = map[string]string{
		"I":    "1",
		"II":   "2",
		"III":  "3",
		"IV":   "4",
		"V":    "5",
		"VI":   "6",
		"VII":  "7",
		"VIII": "8",
		"IX":   "9",
		"X":    "10",
	}

	n := []string{}
	var a string = x
	keys := []string{}
	for key := range rim {
		keys = append(keys, key)
	}

	for i := 0; i <= 9; i++ {
		if a != x {
			break
		}
		if a == keys[i] {
			a = rim[keys[i]]
			x = a
		}
	}

	var b string = y
	for i := 0; i <= 9; i++ {
		if b != y {
			break
		}
		if b == keys[i] {
			b = rim[keys[i]]
			y = b
		}
	}

	n = append(n, x, y)
	return n

}

func coder(x int) string {
	var rimtoInt = map[int]string{
		1:   "I",
		2:   "II",
		3:   "III",
		4:   "IV",
		5:   "V",
		6:   "VI",
		7:   "VII",
		8:   "VIII",
		9:   "IX",
		10:  "X",
		20:  "XX",
		30:  "XXX",
		40:  "XL",
		50:  "L",
		60:  "LX",
		70:  "LXX",
		80:  "LXXX",
		90:  "XC",
		100: "C",
	}

	var a int = x
	keys := []int{}
	for key := range rimtoInt {
		keys = append(keys, key)
	}
	//fmt.Println(keys)
	i := 0
	for range keys {
		if a != x {
			break
		}
		if a == keys[i] {
			a = keys[i]
			x = a
			i++
		}
	}

	x1 := strconv.Itoa(x)
	x1 = rimtoInt[x]
	return x1
}
