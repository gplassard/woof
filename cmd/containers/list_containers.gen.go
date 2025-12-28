package containers

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListContainersCmd = &cobra.Command{
	Use:   "list-containers",
	Aliases: []string{ "list", },
	Short: "Get All Containers",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewContainersApi(client.NewAPIClient())
		res, _, err := api.ListContainers(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-containers")

		cmdutil.PrintJSON(res, "containers")
	},
}

func init() {
	Cmd.AddCommand(ListContainersCmd)
}
