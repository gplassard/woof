package powerpack

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetPowerpackCmd = &cobra.Command{
	Use:   "getpowerpack",
	Short: "Get a Powerpack",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/powerpacks/{powerpack_id}")
		fmt.Println("OperationID: GetPowerpack")
	},
}

func init() {
	Cmd.AddCommand(GetPowerpackCmd)
}
