package rum

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateRUMApplicationCmd = &cobra.Command{
	Use:   "updaterumapplication [id]",
	Short: "Update a RUM application",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewRUMApi(client.NewAPIClient())
		res, _, err := api.UpdateRUMApplication(client.NewContext(apiKey, appKey, site), args[0], datadogV2.RUMApplicationUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to updaterumapplication: %v", err)
		}

		cmdutil.PrintJSON(res, "rum")
	},
}

func init() {
	Cmd.AddCommand(UpdateRUMApplicationCmd)
}
