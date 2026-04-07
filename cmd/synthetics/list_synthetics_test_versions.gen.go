package synthetics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListSyntheticsTestVersionsCmd = &cobra.Command{
	Use:     "list-synthetics-test-versions [public_id]",
	Aliases: []string{"list-test-versions"},
	Short:   "Get version history of a test",
	Long: `Get version history of a test
Documentation: https://docs.datadoghq.com/api/latest/synthetics/#list-synthetics-test-versions`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SyntheticsTestVersionHistoryResponse
		var err error

		api := datadogV2.NewSyntheticsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListSyntheticsTestVersions(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-synthetics-test-versions")

		fmt.Println(cmdutil.FormatJSON(res, "version_metadata"))
	},
}

func init() {

	Cmd.AddCommand(ListSyntheticsTestVersionsCmd)
}
