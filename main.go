package main

import (
	"fmt" 

	"github.com/spf13/cobra"

	tool "github.com/lfranchini31/fake-go/tools"
)

func main() {
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, Cobra!")

			tool.PrintHello()
		},
	}

	fmt.Println("Calling cmd.Execute()!")
	cmd.Execute()
}
