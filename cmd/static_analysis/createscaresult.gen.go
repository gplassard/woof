package static_analysis

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateSCAResultCmd = &cobra.Command{
	Use:   "createscaresult",
	Short: "Post dependencies for analysis",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewStaticAnalysisApi(client.NewAPIClient())
		_, err := api.CreateSCAResult(client.NewContext(apiKey, appKey, site), datadogV2.ScaRequest{})
		if err != nil {
			log.Fatalf("failed to createscaresult: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(CreateSCAResultCmd)
}
