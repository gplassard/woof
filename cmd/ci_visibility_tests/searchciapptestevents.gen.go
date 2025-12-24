package ci_visibility_tests

import (
	"fmt"
	"github.com/spf13/cobra"
)

var SearchCIAppTestEventsCmd = &cobra.Command{
	Use:   "searchciapptestevents",
	Short: "Search tests events",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/ci/tests/events/search")
		fmt.Println("OperationID: SearchCIAppTestEvents")
	},
}

func init() {
	Cmd.AddCommand(SearchCIAppTestEventsCmd)
}
