package data_deletion

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateDataDeletionRequestCmd = &cobra.Command{
	Use:     "create-data-deletion-request [product] [payload]",
	Aliases: []string{"create-request"},
	Short:   "Creates a data deletion request",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.CreateDataDeletionRequestBody
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewDataDeletionApi(client.NewAPIClient())
		res, _, err := api.CreateDataDeletionRequest(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-data-deletion-request")

		cmd.Println(cmdutil.FormatJSON(res, "data_deletion"))
	},
}

func init() {
	Cmd.AddCommand(CreateDataDeletionRequestCmd)
}
