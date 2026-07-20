package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/akshatmanohar21/devbitkit/internal/generators"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: dbk <domain> <capability>")
		os.Exit(1)
	}

	domain := os.Args[1]
	capability := os.Args[2]
	remainingArgs := os.Args[3:]
	
	switch domain{
	case "generate":
		handleGenerate(capability, remainingArgs)
	default:
		fmt.Printf("unknown domain: %s\n", domain)
		os.Exit(1)
	}

}

func handleGenerate(capability string, args []string) {
	switch capability{
	case "password":
		fs := flag.NewFlagSet("length", flag.ExitOnError)
		length := fs.Int("length", 16, "password length")
		noLetters := fs.Bool("no-letters", false, "exclude letter characters")
		noSymbols := fs.Bool("no-symbols", false, "exclude symbol characters")
		noNumbers := fs.Bool("no-numbers", false, "exclude number characters")
		count := fs.Int("count", 1, "number of passwords to be generated")
		fs.Parse(args)

		if len(fs.Args()) > 0 {
			fmt.Printf("unexpected argument passed: %v\n", fs.Args())			
			os.Exit(1)
		}

		if *count <= 0 {
			fmt.Printf("error: count must be positive, got %d\n", *count)
			os.Exit(1)
		}

		for i := 0; i<*count; i++ {
			password, err := generators.GeneratePassword(
			*length,
			*noLetters,
			*noNumbers,
			*noSymbols,
			)				
			if err != nil {
				fmt.Printf("error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(password)
		}

	default:
		fmt.Printf("unknown capability: %s\n", capability)
		os.Exit(1)
	}
}