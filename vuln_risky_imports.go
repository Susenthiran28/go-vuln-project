package main

// Best-effort blank imports for modules with uncertain root package paths
// or environment dependencies. If `go mod tidy` fails on any of these,
// delete this entire file — the other vuln_*.go files remain usable.

import (
	_ "code.gitea.io/gitea/modules/log"
)
