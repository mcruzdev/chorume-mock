package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)


var (
	openApiFile string
)
func main() {
	if err := CreateRootCmd().Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func CreateRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "chorume-mock",
		Short: "Generate a Wiremock definition from OpenAPI specification",
	}

	rootCmd.AddCommand(CreateGenerateCmd())

	return rootCmd
}

func CreateGenerateCmd() *cobra.Command {

	cmd := cobra.Command{
		Use:     "generate",
		Aliases: []string{"s"},
		Short: "Generate the Wiremock definition from OpenAPI specification file",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("[generate]")
			fmt.Println(fmt.Sprintf("file: %s", openApiFile))
		},
	}

	cmd.PersistentFlags().StringVar(&openApiFile, "oapi", "", `The OpenAPI specification file.
Usage: chorume-mock generate --oapi=my-openapi-spec.yaml`)

	return &cmd
}
