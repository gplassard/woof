package domain_allowlist

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var PatchDomainAllowlistCmd = &cobra.Command{
	Use:   "patch_domain_allowlist",
	Short: "Sets Domain Allowlist",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDomainAllowlistApi(client.NewAPIClient())
		res, _, err := api.PatchDomainAllowlist(client.NewContext(apiKey, appKey, site), datadogV2.DomainAllowlistRequest{})
		if err != nil {
			log.Fatalf("failed to patch_domain_allowlist: %v", err)
		}

		cmdutil.PrintJSON(res, "domain_allowlist")
	},
}

func init() {
	Cmd.AddCommand(PatchDomainAllowlistCmd)
}
