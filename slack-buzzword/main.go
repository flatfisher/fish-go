package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
	"unicode/utf8"

	"github.com/bluele/mecab-golang"
)

type Message struct {
	Text string `json:"text"`
}

type Channel struct {
	Messages []Message `json:"messages"`
}
type Buzzword struct {
	Word  string
	Count int
}

type Buzzwords []Buzzword

func (p Buzzwords) Len() int {
	return len(p)
}

func (p Buzzwords) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type ByCount struct {
	Buzzwords
}

func (b ByCount) Less(i, j int) bool {
	return b.Buzzwords[i].Count > b.Buzzwords[j].Count
}

func main() {
	// 控えたURLを入れる
	url := ""

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		return
	}
	defer res.Body.Close()

	var c Channel

	m, err := mecab.New("-Owakati")
	if err != nil {
		fmt.Printf("Mecab instance error. err: %v", err)
	}
	defer m.Destroy()
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	var tmp []string
	for _, msg := range c.Messages {
		parseToNode(m, msg.Text, &tmp)
	}

	var buzzwords []Buzzword
	for word, count := range wordCount(tmp) {
		buzzwords = append(buzzwords, Buzzword{word, count})
	}
	sort.Sort(ByCount{buzzwords})

	for _, s := range buzzwords {
		fmt.Println(s.Word, "=>", s.Count)
	}
}

func parseToNode(m *mecab.MeCab, s string, t *[]string) {
	tg, err := m.NewTagger()
	if err != nil {
		panic(err)
	}
	defer tg.Destroy()
	lt, err := m.NewLattice(s)
	if err != nil {
		panic(err)
	}
	defer lt.Destroy()

	node := tg.ParseToNode(lt)
	for {
		features := strings.Split(node.Feature(), ",")
		if utf8.RuneCountInString(features[6]) > 2 {
			if features[0] == "名詞" {
				*t = append(*t, features[6])
			}
		}
		if node.Next() != nil {
			break
		}
	}
}

func wordCount(a []string) map[string]int {
	c := make(map[string]int)
	for _, word := range a {
		_, ok := c[word]
		if ok {
			c[word]++
		} else {
			c[word] = 1
		}
	}
	return c
}
