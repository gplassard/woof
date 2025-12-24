package powerpack

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListPowerpacksCmd = &cobra.Command{
	Use:   "listpowerpacks",
	Short: "Get all powerpacks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/powerpacks")
		fmt.Println("OperationID: ListPowerpacks")
	},
}

func init() {
	Cmd.AddCommand(ListPowerpacksCmd)
}
