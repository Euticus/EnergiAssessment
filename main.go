package main


import "fmt"


func Generate(out chan<- int) {
	for i := 2; ; i++ {
		out <- i                  
	}
}

func Filter(in <-chan int, out chan<- int, prime int) {
	m := prime * prime              
	for {
		i := <- in               
		for i > m {
			m = m + prime    
		}
		if i < m {
			out <- i         
		}
	}
}

func Sieve(out chan<- int) {
	gen := make(chan int)           
	go Generate(gen)                
	p := <- gen
	out <- p
	p = <- gen         
	out <- p           

	base_primes := make(chan int)     
	go Sieve(base_primes)
	bp := <- base_primes             
	bq := bp * bp              

	for  {
		p = <- gen
		if p == bq {                  
			ft := make(chan int)
			go Filter(gen, ft, bp) 
			gen = ft
			bp = <- base_primes     
			bq = bp * bp            
		} else {
			out <- p
		}
	}
}

func main() {
	sv := make(chan int)          
	go Sieve(sv)                    
	lim := 1000000
	for i := 0; i < lim; i++ {
		prime := <- sv
		if i == (lim-1) {
			fmt.Printf("The %dth prime number is: \t", lim)
			fmt.Printf("%4d ", prime)

		}
	}
}
