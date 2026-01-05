package security_monitoring

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateCasesCmd = &cobra.Command{
	Use: "create-cases",

	Short: "Create cases for security findings",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.CreateCases(client.NewContext(apiKey, appKey, site), datadogV2.CreateCaseRequestArray{})
		cmdutil.HandleError(err, "failed to create-cases")

		cmd.Println(cmdutil.FormatJSON(res, "cases"))
	},
}

func init() {
	Cmd.AddCommand(CreateCasesCmd)
}
