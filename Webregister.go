package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var email string

func webRegister() {
	var begin string
	fmt.Println("hallo gebruiker, wil je inloggen of registeren, druk op 1 voor inloggen, druk op 2 voor registreren")
	fmt.Scanln(&begin)
	if begin == "1" {
		login()
	} else if begin == "2" {
		register()
	} else {
		return
	}
}

type feedbacklogin struct {
	Msg   string `json:"msg"` //Hier wordt de teruggestuurde data opgeslagen
	Token string `json:"token"`
}

func login() {

	var wachtwoord string
	fmt.Println("Hallo gebruiker, wat is je email?")
	fmt.Scanln(&email)
	fmt.Println("Wat is je wachtwoord?")
	fmt.Scanln(&wachtwoord)

	url := "http://localhost:8080/api/signin/"
	method := "POST"

	payload := strings.NewReader(`{
	  "email":"` + email + `",
	  "password":"` + wachtwoord + `"
  }`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(string(body))

	var data feedbacklogin
	json.Unmarshal(body, &data)
	//fmt.Println(data.Msg) //hier wordt de data opgeroepen uit de struct
	if data.Msg == "Successfully SignIN" {
		fmt.Println("Welkom " + email)
	} else {
		fmt.Println("dit is niet goed")
		os.Exit(0)
	}

}

type feedbackregister struct {
	Id    string `json:"ID"`
	Maken string `json:"CreatedAt"`
}

func register() {

	var naam string
	var email string
	var password string

	fmt.Println("Welkom bij je registratie, geef eerst je naam")
	fmt.Scanln(&naam)
	fmt.Println("geef nu je email")
	fmt.Scanln(&email)
	fmt.Println("Geef nu een wachtwoord")
	fmt.Scanln(&password)

	url := "http://localhost:8080/api/register/" //Waar de data heen moet
	method := "POST"                             //Wat er moet gebeuren met de data

	payload := strings.NewReader(`{				
	  "name":"` + naam + `",
	  "role":"gebruiker",
	  "email":"` + email + `",
	  "password":"` + password + `"
  
  }`) //De geleverde informatie

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

	var data feedbackregister
	json.Unmarshal(body, &data)
	fmt.Println(data.Id)
	fmt.Println(data.Maken)
}
