package synthetics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateSyntheticsSuiteCmd = &cobra.Command{
	Use:     "create-synthetics-suite",
	Aliases: []string{"create-suite"},
	Short:   "Create a test suite",
	Long: `Create a test suite
Documentation: https://docs.datadoghq.com/api/latest/synthetics/#create-synthetics-suite`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SyntheticsSuiteResponse
		var err error

		var body datadogV2.SuiteCreateEditRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSyntheticsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateSyntheticsSuite(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-synthetics-suite")

		fmt.Println(cmdutil.FormatJSON(res, "suites"))
	},
}

func init() {

	CreateSyntheticsSuiteCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateSyntheticsSuiteCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateSyntheticsSuiteCmd)
}
