package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateCustomFrameworkCmd = &cobra.Command{
	Use:   "update-custom-framework [handle] [version]",
	
	Short: "Update a custom framework",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.UpdateCustomFramework(client.NewContext(apiKey, appKey, site), args[0], args[1], datadogV2.UpdateCustomFrameworkRequest{})
		if err != nil {
			log.Fatalf("failed to update-custom-framework: %v", err)
		}

		cmdutil.PrintJSON(res, "custom_framework")
	},
}

func init() {
	Cmd.AddCommand(UpdateCustomFrameworkCmd)
}
