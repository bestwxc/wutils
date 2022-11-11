package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println([]byte("\377\330\377"))
	abs := "/srv/app/cc"
	a := "/srv/app/dd/ff"
	fmt.Println(filepath.Rel(abs, a))
	fmt.Println(filepath.Rel(a, abs))
	fmt.Println(filepath.Rel(a, a))

}
