package main

import "fmt"

func main(){
	var a [5]int
	fmt.Println("emp:", a)

	a[2] = 80
	fmt.Println("set:", a)
	fmt.Println("get:", a[2])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

	b = [...]int{100, 2: 400, 4: 500}
    fmt.Println("idx:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i+j
		}
	}
	fmt.Println("2d:", twoD)

	var threeD = [2][2][2]int{
		{
			{1,2},
			{2,3},
		},
		{
			{3,4},
			{4,5},
		},
	}
	fmt.Println("3d:", threeD)
}