package section

import (
	"testing"

	"github.com/daixiang0/gci/pkg/specificity"
)

func TestPrefixSpecificity(t *testing.T) {
	testCases := []specificityTestData{
		{`"foo/pkg/bar"`, Custom{""}, specificity.MisMatch{}},
		{`"foo/pkg/bar"`, Custom{"foo"}, specificity.Match{Length: 3}},
		{`"foo/pkg/bar"`, Custom{"bar"}, specificity.MisMatch{}},
		{`"foo/pkg/bar"`, Custom{"github.com/foo/bar"}, specificity.MisMatch{}},
		{`"foo/pkg/bar"`, Custom{"github.com/foo"}, specificity.MisMatch{}},
		{`"foo/pkg/bar"`, Custom{"github.com/bar"}, specificity.MisMatch{}},
	}
	testSpecificity(t, testCases)
}
