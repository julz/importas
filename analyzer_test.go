package importas

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()

	testCases := []struct {
		desc    string
		pkg     string
		aliases stringMap
	}{
		{
			desc: "Valid imports",
			pkg:  "a",
			aliases: stringMap{
				"fmt": "fff",
				"os":  "stdos",
			},
		},
		{
			desc: "Invalid imports",
			pkg:  "b",
			aliases: stringMap{
				"fmt": "fff",
				"os":  "stdos",
			},
		},
		{
			desc: "external libs",
			pkg:  "c",
			aliases: stringMap{
				"knative.dev/serving/pkg/apis/autoscaling/v1alpha1": "autoscalingv1alpha1",
				"knative.dev/serving/pkg/apis/serving/v1":           "servingv1",
			},
		},
		{
			desc: "regexp",
			pkg:  "d",
			aliases: stringMap{
				"knative.dev/serving/pkg/apis/(\\w+)/(v[\\w\\d]+)": "$1$2",
			},
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {

			dir := filepath.Join(testdata, "src", test.pkg)

			if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
				cmd := exec.Command("go", "mod", "vendor")
				cmd.Dir = dir

				t.Cleanup(func() {
					_ = os.RemoveAll(filepath.Join(testdata, "src", test.pkg, "vendor"))
				})

				if output, err := cmd.CombinedOutput(); err != nil {
					t.Fatal(err, string(output))
				}
			}

			a := Analyzer

			flg := a.Flags.Lookup("alias")
			for k, v := range test.aliases {
				err := flg.Value.Set(fmt.Sprintf("%s:%s", k, v))
				if err != nil {
					t.Fatal(err)
				}
			}

			analysistest.RunWithSuggestedFixes(t, testdata, Analyzer, test.pkg)
		})
	}
}
