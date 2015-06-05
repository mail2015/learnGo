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
	if(a < 20 ){
		fmt.Println(" a is less than 20\n ")
	}else{
		fmt.Println(" a is large than 20\n ")
	}

	if(a == 20 ){
		fmt.Println(" a is equal to 20\n ")
	}else if(a == 50){
		fmt.Println(" a is equal to 50\n ")
	}else if(a == 100){
		fmt.Println(" a is equals to 100\n ")		
	}else if(a == 200){
		fmt.Println(" a is equals to 200\n ")
	}else{
		fmt.Println(" a= ", a, " not a expected value")
	}


	fmt.Printf(" a vlaue is: %d \n ", a)

}


func methodIfMultiple(){
	fmt.Println(" \n\n \t if循环嵌套  ")
	var a int = 100
	var b int = 200
	if(a == 100){
		if(b == 200){
			fmt.Println(" a value is 100 and b vlaue is 200 \n ")
		}
	}

	fmt.Println(" case3 finish \n ")
}


func methodSwitch(){
	fmt.Println(" \n\n \t switch语句 ")

	var grade string  = "B"
	var marks int = 90

	switch marks{
		case 90 : grade = "AA"
		case 80 : grade = "BB"
		case 50, 60, 70 : grade = "CC"
		default : grade = "DD"  
	}
	fmt.Println(" marks is ", marks, "  grade is ", grade)

	switch{
		case grade == "AA" : 
		  fmt.Println(" Excellent! ")
		case grade == "BB" :
		  fmt.Println(" Well done! ")
		case grade == "CC" :
		  fmt.Println(" Normal! ")
		case grade == "DD" :
		  fmt.Println(" Not pass! ")
		default :
		   fmt.Print(" N/A ")        
	}

	// switch type
    var x interface{}

    switch i := x.(type){
    	case nil : 
    	  fmt.Printf(" type of x : %T ", i)
    	case int : 
    	  fmt.Println(" type of x is int ")
    	case float64 : 
    	  fmt.Println(" type of x is float64 ")
    	case func(int) float64 :
    	  fmt.Println(" type of x is func(int) ") 
    	case bool , string : 
    	  fmt.Println(" type of x is bool or string ") 
    	default : 
    	  fmt.Printf(" type is x is %T ", i)  
    }


	fmt.Println(" case4 finish. marks is ", marks, "  grade is ", grade)
}


func methodSelect(){
	fmt.Println(" \n\n \t select语句 ")
}
