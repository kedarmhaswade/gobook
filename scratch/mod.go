package main

import "fmt"

func main() {
	// integer division truncates the result toward 0
	// % returns a number whose sign is that of the dividend
	fmt.Printf("5 / -3 = %d and 5 %% -3 = %d\n-5 / 3 = %d and -5 %% 3 = %d\n-5 / -3 = %d and -5 %% -3 = %d\n",
		    5 / -3,         5 %  -3,      -5 / 3,         -5 %  3,      -5 / -3,         -5 %  -3)
}
