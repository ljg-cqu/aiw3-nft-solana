# AIW3 NFT Documentation Lifecycle Guide

<!-- Document Metadata -->
**Version:** v1.0.0  
**Last Updated:** 2025-08-06  
**Status:** Active  
**Purpose:** Provides guidelines for maintaining and managing the lifecycle of all AIW3 NFT project documentation.

---

## 1. Introduction

This guide establishes the official process for creating, updating, and maintaining all documentation within the AIW3 NFT project. Adhering to these standards ensures our documentation remains accurate, consistent, and valuable for the entire team.

All documentation is subject to the metadata standard defined in `UNIFIED DOCUMENTATION METADATA STANDARD IMPLEMENTED` and validated by the script located at `scripts/validate_docs.sh`.

## 2. Versioning Guidelines

We use GitHub-style semantic versioning (SemVer) in the format `vX.Y.Z` to track document revisions.

- **MAJOR (X)**: Increment for fundamental changes that overhaul the document's core concepts, architecture, or purpose. A major version change often implies that previous versions are obsolete.
  - *Example*: Rewriting the System Design document to use a completely new architecture.

- **MINOR (Y)**: Increment for substantial additions or changes that add new information or sections but do not fundamentally alter the existing content's meaning. The document remains backward-compatible with previous minor versions.
  - *Example*: Adding a new section for a new service, or adding detailed diagrams to an existing section.

- **PATCH (Z)**: Increment for minor corrections, clarifications, typo fixes, or formatting updates that do not change the substance of the document.
  - *Example*: Fixing a broken link, correcting a field name, or clarifying a sentence.

## 3. Document Status Management

The `Status` field in the metadata header tracks a document's current state in its lifecycle.

- **`Active`**: The document is current, accurate, and relevant to the project's present state. This is the default status for most documentation.

- **`In Review`**: The document is undergoing significant revisions or updates. This status indicates that the content may be in flux and should be read with caution.

- **`Deprecated`**: The document contains outdated information but is kept for historical context. It should not be used for current development work. A deprecated document should clearly state what new document supersedes it.

- **`Archived`**: The document is no longer relevant to the project and is preserved for archival purposes only. 

## 4. Document Update Process

Follow these steps when updating any document:

1.  **Identify Required Changes**: Determine the scope of your update.
2.  **Update the Content**: Make the necessary additions, corrections, or revisions to the document body.
3.  **Update the Metadata Header**:
    -   Increment the **Version** field according to the Versioning Guidelines.
    -   Update the **Last Updated** field to the current date (YYYY-MM-DD).
    -   If necessary, change the **Status** field.
4.  **Validate Your Changes**: Before committing, run the validation script from the project root to ensure compliance:
    ```bash
    ./scripts/validate_docs.sh
    ```
5.  **Commit and Push**: Commit your changes with a clear message describing the update.

## 5. Metadata Validation

The `scripts/validate_docs.sh` script is a critical tool for maintaining our documentation standards. It automatically checks all `.md` files (excluding specific non-documentation files) for a correctly formatted metadata header.

- **Purpose**: To enforce consistency and prevent non-compliant documentation from being added to the repository.
- **Usage**: Run it from the project root before any commit that involves changes to markdown files.
- **CI/CD Integration**: This script is designed to be integrated into a CI/CD pipeline to automate validation and block pull requests that fail the check.
