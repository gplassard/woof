package powerpack

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreatePowerpackCmd = &cobra.Command{
	Use:   "createpowerpack",
	Short: "Create a new powerpack",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/powerpacks")
		fmt.Println("OperationID: CreatePowerpack")
	},
}

func init() {
	Cmd.AddCommand(CreatePowerpackCmd)
}
