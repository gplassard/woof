package logs_restriction_queries

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateRestrictionQueryCmd = &cobra.Command{
	Use: "create-restriction-query [payload]",

	Short: "Create a restriction query",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RestrictionQueryWithoutRelationshipsResponse
		var err error

		var body datadogV2.RestrictionQueryCreatePayload
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewLogsRestrictionQueriesApi(client.NewAPIClient())
		res, _, err = api.CreateRestrictionQuery(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-restriction-query")

		cmd.Println(cmdutil.FormatJSON(res, "logs_restriction_queries"))
	},
}

func init() {
	Cmd.AddCommand(CreateRestrictionQueryCmd)
}
