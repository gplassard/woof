package ip_allowlist

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateIPAllowlistCmd = &cobra.Command{
	Use:   "update-i-p-allowlist",
	Short: "Update IP Allowlist",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIPAllowlistApi(client.NewAPIClient())
		res, _, err := api.UpdateIPAllowlist(client.NewContext(apiKey, appKey, site), datadogV2.IPAllowlistUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to update-i-p-allowlist: %v", err)
		}

		cmdutil.PrintJSON(res, "ip_allowlist")
	},
}

func init() {
	Cmd.AddCommand(UpdateIPAllowlistCmd)
}
