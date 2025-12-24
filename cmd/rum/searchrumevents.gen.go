package rum

import (
	"fmt"
	"github.com/spf13/cobra"
)

var SearchRUMEventsCmd = &cobra.Command{
	Use:   "searchrumevents",
	Short: "Search RUM events",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/rum/events/search")
		fmt.Println("OperationID: SearchRUMEvents")
	},
}

func init() {
	Cmd.AddCommand(SearchRUMEventsCmd)
}
