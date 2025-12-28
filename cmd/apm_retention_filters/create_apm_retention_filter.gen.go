package apm_retention_filters

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateApmRetentionFilterCmd = &cobra.Command{
	Use:   "create_apm_retention_filter",
	Short: "Create a retention filter",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAPMRetentionFiltersApi(client.NewAPIClient())
		res, _, err := api.CreateApmRetentionFilter(client.NewContext(apiKey, appKey, site), datadogV2.RetentionFilterCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create_apm_retention_filter: %v", err)
		}

		cmdutil.PrintJSON(res, "apm_retention_filter")
	},
}

func init() {
	Cmd.AddCommand(CreateApmRetentionFilterCmd)
}
