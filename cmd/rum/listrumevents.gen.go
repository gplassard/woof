package rum

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListRUMEventsCmd = &cobra.Command{
	Use:   "listrumevents",
	Short: "Get a list of RUM events",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/rum/events")
		fmt.Println("OperationID: ListRUMEvents")
	},
}

func init() {
	Cmd.AddCommand(ListRUMEventsCmd)
}
