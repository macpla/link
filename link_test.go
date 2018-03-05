package link

import (
	"log"
	"reflect"
	"strings"
	"testing"
)

var tests = []struct {
	input    string
	expected []Link
}{
	{
		ex1HTML,
		[]Link{{Href: "/other-page", Text: "A link to another page"}},
	},
	{
		ex2HTML,
		[]Link{
			{Href: "https://www.twitter.com/joncalhoun", Text: "Check me out on twitter"},
			{Href: "https://github.com/gophercises", Text: "Gophercises is on Github!"},
		},
	},
	{
		ex3HTML,
		[]Link{
			{Href: "#", Text: "Login"},
			{Href: "/lost", Text: "Lost? Need help?"},
			{Href: "https://twitter.com/marcusolsson", Text: "@marcusolsson"},
		},
	},
	{
		ex4HTML,
		[]Link{{Href: "/dog-cat", Text: "dog cat"}},
	},
}

func TestParse(*testing.T) {
	for _, t := range tests {
		actual, err := Parse(strings.NewReader(t.input))
		if err != nil {
			log.Fatal(err)
		}
		if !reflect.DeepEqual(actual, t.expected) {
			log.Fatalf("expected %+v, got %+v", t.expected, actual)
		}

	}
}
