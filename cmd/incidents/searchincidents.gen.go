package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var SearchIncidentsCmd = &cobra.Command{
	Use:   "searchincidents",
	Short: "Search for incidents",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/incidents/search")
		fmt.Println("OperationID: SearchIncidents")
	},
}

func init() {
	Cmd.AddCommand(SearchIncidentsCmd)
}
