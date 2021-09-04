package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello World!") //Ch1
	DisplayTime()               //Ch1
	var x = 9                   //Ch2
	fmt.Println(x + 1)
	var num int = 3
	_ = num //the compiler now is happy
	var y int
	fmt.Println(y)

	var z bool = true
	fmt.Println(z)

	FirstName := "Mahdi"
	LastName, Age, ID := "Rezaei", "18", "1279403829"
	var god string = "in\tthe\tname\tof\tgod"
	fmt.Println(FirstName + LastName + Age + ID)
	fmt.Println(god)

	qoute := `"Hello World"
	Im Mahdi Rezaei
	welcome to babol`

	fmt.Println(qoute)

	/*var input string
	fmt.Print("Plaese Enter your age: ")
	fmt.Scanf("%s", &input)
	age, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("your age is: ", age)
	}*/

	v, err := DoSomething()
	if err {
		// handle the error
	} else {
		fmt.Println(v)
	}
	if v, err := DoSomething(); !err {
		fmt.Println(v)
	}

	DisplayDate("15:30  \"Fri\"  2006  MST   Jan", "Current date and time")
	s := "12345678"
	even, odd := evenORoddd(s)
	fmt.Println(even, odd)
	gen := fib()
	for g := 0; g < 10; g++ {
		fmt.Println(gen())
	}
	a := []int{1, 2, 3, 4, 5}
	evens := filter(a,
		func(val int) bool {
			return val%2 == 0
		})
	fmt.Println(evens)

	answer := filter(a,
		func(val int) bool {
			return val > 3
		})
	fmt.Println(answer)
	var cube [4][3][3]string
	for row := 0; row < 4; row++ {
		for column := 0; column < 3; column++ {
			for depth := 0; depth < 3; depth++ {
				cube[row][column][depth] =
					strconv.Itoa(row) +
						strconv.Itoa(column) +
						strconv.Itoa(depth)
			}
		}
	}
	fmt.Println(cube)
	g := make([]int, 0, 10)
	fmt.Println(g)
	g = append(g, 1, 2, 3, 4)
	fmt.Println(cap(g))
	fmt.Println(len(g))

	var m [3]string
	m[0] = "mahdi"
	m[1] = "ali"
	m[2] = "morteza"
	fmt.Println(m[:])
	t := []int{1, 2, 3, 4, 5}
	t, erro := insert(t, 2, 9)
	if erro == nil {
		fmt.Println(t) // 1 2 9 3 4 5]
	} else {
		fmt.Println(err)
	}
	t, errorr := delete(t, 2)
	if errorr == nil {
		fmt.Println(t) // [1 2 4 5]
	} else {
		fmt.Println(err)
	}
	//structs

	var p1 point
	p1.x = 3.2
	p1.y = 2.4
	p1.z = 4.5

	p2 := point{2.3, 1.3, 5.8}

	p3 := point{x: 2.9, y: 3}
	fmt.Println(p1, p2, p3)
	p4 := newPoint(4, 6.6, 6)
	fmt.Println(p4)
}

// define a struct for three dimonsional cordinate
type point struct {
	x float32
	y float32
	z float32
}

func DoSomething() (int, bool) {
	return 5, false
}

func evenORoddd(no string) (int, int) {
	even, odd := 0, 0
	for _, s := range no {
		if int(s)%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	return even, odd
}
func fib() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		f1, f2 = f2, (f1 + f2)

		return f1
	}
}

func filter(arr []int, cond func(int) bool) []int {
	result := []int{}
	for _, v := range arr {
		if cond(v) {
			result = append(result, v)
		}
	}
	return result
}

func newPoint(x, y, z float32) *point {
	p := point{x: x, y: y, z: z}
	return &p
}
