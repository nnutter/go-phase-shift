package main

import (
	"github.com/phasemerge/phase-shift-go/internal/analysis/nonmutating"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(nonmutating.Analyzer)
}
