package main

import "fmt"

func main(){

	num1 :=5
	num2 :=4

	result := add(num1,num2)

	fmt.Println(result)

	result1,result2 := calc(num1,num2)

	fmt.Println(result1,result2)

	result3,result4 := calc1(num1,num2)

	fmt.Println(result3,result4)

}

// func (r receiver) identifier(params) returns {code}
func add(x int,y int) int{

	var out = x + y
	return out
}

func calc(x int,y int) (int,int){

	var out1 = x + y
	var out2 = x - y
	return out1,out2
}

func calc1(x int,y int) (out1, out2 int){

	out1 = x + y
	out2 = x - y
	return
}