package synthetics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetOnDemandConcurrencyCapCmd = &cobra.Command{
	Use:   "getondemandconcurrencycap",
	Short: "Get the on-demand concurrency cap",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/synthetics/settings/on_demand_concurrency_cap")
		fmt.Println("OperationID: GetOnDemandConcurrencyCap")
	},
}

func init() {
	Cmd.AddCommand(GetOnDemandConcurrencyCapCmd)
}
