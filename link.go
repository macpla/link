package link

import (
	"io"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/net/html"
)

// Link represents HTML link tag < a href="...">
type Link struct {
	Href string
	Text string
}

// Parse will take a HTML document and will return a slice of links
// parsed from it.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse html document")
	}

	nodes := linkNodes(doc)
	var links []Link
	for _, n := range nodes {
		links = append(links, newLink(n))
	}

	return links, nil
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}

func newLink(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	var txt string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		txt += text(c)
	}
	ret.Text = txt
	return ret
}

func text(n *html.Node) (txt string) {
	if n.Type == html.TextNode {
		txt += strings.Join(strings.Fields(n.Data), " ")
	} else {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			txt += " " + text(c)
		}
	}
	return txt
}
