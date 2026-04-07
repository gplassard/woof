# ADR-0001: SDK-Driven Codegen Pipeline

## Status

Accepted

## Context

The previous codegen pipeline consumed a downloaded OpenAPI spec and derived CLI commands from that schema model. This created drift risk with the runtime source used by the CLI (`datadog-api-client-go`) and required network availability for generation.

## Decision

The generator now uses the Datadog Go SDK (`datadogV2`) as the single source of truth.

Implementation highlights:
- Parse SDK API files with `go/parser` + AST.
- Build an intermediate model containing services, operations, request/response types, enums/constants, and extracted docs.
- Extract operation metadata directly from method signatures and bodies:
  - service and operation names from receivers/methods,
  - HTTP method and path from generated request wiring,
  - parameter API names from path/query assignments,
  - request/response types from signatures,
  - summaries from method comments.
- Infer resource type from response struct `data -> type` traversal using parsed model/enum metadata.
- Keep command templates and output format unchanged as much as possible.

## Consequences

Positive:
- One source of truth aligned with the actual SDK called by generated commands.
- No OpenAPI fetch step.
- Better resiliency in constrained or offline environments.

Tradeoffs:
- Parsing generated SDK method bodies is coupled to SDK generator conventions.
- Some bundle naming required compatibility mappings where SDK service names differ from previous tag-based names.

## Follow-ups

- Add snapshot parity checks against selected command fixtures for high-risk APIs.
- Add parser conformance tests for additional SDK method variants.
