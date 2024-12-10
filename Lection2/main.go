package main

import (
	"fmt"
	"math"
)

func main() {
	// Простейший вывод на консоль. println - это вывод аргумента + '\n'
	fmt.Println("Hello world")
	fmt.Println("Second line")
	//Функция print - простой вывод аргумента
	fmt.Print("First")
	fmt.Print("Second")
	fmt.Print("Third")
	//Форматироованный вывод: Printf - стандатный вывод os.Stdout с флагами форматирования
	fmt.Printf("\nHello, my name is %s\nMy age is: %d\n", "Vasyan", 67)
	////////////////////////////////////
	////////////////////////////////////
	//Декларирование переменных
	var age int
	fmt.Println("My age is:", age)
	age = 32
	fmt.Println("Age after assignment:", age)

	//Декларирование и инициализация пользовательским значением
	var height int = 183
	fmt.Println("My height is:", height)

	//В чем "полустрогость" типизации? Можно опускать тип переменной
	var uid = 12345
	fmt.Println("My uid:", uid)
	//Декларирование и инициализация переменных одного типа (множественный случай)
	var firstVar, secondVar int = 20, 30
	fmt.Printf("FirstVar:%d SecondVar:%d\n", firstVar, secondVar)
	//Декларирование блока переменных
	var (
		personName string = "Vasyan"
		personAge  int    = 67
		personUID  int
	)
	fmt.Printf("Name: %s\nAge %d\nUID: %d\n", personName, personAge, personUID)

	//Немного странного. Как только отсутствует явная типизация, то, управление сразу
	//передается компилятору.
	var a, b = 30, "Vasya"
	fmt.Println(a, b)
	a = 300
	//Немного хорошего. Повторное декларирование переменной приводит к ошибке компиляции
	//var a = 200 - будет ошибка компиляции.

	//Короткая декларация (короткое объявление)
	count := 10
	fmt.Println("Count:", count)
	count++
	fmt.Println("Count:", count)
	count = count + 1
	fmt.Println("Count:", count)

	//Множественное присваивание через :=
	//aArg, bArg := 10, "Vasya"
	//fmt.Println(aArg, bArg)
	// 10 Vasya
	aArg, bArg := 10, 30
	fmt.Println(aArg, bArg)
	aArg, bArg = 30, 40 // замещение значений
	fmt.Println(aArg, bArg)
	//aArg, bArg := 10, 30  // Редекларация, повторное объявление через := приводит к ошибке
	//fmt.Println(aArg, bArg)

	//Исключение из этого правила. Слева от := должна быть хоть одна новая переменная
	bArg, cArg := 300, 400
	fmt.Println(aArg, bArg, cArg)

	//Пример
	width, length := 20.5, 30.2
	fmt.Printf("Min dimensional of rectangle is : %.2f\n", math.Min(width, length))
}
