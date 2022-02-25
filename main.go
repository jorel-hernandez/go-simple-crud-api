package main

/*update your REST API to use a gorilla/mux router instead of the traditional net/http router.
Swapping the routers will enable you to more easily perform tasks such as parsing any path or query parameters that may reside
within an incoming HTTP request which we will need later on. */
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	//fmt.Fprintf(w, "key: "+key)
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	//reqBody, _ := ioutil.ReadAll(r.Body)
	//fmt.Fprintf(w, "%+v", string(reqBody))

	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)

	// update our global Articles array to include
	// our new Article
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)

	// we will need to extract the `id` of the article we wish to delete
	id := vars["id"]

	//we then need to loop through al our articles
	for index, article := range Articles {
		// if our path parameter matches one of our articles
		if article.Id == id {
			// updates our Articles array to remove the article
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	reqBody, _ := ioutil.ReadAll(r.Body)
	var articleUpdate Article
	json.Unmarshal(reqBody, &articleUpdate)

	for index, article := range Articles {
		//Articles = append(Articles, article)
		if article.Id == id {
			Articles[index].Title = articleUpdate.Title
			Articles[index].Desc = articleUpdate.Desc
			Articles[index].Content = articleUpdate.Content

			json.NewEncoder(w).Encode(Articles[index])
		}
	}
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	//http.HandleFunc("/", homePage)
	//http.HandleFunc("/articles", returnAllArticles)

	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)

	// NOTE: Ordering is important here! This has to be defined before
	// the other `/article` endpoint.
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")

	// add our new UPDATE endpoint here
	myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")

	// add our new DELETE endpoint here
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")

	myRouter.HandleFunc("/article/{id}", returnSingleArticle)

	//log.Fatal(http.ListenAndServe(":10000", nil))

	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	handleRequests()
}
