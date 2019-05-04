[![Go Report Card](https://goreportcard.com/badge/github.com/hoenn/go-hn)](https://goreportcard.com/report/github.com/hoenn/go-hn)
[![Documentation](https://godoc.org/github.com/Hoenn/go-hn?status.svg)](http://godoc.org/github.com/Hoenn/go-hn)

## go-hn
`go` wrapper for the HackerNews Firebase API.

### Getting started

Add this package to your `go` path.

`go get github.com/Hoenn/go-hn`

then create the client with

```go
 import "github.com/Hoenn/go-hn/pkg/hnapi"
 //...
 c := hnapi.NewHNClient()
```

### Concepts
Stories, comments, jobs, asks and polls are all "items" and `Item(id string) (interface{}, error)` encourages you to assert the type of the item returned:

#### Items

##### Story
Field | Type | Description
------|------|------------
By      | `string`    | The username of the story's author
ID      | `int`       | This story's item id
Descendants | `int`   | The total comment count
Kids    | `[]int`     | The ids of the story's comments, in ranked display order.
Score   | `int`       | The story's score
Time    | `int64`     | Creation date of the story in Unix Time
Title   | `string`    | The story's title
URL     | `string`    | The URL of the story
Type    | `string`    | The type of item ("story")

##### Comment
Field | Type | Description
------|------|------------
By      | `string`    | The username of the comments's author
ID      | `int`       | This comments's item id
Kids    | `[]int`     | The ids of the story's comments, in ranked display order.
Parent  | `int`       | The comment's parent (another comment or the original story)
Time    | `int64`     | Creation date of the comment in Unix Time
Text    | `string`    | The comment's text
Type    | `string`    | The type of item ("comment")

##### Poll
Field | Type | Description
------|------|------------
By      | `string`    | The username of the Poll's author
ID      | `int`       | This poll's item id
Descendants | `int`   | The total comment count
Kids    | `[]int`     | The ids of the poll's comments, in ranked display order.
Score   | `int`       | The poll's score
Time    | `int64`     | Creation date of the poll in Unix Time
Title   | `string`    | The poll's title
URL     | `string`    | The URL of the poll
Type    | `string`    | The type of item ("poll")
Parts   | `[]int`     | The item ids of the poll options

##### PollOpt
Field | Type | Description
------|------|------------
By      | `string`    | The username of the pollopt's author
ID      | `int`       | This pollopt's item id
Poll    | `int`       | The pollopt's parent (the poll it belongs to)
Time    | `int64`     | Creation date of the pollopt in Unix Time
Text    | `string`    | The pollopt's text
Type    | `string`    | The type of item ("pollopt")

#### Users
Users are identified by case-sensitive ids. Only users that have public activity (comments or story submissions) on the site are available through the API.

##### HNUser
Field | Type | Description
------|------|------------
About       | `string`    | The user's 'about' info (HTML)
Created     | `int64`     | Creation date of the user in Unix Time
Delay       | `int`       | Delay in minutes between comments becoming visible after posting
ID          | `string`  | The user's unique username, case-sensitive
Karma       | `int`       | The user's karma
Submitted   | `[]int`     | List of all item submissions by user


#### Additional Functionality

##### Top Stories
There are additional functions in the `hnapi` package for getting the "Top" stories by type as determined by the API itself. You can call `client.TopStoryIDs` passing one of the following `TopType`:
```
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
```
Any of the `TopType` options will return up to ~500 item ids.

##### Updates
The HackerNews API also exposes an endpoint to poll for updates. `client.Updates` will return a single `*Update` containing

Field | Type | Description
------|------|------------
Items | `[]int` | Items that have been recently updated
Profiles | `[]string` | User profiles that have been recently updated

### Example Usage

```go
package main

import (
    "fmt"
    "github.com/Hoenn/go-hn/pkg/hnapi"
)

func main() {
    //Create the client
    c := hnapi.NewHNClient()

    //Display a specific user's karma
    user, err := c.User("someuser")
    if err != nil {
        panic(err)
    }
    fmt.Println(user.Karma)

    //Get the current max item id
    maxID, err := c.MaxItemID()
    if err != nil {
        panic(err)
    }

    //Get the details of the current max item
    maxItem, err := c.Item(maxID)
    //...
    switch maxItem.(type) {
        case hnapi.Story:
            //...
        case hnapi.Comment:
            //...
    }
}
```
