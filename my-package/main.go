package main

import (
	"bondarik.net/slice"
	"fmt"
)

func main() {

	s := slice.Slice{1, 2, 3}
	fmt.Println(s)
	fmt.Println("Сумма слайса: ", slice.SumSlice(s))

	slice.MapSlice(s, func(i slice.Element) slice.Element {
		return i * 2
	})

	fmt.Println("Слайс, умноженный на два: ", s)

	fmt.Println("Сумма слайса: ", slice.SumSlice(s))

	fmt.Println("Свёртка слайса умножением ",
		slice.FoldSlice(s,
			func(x slice.Element, y slice.Element) slice.Element {
				return x * y
			},
			1))

	fmt.Println("Свёртка слайса сложением ",
		slice.FoldSlice(s,
			func(x slice.Element, y slice.Element) slice.Element {
				return x + y
			},
			0))

}
