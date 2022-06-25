package main

import (
	"fmt"
)

type HelloWorldDirector[T any] struct {
	builder Builder[T]
}

func NewHelloWorldDirector[T any](b Builder[T]) HelloWorldDirector[T] {
	return HelloWorldDirector[T]{
		builder: b,
	}
}

func (d HelloWorldDirector[T]) construct() T {
	d.builder.SetTitle("Hello world!!")
	d.builder.SetLang("en")
	return d.builder.Build()
}

type Builder[T any] interface {
	SetTitle(s string)
	SetLang(s string)
	Build() T
}

type HtmlBuilder struct {
	lang  string
	title string
}

func NewHtmlBuilder() HtmlBuilder {
	return HtmlBuilder{}
}

func (b *HtmlBuilder) SetLang(s string) {
	b.lang = s
}

func (b *HtmlBuilder) SetTitle(s string) {
	b.title = s
}

func (b *HtmlBuilder) Build() Html {
	return Html{
		b.lang, b.title,
	}
}

type Html struct {
	lang  string
	title string
}

func (h Html) Output() {
	fmt.Println("<!DOCTYPE html>")
	fmt.Println("<html lang=\"" + h.lang + "\">")
	fmt.Println("<head>")
	fmt.Println("<meta charset=\"utf-8\">")
	fmt.Println("<title>" + h.title + "</title>")
	fmt.Println("</head>")
	fmt.Println("<body>")
	fmt.Println("<h1>" + h.title + "</h1>")
	fmt.Println("</body>")
	fmt.Println("</html>")
}

func main() {
	b := NewHtmlBuilder()
	NewHelloWorldDirector[Html](&b).construct().Output()
}
