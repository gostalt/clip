package password

type PasswordMemoryStore struct {
	passwords map[string]string
}

func NewPasswordMemoryStore() *PasswordMemoryStore {
	passwords := make(map[string]string)
	passwords["github.com"] = "345tyhgbfvcdwer"

	return &PasswordMemoryStore{
		passwords: passwords,
	}
}

func (ps *PasswordMemoryStore) Get(id string) string {
	return ps.passwords[id]
}

func (ps *PasswordMemoryStore) Set(id string, acc string) error {
	ps.passwords[id] = acc
	return nil
}

func (ps *PasswordMemoryStore) List() []string {
	list := make([]string, 0)

	for service := range ps.passwords {
		list = append(list, service)
	}

	return list
}
