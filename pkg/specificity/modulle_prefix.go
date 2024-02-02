package specificity

import "fmt"

type ModulePrefix struct {
	Length int
}

func (m ModulePrefix) IsMoreSpecific(than MatchSpecificity) bool {
	otherMatch, isMatch := than.(ModulePrefix)
	return isMoreSpecific(m, than) || (isMatch && m.Length > otherMatch.Length)
}

func (m ModulePrefix) Equal(to MatchSpecificity) bool {
	return equalSpecificity(m, to)
}

func (m ModulePrefix) class() specificityClass {
	return ModulePrefixClass
}

func (m ModulePrefix) String() string {
	return fmt.Sprintf("ModulePrefix(length: %d)", m.Length)
}
