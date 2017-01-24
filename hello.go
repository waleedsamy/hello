package main

import "fmt"

const (
	/*PI jhg*/
	PI = iota
	/*CI hgk*/
	CI
	/*DI ihjl*/
	DI
	/*RI jvj*/
	RI
)

/*Rectangle Shape of type Rectangle*/
type Rectangle struct {
	width, height int
}

/*Square Shape of type Square*/
type Square struct {
	width int
}

func (r Rectangle) area() int {
	return r.width * r.height
}

func (r *Rectangle) doubleTheWidth() {
	r.width = r.width * 2
}

func (s Square) area() int {
	return s.width * s.width
}

func objori() {
	r := Rectangle{5, 10}
	fmt.Printf("Rectangle %v area is %d\n", r, r.area())
	r.doubleTheWidth()
	fmt.Printf("Doubled Width Rectangle %v area is %d\n", r, r.area())

	s := Square{width: 5}
	fmt.Printf("Square %v area is %d\n", s, s.area())
}

type testint func(int) bool

/*Person is a Guy*/
type Person struct {
	name string
	age  int
}

func older(p1, p2 Person) (Person, int) {
	if p1.age < p2.age {
		return p2, p2.age - p1.age
	}
	return p1, p1.age - p2.age
}

func structs() {
	var a1 Person
	a1.name = "john"
	a1.age = 30

	a2 := Person{name: "marc", age: 40}

	olderPerson, diff := older(a1, a2)

	fmt.Printf("old person is %v by %d\n", olderPerson, diff)
}

func isEven(i int) bool {
	if i%2 == 0 {
		return true
	}
	return false
}

func isOdd(i int) bool {
	if i%2 != 0 {
		return true
	}
	return false
}

func filter(slice []int, f testint) (result []int) {
	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func sumAndProduct(a int, b int) (aPlusb int, aProductb int) {
	return a + b, a * b
}

func cstatments() {
	if c := 55; c > 3 {
		fmt.Printf("x is %d\n", c)
	} else {
		fmt.Printf("x less than 55\n")
	}

	sum := 1

	for sum < 100 {
		sum++
	}

	fmt.Printf("sum is %d\n", sum)
	cds := map[string]int32{"a": 5, "b": 6}
	for k, v := range cds {
		fmt.Printf("key is %s\n", k)
		fmt.Printf("value is %d\n", v)
	}

	aPlusB, aProductB := sumAndProduct(3, 4)
	fmt.Printf("a + b = %d\n", aPlusB)
	fmt.Printf("a * b = %d\n", aProductB)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	oddNumbers := filter(numbers, isOdd)
	EvenNumbers := filter(numbers, isEven)

	fmt.Printf("odd numbers is %v\n", oddNumbers)
	fmt.Printf("even numbers is %v\n", EvenNumbers)

}

func ar() {
	var ff [4]int8

	ff[0] = 3

	zz := [1]string{"fgg"}

	cc := [...]int{1, 3, 444}

	fmt.Printf("array %d %s %d\n", ff[0], zz[0], cc[0])

	zzce := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}

	fmt.Printf("%v\n%+v\n%#v\n", zzce, zzce, zzce)

	cds := map[string]int32{"a": 1, "b": 2, "c": 3, "c++": 4, "c#": 5}

	csharp, ok := cds["c#"]

	if ok {
		fmt.Printf("charp found with value %d\n", csharp)
	} else {
		fmt.Printf("charp not found\n")
	}

	delete(cds, "c")
	fmt.Printf("cds has %v\n", cds)

}

func main1() {
	fmt.Printf("%d, %d, %d, %d", PI, CI, DI, RI)
	var v1, v2, v3 int = 1, 2, 3
	_, b := 34, 35
	_, gg := 55, 77

	mline := `multi
	line`
	s := "hello"
	s = "c" + s[1:] // you cannot change string values by index, but you can get values instead.
	fmt.Printf("%s %s\n", s, mline)

	const PI float32 = 3.145566

	fmt.Printf("%d %d %d %d %d %f\n", v1, v2, v3, b, gg, PI)
	// ar()
	// cstatments()
	// structs()
	objori()
}
