package gopheran

import (
	"net/http"
	"time"

	"appengine"
	"appengine/datastore"
	"github.com/crhym3/go-endpoints/endpoints"
)

type Greeting struct {
	Key     *datastore.Key `json:"id" datastore:"-"`
	Author  string         `json:"author"`
	Content string         `json:"content" datastore:",noindex" endpoints:"req"`
	Date    time.Time      `json:"date"`
}

// bu yollarken gerekli alanlar zorunlu değil
type GreetingReq struct {
	Author  string `json:"author" endpoints:"req"`
	Content string `json:"content"`
}

type GreetingReqShow struct {
	Key *datastore.Key `json:"id"`
}

// geri dönen alanlar
type GreetingResp struct {
	Author  string    `json:"author" endpoints:"req"`
	Content string    `json:"content"`
	Date    time.Time `json:"date"`
}

// GreetingsList is a response type of GreetingService.List method
type GreetingsList struct {
	Items []*Greeting `json:"items"`
}

// Request type for GreetingService.List
type GreetingsListReq struct {
	Limit int `json:"limit" endpoints:"d=10"`
}

func (gr *Greeting) put(c appengine.Context) (err error) {
	key := gr.Key

	if key == nil {
		key = datastore.NewIncompleteKey(c, "greeting", nil)
	}

	key, err = datastore.Put(c, key, gr)

	if err == nil {
		gr.Key = key
	}

	return
}

func newGreeting(author string, content string) *Greeting {
	return &Greeting{Author: author, Content: content, Date: time.Now()}
}

// api
type GreetingService struct {
}

func (gs *GreetingService) Create(r *http.Request, req *GreetingReq, resp *GreetingResp) error {
	c := endpoints.NewContext(r)

	greeting := newGreeting(req.Author, req.Content)

	if err := greeting.put(c); err != nil {
		return err
	}
	resp.Content = greeting.Content
	resp.Author = greeting.Author
	resp.Date = greeting.Date
	return nil
}

func (gs *GreetingService) Show(r *http.Request, req *GreetingReqShow, resp *GreetingResp) error {
	c := endpoints.NewContext(r)

	var g Greeting
	if err := datastore.Get(c, req.Key, &g); err != nil {
		return err
	}

	resp.Content = g.Content
	resp.Author = g.Author
	resp.Date = g.Date
	return nil
}

// List responds with a list of all greetings ordered by Date field.
// Most recent greets come first.
func (gs *GreetingService) List(r *http.Request, req *GreetingsListReq, resp *GreetingsList) error {
	if req.Limit <= 0 {
		req.Limit = 10
	}

	c := endpoints.NewContext(r)

	q := datastore.NewQuery("greeting").Order("-Date").Limit(req.Limit)

	greets := make([]*Greeting, 0, req.Limit)
	keys, err := q.GetAll(c, &greets)

	if err != nil {
		return err
	}

	for i, k := range keys {
		greets[i].Key = k
	}
	resp.Items = greets
	return nil
}

func init() {
	greetService := &GreetingService{}

	api, err := endpoints.RegisterService(greetService, "greeting", "v1", "Greetings API", true)
	if err != nil {
		panic(err.Error())
	}

	info := api.MethodByName("List").Info()
	info.Name, info.HttpMethod, info.Path, info.Desc = "greets.list",
		"GET", "greetings", "List"

	info = api.MethodByName("Create").Info()
	info.Name, info.HttpMethod, info.Path, info.Desc = "greets.create",
		"POST", "greetings", "Create"

	info = api.MethodByName("Show").Info()
	info.Name, info.HttpMethod, info.Path, info.Desc = "greets.show",
		"GET", "greetings/{id}", "Show"

	endpoints.HandleHttp()
}
