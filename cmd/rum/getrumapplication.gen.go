package rum

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetRUMApplicationCmd = &cobra.Command{
	Use:   "getrumapplication",
	Short: "Get a RUM application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/rum/applications/{id}")
		fmt.Println("OperationID: GetRUMApplication")
	},
}

func init() {
	Cmd.AddCommand(GetRUMApplicationCmd)
}
