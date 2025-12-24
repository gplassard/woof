package organizations

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UploadIdPMetadataCmd = &cobra.Command{
	Use:   "uploadidpmetadata",
	Short: "Upload IdP metadata",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewOrganizationsApi(client.NewAPIClient())
		_, err := api.UploadIdPMetadata(client.NewContext(apiKey, appKey, site), *datadogV2.NewUploadIdPMetadataOptionalParameters())
		if err != nil {
			log.Fatalf("failed to uploadidpmetadata: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(UploadIdPMetadataCmd)
}
