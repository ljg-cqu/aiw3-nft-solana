#!/bin/bash

# validate_docs.sh
# Description: Validates that all markdown files in the repository contain the standard
#              AIW3 NFT documentation metadata header.
# Usage: ./scripts/validate_docs.sh

# --- Configuration ---
PROJECT_ROOT=$(git rev-parse --show-toplevel)
EXIT_CODE=0

# --- Helper Functions ---

# Color codes for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

error() {
    echo -e "${RED}[ERROR]${NC} $1"
    EXIT_CODE=1
}

warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

# --- Main Logic ---

info "Starting documentation metadata validation..."

# Find all markdown files, excluding those in node_modules
# Find all markdown files, excluding node_modules and specific chat history files
MD_FILES=$(find "$PROJECT_ROOT" -type f -name "*.md" -not -path "*/node_modules/*" -not -name ".aider.chat.history.md")

if [ -z "$MD_FILES" ]; then
    warn "No markdown files found. Exiting."
    exit 0
fi

for file in $MD_FILES; do
    # Read the first 15 lines of the file, should be enough for the header
    header_content=$(head -n 15 "$file")

    # Check for the presence of the metadata block comment
    if ! echo "$header_content" | grep -q "<!-- Document Metadata -->"; then
        error "Missing metadata block in: $file"
        continue
    fi

    # Check for key fields
    if ! echo "$header_content" | grep -q -E "^\*\*Version:\*\*"; then
        error "Missing 'Version' field in: $file"
    fi
    if ! echo "$header_content" | grep -q -E "^\*\*Last Updated:\*\*"; then
        error "Missing 'Last Updated' field in: $file"
    fi
    if ! echo "$header_content" | grep -q -E "^\*\*Status:\*\*"; then
        error "Missing 'Status' field in: $file"
    fi
    if ! echo "$header_content" | grep -q -E "^\*\*Purpose:\*\*"; then
        error "Missing 'Purpose' field in: $file"
    fi

done

# --- Final Status ---

if [ $EXIT_CODE -eq 0 ]; then
    info "Validation successful! All markdown files contain the required metadata header."
else
    error "Validation failed. Please fix the issues listed above."
fi

exit $EXIT_CODE
