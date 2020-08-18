package model

type Website struct {
	Name   string `xml:"name,attr"`
	Url    string
	Course []string
}