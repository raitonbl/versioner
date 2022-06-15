package pkg

type Editor interface {
	GetVersionFile() string
	IsSupported(value string) bool
	EditVersion(value string, newValue string) error
}
