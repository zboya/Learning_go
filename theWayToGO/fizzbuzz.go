package main

func main(){
    for i:=1;i<=100;i++ {
        switch {
            case (i%3)==0 && (i%5)!=0:
                print("fizz\n")
            case (i%3)!=0 && (i%5)==0:
            	print("buzz\n")
            case (i%3)==0 && (i%5)==0:
            	print("FizzBuzz\n")
            default:
            	print(i,"\n")
            }
    }
}
