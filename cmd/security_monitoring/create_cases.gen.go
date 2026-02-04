package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateCasesCmd = &cobra.Command{
	Use: "create-cases",

	Short: "Create cases for security findings",
	Long: `Create cases for security findings
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#create-cases`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FindingCaseResponseArray
		var err error

		var body datadogV2.CreateCaseRequestArray
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateCases(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-cases")

		fmt.Println(cmdutil.FormatJSON(res, "case"))
	},
}

func init() {

	CreateCasesCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateCasesCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateCasesCmd)
}
