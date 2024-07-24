package main

import "fmt"

func main() {

	// Basics
	/*
			var intNum uint16 = 32767
			fmt.Println(intNum)

			var floatNum float32 = 32765457.09
			fmt.Println(floatNum)

			var floatNum32 float32 = 10.1
			var intNum32 int32 = 2
			var result float32 = floatNum32 + float32(intNum32)

			fmt.Println(result)

			var myString string = `Hello
		WOlrd
		d
		d
		`

			var mySecondString string = "Hello worlds!!!"
			fmt.Println(mySecondString)
			fmt.Println(myString)

			fmt.Println(len("A"))
			fmt.Println(utf8.RuneCountInString(myString))

			var myRune rune = 'B'
			fmt.Println(myRune)

			var myBool bool = true
			fmt.Println(myBool)


	*/

	//VARIABLES IF/SWITCH STaTEMENTS
	/*
		var myInt int8
		fmt.Println(myInt)

		var myVar = "Test"
		fmt.Println(myVar)

		mySecondVar := "Test"
		fmt.Println(mySecondVar)

		var1, var2 := 1, 2
		fmt.Println(var1, "\n", var2)

		const myConst string = "const string"
		fmt.Println(myConst)

		printMe("mememm")

		result, remainder, err := intDivision(10, 8)

		if err != nil {
			fmt.Println(err.Error())
		} else if remainder == 0 {
			fmt.Printf("The result of the integer division is %v.\n", result)
		} else {
			fmt.Printf("The result of the integer division is %v with remainder %v.\n", result, remainder)
		}

		fmt.Println(intDivision(10, 2))

		switch remainder {
		case 0:
			fmt.Printf("The division was exact")
			break
		case 1, 2:
			fmt.Printf("The division was close")
			break
		default:
			fmt.Printf("The division was not close")
			break
		}

	*/

	// Arrays, Slices, Maps, Loops

	var intArr [3]int32
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	//fmt.Println(&intArr[0])  Adresy pamięci
	//fmt.Println(&intArr[1])
	//fmt.Println(&intArr[2])

	var secondIntArr [3]int32 = [3]int32{1, 2, 3}
	fmt.Println(secondIntArr)

	thridIntArr := [...]int32{1, 2, 3}
	fmt.Println(thridIntArr)

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("\n%v", intSlice)
	fmt.Printf("\nThe length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	//fmt.Println(intSlice[4])

	intSlice2 := []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)

	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Jason"])

	// odwołanie do mapy myMap2["value"] zwraca klucz i wartość bool oznaczającą czy szukana wartość znajduje się w mapie
	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Printf("The age is %v\n", age)
	} else {
		fmt.Printf("Invalid Name\n")
	}

	delete(myMap2, "Adam")
	myMap2["Adam"] = 23

	for name, age := range myMap2 {
		fmt.Printf("Name: %v  Age: %v\n", name, age)
	}

	tempIntArr := [...]int8{1, 2, 3, 5}

	for i, v := range tempIntArr {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for j := 0; j < 10; j++ {
		fmt.Print(j)
	}

	fmt.Println(10 % 2)
}

/*
func printMe(printvalue string) {
	fmt.Println("Hello World")
	fmt.Println(printvalue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}
	result := numerator / denominator
	remainder := numerator % denominator
	return result, remainder, err
}

*/
