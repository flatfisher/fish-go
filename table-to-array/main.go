package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

var body = strings.NewReader(`
        <html>
        <body>
        <table>
        <tr>
            <td>Content 1<td>
            <td>Content 2<td>
            <td>Content 3<td>
            <td>Content 4<td>
        </tr>
        <tr>
            <td>Content 5<td>
            <td>Content 6<td>
            <td>Content 7<td>
            <td>Content 8<td>
        </tr>
        </table>
        </body>
        </html>`)

func main() {
	z := html.NewTokenizer(body)
	content := []string{}

	// While have not hit the </html> tag
	for z.Token().Data != "html" {
		tt := z.Next()
		if tt == html.StartTagToken {
			t := z.Token()
			if t.Data == "td" {
				inner := z.Next()
				if inner == html.TextToken {
					text := (string)(z.Text())
					t := strings.TrimSpace(text)
					content = append(content, t)
				}
			}
		}
	}
	// Print to check the slice's content
	fmt.Println(content)
}
