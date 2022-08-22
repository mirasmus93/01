package 01

import "fmt"

func gcd(){
	fmt.Println(" // import ( 
		    "fmt"
		    "os"
		    "strconv"
		   )
	func main() {
	args := os.Args[1:]
	v1, _ := strconv.Atoi(args[0])
	v2, _ := strconv.Atoi(args[1])

	if len(args) != 2 {
		fmt.Println()
	} else {
		gcd(v1, v2)
		}
	}
	
	func gcd(num1 int, num2 int) {
	var gcdnum int
	for i := 1; i <= num1 && i <= num2; i++ {
		if num1%i == 0 && num2%i == 0 {
			gcdnum = i
			}
		}
	fmt.Println(gcdnum)
	}")
