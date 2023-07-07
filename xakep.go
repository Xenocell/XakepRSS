package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

const (
	constUrl     string = "url"
	constTitle   string = "title"
	constContent string = "content"
)

var (
	articlesRegex    string = "(?:<article >)(.*?)(?:<\\/article>)"
	articleDataRegex        = fmt.Sprintf(`<h3 class="entry-title"><a href="(?P<%s>.*?)"><span>(?P<%s>.*?)</span></a></h3>\s+</header>\s+<p class="block-exb">(?P<%s>.*?)</p>`, constUrl, constTitle, constContent)
)

type IXakepParse interface {
	GetPageByNumberPage(page int) ([]Article, error)
	GetFirstPage() ([]Article, error)
}

type XakepParse struct {
	url string
}

func NewXakepParse() *XakepParse {
	return &XakepParse{"https://xakep.ru/"}
}

func (x *XakepParse) GetPageByNumberPage(page int) ([]Article, error) {
	var result []Article

	firstPage, err := x.getDataFirstPage()
	if err != nil {
		return result, fmt.Errorf("xakep - GetPageByNumberPage - x.getDataFirstPage: %w", err)
	}

	ar, err := regexp.Compile(articlesRegex)
	if err != nil {
		return result, fmt.Errorf("xakep - GetPageByNumberPage - regexp.Compile(articlesRegex): %w", err)
	}

	articles := ar.FindAllString(firstPage, -1)

	for _, v := range articles {
		if strings.Contains(v, "Новости") {
			articleRegex, _ := regexp.Compile(articleDataRegex)

			article := articleRegex.FindStringSubmatch(v)

			result = append(result, Article{
				Url:     article[articleRegex.SubexpIndex(constUrl)],
				Title:   article[articleRegex.SubexpIndex(constTitle)],
				Content: article[articleRegex.SubexpIndex(constContent)],
			})
		}
	}

	return result, nil
}

func (x *XakepParse) GetFirstPage() ([]Article, error) {
	return x.GetPageByNumberPage(1)
}

func (x *XakepParse) getDataPageByPage(page int) (string, error) {
	var result string

	res, err := http.Get(fmt.Sprintf("%spage/%d/", x.url, page))
	if err != nil {
		return result, fmt.Errorf("http.Get: %w", err)
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return result, fmt.Errorf("io.ReadAll: %w", err)
	}
	result = strings.Replace(string(resBody), "\n", "", -1)
	return result, err
}

func (x *XakepParse) getDataFirstPage() (string, error) {
	return x.getDataPageByPage(1)
}
