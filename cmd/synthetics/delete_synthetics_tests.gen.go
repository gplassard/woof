package synthetics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteSyntheticsTestsCmd = &cobra.Command{
	Use:     "delete-synthetics-tests",
	Aliases: []string{"delete-tests"},
	Short:   "Delete synthetics tests",
	Long: `Delete synthetics tests
Documentation: https://docs.datadoghq.com/api/latest/synthetics/#delete-synthetics-tests`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DeletedTestsResponse
		var err error

		var body datadogV2.DeletedTestsRequestDeleteRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSyntheticsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.DeleteSyntheticsTests(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to delete-synthetics-tests")

		fmt.Println(cmdutil.FormatJSON(res, "delete_tests"))
	},
}

func init() {

	DeleteSyntheticsTestsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	DeleteSyntheticsTestsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(DeleteSyntheticsTestsCmd)
}
