package container_images

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListContainerImagesCmd = &cobra.Command{
	Use:   "list-container-images",
	Aliases: []string{ "list", },
	Short: "Get all Container Images",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewContainerImagesApi(client.NewAPIClient())
		res, _, err := api.ListContainerImages(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-container-images: %v", err)
		}

		cmdutil.PrintJSON(res, "container_images")
	},
}

func init() {
	Cmd.AddCommand(ListContainerImagesCmd)
}
