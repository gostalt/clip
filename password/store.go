package password

// TODO:
// - History
// - 2FA (depends on the implementation? No idea how it works...)

type PasswordStore interface {
	Get(id string) string
	Set(id string, password string) error
	List() []string
}
