package nonmutating_test

import (
	"testing"

	"github.com/phasemerge/phase-shift-go/internal/analysis/nonmutating"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), nonmutating.Analyzer, "a")
}
