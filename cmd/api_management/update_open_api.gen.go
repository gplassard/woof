package api_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateOpenAPICmd = &cobra.Command{
	Use: "update-open-api [id] [payload]",

	Short: "Update an API",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UpdateOpenAPIResponse
		var err error

		var body datadogV2.UpdateOpenAPIOptionalParameters
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewAPIManagementApi(client.NewAPIClient())
		res, _, err = api.UpdateOpenAPI(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), body)
		cmdutil.HandleError(err, "failed to update-open-api")

		cmd.Println(cmdutil.FormatJSON(res, "api_management"))
	},
}

func init() {
	Cmd.AddCommand(UpdateOpenAPICmd)
}
