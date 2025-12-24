package events

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListEventsCmd = &cobra.Command{
	Use:   "listevents",
	Short: "Get a list of events",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/events")
		fmt.Println("OperationID: ListEvents")
	},
}

func init() {
	Cmd.AddCommand(ListEventsCmd)
}
