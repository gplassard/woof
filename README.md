# ouaf

`ouaf` is a Datadog CLI tool built with Go. It provides a command-line interface to interact with Datadog API endpoints. The project includes a code generator that automatically creates CLI commands based on the Datadog OpenAPI specification.

## Requirements

- **Go**: 1.25+ (as specified in `go.mod`)
- **Datadog Account**: An active Datadog account with API and Application keys.

## Setup

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd ouaf
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Build the CLI**:
   ```bash
   go build -o ouaf main.go
   ```

## Configuration

`ouaf` requires Datadog credentials to interact with the API. You can provide these via environment variables or command-line flags.

### Environment Variables

- `DD_API_KEY`: Your Datadog API key.
- `DD_APP_KEY`: Your Datadog Application key.
- `DD_SITE`: The Datadog site to use (defaults to `datadoghq.com`).

### Command-line Flags

The following flags are available for all commands:

- `--api-key`: Datadog API Key (defaults to `DD_API_KEY`)
- `--app-key`: Datadog App Key (defaults to `DD_APP_KEY`)
- `--site`: Datadog Site (defaults to `DD_SITE`)

## Usage

Once built, you can run `ouaf` followed by a command and subcommand.

```bash
./ouaf [command] [subcommand] [flags]
```

Example:
```bash
./ouaf roles list-roles
```

To see all available commands, run:
```bash
./ouaf --help
```

## Code Generator

The project features a code generator located in the `generator/` directory. It downloads the Datadog OpenAPI V2 specification and generates the corresponding Cobra commands in the `cmd/` directory.

### Running the Generator

To update the generated commands:

```bash
go run generator/*.go
```

The generator uses templates found in `generator/*.tmpl` and configuration from `generator/config.yaml`.

## Project Structure

- `main.go`: Entry point for the CLI application.
- `cmd/`: Contains the CLI command implementations.
  - `root.gen.go`: The root command definition (generated).
  - Subdirectories (e.g., `cmd/roles/`): Generated commands grouped by API tags.
- `generator/`: Logic and templates for the code generator.
  - `generate.go`: Main entry point for the generator.
  - `config.yaml`: Configuration for skipping or customizing operations.
  - `*.tmpl`: Go templates for generating command files.
- `pkg/`: Shared packages and utilities.
  - `client/`: Datadog API client initialization.
  - `cmdutil/`: Utilities for command execution and output formatting.
  - `output/`: Placeholder for output formatting logic.
- `ouaf`: The compiled binary (after building).

## Tests

To run the project tests:

```bash
go test ./...
```

