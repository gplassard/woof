package powerpack

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeletePowerpackCmd = &cobra.Command{
	Use:   "delete-powerpack [powerpack_id]",
	Aliases: []string{ "delete", },
	Short: "Delete a powerpack",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewPowerpackApi(client.NewAPIClient())
		_, err := api.DeletePowerpack(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-powerpack")

		
	},
}

func init() {
	Cmd.AddCommand(DeletePowerpackCmd)
}
