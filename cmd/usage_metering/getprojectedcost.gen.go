package usage_metering

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetProjectedCostCmd = &cobra.Command{
	Use:   "getprojectedcost",
	Short: "Get projected cost across your account",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewUsageMeteringApi(client.NewAPIClient())
		res, _, err := api.GetProjectedCost(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to getprojectedcost: %v", err)
		}

		cmdutil.PrintJSON(res, "usage_metering")
	},
}

func init() {
	Cmd.AddCommand(GetProjectedCostCmd)
}
