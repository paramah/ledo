#!/usr/bin/env bash

set -euo pipefail

echo "🚀 Odświeżam Go vendora i aktualizuję flake.nix..."

# 1. Wyczyść i odśwież Go vendor
go mod tidy
go mod vendor

# 2. Policzenie aktualnego vendorHash
hash=$(nix hash path ./vendor)

echo "📦 Obliczony vendorHash: $hash"

# 3. Aktualizacja flake.nix
echo "✍️ Aktualizuję flake.nix..."

# Zastąp istniejący vendorHash w flake.nix
sed -i.bak "s|vendorHash = .*;|vendorHash = \"$hash\";|" flake.nix

# Usuń backup
rm -f flake.nix.bak

echo "✅ Flake.nix zaktualizowany!"
echo "✅ Gotowe! Teraz możesz robić: nix build"
