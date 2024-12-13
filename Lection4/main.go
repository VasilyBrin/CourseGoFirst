package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	//Boolean => default false
	var firstBoolean bool
	fmt.Println(firstBoolean)
	//Boolean operands
	aBoolean, bBoolean := true, true
	fmt.Println("AND:", aBoolean && bBoolean)
	fmt.Println("OR:", aBoolean || bBoolean)
	fmt.Println("NOT:", !aBoolean)

	//Numerics. Integer
	//int, int8, int16, int32, int64
	//uint8, uint16, uint32, uint64, uint
	var a int = 32
	b := 92
	fmt.Println("Value of a:", a, "Value of b:", b, "Sum of a+b:", a+b)
	//вывод типа через %T форматирование
	fmt.Printf("Type is %T\n", a)
	//Узнаем, сколько байт занимает переменная типа int
	fmt.Printf("Type %T size of %d bytes\n", a, unsafe.Sizeof(a))
	//выдаст - Type int size of 8 bytes (или 64 bit)

	/* Эксперимент 1. При использовании короткого объявления - тип для
	целого числа - int платформо-зависимый. */
	fmt.Printf("Type %T size of %d bytes\n", b, unsafe.Sizeof(b))

	/* Эксперимент 2. Использование явного приведения типов возможно если
	есть уверенность что не произойдет коллизии. */
	var first32 int32 = 12
	var second64 int64 = 13
	fmt.Println(first32 + int32(second64))
	//Приведение к типу int32. Число int64 может не уместиться в int32.
	//Возможна потеря в точности результата.

	fmt.Println(int64(first32) + second64)
	//Более правильный подход, приведение результата к типу int64.

	/* Эксперимент 3. Если проводятся арифметические операции над int и intX ,
	то обязательно нужно использовать механизм приведения. Т.к. int != int64 */
	var third64 int64 = 16123414
	var fourthInt int = 156234
	fmt.Println(third64 + int64(fourthInt))
	// + - * / %  - возможные арифметические операции с int числами(% - это остаток от /)

	//Аналогичным образом устроены unit8, uint16, uint32, uint64, uint
	//Numerics. Float
	//float32, float64
	floatFirst, floatSecond := 5.67, 12.54
	fmt.Printf("type of a %T and type of %T\n", floatFirst, floatSecond)
	sum := floatFirst + floatSecond
	sub := floatFirst - floatSecond
	fmt.Println("Sum:", sum, "Sub:", sub)

	fmt.Printf("Sum: %.3f and Sub: %.3f\n", sum, sub)
	//вывод вещественного числа с указанной точностью

	//Numeric. Complex - комплексные числа
	c1 := complex(5, 7)
	c2 := 12 + 32i
	fmt.Println(c1 + c2)
	fmt.Println(c1 * c2)

	//Strings. Строка - это набор БАЙТ
	name := "Федя"
	lastname := "Pupkin"
	concat := name + " " + lastname
	fmt.Println("Full name:", concat)
	fmt.Println("Length of string :", name, len(name))
	// Функция len() возвращает количество элементов в наборе

	fmt.Println("Amount of chars:", name, utf8.RuneCountInString(name))
	//Rune - руна. Это один utf-ный символ.
	//Поиск подстроки в строке
	totalString, subString := "ABCDEDFG", "asd"
	fmt.Println(strings.Contains(totalString, subString))
	//rune -> alias int32
	var sampleRune rune
	var anotherRune rune = 'Q' // Для инициализации руны символьным значением - используйте ''
	var thirdRune rune = 234
	fmt.Println(sampleRune)
	fmt.Printf("Rune as char %c\n", sampleRune)
	fmt.Printf("Rune as char %c\n", anotherRune)
	fmt.Printf("Rune as char %c\n", thirdRune)
	// "A" < "abcd"
	fmt.Println(strings.Compare("abcd", "a"))
	// -1 if first < second, 0 if first == second, 1 if first > second

	var aByte byte // alias uint8
	fmt.Println("Byte:", aByte)
}
