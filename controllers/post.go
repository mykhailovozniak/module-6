package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Post struct {
	UserId int64 	`json:"userId"`
	Id int64     	`json:"id"`
	Title string 	`json:"title"`
	Body string  	`json:"body"`
}

func IsJSON(s string) bool {
	var js map[string]interface{}

	return json.Unmarshal([]byte(s), &js) == nil

}

func PostController(res http.ResponseWriter, req *http.Request) {
	keys, ok := req.URL.Query()["postId"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url param 'postId' is missing")
		http.Error(res, "Bad request", 400)

		return
	}

	postId := keys[0]

	cachePosts := map[string] string {
		"1": `{ "userId" : 1, "id" : 1, "title" : "sunt aut facere repellat provident occaecati excepturi optio reprehenderit","body":"quia et suscipit\\nsuscipit recusandae consequuntur expedita et cum\\nreprehenderit molestiae ut ut quas totam\\nnostrum rerum est autem sunt rem eveniet architecto" }`,
		"2": "{\"userId\":1,\"id\":2,\"title\":\"qui est esse\",\"body\":\"est rerum tempore vitae nsequi sint nihil reprehenderit dolor beatae ea dolores neque nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis nqui aperiam non debitis possimus qui neque nisi nulla\"}",
		"3": "{\"userId\":1,\"id\":3,\"title\":\"ea molestias quasi exercitationem repellat qui ipsa sit aut\",\"body\":\"et iusto sed quo iure\\nvoluptatem occaecati omnis eligendi aut ad\\nvoluptatem doloribus vel accusantium quis pariatur\\nmolestiae porro eius odio et labore et velit aut\"}",
	}

	cachePost, hasCacheVersion := cachePosts[postId]
	var post Post

	if !hasCacheVersion {
		resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + postId)

		if err != nil {
			log.Println("Error during fetching post with id " + postId)
		}

		body, err := ioutil.ReadAll(resp.Body)
		log.Println("Body from external resource: " + string(body))

		err = json.Unmarshal(body, &post)

		res.Header().Set("Content-Type", "application/json")
		res.Header().Set("cached", "false")

		jsonData, err := json.Marshal(post)
		res.Write(jsonData)

		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("cached", "true")

	log.Println("isJSON", IsJSON(cachePost))

	err := json.Unmarshal([]byte(cachePost), &post)

	if err != nil {
		log.Println("Error during parsing a cache")
	}

	jsonData, err := json.Marshal(post)
	res.Write(jsonData)
}
