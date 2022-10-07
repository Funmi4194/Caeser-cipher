package caesar

import "fmt"

const min byte = 65
func EncryptChar(word string, key byte){//A function that takes the word and key parameter
	
    fmt.Println("Enter Key number:", key) //An fmt package that tells the user to input a Key value
	fmt.Println("Enter word:", word) //An fmt package that tells the user to input a word value

	for i := 0; i < len(word); i++{ // create a for loop for my character
		char := word[i]
		result := (char + key - min)%26 + min
		
		//if result > max {
		//	offset:= result - max //how far away from max is the result
		//	result = offset + min
	
		//}
		//result := word[i] + key // adding my key value declared in byte to the byte value of my string
		str := string(result) // convert byte values to string
		fmt.Println(str, result) // print encrpyted string and value
		
	}

}