#!/usr/bin/env python3
"""
CB-Tumblebug Model Dependency Analyzer

This script analyzes the dependencies between structs in the cloudmodel package
(copied-tb-model.go and model.go) and provides insights into struct relationships
and usage patterns across the entire package.

Usage:
    python3 scripts/analyze_dependencies.py
    python3 scripts/analyze_dependencies.py --verbose
    python3 scripts/analyze_dependencies.py --unused-only
"""

import argparse
import re
import sys
from pathlib import Path


def find_project_root():
    """Find the project root directory by looking for go.mod file."""
    # Start from script directory
    script_dir = Path(__file__).parent.parent

    # Check script's parent directory first (likely project root)
    if (script_dir / "go.mod").exists():
        return script_dir

    # Check current working directory and its parents
    current = Path.cwd()
    for parent in [current] + list(current.parents):
        if (parent / "go.mod").exists():
            return parent

    # Fallback to script's parent directory
    return script_dir


def extract_structs_and_types(content, file_name=""):
    """Extract all struct names and custom type definitions from Go content."""
    # Extract struct names
    struct_pattern = r"type\s+(\w+)\s+struct"
    all_structs = set(re.findall(struct_pattern, content))

    # Extract string type definitions
    string_type_pattern = r"type\s+(\w+)\s+string"
    string_types = set(re.findall(string_type_pattern, content))

    # Combine all custom types
    all_types = all_structs | string_types

    return all_structs, string_types, all_types


def read_cloudmodel_files(project_root):
    """Read all files in the cloudmodel package."""
    cloud_model_dir = project_root / "imdl" / "cloud-model"

    files_data = {}
    all_structs = set()
    all_string_types = set()
    all_types = set()

    # Files to analyze in the cloudmodel package
    target_files = [("copied-tb-model.go", "CB-Tumblebug models"), ("model.go", "CM-Model types"), ("vm-infra-info.go", "VM infrastructure info")]

    for filename, description in target_files:
        file_path = cloud_model_dir / filename
        if file_path.exists():
            try:
                with open(file_path, "r", encoding="utf-8") as f:
                    content = f.read()

                structs, string_types, types = extract_structs_and_types(content, filename)

                files_data[filename] = {"content": content, "description": description, "structs": structs, "string_types": string_types, "types": types}

                all_structs.update(structs)
                all_string_types.update(string_types)
                all_types.update(types)

            except Exception as e:
                print(f"⚠️  Warning: Could not read {filename}: {e}")

    return files_data, all_structs, all_string_types, all_types


def find_struct_dependencies_in_content(content, struct_name, all_types):
    """Find dependencies for a specific struct within given content."""
    # Extract the struct definition
    struct_start = content.find(f"type {struct_name} struct")
    if struct_start == -1:
        return []

    # Find the end of the struct (next 'type' declaration or end of file)
    next_type = content.find("\ntype ", struct_start + 1)
    if next_type == -1:
        struct_content = content[struct_start:]
    else:
        struct_content = content[struct_start:next_type]

    # Find dependencies
    dependencies = []
    for other_type in all_types:
        if other_type == struct_name:
            continue

        # Look for field declarations that use other custom types
        # Enhanced patterns to handle complex type references:
        # - Simple: FieldName Type
        # - Array: FieldName []Type
        # - Pointer: FieldName *Type
        # - Pointer to array: FieldName *[]Type
        # - Array of pointers: FieldName []*Type
        patterns = [
            rf"\s+\w+\s+{re.escape(other_type)}(\s|$|\`)",  # FieldName Type
            rf"\s+\w+\s+\[\]{re.escape(other_type)}(\s|$|\`)",  # FieldName []Type
            rf"\s+\w+\s+\*{re.escape(other_type)}(\s|$|\`)",  # FieldName *Type
            rf"\s+\w+\s+\*\[\]{re.escape(other_type)}(\s|$|\`)",  # FieldName *[]Type
            rf"\s+\w+\s+\[\]\*{re.escape(other_type)}(\s|$|\`)",  # FieldName []*Type
        ]

        for pattern in patterns:
            if re.search(pattern, struct_content):
                dependencies.append(other_type)
                break  # Found this type, no need to check other patterns

    return dependencies


def find_struct_dependencies_across_files(files_data, struct_name, all_types):
    """Find dependencies for a struct across all cloudmodel files."""
    dependencies = []
    source_file = None

    # First, find which file contains this struct
    for filename, data in files_data.items():
        if struct_name in data["structs"]:
            source_file = filename
            dependencies = find_struct_dependencies_in_content(data["content"], struct_name, all_types)
            break

    return dependencies, source_file


def find_references_to_struct_across_files(files_data, target_struct):
    """Find which structs reference the target struct across all files."""
    references = []

    for filename, data in files_data.items():
        for struct_name in data["structs"]:
            if struct_name == target_struct:
                continue

            # Extract the struct definition
            content = data["content"]
            struct_start = content.find(f"type {struct_name} struct")
            if struct_start == -1:
                continue

            next_type = content.find("\ntype ", struct_start + 1)
            if next_type == -1:
                struct_content = content[struct_start:]
            else:
                struct_content = content[struct_start:next_type]

            # Check if this struct references the target struct
            # Enhanced patterns to handle complex type references:
            patterns = [
                rf"\s+\w+\s+{re.escape(target_struct)}(\s|$|\`)",  # FieldName Type
                rf"\s+\w+\s+\[\]{re.escape(target_struct)}(\s|$|\`)",  # FieldName []Type
                rf"\s+\w+\s+\*{re.escape(target_struct)}(\s|$|\`)",  # FieldName *Type
                rf"\s+\w+\s+\*\[\]{re.escape(target_struct)}(\s|$|\`)",  # FieldName *[]Type
                rf"\s+\w+\s+\[\]\*{re.escape(target_struct)}(\s|$|\`)",  # FieldName []*Type
            ]

            found = False
            for pattern in patterns:
                if re.search(pattern, struct_content):
                    references.append((struct_name, filename))
                    found = True
                    break

            if found:
                continue

    return references


def check_usage_in_file(file_path, struct_names):
    """Check which structs are used in a specific file."""
    try:
        with open(file_path, "r", encoding="utf-8") as f:
            content = f.read()

        used_structs = []
        for struct_name in struct_names:
            if struct_name in content:
                used_structs.append(struct_name)

        return used_structs
    except FileNotFoundError:
        return []


def analyze_dependencies(verbose=False, unused_only=False):
    """Main analysis function."""
    project_root = find_project_root()

    # Read all cloudmodel package files
    files_data, all_structs, all_string_types, all_types = read_cloudmodel_files(project_root)

    if not files_data:
        print("❌ Error: No cloudmodel files found!")
        return 1

    print("🔍 CB-Tumblebug Model Dependency Analysis (CloudModel Package)")
    print("=" * 65)

    if not unused_only:
        print("\n📊 Statistics:")
        print(f"   Total files analyzed: {len(files_data)}")
        for filename, data in files_data.items():
            print(f"   {filename}: {len(data['structs'])} structs, {len(data['string_types'])} string types")
        print(f"   Package total: {len(all_structs)} structs, {len(all_string_types)} string types, {len(all_types)} custom types")

        if verbose:
            print("\n📁 Files analyzed:")
            for filename, data in files_data.items():
                print(f"   • {filename} ({data['description']})")
                if data["structs"]:
                    struct_list = ", ".join(sorted(data["structs"]))
                    print(f"     Structs: {struct_list}")
                if data["string_types"]:
                    types_list = ", ".join(sorted(data["string_types"]))
                    print(f"     Types: {types_list}")

    # Analyze dependencies across all files
    struct_dependencies = {}
    struct_files = {}  # Track which file each struct is defined in

    for struct_name in all_structs:
        dependencies, source_file = find_struct_dependencies_across_files(files_data, struct_name, all_types)
        struct_dependencies[struct_name] = dependencies
        struct_files[struct_name] = source_file

    # Find unreferenced structs: Check copied-tb-model.go structs for any references
    # across all files (copied-tb-model.go, model.go, vm-infra-info.go)
    unreferenced_structs = []

    # Get copied-tb-model.go data
    copied_tb_data = files_data.get("copied-tb-model.go")
    if copied_tb_data:
        copied_tb_structs = copied_tb_data["structs"]

        # Check each struct in copied-tb-model.go for any references in the entire package
        for struct_name in copied_tb_structs:
            references = find_references_to_struct_across_files(files_data, struct_name)
            if not references:
                unreferenced_structs.append(struct_name)

    # Note: Only copied-tb-model.go structs are checked for unreferenced status
    # model.go and vm-infra-info.go structs are excluded as they are external API models

    if unused_only:
        print(f"\n⚠️  UNREFERENCED STRUCTS (not used by any other struct) [{len(unreferenced_structs)}]:")
        if unreferenced_structs:
            # Only copied-tb-model.go structs are checked for unreferenced status
            copied_tb_unreferenced = [s for s in unreferenced_structs if struct_files[s] == "copied-tb-model.go"]

            if copied_tb_unreferenced:
                print("\n   📄 From copied-tb-model.go (no references across all files):")
                for struct in sorted(copied_tb_unreferenced):
                    print(f"      • {struct}")
            else:
                print("   No unreferenced structs found in copied-tb-model.go.")
        else:
            print("   No unreferenced structs found.")

        if verbose:
            print("\n💡 Analysis method:")
            print("   - copied-tb-model.go: Check references across all files (copied-tb-model.go, model.go, vm-infra-info.go)")
            print("   - model.go & vm-infra-info.go: Excluded from unreferenced analysis (external API models)")

        return 0

    # Separate all structs into referenced and unreferenced categories
    referenced_structs = []
    truly_unreferenced_structs = []

    for struct_name in all_structs:
        references = find_references_to_struct_across_files(files_data, struct_name)
        source_file = struct_files[struct_name]

        if references:
            referenced_structs.append((struct_name, source_file, references))
        else:
            # Only include copied-tb-model.go structs in truly unreferenced
            # (model.go and vm-infra-info.go structs are external API models)
            if source_file == "copied-tb-model.go":
                truly_unreferenced_structs.append(struct_name)

    # Display referenced structs, prioritizing copied-tb-model.go
    copied_tb_referenced = []
    other_referenced = []

    for struct_name, source_file, references in referenced_structs:
        if source_file == "copied-tb-model.go":
            copied_tb_referenced.append((struct_name, source_file, references))
        else:
            other_referenced.append((struct_name, source_file, references))

    print(f"\n✅ REFERENCED STRUCTS (used by other structs) [{len(copied_tb_referenced)}]:")

    if copied_tb_referenced:
        print("\n   📄 From copied-tb-model.go:")
        for struct_name, source_file, references in copied_tb_referenced:
            ref_details = []
            for ref_struct, ref_file in references:
                if ref_file == "copied-tb-model.go":
                    ref_details.append(ref_struct)
                else:
                    ref_details.append(f"{ref_struct} ({ref_file})")
            ref_str = ", ".join(ref_details)
            print(f"      • {struct_name} ← {ref_str}")
    else:
        print("   No referenced structs found in copied-tb-model.go.")

    print(f"\n🏝️  UNREFERENCED STRUCTS (not used by other structs) [{len(truly_unreferenced_structs)}]:")
    if truly_unreferenced_structs:
        print("   📄 From copied-tb-model.go:")
        for struct in sorted(truly_unreferenced_structs):
            print(f"      • {struct}")
    else:
        print("   No unreferenced structs found in copied-tb-model.go.")

    if verbose:
        print("\n📈 Detailed Dependency Analysis:")

        # Special analysis for copied-tb-model.go structs
        copied_tb_data = files_data.get("copied-tb-model.go")
        if copied_tb_data:
            print("\n   🔍 Internal Dependencies within copied-tb-model.go:")
            copied_tb_structs = copied_tb_data["structs"]
            copied_tb_types = copied_tb_data["types"]

            for struct_name in sorted(copied_tb_structs):
                internal_dependencies = find_struct_dependencies_in_content(copied_tb_data["content"], struct_name, copied_tb_types)
                external_references = find_references_to_struct_across_files(files_data, struct_name)

                print(f"\n      🔸 {struct_name}:")
                print(f"         Internal dependencies: {internal_dependencies if internal_dependencies else 'None'}")
                if external_references:
                    ref_details = [f"{ref_struct} ({ref_file})" for ref_struct, ref_file in external_references]
                    print(f"         External references: {ref_details}")
                else:
                    print("         External references: None")

        # Regular analysis for all other structs
        print("\n   📊 All Structs Analysis:")
        for struct_name in sorted(all_structs):
            if struct_files[struct_name] == "copied-tb-model.go":
                continue  # Already analyzed above

            dependencies = struct_dependencies[struct_name]
            references = find_references_to_struct_across_files(files_data, struct_name)
            source_file = struct_files[struct_name]

            print(f"\n      🔸 {struct_name} (defined in {source_file}):")
            print(f"         Dependencies: {dependencies if dependencies else 'None'}")
            if references:
                ref_details = [f"{ref_struct} ({ref_file})" for ref_struct, ref_file in references]
                print(f"         Referenced by: {ref_details}")
            else:
                print("         Referenced by: None")

    return 0


def main():
    parser = argparse.ArgumentParser(
        description="Analyze struct dependencies in CB-Tumblebug model files",
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog="""
Examples:
  python3 scripts/analyze_dependencies.py                    # Basic analysis
  python3 scripts/analyze_dependencies.py --verbose          # Detailed analysis
  python3 scripts/analyze_dependencies.py --unused-only      # Show only unreferenced structs
        """,
    )

    parser.add_argument("--verbose", "-v", action="store_true", help="Show detailed dependency information")

    parser.add_argument("--unused-only", "-u", action="store_true", help="Show only unreferenced structs (not used by other structs)")

    args = parser.parse_args()

    try:
        return analyze_dependencies(verbose=args.verbose, unused_only=args.unused_only)
    except KeyboardInterrupt:
        print("\n\n⚠️  Analysis interrupted by user")
        return 1
    except Exception as e:
        print(f"\n❌ Error during analysis: {e}")
        return 1


if __name__ == "__main__":
    sys.exit(main())
