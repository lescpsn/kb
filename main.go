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

var (
	home   = os.Getenv("HOME")
	prefix = fmt.Sprintf("%s/.kb", home)
)

// menu prints available commands
func menu() {

	menu := `Usage:
  kb COMMAND

Commands:
  init             creates a keystore
  set <key>        save a key
  get <key>        loads value of a key
  generate <key>   generates & saves 12 character
                   random value for a key
  search <string>  lists all keys with substring
  ls               lists all available keys
  rm <key>         removes a key

Example:
  - set the key github.com
      kb set github.com

  - get value of the key github.com
      kb get github.com
`

	fmt.Printf(menu)
}

// rm deletes a key
func rm(key string) error {
	err := os.Remove(strings.Join([]string{prefix, "/", key}, ""))

	if err != nil {
		return err
	}

	return nil
}

// list prints all available keys
func list() error {

	files, _ := ioutil.ReadDir(prefix)
	fmt.Printf("\nAvalable keys:\n\n")
	for _, f := range files {
		if f.Name() != "username" {
			fmt.Printf("    %s\n", f.Name())
		}
	}
	fmt.Println()

	return nil
}

// encrypt passes a string to the keybase cli, returning encrypted bytes
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

// decrypt decrypts ciphertext with the keybase cli, returning plaintext
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

// search filters keys by a provided substring
func search(s string) error {

	files, _ := ioutil.ReadDir(prefix)
	fmt.Printf("\nAvalable keys: \n")
	for _, f := range files {

		if strings.Contains(f.Name(), s) {
			fmt.Printf("\t%s", f.Name())
		}
	}
	fmt.Println()

	return nil

}

// generate creates and saves a 12-character cryptographicallyrandom string
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

	fmt.Printf("\n\tSaving %s\n", key)

	return nil

}

// user fetches the caller's keybase username
func user() (string, error) {
	home := os.Getenv("HOME")
	u := strings.Join([]string{home, "/.kb/username"}, "")
	s, err := ioutil.ReadFile(u)

	if err != nil {
		return "", err
	}

	return string(s), nil
}

// save encrypts a value and saves the ciphertext at $HOME/.kb/<key>
func save(key, val string) error {

	ctxt, err := encrypt(val)

	if err != nil {
		return err
	}

	path := strings.Join([]string{prefix, "/", key}, "")
	err = ioutil.WriteFile(path, ctxt, 0600)

	if err != nil {
		return err
	}

	return nil

}

// get reads the ciphertext of a key and prints it to the console
func get(key string) (string, error) {

	path := strings.Join([]string{prefix, "/", key}, "")
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

// create makes a keystore in $HOME/.kb
func create() error {
	home := os.Getenv("HOME")

	// read user's keybase name
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Keybase username: ")
	text, _ := reader.ReadString('\n')

	// remove whitespace and newlines
	f := strings.Trim(strings.Trim(text, "\n"), " ")

	// save user's keybase username in ~/.kb
	// for accessing private keybase folder later
	dir := strings.Join([]string{home, "/.kb"}, "")
	_ = os.Mkdir(dir, 0700)

	file := strings.Join([]string{dir, "/username"}, "")

	fmt.Println("\n\tWriting keybase username to ~/.kb/username")
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
				fmt.Println("please provide a key to create.")
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
