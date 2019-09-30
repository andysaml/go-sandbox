package main

func factor(x int) int {
	if x == 1 {
		return x
	}
	return x * factor(x-1)
}

func factorTail(x, acc int) int {
	acc *= x
	if x == 1 {
		return acc
	}
	return factorTail(x-1, acc)

}

func factorTailWrapper(n int) int {
	factorTail := func(x, acc int) int {
		acc *= x
		if x == 1 {
			return acc
		}
		return factorTail(x-1, acc)
	}
	return factorTail(n, 1)
}
