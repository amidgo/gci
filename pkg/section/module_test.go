package section

import (
	"testing"

	"github.com/daixiang0/gci/pkg/specificity"
)

func TestModule(t *testing.T) {
	testCases := []specificityTestData{
		{`"foo/pkg/bar"`, Module{""}, specificity.MisMatch{}},
		{`"foo/pkg/bar"`, Module{"foo"}, specificity.Module{Length: 3}},
		{`"foo/pkg/bar"`, Module{"bar"}, specificity.MisMatch{}},
		{`"foo/pkg/bar"`, Module{"github.com/foo/bar"}, specificity.MisMatch{}},
		{`"foo/pkg/bar"`, Module{"github.com/foo"}, specificity.MisMatch{}},
		{`"foo/pkg/bar"`, Module{"github.com/bar"}, specificity.MisMatch{}},
	}
	testSpecificity(t, testCases)
}
