package synthetics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var EditSyntheticsSuiteCmd = &cobra.Command{
	Use:     "edit-synthetics-suite [public_id]",
	Aliases: []string{"edit-suite"},
	Short:   "Edit a test suite",
	Long: `Edit a test suite
Documentation: https://docs.datadoghq.com/api/latest/synthetics/#edit-synthetics-suite`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SyntheticsSuiteResponse
		var err error

		var body datadogV2.SuiteCreateEditRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSyntheticsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.EditSyntheticsSuite(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to edit-synthetics-suite")

		fmt.Println(cmdutil.FormatJSON(res, "suites"))
	},
}

func init() {

	EditSyntheticsSuiteCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	EditSyntheticsSuiteCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(EditSyntheticsSuiteCmd)
}
