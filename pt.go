package main

import (
	"bytes"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"text/template"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/feeds"
	log "github.com/sirupsen/logrus"
	"gopkg.in/russross/blackfriday.v2"
)

type Config struct {
	Author     string `toml:"author"`
	Email      string `toml:"email"`
	DateFormat string `toml:"dateFormat"`
	BaseURL    string `toml:"baseURL"`
}

type FrontMatter struct {
	Title string `toml:"title"`
	Date  string `toml:"date"`
}

type Post struct {
	Title         string
	Date          time.Time
	FormattedDate string
	Path          string
	Content       string
}

func replaceExtension(path, ext string) string {
	return path[0:len(path)-len(filepath.Ext(path))] + ext
}

func executeTemplate(source, target string, data interface{}) error {
	f, err := os.Create(target)
	if err != nil {
		return err
	}
	defer f.Close()
	tmpl, err := template.ParseFiles("layout.tmpl", source)
	if err != nil {
		return err
	}
	return tmpl.ExecuteTemplate(f, "layout", data)
}

func separateFrontMatter(b []byte) ([]byte, []byte) {
	i := bytes.Index(b[3:], []byte("+++"))
	if i == -1 {
		// Assume everything is Markdown
		return nil, b
	}
	return b[3 : i+3], b[i+6:]
}

func main() {
	var config Config
	_, err := toml.DecodeFile("pt.toml", &config)
	if err != nil {
		log.Fatal(err)
	}
	var posts []Post
	if err := filepath.Walk(".", func(path string, f os.FileInfo, err error) error {
		if filepath.Ext(path) != ".md" {
			return nil
		}
		log.Println(path)
		target := replaceExtension(path, ".html")
		b, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		var frontMatter FrontMatter
		var date time.Time
		fm, md := separateFrontMatter(b)
		if fm != nil {
			if err := toml.Unmarshal(fm, &frontMatter); err != nil {
				return err
			}
			date, err = time.Parse("2006-01-02", frontMatter.Date)
			if err != nil {
				return err
			}
		} else {
			log.Warn("missing front matter")
		}
		content := string(blackfriday.Run(md))
		post := Post{
			Title:         frontMatter.Title,
			Date:          date,
			FormattedDate: date.Format(config.DateFormat),
			Path:          target,
			Content:       content,
		}
		posts = append(posts, post)
		return executeTemplate("post.tmpl", target, post)
	}); err != nil {
		log.Fatal(err)
	}
	sort.Slice(posts, func(i, j int) bool { return posts[i].Date.After(posts[j].Date) })
	executeTemplate("index.tmpl", "index.html", posts)

	author := &feeds.Author{Name: config.Author, Email: config.Email}
	feed := &feeds.Feed{
		Title:  config.Author,
		Link:   &feeds.Link{Href: config.BaseURL},
		Author: author,
	}
	base, err := url.Parse(config.BaseURL)
	if err != nil {
		log.Fatal(err)
	}
	var items []*feeds.Item
	for _, post := range posts {
		if post.Title == "" {
			continue
		}
		u, err := url.Parse(post.Path)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, &feeds.Item{
			Title:   post.Title,
			Author:  author,
			Link:    &feeds.Link{Href: base.ResolveReference(u).String()},
			Created: post.Date,
		})
	}
	feed.Items = items
	f, err := os.Create("feed.xml")
	if err != nil {
		log.Fatal(err)
	}
	err = feed.WriteRss(f)
	if err != nil {
		log.Fatal(err)
	}
}
