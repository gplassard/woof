package application_security

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteApplicationSecurityWafExclusionFilterCmd = &cobra.Command{
	Use:     "delete-application-security-waf-exclusion-filter [exclusion_filter_id]",
	Aliases: []string{"delete-waf-exclusion-filter"},
	Short:   "Delete a WAF exclusion filter",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewApplicationSecurityApi(client.NewAPIClient())
		_, err := api.DeleteApplicationSecurityWafExclusionFilter(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-application-security-waf-exclusion-filter")

	},
}

func init() {
	Cmd.AddCommand(DeleteApplicationSecurityWafExclusionFilterCmd)
}
