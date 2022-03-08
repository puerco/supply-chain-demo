package main

import (
	"fmt"
	"os"

	cowsay "github.com/Code-Hex/Neo-cowsay/v2"
)

func main() {
	os.Setenv("COWPATH", os.Getenv("KO_DATA_PATH"))

	say, err := cowsay.Say(
		"Hello Cloud Native @ Scale",
		cowsay.Type("default"),
		cowsay.BallonWidth(40),
		cowsay.Type("octo"),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(say)
}
