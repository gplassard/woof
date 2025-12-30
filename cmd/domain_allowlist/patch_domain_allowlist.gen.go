package domain_allowlist

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var PatchDomainAllowlistCmd = &cobra.Command{
	Use:     "patch-domain-allowlist",
	Aliases: []string{"patch"},
	Short:   "Sets Domain Allowlist",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDomainAllowlistApi(client.NewAPIClient())
		res, _, err := api.PatchDomainAllowlist(client.NewContext(apiKey, appKey, site), datadogV2.DomainAllowlistRequest{})
		cmdutil.HandleError(err, "failed to patch-domain-allowlist")

		cmd.Println(cmdutil.FormatJSON(res, "domain_allowlist"))
	},
}

func init() {
	Cmd.AddCommand(PatchDomainAllowlistCmd)
}
