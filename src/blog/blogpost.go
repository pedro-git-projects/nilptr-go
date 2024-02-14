package blog

import (
	"html/template"
	"time"
)

type BlogPost struct {
	Title       string
	Content     template.HTML
	Day         int
	Month       time.Month
	Year        int
	Description string
}
