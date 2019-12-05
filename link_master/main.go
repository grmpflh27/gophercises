package link

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	htmlFile := flag.String("file", "ex1.html", "file you'd wish to parse")
	flag.Parse()

	fmt.Printf("Parsing %v\n", *htmlFile)
	fp, err := os.Open(*htmlFile)

	if err != nil {
		log.Fatal("Could not open ", *htmlFile)
	}

	rootNode, err := html.Parse(fp)
	if err != nil {
		log.Fatal("Could not parse")
	}

	allLinkNodes := getAllLinkNodes(rootNode)
	fmt.Printf("Found %v links\n", len(allLinkNodes))

	getLinks(allLinkNodes)
}

func getAllLinkNodes(curNode *html.Node) []*html.Node {

	if isLink(curNode) {
		return []*html.Node{curNode}
	}

	var ret []*html.Node
	for c := curNode.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, getAllLinkNodes(c)...)
	}

	return ret
}

func isLink(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "a"
}

func getLinks(allLinkNodes []*html.Node) {
	var links []Link
	for _, node := range allLinkNodes {
		href, err := getHref(node)
		if err != nil {
			log.Fatal("ERRRRR")
		}
		links = append(links, Link{href, getAllText(node)})
	}

	fmt.Printf("%+v\n", links)
}

func getHref(n *html.Node) (string, error) {
	for _, a := range n.Attr {
		if a.Key == "href" {
			return a.Val, nil
		}
	}
	return "", errors.New("no dice")
}

func getAllText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}

	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += getAllText(c)
	}

	return ret
}

type Link struct {
	Href string
	Text string
}
