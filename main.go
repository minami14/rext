package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/minami14/rext/ext"
)

const usage = `
Usage
rext [pattern]

e.g.
rext Found.000/*.chk
`

func main() {
	if len(os.Args) != 2 {
		log.Fatal(usage)
	}

	filenames, err := filepath.Glob(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	for _, name := range filenames {
		if !isFile(name) {
			continue
		}

		extensions, err := ext.ExtensionFromFile(name)
		if err != nil {
			fmt.Printf("skip: %v\n", name)
			continue
		}

		extension := extensions[0]
		if extension == filepath.Ext(name) {
			fmt.Printf("skip: %v\n", name)
			continue
		}

		newName := name + extension
		if err := os.Rename(name, newName); err != nil {
			log.Println(err)
			continue
		}

		fmt.Printf("changed: %v\n", newName)
	}
}

func isFile(path string) bool {
	stat, err := os.Stat(path)
	return err == nil && !stat.IsDir()
}
