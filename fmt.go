package main
import "fmt"

func main() {
	firstname := "Fardan"
	lastname := "Nozami"
	fmt.Println("Hello '", firstname, lastname, "'")

	fmt.Printf("Hello '%s  %s'\n", firstname, lastname) // using printf to print formatted strings
}