package piscine

func RecursiveFactorial(nb int) int {
	Result := 1
	if nb < 0 || nb > 26 {
		Result = 0
	}
	if nb <= 26 {
		Result = nb * RecursiveFactorial(nb-1)
	}
	return Result
}
