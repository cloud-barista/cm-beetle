---
description: "CB-Tumblebug model synchronization guidelines and validation rules"
applyTo: "imdl/**/copied-tb-model.go"
---

# CB-Tumblebug Model Synchronization Rules

## Version Management Protocols

### Target Version Specification

- Always specify complete version strings (e.g., `v0.11.2`, `v0.12.0`)
- Use `latest` only for development/testing purposes
- Include PR references for significant feature changes

### Version Header Format

```go
// * Version: CB-Tumblebug v[VERSION] (include notable features/PRs when relevant)
```

## Struct Synchronization Standards

### Field Documentation Preservation

- Maintain all existing cm-model field comments
- **CRITICAL**: **NEVER DELETE** existing Tumblebug-synchronized field comments, examples, and documentation
- Enhance documentation with TB patterns where beneficial
- Preserve validation tag explanations and examples
- **ABSOLUTE REQUIREMENT**: Preserve ALL field examples and documentation from CB-Tumblebug source

### Source Path Accuracy

- Update path comments to match exact TB source locations
- Include accurate line number ranges **synchronized with target CB-Tumblebug version**
- **CRITICAL**: Verify Path line numbers against actual CB-Tumblebug source files
- **REQUIREMENT**: "// \* Path:" comments must reflect current TB source locations, not outdated references
- Verify paths against target TB version

### Validation Tag Synchronization

- Copy validation tags exactly from TB source
- Maintain `validate:"required"` patterns
- Preserve `example:` and `default:` attributes
- Keep `enums:` specifications accurate

## Change Analysis Requirements

### Breaking Change Detection

- Identify removed or renamed fields
- Note type changes that affect serialization
- Document validation constraint modifications
- Flag changes to required/optional field status

### Compatibility Assessment

- Evaluate impact on existing cm-model consumers
- Document workarounds for breaking changes
- Suggest deprecation strategies for removed fields
- Recommend migration paths for major changes

## Quality Assurance Checklist

### Pre-Synchronization

- [ ] Target TB version accessibility verified
- [ ] Current model inventory documented
- [ ] Backup of current `copied-tb-model.go` created

### During Synchronization

- [ ] All struct definitions updated to match TB source
- [ ] Source path comments reflect accurate locations **with current line numbers**
- [ ] **CRITICAL**: All Tumblebug-synchronized field comments and examples preserved
- [ ] **CRITICAL**: Path line numbers verified against actual CB-Tumblebug source files
- [ ] Version header updated with target version
- [ ] Field documentation enhanced appropriately

### Post-Synchronization Validation

- [ ] Go compilation successful
- [ ] JSON serialization/deserialization tests pass
- [ ] No circular dependency issues introduced
- [ ] All validation constraints function correctly
- [ ] Documentation consistency maintained

## Error Prevention Guidelines

### Common Pitfalls to Avoid

- **Incomplete Updates**: Ensure all related structs are synchronized
- **Path Inaccuracy**: Verify source paths against actual TB files
- **Documentation Loss**: Preserve valuable cm-model-specific comments
- **Validation Inconsistency**: Match TB validation patterns exactly
- **Version Mismatch**: Ensure header reflects actual synchronized version
- **🚨 CRITICAL**: **Comment Deletion** - Never delete existing Tumblebug-synchronized field documentation
- **🚨 CRITICAL**: **Path Desynchronization** - Always verify Path line numbers match CB-Tumblebug source

### Testing Requirements

- Compile cm-model after synchronization
- Verify JSON marshaling/unmarshaling works correctly
- Test with realistic data structures
- Validate against existing cm-model test suites
- Check integration points with dependent systems

## Maintenance Best Practices

### Documentation Standards

- Include rationale for significant structural changes
- Document any cm-model-specific enhancements kept
- Explain compatibility decisions and trade-offs
- Reference relevant TB PRs or issues when applicable

### Change Communication

- Summarize all modifications in PR descriptions
- Highlight breaking changes prominently
- Provide migration guidance for dependent systems
- Update related documentation files as needed
