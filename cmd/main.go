package main

import (
	"fmt"
	"log"
	"os"

	"glox/scanner"
)

func run(source string) error {
	s := scanner.NewScanner(source)
	for _, t := range s.ScanTokens() {
		fmt.Printf("%s\n", t.String())
	}
	return nil
}

func main() {
	// if len(os.Args) < 2 {
	// 	log.Fatalf("glox\nUsage: %s <script>", os.Args[0])
	// }
	//
	// content, err := ioutil.ReadFile(os.Args[1])
	// if err != nil {
	// 	log.Fatal(err)
	// }

	content := `
// this is a comment
(( )){} // grouping stuff
!*+-/=<> <= == // operators
"foobar"`

	err := run(string(content))
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
