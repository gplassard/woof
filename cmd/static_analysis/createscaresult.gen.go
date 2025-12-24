package static_analysis

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateSCAResultCmd = &cobra.Command{
	Use:   "createscaresult",
	Short: "Post dependencies for analysis",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/static-analysis-sca/dependencies")
		fmt.Println("OperationID: CreateSCAResult")
	},
}

func init() {
	Cmd.AddCommand(CreateSCAResultCmd)
}
