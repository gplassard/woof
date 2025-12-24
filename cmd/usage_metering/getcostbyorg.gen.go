package usage_metering

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	"time"
	
	"github.com/spf13/cobra"
	
)

var GetCostByOrgCmd = &cobra.Command{
	Use:   "getcostbyorg [start_month]",
	Short: "Get cost across multi-org account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewUsageMeteringApi(client.NewAPIClient())
		res, _, err := api.GetCostByOrg(client.NewContext(apiKey, appKey, site), func() time.Time { t, _ := time.Parse(time.RFC3339, args[0]); return t }())
		if err != nil {
			log.Fatalf("failed to getcostbyorg: %v", err)
		}

		cmdutil.PrintJSON(res, "usage_metering")
	},
}

func init() {
	Cmd.AddCommand(GetCostByOrgCmd)
}
