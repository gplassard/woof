package security_monitoring

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateCustomFrameworkCmd = &cobra.Command{
	Use:   "create-custom-framework",
	
	Short: "Create a custom framework",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.CreateCustomFramework(client.NewContext(apiKey, appKey, site), datadogV2.CreateCustomFrameworkRequest{})
		cmdutil.HandleError(err, "failed to create-custom-framework")

		cmdutil.PrintJSON(res, "custom_framework")
	},
}

func init() {
	Cmd.AddCommand(CreateCustomFrameworkCmd)
}
