package password

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type PasswordFileStore struct {
	passwords map[string]string
}

// initialize ensures a password file exists.
func initialize(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Println("File does not exist")
		os.Create(path)
	}
}

func NewPasswordFileStore() *PasswordFileStore {
	passwords := make(map[string]string)

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("Unable to find password repository")
	}

	path := home + "/.pw"

	initialize(path)

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ",")
		passwords[split[0]] = split[1]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &PasswordFileStore{
		passwords: passwords,
	}
}

func (ps *PasswordFileStore) Get(id string) string {
	r := ps.passwords[id]

	if r == "" {
		fmt.Println("❌ No record for", id)
		os.Exit(1)
	}

	return ps.passwords[id]
}

func (ps *PasswordFileStore) Set(id string, acc string) error {
	ps.passwords[id] = acc

	return ps.save()
}

func (ps *PasswordFileStore) List() []string {
	list := make([]string, 0)

	for service := range ps.passwords {
		list = append(list, service)
	}

	return list
}

func (ps *PasswordFileStore) save() error {
	f, err := os.Create("./pw")
	if err != nil {
		return err
	}

	for acc, pw := range ps.passwords {
		_, err := fmt.Fprintf(f, "%s,%s\n", acc, pw)
		if err != nil {
			return err
		}
	}

	return nil
}
