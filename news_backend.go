package news

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

//Api is the client api
type Api struct {
	client  *http.Client
	apiKey  string
	baseUrl string
}

//New instanciate the newapis
func New(key string) *Api {
	return &Api{&http.Client{}, key, "https://newsapi.org/v2/"}
}

// Request interface implements what is required for the interface
type Request interface {
	SerializeUrl() string
}

//ReqArticles describes the request api visit the new api documentation
//for more infos https://newsapi.org/docs/endpoints/everything
type ReqArticles struct {
	baseUrl  string
	q        string
	sources  []string
	domains  []string
	from     string
	to       string
	sortBy   string
	language string
	pageSize int
	page     int
}

// SerializeUrl implements the required interface
func (r ReqArticles) SerializeUrl() string {
	return r.baseUrl + "everything?q=" + r.q
}

//ReqHeadlines describes the request api visit the new api documentation
//for more infos https://newsapi.org/docs/endpoints/everything
type ReqHeadlines struct {
	baseUrl  string
	country  string
	category string
	sources  []string
	q        string
	pageSize int
	page     int
}

// SerializeUrl implements the required interface
func (r ReqHeadlines) SerializeUrl() string {
	return r.baseUrl + "top-headlines?category=" + r.category + "&pageSize=" + strconv.Itoa(r.pageSize) + "&country=" + r.country
}

//Response represent the response to every request
type Response struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []struct {
		Source struct {
			ID   interface{} `json:"id"`
			Name string      `json:"name"`
		} `json:"source"`
		Author      string    `json:"author"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		URL         string    `json:"url"`
		URLToImage  string    `json:"urlToImage"`
		PublishedAt time.Time `json:"publishedAt"`
	} `json:"articles"`
}

//	if category == "" {
//	url = a.baseUrl + reqType + "?q=" + keyword
// { else {
//		url = a.baseUrl + reqType + "?category=" + category + "&country=us" + "&pageSize=100"
// }

func (a *Api) newRequest(r Request) (*http.Request, error) {
	req, err := http.NewRequest("GET", r.SerializeUrl(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Api-Key", a.apiKey)

	return req, nil
}

func (a *Api) do(r Request) (*Response, error) {
	req, err := a.newRequest(r)
	if err != nil {
		return nil, err
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res Response
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil

}

// GetArticles returns all the articles in given a request, a keyword can
// be anything
func (a *Api) GetArticles(keyword string) (*Response, error) {
	req := ReqArticles{baseUrl: a.baseUrl, q: keyword}
	return a.do(req)
}

// GetHeadlines returns all the articles in given a request, a category can
// be one of the following:
// entertainement, general, health, science, sports, technology, business
func (a *Api) GetHeadlines(category string) (*Response, error) {
	req := ReqHeadlines{baseUrl: a.baseUrl, category: category}
	return a.do(req)
}
