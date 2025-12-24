package powerpack

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeletePowerpackCmd = &cobra.Command{
	Use:   "deletepowerpack",
	Short: "Delete a powerpack",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/powerpacks/{powerpack_id}")
		fmt.Println("OperationID: DeletePowerpack")
	},
}

func init() {
	Cmd.AddCommand(DeletePowerpackCmd)
}
