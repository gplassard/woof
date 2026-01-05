package api_management

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var UpdateOpenAPICmd = &cobra.Command{
	Use: "update-open-api [id]",

	Short: "Update an API",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAPIManagementApi(client.NewAPIClient())
		res, _, err := api.UpdateOpenAPI(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), *datadogV2.NewUpdateOpenAPIOptionalParameters())
		cmdutil.HandleError(err, "failed to update-open-api")

		cmd.Println(cmdutil.FormatJSON(res, "api_management"))
	},
}

func init() {
	Cmd.AddCommand(UpdateOpenAPICmd)
}
