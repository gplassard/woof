package metrics

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteTagConfigurationCmd = &cobra.Command{
	Use:   "delete-tag-configuration [metric_name]",
	
	Short: "Delete a tag configuration",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		_, err := api.DeleteTagConfiguration(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-tag-configuration")

		
	},
}

func init() {
	Cmd.AddCommand(DeleteTagConfigurationCmd)
}
