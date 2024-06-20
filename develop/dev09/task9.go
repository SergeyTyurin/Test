package main

/*
	Утилита wget
    Реализовать утилиту wget с возможностью скачивать сайты целиком.

*/
import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/gocolly/colly"
)

var urlMy string
var pathName string

func download(bytes []byte, filename string) error {
	//создание директории(поддиректории), если такая директория не создана
	err := os.MkdirAll(pathName+filepath.Dir(filename), 0755)
	if err != nil {
		log.Fatal(err)
	}
	// создание файла и открытие на запись
	file, err := os.OpenFile(pathName+filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	// запись полученный байтов
	_, err = file.Write(bytes)

	if err != nil {
		return err
	}

	return err
}

func main() {
	// парсинг флага url
	flag.StringVar(&urlMy, "u", "", "'url' - ссылка на html-страницу")
	flag.Parse()

	if len(urlMy) == 0 {
		log.Fatal("url is empty")
	}

	urlPath, err := url.Parse(urlMy)
	if err != nil {
		log.Fatal(err)
	}

	// получение текущей директори и создание директории для хранения результатов
	pathName, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	pathName = pathName + "/" + urlPath.Host
	err = os.MkdirAll(pathName, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pathName)

	// для обхода всех ссылок сайта, для получения всех страниц используется веб скраппер colly
	c := colly.NewCollector()

	// обработчик нахождения html элемента
	c.OnHTML("a", func(e *colly.HTMLElement) {
		// парсинг ссылки
		href := e.Attr("href")
		link, err := url.Parse(href)
		if err != nil {
			log.Fatal(err)
		}
		// проверка, что ссылка не ведет на сторонний сайт
		host := link.Hostname()
		if (host == urlPath.Host || len(host) == 0) && (len(link.Scheme) == 0 || link.Scheme == "http" || link.Scheme == "https") { //!strings.HasPrefix(href, "mailto") {
			// посещение ссылки
			e.Request.Visit(href)
		}
	})

	// обработка получения ответа
	c.OnResponse(func(r *colly.Response) {
		path := r.Request.URL.Path
		// загрузка контента в файл с именем path
		download(r.Body, path)
	})

	// обработка запроса
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// посещение указанного url
	c.Visit(urlMy)
}
