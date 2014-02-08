package gopheran

import (
	"errors"
	"net/http"
	"time"

	_ "appengine"
	"appengine/datastore"
	"appengine/user"
	"github.com/crhym3/go-endpoints/endpoints"
)

const clientId = "YOUR-CLIENT-ID"

var (
	scopes    = []string{endpoints.EmailScope}
	clientIds = []string{clientId, endpoints.ApiExplorerClientId}
	// from android
	audiences = []string{clientId}
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

type GreetingsList struct {
	Items []*Greeting `json:"items"`
}

type GreetingsListReq struct {
	Limit int `json:"limit" endpoints:"d=10"`
}

type GreetingsD struct {
}

func newGreeting(author string, content string) *Greeting {
	return &Greeting{Author: author, Content: content, Date: time.Now()}
}

func getCurrentUser(c endpoints.Context) (*user.User, error) {
	u, err := endpoints.CurrentUser(c, scopes, audiences, clientIds)

	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, errors.New("Unauthorized: Please, sign in.")
	}
	c.Debugf("Current user %#v", u)
	return u, nil
}

// api
type GreetingService struct {
}

func (gs *GreetingService) Create(r *http.Request, req *Greeting, resp *Greeting) error {
	c := endpoints.NewContext(r)

	_, err := getCurrentUser(c)
	if err != nil {
		return err
	}

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

	_, err := getCurrentUser(c)
	if err != nil {
		return err
	}

	if err := datastore.Get(c, req.Key, resp); err != nil {
		return err
	}

	return nil
}

func (gs *GreetingService) Destroy(r *http.Request, req *GreetingReqId, resp *GreetingsD) error {
	c := endpoints.NewContext(r)

	_, err := getCurrentUser(c)
	if err != nil {
		return err
	}

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
	_, err := getCurrentUser(c)
	if err != nil {
		return err
	}

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
	info.Scopes, info.ClientIds, info.Audiences = scopes, clientIds, audiences

	info = api.MethodByName("Create").Info()
	info.Name, info.HttpMethod, info.Path, info.Desc = "greets.create",
		"POST", "greetings", "Create"
	info.Scopes, info.ClientIds, info.Audiences = scopes, clientIds, audiences

	info = api.MethodByName("Show").Info()
	info.Name, info.HttpMethod, info.Path, info.Desc = "greets.show",
		"GET", "greetings/{id}", "Show"
	info.Scopes, info.ClientIds, info.Audiences = scopes, clientIds, audiences

	info = api.MethodByName("Destroy").Info()
	info.Name, info.HttpMethod, info.Path, info.Desc = "greets.destroy",
		"DELETE", "greetings/{id}", "delete"
	info.Scopes, info.ClientIds, info.Audiences = scopes, clientIds, audiences

	endpoints.HandleHttp()
}
