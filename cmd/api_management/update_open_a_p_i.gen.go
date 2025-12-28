package api_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
	
	
	"github.com/spf13/cobra"
	
)

var UpdateOpenAPICmd = &cobra.Command{
	Use:   "update-open-a-p-i [id]",
	Short: "Update an API",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAPIManagementApi(client.NewAPIClient())
		res, _, err := api.UpdateOpenAPI(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), *datadogV2.NewUpdateOpenAPIOptionalParameters())
		if err != nil {
			log.Fatalf("failed to update-open-a-p-i: %v", err)
		}

		cmdutil.PrintJSON(res, "api_management")
	},
}

func init() {
	Cmd.AddCommand(UpdateOpenAPICmd)
}
