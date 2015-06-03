package main

import "fmt"

func main() {
	booleanType()
	printNewLine()

	numericType()
	printNewLine()

	stringType()
	printNewLine()
	derivedType()

	printNewLine()
	printNewLine()

	//
	intType()
	printNewLine()

	floatType()
	printNewLine()

	numberType()
	printNewLine()
}

func printNewLine() {
	fmt.Println("\n")
}

func booleanType() {
	fmt.Println("Boolean Types, 两个预定义常量 true false")
	var var1 = false
	var var2 = true
	fmt.Println("var1=", var1)
	fmt.Println("var2=", var2)
}

func numericType() {
	fmt.Println("Numeric Types 表示整数或boolean")
	var num1 = 0
	fmt.Println("num1=", num1)

	var num2 = 12.234
	fmt.Println("num2=", num2)

	var num3 = false
	fmt.Println("num3=", num3)

}

func stringType() {
	fmt.Println("字符串, 一次创建稳定的字符串, 不能改变")
	var test1 = "str1"
	fmt.Println("test1=", test1)
}

func derivedType() {
	fmt.Println(" (a)指针类型  (b)数据类型  (c)结构类型  (d)联盟类型 ")
	fmt.Println(" (e)函数类型  (f)切片类型  (h)接口类型  (i)映射类型 ")
	fmt.Println(" (j)管道类型 ")
}

func intType() {

	fmt.Println(" 整数类型测试 ")

	var uint8_1 uint8
	fmt.Println("uint8_1=", uint8_1)

	var uint16_1 uint16
	fmt.Println("uint16_1=", uint16_1)

	var uint32_1 uint32
	fmt.Println("uint32_1=", uint32_1)

	var uint64_1 uint64
	fmt.Println("uint64_1=", uint64_1)

	var int8_1 int8
	fmt.Println("int8_1=", int8_1)

	var int16_1 int16 = 1600
	fmt.Println("int16_1=", int16_1)

	var int32_1 int32 = 3200
	fmt.Println("int32_1=", int32_1)

	var int64_1 int64
	fmt.Println("int64_1=", int64_1)

}

func floatType() {
	fmt.Println(" 浮点类型 float32  float64 complex64  complex128")

	var float32_1 float32
	fmt.Println("float32_1=", float32_1)

	var float64_1 float64 = 123.122
	fmt.Println("float64_1=", float64_1)

	var complex64_1 complex64
	fmt.Println("complex64_1=", complex64_1)

	var complex128_1 complex128 = 112312.123
	fmt.Println("complex128_1=", complex128_1)

}

func numberType() {
	fmt.Println(" 其它数据类型 byte  rune  uint  int  uintptr")

	var byte1 byte
	fmt.Println("byte1=", byte1)

	var rune1 rune
	fmt.Println("rune1=", rune1)

	var uint1 uint
	fmt.Println("unit1=", uint1)

	var int1 int
	fmt.Println("int1=", int1)

	var uintptr1 uintptr
	fmt.Println("uintptr1=", uintptr1)

}
