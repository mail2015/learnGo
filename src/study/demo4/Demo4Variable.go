package main

import "fmt"

func main() {
	fmt.Println("\n\n 定义变量")
	var i, j, k int
	var c, ch byte
	var f, salary float32
	// d = 42 error
	d := 42
	fmt.Println("i=", i, " j=", j, " k=", k)
	fmt.Println("c=", c, " ch=", ch)
	fmt.Println("f=", f, " salary=", salary)
	fmt.Println("d=", d)

	// 静态类型
	fmt.Println("\n\n 静态类型")
	var x float64
	x = 20.123
	fmt.Println("x=", x)
	fmt.Printf("x type is %T\n ", x)

	// 动态类型
	fmt.Println("\n\n 动态类型")
	var v1 float64 = 20.222
	v2 := 42
	fmt.Println("v1=", v1)
	fmt.Println("v2=", v2)
	fmt.Printf("v1 type is %T\n", v1)
	fmt.Printf("v2 type is %T\n", v2)

	// 混合类型
	fmt.Println("\n\n 混合类型")
	var m1, m2, m3 = 3, 4, "foo"
	fmt.Println("m1=", m1)
	fmt.Println("m2=", m2)
	fmt.Println("m3=", m3)
	fmt.Printf("m1 type is %T\n ", m1)
	fmt.Printf("m2 type is %T\n ", m2)
	fmt.Printf("m3 type is %T\n ", m3)

}
