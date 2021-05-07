#!/usr/bin/env bash
set -Eeuxo pipefail
cd "$(dirname ${BASH_SOURCE[0]})"


cat head.md > ../README.md
cat ../main.go >> ../README.md
cat foot.md >> ../README.md
