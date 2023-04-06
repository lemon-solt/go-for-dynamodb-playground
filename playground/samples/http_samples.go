package samples

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func SimpleHttp() {

	res, _ := http.Get("https://jp.mercari.com/item/m64717066030")
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	// ioutil.WriteFile("hoge.html", body, 0666)
}

func SimpleParseUrl() {
	base, _ := url.Parse("https://www.lemon-solt-works.com/")
	reference, _ := url.Parse("?hoge=hei")
	url := base.ResolveReference(reference).String()

	fmt.Println(url)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("If-None-Match", "x/xyz")
	q := req.URL.Query()
	q.Add("c", "[]")
	req.URL.RawQuery = q.Encode()

	var client *http.Client = &http.Client{}
	response, _ := client.Do(req)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

}
