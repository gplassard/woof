package domain_allowlist

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetDomainAllowlistCmd = &cobra.Command{
	Use:     "get-domain-allowlist",
	Aliases: []string{"get"},
	Short:   "Get Domain Allowlist",
	Long: `Get Domain Allowlist
Documentation: https://docs.datadoghq.com/api/latest/domain-allowlist/#get-domain-allowlist`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DomainAllowlistResponse
		var err error

		api := datadogV2.NewDomainAllowlistApi(client.NewAPIClient())
		res, _, err = api.GetDomainAllowlist(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-domain-allowlist")

		cmd.Println(cmdutil.FormatJSON(res, "domain_allowlist"))
	},
}

func init() {
	Cmd.AddCommand(GetDomainAllowlistCmd)
}
