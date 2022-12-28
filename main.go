package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Provide input")
		return
	}
	emptyFinder(args[0])
}

func emptyFinder(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	var total int

	for _, file := range files {
		info, _ := file.Info()
		if info.Size() == 0 {
			total += len(info.Name()) + 1
		}
	}

	names := make([]byte, 0, total)

	for _, file := range files {
		info, _ := file.Info()
		if info.Size() == 0 {
			name := info.Name()
			names = append(names, name...)
			names = append(names, "\n"...)
		}
	}
	err = os.WriteFile(".gitignore", names, 744)
	if err != nil {
		log.Fatal(err)
		return
	}
}
