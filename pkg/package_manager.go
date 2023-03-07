package pkg

type PackageManager interface {
	GetSupportTypes() []string
	GetVersion(object string, filename string) (string, error)
	SetVersion(object string, filename string, value string) error
}
