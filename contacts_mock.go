package main

import "errors"

type contactsMock struct {}


func NewContactsMock() (Contacts, error) {
	return &contactsMock{}, nil
}

func (contacts *contactsMock) GetIdMap() (IdMap, error) {
	idMap := make(map[string]string)
	idMap["a:0"] = "0"
	idMap["a:1"] = "1"
	idMap["b:0"] = "0"
	return idMap, nil
}

func (contacts *contactsMock) GetContact(id string) (Contact, error) {
	if id == "0" {
		return Contact{
			id: "0",
			name: "Johnny",
			tags: []string{"friends"},
		}, nil
	}
	if id == "1" {
		return Contact{
			id: "1",
			name: "Andrew",
			tags: []string{"friends"},
		}, nil
	}
	return Contact{}, errors.New("invalid contact id")
}