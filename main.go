package main

import "fmt"

func main() {
	inp := [12]int{366, 456, 163, 673, 127, 387, 973, 824, 367, 766, 357, 287}
	exp := [12]int{127, 163, 287, 357, 366, 367, 387, 456, 673, 766, 824, 973}

	Sort(&inp)

	fmt.Printf("res: %t\n", inp == exp)
}

func Sort(inp *[12]int) {
	var inpMax int
	for _, num := range inp {
		if num > inpMax {
			inpMax = num
		}
	}
	rndCnt := digitsCount(inpMax)

	for rnd := 0; rnd < rndCnt; rnd++ {
		var out [12]int
		var cts [10]int
		var pxs [10]int

		for idx, num := range inp {
			_ = idx
			_ = num

			dig := nthDigit(num, rnd)
			cts[dig]++
		}

		var sum int
		for idx, cnt := range cts {
			pxs[idx] += sum + cnt
			sum += cnt
		}

		for idx := len(inp) - 1; idx >= 0; idx-- {
			num := inp[idx]
			dig := nthDigit(num, rnd)
			nid := pxs[dig] - 1

			pxs[dig] = nid
			out[pxs[dig]] = num
		}

		*inp = out
	}
}

func nthDigit(num int, n int) int {
	for ; n > 0; n-- {
		num /= 10
	}
	return num % 10
}

func digitsCount(num int) (count int) {
	for num != 0 {
		num /= 10
		count += 1
	}
	return
}
