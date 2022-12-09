package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type book struct {
	ID       string `json:"id"` //when serialising with json, convert these into lower case
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}
