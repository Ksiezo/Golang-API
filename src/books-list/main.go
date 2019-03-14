package main

//import dependencies
import(
    "encoding/json"
  //  "fmt"
  "strconv"
    "log" //logging stuff
    "net/http" //used to setup localhost port
    "github.com/gorilla/mux" //router for handling urls
    //"reflect"
    //"database/sql" //allows sql functions within go
    //"github.com/lib/pq" // bridge client used to connect to Postgres/ElephantSQL
    //"github.com/subosito/gotenv" //allows the ability to store configs securely

)

//Create book resource model
type Book struct {
    ID int `json:id`
    Title string `json:title`
    Author string `json:author`
    Year string `json:year`
}


var books []Book
//create main
func main(){
  router := mux.NewRouter()

  books =  append(books,Book{ID:1, Title:"Golang pointers", Author: "Mr. Golang", Year: "2010"},
               Book{ID:2, Title:"Goroutines", Author: "Mr. Goroutine", Year:"2011"},
               Book{ID:3, Title:"Go routers", Author: "Mr. Router", Year:"2012"},
               Book{ID:4, Title:"Golang concurrency", Author: "Mr. Currency", Year:"2013"},
               Book{ID:5, Title:"Golang good parts", Author: "Mr. Good", Year:"2014"})

  router.HandleFunc("/books",getBooks).Methods("GET") //get the collection of books
  router.HandleFunc("/books/{id}",getBook).Methods("GET")
  router.HandleFunc("/books",addBook).Methods("POST")
  router.HandleFunc("/books",updateBook).Methods("PUT")
  router.HandleFunc("/books/{id}",removeBook).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":8000",router)) //set router up to localhost to use handlefuncs //if there is any error check now

}

//API calls

func getBooks(w http.ResponseWriter,r *http.Request){
  json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter,r *http.Request){
  params := mux.Vars(r) //creates a map of route variables stemming from the sent URL request

  //log.Print(reflect.TypeOf(params["id"]) //check what type of data is passed
 //string was passed, use strconv lib to convert the data
// to get a single book get book
  i, _ := strconv.Atoi(params["id"]) // from the params we want the last variable {id}

// in the available book objects we want the one that matches our current request
  for _,book := range books {
    if book.ID == i{
      json.NewEncoder(w).Encode(&book)
    }

  }
  //log.Println(params)
}

func addBook(w http.ResponseWriter,r *http.Request){
  var book Book
  _ = json.NewDecoder(r.Body).Decode(&book)
  //log.Println(book) //check object being posted , (i could also just check in postman but hey whatever )
  books = append(books,book)

  json.NewEncoder(w).Encode(books)



}

func updateBook(w http.ResponseWriter,r *http.Request){
  var book Book
  json.NewDecoder(r.Body).Decode(&book) //decode mapped body and add it to the book variable

  for i, item := range books {
    if item.ID == book.ID {
      books[i] = book
    }
  }

  json.NewEncoder(w).Encode(books)

}
func removeBook(w http.ResponseWriter,r *http.Request){
  params := mux.Vars(r)
  id, _ := strconv .Atoi(params["id"])

  for i, item := range books {
    if item.ID == id {
      books = append(books[:i],books[i+1:]...)
    }

  }

}
