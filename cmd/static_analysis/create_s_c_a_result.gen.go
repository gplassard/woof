package static_analysis

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateSCAResultCmd = &cobra.Command{
	Use:   "create-s-c-a-result",
	
	Short: "Post dependencies for analysis",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewStaticAnalysisApi(client.NewAPIClient())
		_, err := api.CreateSCAResult(client.NewContext(apiKey, appKey, site), datadogV2.ScaRequest{})
		if err != nil {
			log.Fatalf("failed to create-s-c-a-result: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(CreateSCAResultCmd)
}
