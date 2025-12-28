package aws_integration

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListAWSNamespacesCmd = &cobra.Command{
	Use:   "list-aws-namespaces",
	
	Short: "List available namespaces",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		res, _, err := api.ListAWSNamespaces(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-aws-namespaces: %v", err)
		}

		cmdutil.PrintJSON(res, "namespaces")
	},
}

func init() {
	Cmd.AddCommand(ListAWSNamespacesCmd)
}
