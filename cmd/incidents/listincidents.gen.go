package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListIncidentsCmd = &cobra.Command{
	Use:   "listincidents",
	Short: "Get a list of incidents",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/incidents")
		fmt.Println("OperationID: ListIncidents")
	},
}

func init() {
	Cmd.AddCommand(ListIncidentsCmd)
}
