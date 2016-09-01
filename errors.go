package main

import "errors"
import "fmt"

func f1(arg string) (string, error) {
	if arg == "sexy"{
	    return "sexy", errors.New("That's what you told me anyway...")
	}

	return "Could you repeat that?", nil
}

/*
Exlaining the above function

func f1(arg int) (int, error){

func         -> tells that it is a function.
f1           -> tells the name of the function.
(arg int)    -> tells that it accepts one variable and should be of integer
                type.
(int, error) -> tells that it returns 2 values first one is a int and the
                other is an error.
}
*/



func main() {
	fmt.Println(f1("sexy"))
}