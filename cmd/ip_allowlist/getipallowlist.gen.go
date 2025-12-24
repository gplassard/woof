package ip_allowlist

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetIPAllowlistCmd = &cobra.Command{
	Use:   "getipallowlist",
	Short: "Get IP Allowlist",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewIPAllowlistApi(client.NewAPIClient())
		res, _, err := api.GetIPAllowlist(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to getipallowlist: %v", err)
		}

		cmdutil.PrintJSON(res, "ip_allowlist")
	},
}

func init() {
	Cmd.AddCommand(GetIPAllowlistCmd)
}
