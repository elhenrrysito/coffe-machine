package machine

import "fmt"

func MakeCoffe(name string) {
	fmt.Println("Making coffee...")
	coffee := fmt.Sprintf("Here is coffee for %s", name)
	fmt.Println(coffee)
}
