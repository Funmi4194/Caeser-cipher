package caesar
import "fmt" 
const max byte = 97
func DecryptChar(word2 string, key2 byte){
	
    fmt.Println("Enter Key number:", key2) //An fmt package that tells the user to input a Key value
	fmt.Println("Enter word:", word2) //An fmt package that tells the user to input a encrpted value
	for i := 0; i < len(word2); i++{ // create a for loop for my character
		char := word2[i]
        result := (char + key2 - max)%26 + max
	
		// result := char +  key2
        // result := word2[i]- key2 // substrate my key value to get encrypted value back
		//(declared in byte to the byte value of my string)

		str := string(result) // convert byte values to string
		fmt.Println(str , result) // print encrpyted string and value
	}

}