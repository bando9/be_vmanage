package controllers

import (
	"bevmanage/internal/models"

	"github.com/sev-2/raiden"
)

type BooksController struct {
	raiden.ControllerBase
	Http  string `path:"/books" type:"rest"`
	Model models.Books
}