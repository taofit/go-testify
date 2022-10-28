package starter

import (
	"fmt"
	"net/http"
)

func SayHello(name string) string {
	return fmt.Sprintf("Hello %v. Welcome!", name)
}

func OddOrEven(num int) string {
	if num%2 == 0 {
		return fmt.Sprintf("%d is an even number", num)
	}
	return fmt.Sprintf("%d is an odd number", num)
}

func CheckHealth(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "http check passed")
}