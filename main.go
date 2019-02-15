// News agregator

package main

import ("fmt"
        "net/http"
        "io/ioutil"
        "os"
        "bufio"
        "encoding/xml"
		    "html/template"
		    "regexp"
		    "strings")

type News struct {
		Titles []string `xml:"channel>item>title"`
		Locations []string `xml:"channel>item>link"`
		Descriptions []string `xml:"channel>item>description"`
}

type NewsMap struct {
	Location string
	Description string
	Source string
}

type NewsAgg struct {
	Title string
	News map[string]NewsMap
}

func index_handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, `<html>
    <head><title>Index | News Aggregator</title>
    <body>
    <center>
    <p><a href="/news/">Recent News</a></p>
    <p><a href="/about/">About</a></p>
    </center>
    </body>
    </html>`)
}

func newsAgg_handler(w http.ResponseWriter, r *http.Request) {
	var n News

	regexpSource, _ := regexp.Compile("(www.|)([a-zA-Z0-9+]+)\\.([a-zA-Z]+)")
	regexpDescription1, _ := regexp.Compile("<br.*>")
	regexpDescription2, _ := regexp.Compile("<img.*>")

	news_map := make(map[string]NewsMap)
	file_location := os.Args[1]	// Text file location (First argument)
	file, _ := os.Open(file_location)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
					resp, _ := http.Get(scanner.Text())
					bytes, _ := ioutil.ReadAll(resp.Body)
					xml.Unmarshal(bytes, &n)
					for idx, _ := range n.Locations {
						news_map[n.Titles[idx]] = NewsMap{n.Locations[idx], regexpDescription2.ReplaceAllString(regexpDescription1.ReplaceAllString(n.Descriptions[idx], ""), ""), strings.ToUpper(regexpSource.FindString(n.Locations[idx]))}
					}
	}

	p := NewsAgg{Title: "Recent News | News Aggregator", News: news_map }
	t, _ := template.ParseFiles("newstemplate.html")
	t.Execute(w,p)
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<html>
    <head><title>About | News Aggregator</title>
    <body>
    <center>
    <p><a href="https://github.com/mrlibertarian/NewsAggregator">News Aggregator</a> is a simple RSS news aggregator written in Go.</p>
    <p>Please, donate: <b>BTC <a href="bitcoin:15EmrTsRjFuuiRgohSPKqDjjAXdisWULbs">15EmrTsRjFuuiRgohSPKqDjjAXdisWULbs</a></b></p>
    <p><a href="bitcoin:15EmrTsRjFuuiRgohSPKqDjjAXdisWULbs"><img src="/img/Bitcoin-QR.png"></a></p>
    <p><a href="/news/">Recent News</a></p>
    <p><a href="/">Index</a></p>
    </center>
    </body>
    </html>`)
}

func main() {
	var Port string = "800"

	var ServePort string = ":" + Port

	fmt.Printf("Serving at http://127.0.0.1%s", ServePort)
  http.Handle("/img/", http.StripPrefix("/img/",http.FileServer(http.Dir("Images"))))
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler)
	http.HandleFunc("/about/", about_handler)
	http.HandleFunc("/news", newsAgg_handler)
	http.HandleFunc("/news/", newsAgg_handler)
	http.ListenAndServe(ServePort, nil)
}
