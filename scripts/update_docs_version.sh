#!/bin/bash

# Script to update all NFT documentation to align with v8.0.0 business rules
# This updates version headers and key references across all documentation

echo "Updating AIW3 NFT documentation to v8.0.0..."

# Find all markdown files in docs directory
find /home/zealy/aiw3/aiw3-nft-solana/docs -name "*.md" -type f | while read file; do
    echo "Processing: $file"
    
    # Update version headers from various older versions to v8.0.0
    sed -i 's/\*\*Version:\*\* v[0-9]\+\.[0-9]\+\.[0-9]\+/\*\*Version:\*\* v8.0.0/g' "$file"
    
    # Update Last Updated to current date
    sed -i 's/\*\*Last Updated:\*\* [0-9]\{4\}-[0-9]\{2\}-[0-9]\{2\}/\*\*Last Updated:\*\* 2025-08-09/g' "$file"
    
    # Update business rules version references
    sed -i 's/business rules v[0-9]\+\.[0-9]\+\.[0-9]\+/business rules v8.0.0/g' "$file"
    sed -i 's/v[0-9]\+\.[0-9]\+\.[0-9]\+ business rules/v8.0.0 business rules/g' "$file"
    sed -i 's/AIW3-NFT-Business-Rules-and-Flows\.md** v[0-9]\+\.[0-9]\+\.[0-9]\+/AIW3-NFT-Business-Rules-and-Flows.md** v8.0.0/g' "$file"
    
done

echo "Documentation version update completed!"
echo "All documents now reference v8.0.0 business rules."
