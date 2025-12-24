package powerpack

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdatePowerpackCmd = &cobra.Command{
	Use:   "updatepowerpack",
	Short: "Update a powerpack",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/powerpacks/{powerpack_id}")
		fmt.Println("OperationID: UpdatePowerpack")
	},
}

func init() {
	Cmd.AddCommand(UpdatePowerpackCmd)
}
