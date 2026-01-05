package container_images

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListContainerImagesCmd = &cobra.Command{
	Use:     "list-container-images",
	Aliases: []string{"list"},
	Short:   "Get all Container Images",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewContainerImagesApi(client.NewAPIClient())
		res, _, err := api.ListContainerImages(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-container-images")

		cmd.Println(cmdutil.FormatJSON(res, "container_images"))
	},
}

func init() {
	Cmd.AddCommand(ListContainerImagesCmd)
}
