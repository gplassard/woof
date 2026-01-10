package api_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var UpdateOpenAPICmd = &cobra.Command{
	Use: "update-open-api [id]",

	Short: "Update an API",
	Long: `Update an API
Documentation: https://docs.datadoghq.com/api/latest/api-management/#update-open-api`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UpdateOpenAPIResponse
		var err error

		optionalParams := datadogV2.NewUpdateOpenAPIOptionalParameters()

		if cmd.Flags().Changed("payload") || cmd.Flags().Changed("payload-file") {
			err = cmdutil.UnmarshalPayload(cmd, optionalParams)
			cmdutil.HandleError(err, "failed to read payload")
		}

		api := datadogV2.NewAPIManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateOpenAPI(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), *optionalParams)
		cmdutil.HandleError(err, "failed to update-open-api")

		fmt.Println(cmdutil.FormatJSON(res, "api_management"))
	},
}

func init() {

	UpdateOpenAPICmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateOpenAPICmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateOpenAPICmd)
}
