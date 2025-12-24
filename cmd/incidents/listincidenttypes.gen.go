package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListIncidentTypesCmd = &cobra.Command{
	Use:   "listincidenttypes",
	Short: "Get a list of incident types",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/incidents/config/types")
		fmt.Println("OperationID: ListIncidentTypes")
	},
}

func init() {
	Cmd.AddCommand(ListIncidentTypesCmd)
}
