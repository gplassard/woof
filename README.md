# woof

`woof` is a Datadog CLI tool built with Go. It provides a command-line interface to interact with Datadog API endpoints. The project includes a code generator that automatically creates CLI commands from the Datadog Go SDK (`datadogV2`).

## Disclaimer

This project was heavily built using AI.

## Requirements

- **Go**: 1.25+ (as specified in `go.mod`)
- **Datadog Account**: An active Datadog account with API and Application keys.

## Setup

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd woof
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Build the CLI**:
   ```bash
   go build .
   ```

## Configuration

`woof` requires Datadog credentials to interact with the API. You can provide these via environment variables or command-line flags.

### Environment Variables

- `DD_API_KEY`: Your Datadog API key.
- `DD_APP_KEY`: Your Datadog Application key.
- `DD_SITE`: The Datadog site to use (defaults to `datadoghq.com`).
- `DD_CLOUD_AUTH_ORG_UUID`: Datadog organization UUID for cloud auth delegated token flow.
- `DD_CLOUD_AUTH_AWS_REGION`: Optional AWS region for STS signing. If empty, Datadog SDK defaults to `us-east-1`.
- `AWS_ACCESS_KEY_ID`: AWS access key ID for cloud auth.
- `AWS_SECRET_ACCESS_KEY`: AWS secret access key for cloud auth.
- `AWS_SESSION_TOKEN`: AWS session token for cloud auth.

### Command-line Flags

The following flags are available for all commands:

- `--api-key`: Datadog API Key (defaults to `DD_API_KEY`)
- `--app-key`: Datadog App Key (defaults to `DD_APP_KEY`)
- `--site`: Datadog Site (defaults to `DD_SITE`)
- `--cloud-auth-org-uuid`: Datadog org UUID for cloud auth (defaults to `DD_CLOUD_AUTH_ORG_UUID`)
- `--cloud-auth-aws-region`: AWS region for cloud auth STS signing (defaults to `DD_CLOUD_AUTH_AWS_REGION`)
- `--aws-access-key-id`: AWS access key ID for cloud auth (defaults to `AWS_ACCESS_KEY_ID`)
- `--aws-secret-access-key`: AWS secret access key for cloud auth (defaults to `AWS_SECRET_ACCESS_KEY`)
- `--aws-session-token`: AWS session token for cloud auth (defaults to `AWS_SESSION_TOKEN`)

## Usage

Once built, you can run `woof` followed by a command and subcommand.

```bash
./woof [command] [subcommand] [flags]
```

Example:
```bash
./woof roles list-roles
```

To see all available commands, run:
```bash
./woof --help
```

## Code Generator

The project features a code generator located in the `generator/` directory. It parses the Datadog Go SDK package `github.com/DataDog/datadog-api-client-go/v2/api/datadogV2` with Go AST and produces Cobra commands in the `cmd/` directory.

### Running the Generator

To update the generated commands:

```bash
go run ./generator
```

The generator uses templates found in `generator/*.tmpl` and configuration from `generator/config.yaml`.

## Project Structure

- `main.go`: Entry point for the CLI application.
- `cmd/`: Contains generated CLI command implementations.
  - `root.gen.go`: The root command definition (generated).
  - Subdirectories (for example `cmd/roles/`): generated commands grouped by SDK services.
- `generator/`: SDK parser + generation logic and templates.
  - `generate.go`: Main entry point for the generator.
  - `sdk_source.go`: SDK model extraction and operation parsing.
  - `config.yaml`: Configuration for skipping or customizing operations.
  - `*.tmpl`: Go templates for generating command files.
- `pkg/`: Shared packages and utilities.
  - `config/`: Configuration and credential management.
  - `client/`: Datadog API client initialization.
  - `cmdutil/`: Utilities for command execution and output formatting.

## Tests

To run project tests:

```bash
go test ./...
```
