package iteration

import "fmt"

func ForLoops() {

	var i int = 0

	for i <= 3 {
		fmt.Printf("%d", i)
		i++
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
