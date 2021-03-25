package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func DCIM_req(Server string, action string) []byte {

	var username string
	var password string

	if _, err := os.Stat("/Users/sgorini/.credential"); os.IsNotExist(err) {

		//Collectin credential to be used for future curl call
		username, password, _ := Get_credentials()
		fmt.Printf("Username: %s, Password: %s\n", username, password)

	} else {

		file, err := os.Open("/Users/sgorini/.credential")
		check(err)
		defer file.Close()
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)
		count := 1
		for scanner.Scan() {
			if count == 1 {
				username = scanner.Text()
			}
			if count == 2 {
				password = scanner.Text()
			}
			count++
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		log.Println("letto da file")
	}

	var url string = "https://" + Server + "/api/v1/" + action
	//fmt.Println(url)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(username, password)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		log.Println("HTTP Status is in the 2xx range")
	} else {
		log.Println("Argh! Broken")
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))
	return body
}

func Get_credentials() (string, string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}

	fmt.Print("Enter Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", "", err
	}
	password := string(bytePassword)
	return strings.TrimSpace(username), strings.TrimSpace(password), nil
}
