package powerpack

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeletePowerpackCmd = &cobra.Command{
	Use:   "deletepowerpack [powerpack_id]",
	Short: "Delete a powerpack",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewPowerpackApi(client.NewAPIClient())
		_, err := api.DeletePowerpack(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to deletepowerpack: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeletePowerpackCmd)
}
