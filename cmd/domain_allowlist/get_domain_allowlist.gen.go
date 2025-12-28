package domain_allowlist

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetDomainAllowlistCmd = &cobra.Command{
	Use:   "get-domain-allowlist",
	Short: "Get Domain Allowlist",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDomainAllowlistApi(client.NewAPIClient())
		res, _, err := api.GetDomainAllowlist(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to get-domain-allowlist: %v", err)
		}

		cmdutil.PrintJSON(res, "domain_allowlist")
	},
}

func init() {
	Cmd.AddCommand(GetDomainAllowlistCmd)
}
