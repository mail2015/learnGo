package main
	
import "fmt"

func main(){

   forMethod();

   forMultipleMethod();

   breakStatement();

   continueStatement();

   gotoStatement();

}



func forMethod(){
	fmt.Println(" \n\n \t for 循环 \n ")
	var b int = 10
	var a int

	numbers := [6]int{1, 2, 3, 6}
	/* for loop execution */
	fmt.Println(" \n for case 1 ")
	for i := 0; i < 10; i++{
		fmt.Printf(" \n value of i is: %d ", i)
	}
	fmt.Println(" \n for case 2 -----")
	for a < b {
		a++
		fmt.Println(" for a value is ", a, " for a value is ", b)
	}
	fmt.Println(" \n for case3 ")
	for i, x :=range numbers{
		fmt.Printf(" \n value od x = %d and i =  %d ", x, i)
	}


	fmt.Println(" \n\n a value is ", a, " b value is ", b, " numbers value is ", numbers)
}

func forMultipleMethod(){
	fmt.Println(" \n\n \t 嵌套循环  \n")

	var i, j int
	for i = 2; i <= 100; i++ {
		for j = 2; j <= (i / j); j++{
			if( i % j == 0){
				fmt.Println(" i value is ", i , " and j value is ", j , " i % j == 0 ")
				break;
			}
		}
		if(j > ( i / j)){
			fmt.Printf(" \n  i = %d,  j = %d is prime \n ", i, j)
		}

	}
}


func breakStatement(){
	fmt.Println(" \n\n \t break语句 ")
	var a int = 10

	/* for loop execution */
	for a < 25 {
		fmt.Printf(" \n value of a: %d ", a)
		a++
		if a > 15{
			fmt.Println("\n\n break when a > 15, a =", a)
			break
		}
	}
}


func continueStatement(){
	fmt.Println(" \n\n \t  continue 语句 ")

	var a int = 10

	for a < 25 {
		if a == 15 || a == 17 || a == 18 {
			a = a + 1
			continue
		}
		fmt.Println(" value a is : ", a)
		a++
	}

}

func gotoStatement(){
	fmt.Println(" \n\n \t goto语句 ")

}