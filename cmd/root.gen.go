package cmd

import (
	"github.com/gplassard/ouaf/cmd/action_connection"
	"github.com/gplassard/ouaf/cmd/actions_datastores"
	"github.com/gplassard/ouaf/cmd/agentless_scanning"
	"github.com/gplassard/ouaf/cmd/api_management"
	"github.com/gplassard/ouaf/cmd/apm"
	"github.com/gplassard/ouaf/cmd/apm_retention_filters"
	"github.com/gplassard/ouaf/cmd/app_builder"
	"github.com/gplassard/ouaf/cmd/application_security"
	"github.com/gplassard/ouaf/cmd/audit"
	"github.com/gplassard/ouaf/cmd/authn_mappings"
	"github.com/gplassard/ouaf/cmd/aws_integration"
	"github.com/gplassard/ouaf/cmd/aws_logs_integration"
	"github.com/gplassard/ouaf/cmd/case_management"
	"github.com/gplassard/ouaf/cmd/case_management_attribute"
	"github.com/gplassard/ouaf/cmd/case_management_type"
	"github.com/gplassard/ouaf/cmd/ci_visibility_pipelines"
	"github.com/gplassard/ouaf/cmd/ci_visibility_tests"
	"github.com/gplassard/ouaf/cmd/cloud_cost_management"
	"github.com/gplassard/ouaf/cmd/cloud_network_monitoring"
	"github.com/gplassard/ouaf/cmd/cloudflare_integration"
	"github.com/gplassard/ouaf/cmd/confluent_cloud"
	"github.com/gplassard/ouaf/cmd/container_images"
	"github.com/gplassard/ouaf/cmd/containers"
	"github.com/gplassard/ouaf/cmd/csm_agents"
	"github.com/gplassard/ouaf/cmd/csm_coverage_analysis"
	"github.com/gplassard/ouaf/cmd/csm_threats"
	"github.com/gplassard/ouaf/cmd/dashboard_lists"
	"github.com/gplassard/ouaf/cmd/data_deletion"
	"github.com/gplassard/ouaf/cmd/datasets"
	"github.com/gplassard/ouaf/cmd/deployment_gates"
	"github.com/gplassard/ouaf/cmd/domain_allowlist"
	"github.com/gplassard/ouaf/cmd/dora_metrics"
	"github.com/gplassard/ouaf/cmd/downtimes"
	"github.com/gplassard/ouaf/cmd/error_tracking"
	"github.com/gplassard/ouaf/cmd/events"
	"github.com/gplassard/ouaf/cmd/fastly_integration"
	"github.com/gplassard/ouaf/cmd/fleet_automation"
	"github.com/gplassard/ouaf/cmd/gcp_integration"
	"github.com/gplassard/ouaf/cmd/incident_services"
	"github.com/gplassard/ouaf/cmd/incident_teams"
	"github.com/gplassard/ouaf/cmd/incidents"
	"github.com/gplassard/ouaf/cmd/ip_allowlist"
	"github.com/gplassard/ouaf/cmd/key_management"
	"github.com/gplassard/ouaf/cmd/logs"
	"github.com/gplassard/ouaf/cmd/logs_archives"
	"github.com/gplassard/ouaf/cmd/logs_custom_destinations"
	"github.com/gplassard/ouaf/cmd/logs_metrics"
	"github.com/gplassard/ouaf/cmd/logs_restriction_queries"
	"github.com/gplassard/ouaf/cmd/metrics"
	"github.com/gplassard/ouaf/cmd/microsoft_teams_integration"
	"github.com/gplassard/ouaf/cmd/monitors"
	"github.com/gplassard/ouaf/cmd/network_device_monitoring"
	"github.com/gplassard/ouaf/cmd/observability_pipelines"
	"github.com/gplassard/ouaf/cmd/okta_integration"
	"github.com/gplassard/ouaf/cmd/on_call"
	"github.com/gplassard/ouaf/cmd/on_call_paging"
	"github.com/gplassard/ouaf/cmd/opsgenie_integration"
	"github.com/gplassard/ouaf/cmd/org_connections"
	"github.com/gplassard/ouaf/cmd/organizations"
	"github.com/gplassard/ouaf/cmd/powerpack"
	"github.com/gplassard/ouaf/cmd/processes"
	"github.com/gplassard/ouaf/cmd/reference_tables"
	"github.com/gplassard/ouaf/cmd/restriction_policies"
	"github.com/gplassard/ouaf/cmd/roles"
	"github.com/gplassard/ouaf/cmd/rum"
	"github.com/gplassard/ouaf/cmd/rum_audience_management"
	"github.com/gplassard/ouaf/cmd/rum_metrics"
	"github.com/gplassard/ouaf/cmd/rum_retention_filters"
	"github.com/gplassard/ouaf/cmd/security_monitoring"
	"github.com/gplassard/ouaf/cmd/sensitive_data_scanner"
	"github.com/gplassard/ouaf/cmd/service_accounts"
	"github.com/gplassard/ouaf/cmd/service_definition"
	"github.com/gplassard/ouaf/cmd/service_level_objectives"
	"github.com/gplassard/ouaf/cmd/service_scorecards"
	"github.com/gplassard/ouaf/cmd/software_catalog"
	"github.com/gplassard/ouaf/cmd/spa"
	"github.com/gplassard/ouaf/cmd/spans"
	"github.com/gplassard/ouaf/cmd/spans_metrics"
	"github.com/gplassard/ouaf/cmd/static_analysis"
	"github.com/gplassard/ouaf/cmd/synthetics"
	"github.com/gplassard/ouaf/cmd/teams"
	"github.com/gplassard/ouaf/cmd/test_optimization"
	"github.com/gplassard/ouaf/cmd/usage_metering"
	"github.com/gplassard/ouaf/cmd/users"
	"github.com/gplassard/ouaf/cmd/workflow_automation"
	"github.com/gplassard/ouaf/pkg/config"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ouaf",
	Short: "ouaf is a Datadog CLI",
	Long:  `ouaf is a Datadog CLI built with Cobra.`,
}

func Execute() {
	defaultSite := os.Getenv("DD_SITE")
	if defaultSite == "" {
		defaultSite = "datadoghq.com"
	}

	rootCmd.PersistentFlags().StringVar(&config.ApiKey, "api-key", os.Getenv("DD_API_KEY"), "Datadog API Key")
	rootCmd.PersistentFlags().StringVar(&config.AppKey, "app-key", os.Getenv("DD_APP_KEY"), "Datadog App Key")
	rootCmd.PersistentFlags().StringVar(&config.Site, "site", defaultSite, "Datadog Site")
	rootCmd.PersistentFlags().BoolVar(&config.Debug, "debug", false, "Enable debug logging")

	rootCmd.AddCommand(
		action_connection.Cmd,
		actions_datastores.Cmd,
		agentless_scanning.Cmd,
		api_management.Cmd,
		apm.Cmd,
		apm_retention_filters.Cmd,
		app_builder.Cmd,
		application_security.Cmd,
		audit.Cmd,
		authn_mappings.Cmd,
		aws_integration.Cmd,
		aws_logs_integration.Cmd,
		case_management.Cmd,
		case_management_attribute.Cmd,
		case_management_type.Cmd,
		ci_visibility_pipelines.Cmd,
		ci_visibility_tests.Cmd,
		cloud_cost_management.Cmd,
		cloud_network_monitoring.Cmd,
		cloudflare_integration.Cmd,
		confluent_cloud.Cmd,
		container_images.Cmd,
		containers.Cmd,
		csm_agents.Cmd,
		csm_coverage_analysis.Cmd,
		csm_threats.Cmd,
		dashboard_lists.Cmd,
		data_deletion.Cmd,
		datasets.Cmd,
		deployment_gates.Cmd,
		domain_allowlist.Cmd,
		dora_metrics.Cmd,
		downtimes.Cmd,
		error_tracking.Cmd,
		events.Cmd,
		fastly_integration.Cmd,
		fleet_automation.Cmd,
		gcp_integration.Cmd,
		incident_services.Cmd,
		incident_teams.Cmd,
		incidents.Cmd,
		ip_allowlist.Cmd,
		key_management.Cmd,
		logs.Cmd,
		logs_archives.Cmd,
		logs_custom_destinations.Cmd,
		logs_metrics.Cmd,
		logs_restriction_queries.Cmd,
		metrics.Cmd,
		microsoft_teams_integration.Cmd,
		monitors.Cmd,
		network_device_monitoring.Cmd,
		observability_pipelines.Cmd,
		okta_integration.Cmd,
		on_call.Cmd,
		on_call_paging.Cmd,
		opsgenie_integration.Cmd,
		org_connections.Cmd,
		organizations.Cmd,
		powerpack.Cmd,
		processes.Cmd,
		reference_tables.Cmd,
		restriction_policies.Cmd,
		roles.Cmd,
		rum.Cmd,
		rum_audience_management.Cmd,
		rum_metrics.Cmd,
		rum_retention_filters.Cmd,
		security_monitoring.Cmd,
		sensitive_data_scanner.Cmd,
		service_accounts.Cmd,
		service_definition.Cmd,
		service_level_objectives.Cmd,
		service_scorecards.Cmd,
		software_catalog.Cmd,
		spa.Cmd,
		spans.Cmd,
		spans_metrics.Cmd,
		static_analysis.Cmd,
		synthetics.Cmd,
		teams.Cmd,
		test_optimization.Cmd,
		usage_metering.Cmd,
		users.Cmd,
		workflow_automation.Cmd,
	)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
