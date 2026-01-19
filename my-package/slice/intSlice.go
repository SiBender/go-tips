package slice

type Slice []Element

type Element int

func SumSlice(s []Element) (sum Element) {
	for _, value := range s {
		sum += value
	}
	return
}

func MapSlice(s []Element, op func(Element) Element) {
	for i, value := range s {
		s[i] = op(value)
	}
}
func FoldSlice(s []Element, op func(Element, Element) Element, init Element) Element {
	subValue := init
	for _, value := range s {
		subValue = op(subValue, value)
	}

	return subValue
}
