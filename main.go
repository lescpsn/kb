package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	prefix = "/keybase/private/%s/credstore"
)

func menu() {

	menu := `
An encrypted key, value store for saving and loading secrets.

Usage:
  kb COMMAND

Commands:
  init             creates a keystore
  save <key>       save a key
  get <key>        loads value of a key
  generate <key>   auto generates a 12 character random value
  list             lists all available keys

Example:
  - save key github.com
      kb save github.com

  - get key github.com
      kb get github.com

	`

	fmt.Printf(menu)
}

func list() error {
	c, err := credstore()

	if err != nil {
		return err
	}

	files, _ := ioutil.ReadDir(c)
	for _, f := range files {
		fmt.Println(f.Name())
	}
}

func user() (string, error) {
	home := os.Getenv("HOME")
	u := strings.Join([]string{home, "/.kbpass/username"}, "")
	s, err := ioutil.ReadFile(u)

	if err != nil {
		return "", err
	}

	return string(s), nil
}

func credstore() (string, error) {
	u, err := user()

	if err != nil {
		return "", err
	}

	return fmt.Sprintf(prefix, u), nil
}

func save(key, val string) error {
	c, err := credstore()

	if err != nil {
		return err
	}

	path := strings.Join([]string{c, "/", key}, "")
	err = ioutil.WriteFile(path, []byte(val), 0600)

	if err != nil {
		return err
	}

	return nil

}

func get(key string) (string, error) {

	c, err := credstore()

	if err != nil {
		return "", err
	}

	path := strings.Join([]string{c, "/", key}, "")
	b, err := ioutil.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(b), nil
}

func create() error {
	home := os.Getenv("HOME")

	// read user's keybase name
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Keybase username: ")
	text, _ := reader.ReadString('\n')
	f := strings.Trim(strings.Trim(text, "\n"), " ")

	// save user's keybase username in ~/.kbpass
	// for accessing private keybase folder later
	dir := strings.Join([]string{home, "/.kbpass"}, "")
	_ = os.Mkdir(dir, 0700)

	file := strings.Join([]string{dir, "/username"}, "")

	fmt.Println("\n\tWriting keybase username to ~/.kbpass/username")
	err := ioutil.WriteFile(file, []byte(f), 0600)

	if err != nil {
		return err
	}

	// make keystore in keybase private folder
	path := fmt.Sprintf(prefix, f)
	fmt.Printf("\tBuilding keystore in %s\n\n", path)
	_ = os.Mkdir(path, 0700)

	return nil

}

func main() {

	if len(os.Args) > 1 {
		args := os.Args[1:]

		switch args[0] {
		case "init":
			err := create()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Done! use 'kbpass save <identifier> <password>' to save your first password. ")

		case "save":

			if len(args) < 2 {
				fmt.Println("Please provide a key to save.")
			}

			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Value: ")
			text, _ := reader.ReadString('\n')
			val := strings.Trim(strings.Trim(text, "\n"), " ")

			err := save(args[1], val)

			if err != nil {
				fmt.Println("Saving failed")
				log.Fatal(err)
			}

			fmt.Println("Saved.")

		case "get":
			if len(args) < 2 {
				fmt.Println("Please provide a key to retrieve")
			}

			val, err := get(args[1])

			if err != nil {
				fmt.Println("Invalid key!")
				return
			}

			fmt.Printf("%s:\n\t%s\n", args[1], val)
		case "list":

			err := list()

			if err != nil {
				log.Fatal(err)
			}

		default:

			menu()
		}
	} else {
		menu()
	}
}
