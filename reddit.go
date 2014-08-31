//package reddit implements a basic client for the Reddit API.
package reddit

import (
	"net/http"
	"encoding/json"
	"errors"
	"fmt"
)

//Item describe a Reddit item
type Item struct {
	Domain   string
	Title    string
	URL      string
	Comments int `json:"num_comments"`
}
type response struct {
	Kind string
	Data struct {
		Children []struct {
			Kind string
			Data Item
		}
	}
}

//Get fetches the most recent Items posted to the specified subreddit
func Get(reddit string) ([]Item, error) {
	url := fmt.Sprintf("http://reddit.com/r/%s.json", reddit)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	r := new(response)
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	fmt.Println(r.Kind)
	items := make([]Item, len(r.Data.Children))
	for i, child := range r.Data.Children {
		items[i] = child.Data
	}
	return items, nil
}

func (i Item) String() string {
	com := " com "
	switch i.Comments {
	case 0:
		//nothing
	case 1:
		com = "1 comments"
	default:
		com = fmt.Sprintf("%d Comments", i.Comments)
	}
	//return fmt.Sprintf("%s %s\n%s", i.Title, com, i.URL)
	return fmt.Sprintf("%s %s", i.Title, com)
}

//func main() {
//	items, err := Get("golang")
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, item := range items {
//		fmt.Println(item)
//	}
//	resp, err := http.Get("http://reddit.com/r/golang.json")
//	if err != nil {
//		log.Fatal(err)
//	}
//	if resp.StatusCode != http.StatusOK {
//		log.Fatal(resp.Status)
//	}
////	_, err = io.Copy(os.Stdout, resp.Body)
//	r := new(response)
//	err = json.NewDecoder(resp.Body).Decode(r)
//	for _, child := range r.Data.Children {
//		fmt.Println(child.Data.Title)
//	//	fmt.Println(child.Data.URL)
//	}
//	if err != nil {
//		log.Fatal(err)
//	}
//}
