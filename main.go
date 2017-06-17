package main

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

const (
	prefix = "/keybase/private/%s/credstore"
)

func menu() {

	menu := `
A key, value store for saving and loading secrets in keybase.

Usage:
  kb COMMAND

Commands:
  init             creates a keystore
  set <key>        save a key
  get <key>        loads value of a key
  generate <key>   auto generates a 12 character random value
  search <string>  lists all keys with matching substring
  ls               lists all available keys
  rm <key>         removes a key

Example:
  - set key github.com
      kb set github.com

  - get key github.com
      kb get github.com

	`

	fmt.Printf(menu)
}

func rm(key string) error {
	c, err := credstore()

	if err != nil {
		return err
	}

	err = os.Remove(strings.Join([]string{c, "/", key}, ""))

	if err != nil {
		return err
	}

	return nil
}

func list() error {
	c, err := credstore()

	if err != nil {
		return err
	}

	files, _ := ioutil.ReadDir(c)
	fmt.Printf("\nAvalable keys: \n")
	for _, f := range files {
		fmt.Printf("\t%s", f.Name())
	}
	fmt.Println()

	return nil
}

func encrypt(s string) ([]byte, error) {

	u, err := user()
	if err != nil {
		return []byte(""), err
	}

	args := []string{"encrypt", "-m", s, u}

	out, err := exec.Command("keybase", args...).Output()

	if err != nil {
		return []byte(""), err
	}

	return out, nil
}

func decrypt(b []byte) (string, error) {

	msg := fmt.Sprintf("%s", string(b))
	s := strings.Replace(msg, "\n", "", -1)
	args := []string{"decrypt", "-m", s}
	out, err := exec.Command("keybase", args...).Output()

	if err != nil {
		return "", err
	}

	return string(out), nil
}

func search(s string) error {
	c, err := credstore()

	if err != nil {
		return err
	}

	files, _ := ioutil.ReadDir(c)
	fmt.Printf("\nAvalable keys: \n")
	for _, f := range files {

		if strings.Contains(f.Name(), s) {
			fmt.Printf("\t%s", f.Name())
		}
	}
	fmt.Println()

	return nil

}

func generate(key string) error {
	c := 12
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		return err
	}

	val := base64.StdEncoding.EncodeToString(b)

	err = save(key, val)

	if err != nil {
		return err
	}

	fmt.Printf("\n\tSaving %s", key)

	return nil

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

	ctxt, err := encrypt(val)

	if err != nil {
		return err
	}

	path := strings.Join([]string{c, "/", key}, "")
	err = ioutil.WriteFile(path, ctxt, 0600)

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

	txt, err := decrypt(b)

	if err != nil {
		return "", err
	}

	return txt, nil
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
			fmt.Println("Done! Type kb to see available commands.")

		case "set":

			if len(args) < 2 {
				fmt.Println("Please provide a key to save.")
				return
			}
			fmt.Print("\n\tEnter Value: ")
			bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil {
				log.Fatal(err)
			}
			password := string(bytePassword)
			val := strings.Trim(strings.Trim(password, "\n"), " ")

			err = save(args[1], val)

			if err != nil {
				fmt.Println("Saving failed")
				log.Fatal(err)
			}

			fmt.Printf("\n")

		case "get":
			if len(args) < 2 {
				fmt.Println("Please provide a key to retrieve")
				return
			}

			val, err := get(args[1])

			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("\n\t%s\n", val)
		case "ls":

			err := list()

			if err != nil {
				log.Fatal(err)
			}

		case "search":

			if len(args) < 2 {
				fmt.Println("please provide a key to search.")
				return
			}

			err := search(args[1])

			if err != nil {
				log.Fatal(err)
			}
		case "generate":

			if len(args) < 2 {
				fmt.Println("please provide a key to search.")
				return
			}

			err := generate(args[1])
			if err != nil {
				log.Fatal(err)
			}

		case "rm":

			if len(args) < 2 {
				fmt.Println("please provide a key to remove.")
				return
			}

			err := rm(args[1])
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
