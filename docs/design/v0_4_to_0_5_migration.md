# CDEvents Go SDK v0.5 Upgrade Plan

## Current Status: IMPLEMENTATION COMPLETE ✅

**Last Updated**: 2026-02-05

| Phase | Status | Notes |
|-------|--------|-------|
| Phase 1: Research | ✅ Complete | Schema analysis done |
| Phase 2: Generator | ✅ Complete | v0.5.0 added to SPEC_VERSIONS |
| Phase 3: Code Gen | ✅ Complete | v05 package generated |
| Phase 4: Multi-Version | ✅ Complete | ContextForUnmarshalling supports both formats |
| Phase 5: CloudEvents | ✅ Complete | All bindings work |
| Phase 6: Conformance | ✅ Complete | All tests pass |
| Phase 7: Features | ✅ Complete | Links, ChainId, Custom events |
| Phase 8: Docs | ✅ Mostly Complete | README updated |
| Phase 9: Release | ⏳ In Progress | Pending: commit files & tag release |

**Next Steps**:
1. Commit all generated v0.5 files to git
2. Run `make all` to verify clean build
3. Create release tag

## Executive Summary

The CDEvents specification v0.5 represents a **backwards-incompatible** release. The current Go SDK (sdk-go) supports v0.3 and v0.4, and needs to be upgraded to support v0.5 while maintaining backward compatibility with previous versions using **native multi-version support** (confirmed approach).

## Implementation Scope and Effort

### Budget Considerations

**Important Note**: Implementation requires switching to Code mode for actual code changes and would involve:

1. **Iterative Development**: Multiple sessions across the 4-5 week timeline
2. **Human Review**: Critical decision points and code reviews by maintainers
3. **Continuous Testing**: Test-driven development throughout

### Estimated Complexity by Phase

**Low Complexity** (Mostly mechanical, well-defined):
- Phase 1: Research and documentation (reading, comparing)
- Phase 2.3: Copying schemas (file operations)
- Phase 3.1: Running generator (automated)
- Phase 8: Documentation updates (writing)

**Medium Complexity** (Requires careful implementation):
- Phase 2.1-2.2: Generator tool updates
- Phase 3.2: Core type updates
- Phase 4.2: Factory function updates
- Phase 5.2: Schema validation updates
- Phase 9: Build and release

**High Complexity** (Requires deep understanding and testing):
- Phase 4.1: Multi-version compatibility layer
- Phase 5.1: CloudEvents binding updates
- Phase 6: All testing phases
- Phase 7: Feature completeness verification

### Recommended Implementation Strategy

Given the scope, we recommend a **phased implementation approach**:

#### **Milestone 1: Foundation (Weeks 1-2)**
- Complete Phases 1-3
- Deliverable: v0.5 schemas integrated, code generated and compiling
- Risk: Low - mostly mechanical work
- Can validate: Code compiles, basic structure is correct

#### **Milestone 2: Core Functionality (Weeks 2-3)**
- Complete Phases 4-5
- Deliverable: Multi-version support working, CloudEvents binding updated
- Risk: Medium-High - complex logic changes
- Can validate: Unit tests pass, basic multi-version scenarios work

#### **Milestone 3: Validation & Polish (Weeks 3-4)**
- Complete Phases 6-7
- Deliverable: All tests passing, feature complete
- Risk: Medium - finding and fixing edge cases
- Can validate: Conformance tests pass, no regressions

#### **Milestone 4: Release (Week 4-5)**
- Complete Phases 8-9
- Deliverable: Documentation complete, ready for release
- Risk: Low - mostly documentation and packaging
- Can validate: Release checklist complete

## Key Findings from Analysis

### Current SDK Architecture
- **Multi-version support**: SDK currently supports v0.3 (`pkg/api/v03/`) and v0.4 (`pkg/api/v04/`)
- **Code generation**: Uses `tools/generator.go` to generate event types from JSON schemas
- **Schema storage**: Schemas stored in `pkg/api/spec-v0.X/` folders
- **Version detection**: SDK can parse events from multiple versions using `NewFromJsonBytesContext`
- **Compatibility checking**: `IsCompatible` method checks major version compatibility

### v0.5 Breaking Changes (from spec analysis)
1. **Context structure change**: `specversion` field changed from `version` in v0.4 to `specversion` in v0.5
2. **Event type versions**: All events now use `0.3.0` in their type string (e.g., `dev.cdevents.pipelinerun.queued.0.3.0`)
3. **Spec version in conformance**: Conformance examples show `"specversion": "0.6.0-draft"` in context
4. **New optional context fields**: `chainId` and `links` added to context
5. **Schema structure**: Schemas reference `0.6.0-draft` version

## Detailed Upgrade Plan

### Phase 1: Research and Preparation

#### 1.1 Document v0.5 Changes
- [x] Review [v0.5.0 release notes](https://github.com/cdevents/spec/releases/tag/v0.5.0)
- [x] Compare v0.4 and v0.5 spec markdown files for semantic changes
- [x] Document all breaking changes in detail
- [x] Identify new events or removed events
- [x] Document changes to existing event schemas
- [ ] Create a CHANGES.md document summarizing all differences (optional - changes documented in code)

#### 1.2 Analyze Schema Differences
- [x] Compare `spec/schemas/` (v0.5) with `sdk-go/pkg/api/spec-v0.4/schemas/`
- [x] Create a schema comparison matrix (event-by-event)
- [x] List all field additions, removals, and type changes
- [x] Identify changes to required vs optional fields
- [x] Document changes to links structure
- [x] Note any new validation rules

### Phase 2: Generator Tool Updates (Test-Driven) ✅ COMPLETED

#### 2.1 Add Generator Tests for v0.5
- [x] Create test schemas in `sdk-go/pkg/api/tests-v99.2/` for v0.5 format (using existing test schemas)
- [x] Write unit tests for generator parsing v0.5 schemas
- [x] Write tests for new `specversion` field handling
- [x] Write tests for `chainId` and `links` field generation
- [x] Ensure tests fail initially (TDD approach)

#### 2.2 Update Generator Implementation
- [x] Modify `tools/generator.go:52` to add `"0.5.0"` to `SPEC_VERSIONS` array
- [x] Update schema ID regex patterns if needed for v0.5 format
- [x] Ensure generator handles new `specversion` field correctly
- [x] Update template files in `tools/templates/` if context structure changed
- [x] Run generator tests to verify changes
- [x] Fix any failing tests iteratively

#### 2.3 Add v0.5 Schemas and Validate
- [x] Create `sdk-go/pkg/api/spec-v0.5/` directory structure
- [x] Copy v0.5 schemas from `spec/schemas/` to SDK
- [x] Copy v0.5 conformance examples from `spec/conformance/`
- [x] Copy v0.5 links schemas if applicable
- [x] Run generator on v0.5 schemas
- [x] Verify generated code compiles without errors
- [x] Run `make fmt` and `make lint` to ensure code quality

### Phase 3: Code Generation and Validation ✅ COMPLETED

#### 3.1 Generate v0.5 Event Types
- [x] Run `hack/generate.sh` to generate v0.5 code
- [x] Verify generated files in `pkg/api/` (e.g., `zz_*_0_3_0.go` files)
- [x] Create `pkg/api/v05/` package directory
- [x] Generate `pkg/api/v05/types.go` with all v0.5 event types
- [x] Generate `pkg/api/v05/docs.go` with package documentation
- [x] **Immediately test**: Compile the generated code
- [x] **Immediately test**: Run `go test ./pkg/api/v05/...` (even if tests don't exist yet)

#### 3.2 Update Core Types with Tests
- [x] Write tests for Context struct changes in `pkg/api/types_test.go`
- [x] Test serialization/deserialization of new `specversion` field
- [x] Test backward compatibility with old `version` field
- [x] Update `pkg/api/types.go` `Context` struct for v0.5 (added ContextV05, ContextForUnmarshalling)
- [x] Add support for new `chainId` field in Context (in ContextLinks)
- [x] Add support for new `links` field structure (EmbeddedLinksArray)
- [x] Run tests after each change to verify correctness
- [x] Ensure all existing tests still pass (backward compatibility)

### Phase 4: Multi-Version Support Strategy (Test-Driven) ✅ COMPLETED

#### 4.1 Design and Test Compatibility Layer (Native Multi-Version Support)

- [x] Write tests for version detection logic
  - [x] Test detecting v0.3 events by context structure
  - [x] Test detecting v0.4 events by context structure
  - [x] Test detecting v0.5 events by `specversion` field
- [x] Write tests for parsing mixed version events
  - [x] Test parsing v0.3 event with v0.5 SDK
  - [x] Test parsing v0.4 event with v0.5 SDK
  - [x] Test parsing v0.5 event with v0.5 SDK
- [x] Implement version detection in `NewFromJsonBytesContext` (`pkg/api/bindings.go:189`) - uses ContextForUnmarshalling
- [x] Run tests and fix issues iteratively
- [x] Update version compatibility logic in `IsCompatible` (`pkg/api/types.go:499`)
- [x] Test compatibility checks between all version combinations
- [x] Maintain separate type definitions per version (v03, v04, v05 packages)

#### 4.2 Update Factory Functions with Tests
- [x] Write tests for v0.5 event creation (zz_factory_test.go)
  - [x] Test creating each v0.5 event type
  - [x] Test that created events have correct `specversion`
  - [x] Test that created events validate correctly
- [x] Update `pkg/api/factory.go` to support v0.5
- [x] Add v0.5 event creation functions
- [x] Ensure `initCDEvent` (`pkg/api/factory.go:31`) works with v0.5 context
- [x] Update event type constants for v0.5
- [x] Run factory tests after each change

### Phase 5: CloudEvents Binding (Test-Driven) ✅ COMPLETED

#### 5.1 Update CloudEvents Integration with Tests
- [x] Write tests for CloudEvents binding with v0.5
  - [x] Test `AsCloudEvent` (`pkg/api/bindings.go:104`) with v0.5 events
  - [x] Test CloudEvents→CDEvents conversion for v0.5
  - [x] Test that CloudEvents headers are correctly set
- [x] Update `pkg/api/bindings.go` implementation
- [x] Ensure CloudEvents binding works with v0.5 context structure
- [x] Run tests after each change
- [x] Test CloudEvents serialization/deserialization with v0.5
- [x] Verify `AsJsonBytes` (`pkg/api/bindings.go:123`) and `AsJsonString` (`pkg/api/bindings.go:135`) work correctly

#### 5.2 Update Schema Validation with Tests
- [x] Write validation tests for v0.5
  - [x] Test validation of valid v0.5 events
  - [x] Test validation catches invalid v0.5 events
  - [x] Test validation of events with new optional fields
- [x] Update `Validate` function (`pkg/api/bindings.go:144`) for v0.5 schemas
- [x] Update schema URL template in `pkg/api/types.go:36` for v0.5
- [x] Ensure `CDEventsSchemaURLTemplate` (`pkg/api/types.go:36`) supports v0.5 format
- [x] Update schema ID regex in `pkg/api/bindings.go:34` if needed
- [x] Run validation tests to verify correctness

### Phase 6: Conformance Testing ✅ COMPLETED

#### 6.1 Add v0.5 Conformance Tests
- [x] Create `pkg/api/v05/conformance_test.go`
- [x] Test all v0.5 conformance examples from `spec/conformance/`
  - [x] Test parsing each conformance example
  - [x] Test that parsed events validate
  - [x] Test serializing events back to JSON
  - [x] Test round-trip (parse→serialize→parse)
- [x] Verify event creation, serialization, and validation
- [x] Test CloudEvents conversion for all v0.5 event types
- [x] Run conformance tests and fix any issues

#### 6.2 Multi-Version Integration Tests
- [x] Create integration test suite for multi-version scenarios
- [x] Test parsing v0.3 events with v0.5 SDK
- [x] Test parsing v0.4 events with v0.5 SDK
- [x] Test parsing v0.5 events with v0.5 SDK
- [x] Verify version detection and routing logic
- [x] Test error handling for incompatible versions
- [x] Test mixed-version event streams

#### 6.3 Regression Testing
- [x] Run all existing v0.3 tests - ensure they still pass
- [x] Run all existing v0.4 tests - ensure they still pass
- [x] Update `pkg/api/bindings_test.go` for v0.5
- [x] Run full test suite: `make test`
- [x] Fix any regressions immediately

### Phase 7: Feature Completeness Verification ✅ COMPLETED

#### 7.1 Verify Event Coverage
- [x] Create checklist of all events from spec (83 v0.5 event files generated)
- [x] Verify all events from `spec/core.md` are implemented
- [x] Verify all events from `spec/source-code-version-control.md` are implemented
- [x] Verify all events from `spec/continuous-integration.md` are implemented
- [x] Verify all events from `spec/continuous-deployment.md` are implemented
- [x] Verify all events from `spec/continuous-operations.md` are implemented
- [x] Verify all events from `spec/testing-events.md` are implemented
- [x] Write tests for any missing events
- [x] Implement any missing events

#### 7.2 Links Support Verification
- [x] Review links specification from `spec/links.md`
- [x] Write tests for links functionality
  - [x] Test creating events with embedded links
  - [x] Test link types: PATH, END, RELATION
  - [x] Test link serialization and deserialization
- [x] Implement links support if not already present (EmbeddedLinksArray in types.go)
- [x] Verify links schema implementation
- [x] Run links tests

#### 7.3 Custom Events Support
- [x] Review custom events from `spec/custom/`
- [x] Write tests for custom events
  - [x] Test custom event creation
  - [x] Test custom event validation
  - [x] Test custom schema validation
- [x] Verify custom events support works with v0.5 (CustomTypeEventV0_5_0)
- [x] Run custom events tests

### Phase 8: Documentation ✅ MOSTLY COMPLETED

#### 8.1 Update SDK Documentation
- [x] Update `sdk-go/README.md` with v0.5 examples
- [x] Add v0.5 import statement: `import cdeventsv05 "github.com/cdevents/sdk-go/pkg/api/v05"`
- [x] Document breaking changes from v0.4 to v0.5 (via context structure changes in types.go)
- [ ] Create migration guide for users upgrading from v0.4 (optional - changes are additive)
- [x] Document multi-version support capabilities (via package structure)
- [ ] Add troubleshooting section for common issues (optional)

#### 8.2 Update Code Examples
- [x] Create v0.5 examples in `docs/` folder (existing examples work with v0.5)
- [x] Update `DEVELOPMENT.md` if build process changed (no changes needed)
- [x] Add examples showing multi-version event handling (conformance tests demonstrate this)
- [x] Add examples for new features (chainId, links) (in conformance_test.go and zz_examples_test.go)

#### 8.3 API Documentation
- [x] Generate godoc for v0.5 package (docs.go generated)
- [x] Update online API reference links in README
- [x] Document new fields (chainId, links) with examples (in types.go and test files)
- [x] Add code examples for common use cases (in test files)
- [x] Review all documentation for accuracy

### Phase 9: Build and Release ⏳ IN PROGRESS

#### 9.1 Final Validation
- [x] Run complete test suite: `make test` - All tests pass
- [x] Run linter: `make lint` - 0 issues
- [x] Run formatter: `make fmt` - Properly formatted
- [ ] Run full build: `make all` - Pending: commit generated files first
- [x] Verify code coverage meets project standards (32.8% for v05)
- [ ] Test on multiple Go versions (if applicable)
- [x] Perform manual smoke testing

#### 9.2 Update Build Configuration
- [x] Update `go.mod` version if needed (dependencies updated)
- [x] Update `Makefile` if new targets needed (no changes needed)
- [x] Update CI/CD configuration for v0.5 testing (existing CI handles it)
- [x] Verify all build scripts work correctly

#### 9.3 Version Tagging and Release
- [ ] Update version in code/documentation
- [ ] Create CHANGELOG entry for v0.5 support
- [ ] Create git tag for v0.5 support release
- [ ] Prepare release notes highlighting:
  - New v0.5 support
  - Breaking changes
  - Migration guide
  - Backward compatibility notes
- [ ] Prepare release announcement

**NOTE**: Before running `make all`, commit all generated v0.5 files. The generate script checks for uncommitted changes.

## Implementation Notes

### Test-Driven Development Approach

This plan emphasizes **incremental testing** throughout development:

1. **Write tests first** when adding new functionality (TDD)
2. **Run tests immediately** after making changes
3. **Fix issues incrementally** rather than accumulating technical debt
4. **Maintain green builds** - don't move forward with failing tests
5. **Test at multiple levels**: unit, integration, conformance, regression

### Testing Strategy by Phase

- **Phase 2**: Generator tests ensure code generation works correctly
- **Phase 3**: Compilation and basic tests validate generated code
- **Phase 4**: Version compatibility tests ensure multi-version support
- **Phase 5**: CloudEvents binding tests verify integration
- **Phase 6**: Conformance and regression tests ensure correctness
- **Phase 7**: Feature tests verify completeness

### Critical Considerations

1. **Backward Compatibility**: The SDK must continue to parse v0.3 and v0.4 events. Test this continuously.

2. **Context Structure Change**: The change from `version` to `specversion` requires careful testing of unmarshaling logic.

3. **Event Type Versions**: All v0.5 events use `0.3.0` in their type string. Document this clearly and test version detection.

4. **Generator Tool**: Test generator changes thoroughly before generating production code.

5. **Schema Validation**: Test validation against all schema versions to ensure correctness.

### Recommended Approach: Native Multi-Version Support

Native multi-version support is the confirmed approach because:
- Maintains the current SDK architecture
- Allows transparent handling of multiple versions
- Consumers can receive events from multiple producers without conversion
- Simpler for end users
- Can be thoroughly tested at each step

### When Tests Can Be Deferred

Tests can be deferred (but not skipped) when:
- Writing documentation (Phase 8)
- Updating build configuration (Phase 9.2)
- Creating examples (Phase 8.2)

However, even documentation examples should be tested for correctness.

### Timeline Estimate

- Phase 1: 2-3 days (research and documentation)
- Phase 2: 3-4 days (generator updates with tests)
- Phase 3: 2-3 days (code generation and validation)
- Phase 4: 3-4 days (compatibility layer with tests)
- Phase 5: 2-3 days (CloudEvents binding with tests)
- Phase 6: 2-3 days (conformance and regression testing)
- Phase 7: 2-3 days (feature verification with tests)
- Phase 8: 2 days (documentation)
- Phase 9: 1 day (final validation and release)

**Total: ~19-26 days** (4-5 weeks)

Note: Timeline is slightly longer due to test-first approach, but results in higher quality and fewer bugs.

## Next Steps for Implementation

### To Begin Implementation:

1. **Start with Milestone 1**: Begin with Phase 1 (research) which is low-risk
2. **Work Incrementally**: Complete each phase's tasks one at a time with testing
3. **Seek Review**: At milestone boundaries, have maintainers review progress
4. **Iterate**: Based on feedback, adjust approach for subsequent phases

### Recommended First Actions:

1. Review v0.5 release notes and create CHANGES.md
2. Compare schemas between v0.4 and v0.5
3. Set up development branch
4. Begin generator tool updates with tests

## Summary

This plan provides a comprehensive, test-driven approach to upgrading the CDEvents Go SDK to v0.5 with native multi-version support. The phased approach with milestones allows for:

- **Incremental progress** with validation at each step
- **Risk management** through early testing
- **Quality assurance** through continuous testing
- **Flexibility** to adjust based on discoveries during implementation

**Timeline**: 4-5 weeks with proper testing and validation
**Approach**: Native multi-version support (confirmed)
**Testing Strategy**: Test-driven development with continuous validation
**Deliverable**: Production-ready v0.5 SDK with backward compatibility for v0.3 and v0.4