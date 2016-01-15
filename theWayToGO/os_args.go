package main
import (
	"os"
	"strings"
)

func main() {
	who := "alice "
	if len(os.Args) > 1 {
		who +=strings.Join(os.Args[1:]," ")
	}
	println("Good Morning",who)
	
}
