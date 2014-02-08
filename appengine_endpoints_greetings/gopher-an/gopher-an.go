package gopheran

import (
	"net/http"
	"time"

	_ "appengine"
	"appengine/datastore"
	"github.com/crhym3/go-endpoints/endpoints"
)

type Greeting struct {
	Key     *datastore.Key `json:"id" datastore:"-"`
	Author  string         `json:"author"`
	Content string         `json:"content" datastore:",noindex" endpoints:"req"`
	Date    time.Time      `json:"date"`
}

type GreetingReqId struct {
	Key *datastore.Key `json:"id"`
}

// GreetingsList is a response type of GreetingService.List method
type GreetingsList struct {
	Items []*Greeting `json:"items"`
}

type GreetingsD struct {
}

// Request type for GreetingService.List
type GreetingsListReq struct {
	Limit int `json:"limit" endpoints:"d=10"`
}

func newGreeting(author string, content string) *Greeting {
	return &Greeting{Author: author, Content: content, Date: time.Now()}
}

// api
type GreetingService struct {
}

func (gs *GreetingService) Create(r *http.Request, req *Greeting, resp *Greeting) error {
	c := endpoints.NewContext(r)

	resp.Author, resp.Content = req.Author, req.Content

	key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "greeting", nil), resp)

	if err != nil {
		return err
	}
	resp.Key = key
	resp.Date = time.Now()

	return nil
}

func (gs *GreetingService) Show(r *http.Request, req *GreetingReqId, resp *Greeting) error {
	c := endpoints.NewContext(r)

	if err := datastore.Get(c, req.Key, resp); err != nil {
		return err
	}

	return nil
}

func (gs *GreetingService) Destroy(r *http.Request, req *GreetingReqId, resp *GreetingsD) error {
	c := endpoints.NewContext(r)

	if err := datastore.Delete(c, req.Key); err != nil {
		return err
	}

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

	info = api.MethodByName("Destroy").Info()
	info.Name, info.HttpMethod, info.Path, info.Desc = "greets.destroy",
		"DELETE", "greetings/{id}", "delete"
	endpoints.HandleHttp()
}
