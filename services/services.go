package services

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Account struct {
	ID      int    `json:"Id"`
	Name    string `json:"Name"`
	Balance int    `json:"Balance"`
}

type User struct {
	NextID   int       `json:"NextId"`
	Accounts []Account `json:"Accounts"`
}

func HomePage(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello from a HandleFunc #1!\n")
}

func GetAccounts(w http.ResponseWriter, _ *http.Request) {
	var data User

	file, err := os.Open("./accounts.json")
	check(err)
	defer file.Close()

	dat, err := ioutil.ReadAll(file)
	check(err)

	json.Unmarshal(dat, &data)

	io.WriteString(w, fmt.Sprintf("%+v", data.Accounts))
}

func AddAccounts(w http.ResponseWriter, _ *http.Request) {
	var data User

	file, err := os.Open("./accounts.json")
	check(err)
	defer file.Close()

	dat, err := ioutil.ReadAll(file)
	check(err)

	addAcc := Account{
		ID:      18,
		Name:    "Lucas Raoni",
		Balance: 1200,
	}

	json.Unmarshal(dat, &data)

	data.Accounts = append(data.Accounts, addAcc)
	data.NextID++

	bd, err := json.Marshal(data)
	check(err)

	err = os.WriteFile("./accounts.json", bd, 0644)
	check(err)

	io.WriteString(w, fmt.Sprintf("%+v", data.Accounts))
}
