package pkg

type Reader interface {
	IsSupported(value string) bool
	ReadVersion(value string) (string, error)
	GetVersionFile() string
}
