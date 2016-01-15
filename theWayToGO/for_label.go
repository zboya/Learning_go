package main

func main() {

	LABEL:
	for i := 0; i < 3; i++ {

		for j := 0; j < 10; j++ {
			if j==5 {
				// continue LABEL
				break LABEL
			}
			print(i,j," ")
		}

	}
}
//continue 00 01 02 03 04 10 11 12 13 14 20 21 22 23 24
//break 00 01 02 03 04
