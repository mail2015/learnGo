package main

import "fmt"

func main(){
	methodIf()

	methodIfElse()

	methodIfMultiple()

	methodSwitch()

	methodSelect()

}

func methodIf(){
	fmt.Println(" \n\n \t if语句 ")

	var a int = 10
	if(a < 20){
		fmt.Println(" a is less than 20 \n ")
	}
	fmt.Printf(" case1 finish, a vlaue is: %d \n ", a)
}


func methodIfElse(){
	fmt.Println(" \n\n \t if...else语句 ")

	var a int = 100

}


func methodIfMultiple(){
	fmt.Println(" \n\n \t if循环嵌套  ")
}


func methodSwitch(){
	fmt.Println(" \n\n \t switch语句 ")
}


func methodSelect(){
	fmt.Println(" \n\n \t select语句 ")
}
