### Build and Configuration Instructions

#### Build Instructions
The project is written in Go and uses a standard build process.
- **Build the binary**:
  ```bash
  go build .
  ```
- **Install dependencies**:
  ```bash
  go mod download
  ```

#### Configuration
`ouaf` requires Datadog credentials. These can be provided via environment variables or command-line flags:
- `DD_API_KEY`: Your Datadog API key (or `--api-key` flag).
- `DD_APP_KEY`: Your Datadog Application key (or `--app-key` flag).
- `DD_SITE`: The Datadog site to use, defaults to `datadoghq.eu` (or `--site` flag).

### Testing Information

#### Running Tests
To run all tests in the project:
```bash
go test ./...
```

#### Adding New Tests
When adding new features or fixing bugs, add tests in the corresponding package.
- **Package Utilities**: Add `_test.go` files in the same directory.
- **Generated Commands**: These are mostly boilerplate; focus testing on shared logic in `pkg/`.
- **Generator**: Add tests in `generator/` to verify code generation logic.

### Additional Development Information

#### Code Style
- Follow standard Go conventions (`gofmt`, `go vet`).
- CLI commands use the [Cobra](https://github.com/spf13/cobra) library.
- API interactions use the [Datadog Go API Client V2](https://github.com/DataDog/datadog-api-client-go).

#### Code Generator
The `cmd/` directory is fully generated. Do not manually edit any files in this directory.
- **Run Generator**: `go run ./generator`
- **Templates**: Located in `generator/*.tmpl`.
- **Configuration**: `generator/config.yaml` allows skipping specific operations.
- **Important**: If you modify templates or the generator logic, run the generator and ensure `TestGenerationNoChanges` in `generator/generator_test.go` passes (it checks if `git status` is clean after generation).

#### Directory Structure
- `cmd/`: Generated Cobra commands. Everything under this directory should be generated; do not add manual changes here.
- `pkg/config/`: Configuration and credential management.
- `pkg/client/`: API client setup.
- `pkg/cmdutil/`: Command utilities and output formatting.
- `generator/`: Code generation logic.
