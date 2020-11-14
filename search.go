package main

import (
	"fmt"
	"io/ioutil"
)

type Article struct {
	Title string
	Description []byte
	Image []byte
}

func (a *Article) save() error {
	filename := a.Title + ".txt"
	return ioutil.WriteFile(filename, a.Description, 0600)
}

func readArticle(title string) (*Article, error) {
	filename := title + ".txt"
	description, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Article{Title: title, Description: description}, nil
}

func main() {
	article := &Article{
		Title: "conceived",
		Description: []byte("It's about something you may be familiar with")}
	article.save()

	justRead, err := readArticle("conceived")
	if err != nil {
		fmt.Println("Something was wrong")
	}

	fmt.Println(justRead.Title)
	fmt.Println(string(justRead.Description))
}