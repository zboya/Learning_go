package main
import "fmt"
func main() {
	for i := 0; i < 3; i++ {

		for j := 0; j < 10; j++ {
			if j>5 {
				break
			}
			print("j",j)
		}

		print(" ")

	}

	/*i := 0
	for { //since there are no checks, this is an infinite loop
	    if i >= 3 { break }
	    //break out of this for loop when this condition is met
	    fmt.Println("Value of i is:", i)
	    i++;
	}
	fmt.Println("A statement just after for loop.")*/
}
// j0j1j2j3j4j5 j0j1j2j3j4j5 j0j1j2j3j4j5