package main

func main() {
	for i := 0; i < 10; i++ {
		if i==5 {
			continue
		}
		print(i," ")
	}
}

// 0 1 2 3 4 6 7 8 9