package password

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func initialize(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Println("File does not exist")
		os.Create(path)
	}
}

type PasswordJsonStore struct {
	passwords map[string]Account
}

func NewPasswordJsonStore() *PasswordJsonStore {
	passwords := make(map[string]Account)

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("Unable to find password repository")
	}

	path := home + "/.pw"
	initialize(path)

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln("Unable to read contents of password file.")
	}
	json.Unmarshal(content, &passwords)

	return &PasswordJsonStore{
		passwords: passwords,
	}
}

func (ps *PasswordJsonStore) Get(id string) Account {
	return ps.passwords[id]
}

func (ps *PasswordJsonStore) Set(id string, acc Account) error {
	ps.passwords[id] = acc

	return ps.save()
}

func (ps *PasswordJsonStore) List() []string {
	list := make([]string, 0)

	for service := range ps.passwords {
		list = append(list, service)
	}

	return list
}

func (ps *PasswordJsonStore) save() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	f, err := os.Create(home + "/.pw")
	if err != nil {
		return err
	}

	content, err := json.Marshal(ps.passwords)
	if err != nil {
		log.Fatalln("Cannot save JSON")
	}

	if _, err = f.Write(content); err != nil {
		log.Fatalln("Unable to save passwords")
	}

	return nil
}
