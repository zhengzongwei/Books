package main

import "fmt"

func main() {
	// var team [3]string
	// team[0] = "hammer"
	// team[1] = "soldier"
	// team[2] = "mum"
	var team = [...]string{"hammer", "soldier", "mum"}
	for k, v := range team {
		fmt.Printf("k %d v %s\n", k, v)
	}

}
