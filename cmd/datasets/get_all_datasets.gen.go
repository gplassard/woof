package datasets

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetAllDatasetsCmd = &cobra.Command{
	Use:     "get-all-datasets",
	Aliases: []string{"get-all"},
	Short:   "Get all datasets",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDatasetsApi(client.NewAPIClient())
		res, _, err := api.GetAllDatasets(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-all-datasets")

		cmd.Println(cmdutil.FormatJSON(res, "dataset"))
	},
}

func init() {
	Cmd.AddCommand(GetAllDatasetsCmd)
}
