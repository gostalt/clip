package password

// TODO:
// - History
// - 2FA (depends on the implementation? No idea how it works...)

type PasswordStore interface {
	Get(id string) Account
	Set(id string, acc Account) error
	List() []string
}
