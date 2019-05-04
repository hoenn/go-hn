package hnapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

const hnAddr = "https://hacker-news.firebaseio.com/v0/"

// HNClient wraps an http client with methods to query the hacker-news firebaseio API
type HNClient struct {
	httpClient *http.Client
}

// NewHNClient returns an HNClient with default settings
func NewHNClient() *HNClient {
	return &HNClient{
		httpClient: &http.Client{},
	}
}

// Item issues a GET request on the /item/<id> path and creates a struct
// from the response body
// It returns an empty interface to be asserted into one of the hnapi types
// Story, Comment, Poll, PollOpt or an error
func (h *HNClient) Item(id string) (interface{}, error) {
	url := hnObjURLString("item", id)
	resp, err := h.httpClient.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "could not make GET request for item")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.Wrap(err, "non 200 response code")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "could not read response body")
	}

	var t msgWrapper
	if err := json.Unmarshal(body, &t); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal response body")
	}
	var res interface{}
	switch t.Type {
	case "story", "job", "ask":
		s := &Story{}
		if err := json.Unmarshal(body, s); err != nil {
			return nil, err
		}
		res = s
	case "comment":
		c := &Comment{}
		if err := json.Unmarshal(body, c); err != nil {
			return nil, err
		}
		res = c
	case "poll":
		p := &Poll{}
		if err := json.Unmarshal(body, p); err != nil {
			return nil, err
		}
		res = p
	case "pollopt":
		p := &PollOpt{}
		if err := json.Unmarshal(body, p); err != nil {
			return nil, err
		}
		res = p
	default:
		return nil, errors.New("unknown item type, open a github issue")
	}

	return res, nil
}

// User issues a GET request on the /user/<id> path and
// returns a HNUser struct containing the details from
// the response
func (h *HNClient) User(id string) (*HNUser, error) {
	url := hnObjURLString("user", id)
	resp, err := h.httpClient.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "could not make GET request")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.Wrap(err, "non 200 response code")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "could not read response body")
	}
	u := &HNUser{}
	err = json.Unmarshal(body, u)
	if err != nil {
		return nil, errors.Wrap(err, "could not unmarshal response body")
	}

	return u, nil
}

//TopType is a type alias used to represent an enum
type TopType string

const (
	//Top is for the top ~500 stories
	Top TopType = "topstories"
	//New is for the new stories
	New TopType = "newstories"
	//Best is for the highest ranking stories
	Best TopType = "beststories"
	//Show is for stories categorized as 'Show'
	Show TopType = "showstories"
	//Job is for stories categorized as 'Jobs'
	Job TopType = "jobstories"
)

// TopStoryIDs issues a GET request to the /<TopType> path where TopType
// is one of the defined enum values. The API will return up to ~500 results
// for the Top and New categories
func (h *HNClient) TopStoryIDs(t TopType) ([]int, error) {
	url := hnURLString(fmt.Sprint(t))
	resp, err := h.httpClient.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "could not make GET request for top stories")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.Wrap(err, "non 200 response code")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "could not read response body")
	}
	var top []int
	err = json.Unmarshal(body, &top)
	if err != nil {
		return nil, errors.Wrap(err, "could not unmarshal top stories ids")
	}

	return top, nil
}

// MaxItemID issues a GET request to the /maxitem path and returns the current
// largest item id. This can be used to request information for all items by
// walking backwards
func (h *HNClient) MaxItemID() (int, error) {
	url := hnURLString("maxitem")
	resp, err := h.httpClient.Get(url)
	if err != nil {
		return -1, errors.Wrap(err, "could not make GET request for max item id")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return -1, errors.Wrap(err, "non 200 response code")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, errors.Wrap(err, "could not read response body")
	}
	var max int
	err = json.Unmarshal(body, &max)
	if err != nil {
		return -1, errors.Wrap(err, "could not unmarshal max item id")
	}

	return max, nil
}

// Updates issues a GET request to the /updates path and returns the latest item
// and profile changes. Items are updates to posts and comments and profiles are
// the IDs of the profiles that have recently changed
func (h *HNClient) Updates() (*Update, error) {
	url := hnURLString("updates")
	resp, err := h.httpClient.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "could not make GET request for updates")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.Wrap(err, "non 200 response code")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "could not read response body")
	}
	u := &Update{}
	err = json.Unmarshal(body, u)
	if err != nil {
		return nil, errors.Wrap(err, "could not unmarshal updates")
	}

	return u, nil
}

// hnObjURLString is a helper function that returns an http URL for an api path like
// .../v0/<obj>/<id>.json
func hnObjURLString(obj, id string) string {
	return hnAddr + obj + "/" + id + ".json"
}

// hnURLString is a helper function that returns an http URL for an api path like
// .../v0/<path>.json
func hnURLString(path string) string {
	return hnAddr + path + ".json"
}

// Update represents the response from the 'updates' path
type Update struct {
	Items    []int
	Profiles []string
}

// HNUser represents a HackerNews user profile
type HNUser struct {
	About     string
	Created   int64
	Delay     int
	ID        string
	Karma     int
	Submitted []int
}

type msgWrapper struct {
	Type string
}

// Comment represents a HackerNews comment on a story
type Comment struct {
	By     string
	ID     int
	Kids   []int
	Parent int
	Text   string
	Time   int64
	Type   string
}

// Story represents a HackerNews submitted story
type Story struct {
	By          string
	Descendants int
	ID          int
	Kids        []int
	Score       int
	Time        int64
	Title       string
	Type        string
	URL         string
}

// Poll represents a HackerNews poll. It contains a Parts field that describes the related
// poll options
type Poll struct {
	By          string
	Descendants int
	ID          int
	Kids        []int
	Parts       []int
	Score       int
	Text        string
	Time        int64
	Title       string
	Type        string
}

// PollOpt represents the poll options from a given poll
type PollOpt struct {
	By    string
	ID    int
	Poll  int
	Score int
	Text  string
	Time  int64
	Type  string
}
