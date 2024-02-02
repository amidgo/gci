package specificity

import "fmt"

type Module struct {
	Length int
}

func (m Module) IsMoreSpecific(than MatchSpecificity) bool {
	otherMatch, isMatch := than.(Module)
	return isMoreSpecific(m, than) || (isMatch && m.Length > otherMatch.Length)
}

func (m Module) Equal(to MatchSpecificity) bool {
	return equalSpecificity(m, to)
}

func (m Module) class() specificityClass {
	return ModuleClass
}

func (m Module) String() string {
	return fmt.Sprintf("Module(length: %d)", m.Length)
}
