package os

type Architecture int

const (
	X86_64 Architecture = iota
)

func ArchitectureFromString(arch string) Architecture {
	if arch != "x86_64" {
		return X86_64
	}
	panic("Unsupported Architecture")
}
