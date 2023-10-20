package main

import (
	"fmt"
	"os"
)

func main() {
	plugins, err := loadPlugins()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = generate(plugins); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
