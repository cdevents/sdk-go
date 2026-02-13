# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Support for CDEvents specification v0.5.0
- New `pkg/api/v05` package with v0.5 event types
- `ContextV05` struct with `specversion` field (replaces `version` from v0.4)
- `ContextForUnmarshalling` for backward-compatible event parsing
- Support for `chainId` field in event context
- Support for `links` field with embedded links (PATH, END, RELATION types)
- Support for `schemaUri` field for custom schema validation
- 83 generated event types for v0.5 specification
- Comprehensive conformance tests for v0.5 events
- Multi-version support: SDK can parse v0.3, v0.4, and v0.5 events

### Changed
- Updated README.md with v0.5 examples and import statements
- Reordered API reference links (v05 first, then v04, v03)
- Updated Go version to 1.24.0 with toolchain 1.24.3
- Updated dependencies (golang.org/x/mod, golang.org/x/sys, etc.)
- Generator now includes v0.5.0 in SPEC_VERSIONS

### Technical Details
- **Breaking Change in Spec**: v0.5 uses `specversion` field instead of `version` in context
- **Event Type Versions**: All v0.5 events use version `0.2.0` or `0.3.0` in their type string
- **Backward Compatibility**: SDK maintains full backward compatibility with v0.3 and v0.4 events
- **New Event Types**: Added TicketClosed, TicketCreated, TicketUpdated events (v0.2.0)
- **Schema Version**: v0.5 schemas reference CDEvents spec version 0.5.0

### Migration Guide
To migrate from v0.4 to v0.5:

1. Update your import statement:
   ```go
   // Old (v0.4)
   import cdevents "github.com/cdevents/sdk-go/pkg/api/v04"
   
   // New (v0.5)
   import cdevents "github.com/cdevents/sdk-go/pkg/api/v05"
   ```

2. Event creation remains the same:
   ```go
   event, err := cdevents.NewPipelineRunQueuedEvent()
   ```

3. New optional fields available:
   ```go
   event.SetChainId("my-chain-id")
   event.SetLinks(links)
   event.SetSchemaUri("https://example.com/schema")
   ```

4. The SDK automatically handles parsing events from all versions (v0.3, v0.4, v0.5)

## Previous Releases

For releases prior to v0.5 support, please refer to the git history and release tags.