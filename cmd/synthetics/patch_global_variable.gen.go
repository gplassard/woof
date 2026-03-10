package synthetics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var PatchGlobalVariableCmd = &cobra.Command{
	Use: "patch-global-variable [variable_id]",

	Short: "Patch a global variable",
	Long: `Patch a global variable
Documentation: https://docs.datadoghq.com/api/latest/synthetics/#patch-global-variable`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GlobalVariableResponse
		var err error

		var body datadogV2.GlobalVariableJsonPatchRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSyntheticsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.PatchGlobalVariable(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to patch-global-variable")

		fmt.Println(cmdutil.FormatJSON(res, "global_variables"))
	},
}

func init() {

	PatchGlobalVariableCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	PatchGlobalVariableCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(PatchGlobalVariableCmd)
}
