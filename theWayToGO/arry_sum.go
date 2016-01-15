package main

func main() {
	arr :=[3]float64{2.4,6.9,2.0}
	res:=sum(&arr)
	println(res)
	
}

func sum(a *[3]float64) (s float64) {
	for _,v := range a {
		s+=v
	}
	return
}