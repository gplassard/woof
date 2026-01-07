package application_security

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateApplicationSecurityWafExclusionFilterCmd = &cobra.Command{
	Use:     "create-application-security-waf-exclusion-filter",
	Aliases: []string{"create-waf-exclusion-filter"},
	Short:   "Create a WAF exclusion filter",
	Long: `Create a WAF exclusion filter
Documentation: https://docs.datadoghq.com/api/latest/application-security/#create-application-security-waf-exclusion-filter`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ApplicationSecurityWafExclusionFilterResponse
		var err error

		var body datadogV2.ApplicationSecurityWafExclusionFilterCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewApplicationSecurityApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateApplicationSecurityWafExclusionFilter(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-application-security-waf-exclusion-filter")

		fmt.Println(cmdutil.FormatJSON(res, "exclusion_filter"))
	},
}

func init() {

	CreateApplicationSecurityWafExclusionFilterCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateApplicationSecurityWafExclusionFilterCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateApplicationSecurityWafExclusionFilterCmd)
}
