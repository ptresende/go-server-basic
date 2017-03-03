package main


type Rectangle struct {
	number_one int
	number_two int
}

func (tis *Rectangle) perimeter() int {
	return (tis.number_one + tis.number_two) * 2
}
