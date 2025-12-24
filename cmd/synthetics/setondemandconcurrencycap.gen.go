package synthetics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var SetOnDemandConcurrencyCapCmd = &cobra.Command{
	Use:   "setondemandconcurrencycap",
	Short: "Save new value for on-demand concurrency cap",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/synthetics/settings/on_demand_concurrency_cap")
		fmt.Println("OperationID: SetOnDemandConcurrencyCap")
	},
}

func init() {
	Cmd.AddCommand(SetOnDemandConcurrencyCapCmd)
}
