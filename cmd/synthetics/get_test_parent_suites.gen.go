package synthetics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetTestParentSuitesCmd = &cobra.Command{
	Use: "get-test-parent-suites [public_id]",

	Short: "Get parent suites for a test",
	Long: `Get parent suites for a test
Documentation: https://docs.datadoghq.com/api/latest/synthetics/#get-test-parent-suites`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SyntheticsTestParentSuitesResponse
		var err error

		api := datadogV2.NewSyntheticsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetTestParentSuites(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-test-parent-suites")

		fmt.Println(cmdutil.FormatJSON(res, "parent_suite"))
	},
}

func init() {

	Cmd.AddCommand(GetTestParentSuitesCmd)
}
