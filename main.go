package main

import (
	"fmt"
  "strconv"
  "strings"
)

func main() {
  
	var primer string
  fmt.Println("Введите выражение")
  fmt.Scan(&primer)

  //per1 := strings.Split(primer, "")
  //n, err := strconv.Atoi(per1[0])
  n, err := strconv.Atoi(strings.Split(primer, "")[0])
  if err != nil {
    panic(err)
  }

  //per2 := strings.Split(primer, "")
  d, err := strconv.Atoi(strings.Split(primer, "")[2])
  if err != nil {
    panic(err)
  }

  
  //per2 := strings.Split(primer, "")
  var operand string = (strings.Split(primer, "")[1])

  if operand == "+" {
    fmt.Println(n + d)
  } else if operand == "-" {
    fmt.Println(n - d)
  } else if operand == "*" {
    fmt.Println(n * d)
  } else if operand == "/" {
    fmt.Println(n/d)
  }


  
  
}