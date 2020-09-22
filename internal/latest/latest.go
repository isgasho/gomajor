package latest

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/andybalholm/cascadia"
	"golang.org/x/mod/semver"
	"golang.org/x/net/html"
)

// Version returns the latest version of the package
func Version(pkgpath string) (string, error) {
	vv, err := Versions(pkgpath)
	if err != nil {
		return "", err
	}
	var newest string
	for _, s := range vv {
		if !semver.IsValid(s) || semver.Prerelease(s) != "" {
			continue
		}
		if newest == "" {
			newest = s
		}
		if semver.Compare(s, newest) > 0 {
			newest = s
		}
	}
	if newest == "" {
		return "", errors.New("no valid versions")
	}
	return newest, nil
}

// Versions returns all versions of a package
func Versions(pkgpath string) ([]string, error) {
	url := fmt.Sprintf("https://pkg.go.dev/%s?tab=versions", pkgpath)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	// extract versions from the html
	var versions []string
	doc, err := html.Parse(res.Body)
	if err != nil {
		return nil, err
	}
	sel := cascadia.MustCompile(".Versions-item>a")
	for _, node := range cascadia.QueryAll(doc, sel) {
		walk(node, func(n *html.Node) {
			if n.Type == html.TextNode {
				versions = append(versions, n.Data)
			}
		})
	}
	return versions, nil
}

// walk the node using depth first
func walk(n *html.Node, f func(*html.Node)) {
	f(n)
	if n.FirstChild != nil {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c, f)
		}
	}
}
