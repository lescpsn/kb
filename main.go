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

func user() (string, error) {
	home := os.Getenv("HOME")
	u := strings.Join([]string{home, "/.kbpass/username"}, "")
	s, err := ioutil.ReadFile(u)

	if err != nil {
		return "", err
	}

	return string(s), nil
}

// func credstore() string

func createKeystore() error {
	home := os.Getenv("HOME")

	// read user's keybase name
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Keybase username: ")
	text, _ := reader.ReadString('\n')
	f := strings.Trim(strings.Trim(text, "\n"), " ")

	// save user's keybase username in ~/.kbpass
	// for accessing private keybase folder later
	dir := strings.Join([]string{home, "/.kbpass"}, "")
	err := os.MkdirAll(dir, 0700)

	if err != nil {
		return err
	}

	file := strings.Join([]string{dir, "/username"}, "")

	fmt.Println("\n\tWriting keybase username to ~/.kbpass/username")
	err = ioutil.WriteFile(file, []byte(f), 0600)

	if err != nil {
		return err
	}

	// make keystore in keybase private folder
	path := fmt.Sprintf(prefix, f)
	fmt.Printf("\tBuilding keystore in %s\n\n", path)
	err = os.Mkdir(path, 0700)

	if err != nil {
		return err
	}

	return nil

}

func main() {
	args := os.Args[1:]

	switch args[0] {
	case "init":
		err := createKeystore()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Done! use 'kbpass save <identifier> <password>' to save your first password. ")

	case "save":
		fmt.Println("saving so hard")
	}

	u, err := user()

	if err != nil {
		log.Fatal(err)
	}
	log.Printf(u)
}
