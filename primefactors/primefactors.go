package primefactors


func GeneratePrimeCandidates2(done <- chan struct{}) <- chan int {
	c := make(chan int)

	go func() {

		var next = -1
		for {
			switch next {
				case -1: next = 2
				case 2: next = 3
				default: next += 2
			}

			select {
				case c <- next:
				case <- done:
					return
			}
		}
	}()
	return c
}

func Sieve(in <- chan int, done <- chan struct{}, val int) <- chan int {
	out := make(chan int)

	go func() {
		for {
			select {
				case n := <- in:
					if n % val != 0 {
						select {
						case out <- n:
						case <- done:
							return
						}
					}
				case <- done:
					return

			}
		}
	}()

	return out
}

func PrimeFactors(n int) []int {
	primeFactors := []int{}
	done := make(chan struct{})
	defer close(done)
	c2 := GeneratePrimeCandidates2(done)
	for ; n > 1; {
		prime := <- c2
		c2 = Sieve(c2, done, prime)

		for n % prime == 0 {
			n = n / prime
			primeFactors = append(primeFactors, prime)
		}
	}

	return primeFactors
}
