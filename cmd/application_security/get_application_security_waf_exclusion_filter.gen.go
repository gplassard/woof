package application_security

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetApplicationSecurityWafExclusionFilterCmd = &cobra.Command{
	Use:     "get-application-security-waf-exclusion-filter [exclusion_filter_id]",
	Aliases: []string{"get-waf-exclusion-filter"},
	Short:   "Get a WAF exclusion filter",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ApplicationSecurityWafExclusionFilterResponse
		var err error

		api := datadogV2.NewApplicationSecurityApi(client.NewAPIClient())
		res, _, err = api.GetApplicationSecurityWafExclusionFilter(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-application-security-waf-exclusion-filter")

		cmd.Println(cmdutil.FormatJSON(res, "exclusion_filter"))
	},
}

func init() {
	Cmd.AddCommand(GetApplicationSecurityWafExclusionFilterCmd)
}
