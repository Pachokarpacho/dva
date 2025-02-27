package main

import (
 "fmt"
 "math"
 "strconv"
)

// Функция для проверки, является ли число простым/
func isPrime(n int) bool {
 if n < 1 {
  return false
 }
 if n == 2 || n == 3  || n == 1{
  return true
 }
 if n%2 == 0 {
  return false
 }
 for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
  if n%i == 0 {
   return false
  }
 }
 return true
}

func main() {
 var input string
 fmt.Print("Введите число: ")
 fmt.Scanln(&input)

 num, err := strconv.Atoi(input)
 if err != nil {
  fmt.Println("Ошибка: введенное значение не является целым числом.")
  return
 }

 if num < 1 {
  fmt.Println("Ошибка: число должно быть положительным и >= 1.")
  return
 }

 if isPrime(num) {
  fmt.Printf("%d - простое число.\n", num)
 } else {
  fmt.Printf("%d - не является простым числом.\n", num)
 }
}