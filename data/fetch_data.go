package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var baseURL = "https://reqres.in/api/users/"

// User represents User data
type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

// SingleData represents a single User
type SingleData struct {
	User User `json:"data"`
}

// CollectionData represents collection of Users
type CollectionData struct {
	Users []User `json:"data"`
}

//FetchUser fetchs a single user by its id
func FetchUser(id int) (*User, error) {
	client := &http.Client{}
	URL := fmt.Sprintf("%s%d", baseURL, id)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	//res, err := client.Get(URL)
	if err != nil {
		return nil, err
	}
	userdata := &SingleData{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, userdata)
	if err != nil {
		return nil, err
	}
	return &userdata.User, nil
}

// FetchUsers fetchs all users on a given page
func FetchUsers(page int) ([]User, error) {
	client := &http.Client{}
	URL := fmt.Sprintf("%s?page=%d", baseURL, page)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	//res, err := client.Get(URL)
	if err != nil {
		return nil, err
	}
	usdata := &CollectionData{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, usdata)
	if err != nil {
		return nil, err
	}
	return usdata.Users, nil
}
