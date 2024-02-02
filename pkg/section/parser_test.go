package section

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type sectionTestData struct {
	input           []string
	expectedSection SectionList
	expectedError   error
	goModulePath    string
}

func TestParse(t *testing.T) {
	testCases := []sectionTestData{
		{
			input:           []string{""},
			expectedSection: nil,
			expectedError:   nil,
		},
		{
			input:           []string{"prefix(go)"},
			expectedSection: SectionList{Custom{"go"}},
			expectedError:   nil,
		},
		{
			input:           []string{"prefix(go-UPPER-case)"},
			expectedSection: SectionList{Custom{"go-UPPER-case"}},
			expectedError:   nil,
		},
		{
			input:           []string{"PREFIX(go-UPPER-case)"},
			expectedSection: SectionList{Custom{"go-UPPER-case"}},
			expectedError:   nil,
		},
		{
			input:           []string{"prefix("},
			expectedSection: nil,
			expectedError:   errors.New("invalid params: prefix("),
		},
		{
			input:           []string{"prefix(domainA,domainB)"},
			expectedSection: SectionList{Custom{"domainA,domainB"}},
			expectedError:   nil,
		},
		{
			input: []string{"module"},
			expectedSection: SectionList{
				Module{Pkg: "github.com/hexops/gotextdiff"},
				Module{Pkg: "github.com/pmezard/go-difflib"},
				Module{Pkg: "github.com/spf13/cobra"},
				Module{Pkg: "github.com/stretchr/testify"},
				Module{Pkg: "go.uber.org/zap"},
				Module{Pkg: "golang.org/x/mod"},
				Module{Pkg: "golang.org/x/sync"},
				Module{Pkg: "golang.org/x/tools"},
				Module{Pkg: "gopkg.in/yaml.v3"},
			},
			expectedError: nil,
			goModulePath:  "testdata/go.mod.test",
		},
		{
			input:           []string{"module"},
			expectedSection: nil,
			expectedError:   errors.New("invalid params: failed parse go.mod file, wrong format"),
			goModulePath:    "testdata/go.mod.wrong_format.test",
		},
		{
			input:           []string{"module"},
			expectedSection: nil,
			expectedError:   errors.New("invalid params: go module file not found"),
		},
	}
	for _, test := range testCases {
		parsedSection, err := Parse(test.input, test.goModulePath)
		assert.Equal(t, test.expectedSection, parsedSection)
		assert.Equal(t, test.expectedError, err)
	}
}
