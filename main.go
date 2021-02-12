package main

import (
   "context"
   "fmt"
   "log"
   "net/http"
   "encoding/json"
   "github.com/gorilla/mux"
   "github.com/eefret/gomdb"
   "strconv"
   "time"
   "go.mongodb.org/mongo-driver/bson"
   "go.mongodb.org/mongo-driver/mongo"
   "go.mongodb.org/mongo-driver/mongo/options"
   "go.mongodb.org/mongo-driver/mongo/readpref"
)

func handleRequests() {
   // creates a new instance of a mux router
   myRouter := mux.NewRouter().StrictSlash(true)
   // replace http.HandleFunc with myRouter.HandleFunc
   myRouter.HandleFunc("/queryfrominternet", queryFromInternet)
   myRouter.HandleFunc("/searchbyid/{id}", searchbyid)
   myRouter.HandleFunc("/searchbyyear/{year}", searchbyYear)
   myRouter.HandleFunc("/searchbyperiod/{startYear}/{endYear}", searchbyPeriod)
   myRouter.HandleFunc("/searchbyrating/{rating}", searchbyRating)
   myRouter.HandleFunc("/searchbygenre/{genre}", searchbyGenre)

   log.Fatal(http.ListenAndServe(":10013", myRouter))
}

type Movie struct {
   Id  string `json:"Id"`
}

func searchbyid(w http.ResponseWriter, r *http.Request) {
   vars := mux.Vars(r)
   key := vars["id"]

   client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://tester:Shipra@cluster0.7wpka.mongodb.net/movies"))
   if err != nil {
      log.Fatal(err)
   }
   ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
   err = client.Connect(ctx)
   if err != nil {
      log.Fatal(err)
   }

   defer client.Disconnect(ctx)
   err = client.Ping(ctx, readpref.Primary())
   if err != nil {
      log.Fatal(err)
   }
   databases, err := client.ListDatabaseNames(ctx, bson.M{})
   if err != nil {
      log.Fatal(err)
   }
   fmt.Println(databases)

   col := client. Database("movies"). Collection("list")


   cursor, err := col.Find(ctx, bson.M{"id" : key})
   if err != nil {
      log.Fatal(err)
   }
   var episodes []bson.M
   if err = cursor.All(ctx, &episodes); err != nil {
      log.Fatal(err)
   }
   fmt.Println(episodes)

   json.NewEncoder(w).Encode(episodes)

}

func searchbyYear(w http.ResponseWriter, r *http.Request) {
   vars := mux.Vars(r)
   key ,err := strconv.Atoi(vars["year"])

   client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://tester:Shipra@cluster0.7wpka.mongodb.net/movies"))
   if err != nil {
      log.Fatal(err)
   }
   ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
   err = client.Connect(ctx)
   if err != nil {
      log.Fatal(err)
   }

   defer client.Disconnect(ctx)
   err = client.Ping(ctx, readpref.Primary())
   if err != nil {
      log.Fatal(err)
   }
   databases, err := client.ListDatabaseNames(ctx, bson.M{})
   if err != nil {
      log.Fatal(err)
   }
   fmt.Println(databases)

   col := client. Database("movies"). Collection("list")


   cursor, err := col.Find(ctx, bson.M{"released year" : key})
   if err != nil {
      log.Fatal(err)
   }
   var episodes []bson.M
   if err = cursor.All(ctx, &episodes); err != nil {
      log.Fatal(err)
   }
   fmt.Println(episodes)

   json.NewEncoder(w).Encode(episodes)

}

func searchbyPeriod(w http.ResponseWriter, r *http.Request) {
   vars := mux.Vars(r)
   key1 ,err := strconv.Atoi(vars["startYear"])
   key2 ,err := strconv.Atoi(vars["endYear"])

   client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://tester:Shipra@cluster0.7wpka.mongodb.net/movies"))
   if err != nil {
      log.Fatal(err)
   }
   ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
   err = client.Connect(ctx)
   if err != nil {
      log.Fatal(err)
   }

   defer client.Disconnect(ctx)
   err = client.Ping(ctx, readpref.Primary())
   if err != nil {
      log.Fatal(err)
   }
   databases, err := client.ListDatabaseNames(ctx, bson.M{})
   if err != nil {
      log.Fatal(err)
   }
   fmt.Println(databases)

   col := client. Database("movies"). Collection("list")


   cursor, err := col.Find(ctx, bson.M{"released year" : bson.M{ "$gte" : key1 , "$lte" : key2 } })
   if err != nil {
      log.Fatal(err)
   }
   var episodes []bson.M
   if err = cursor.All(ctx, &episodes); err != nil {
      log.Fatal(err)
   }
   fmt.Println(episodes)

   json.NewEncoder(w).Encode(episodes)

}

func searchbyRating(w http.ResponseWriter, r *http.Request) {
   vars := mux.Vars(r)
   key,err := strconv.Atoi(vars["rating"])

   client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://tester:Shipra@cluster0.7wpka.mongodb.net/movies"))
   if err != nil {
      log.Fatal(err)
   }
   ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
   err = client.Connect(ctx)
   if err != nil {
      log.Fatal(err)
   }

   defer client.Disconnect(ctx)
   err = client.Ping(ctx, readpref.Primary())
   if err != nil {
      log.Fatal(err)
   }
   databases, err := client.ListDatabaseNames(ctx, bson.M{})
   if err != nil {
      log.Fatal(err)
   }
   fmt.Println(databases)

   col := client. Database("movies"). Collection("list")


   cursor, err := col.Find(ctx, bson.M{"rating" : key})
   if err != nil {
      log.Fatal(err)
   }
   var episodes []bson.M
   if err = cursor.All(ctx, &episodes); err != nil {
      log.Fatal(err)
   }
   fmt.Println(episodes)

   json.NewEncoder(w).Encode(episodes)

}

func searchbyGenre(w http.ResponseWriter, r *http.Request) {
   vars := mux.Vars(r)
   key := vars["genre"]

   client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://tester:Shipra@cluster0.7wpka.mongodb.net/movies"))
   if err != nil {
      log.Fatal(err)
   }
   ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
   err = client.Connect(ctx)
   if err != nil {
      log.Fatal(err)
   }

   defer client.Disconnect(ctx)
   err = client.Ping(ctx, readpref.Primary())
   if err != nil {
      log.Fatal(err)
   }
   databases, err := client.ListDatabaseNames(ctx, bson.M{})
   if err != nil {
      log.Fatal(err)
   }
   fmt.Println(databases)

   col := client. Database("movies"). Collection("list")


   cursor, err := col.Find(ctx, bson.M{"genres" : key})
   if err != nil {
      log.Fatal(err)
   }
   var episodes []bson.M
   if err = cursor.All(ctx, &episodes); err != nil {
      log.Fatal(err)
   }
   fmt.Println(episodes)

   json.NewEncoder(w).Encode(episodes)

}

func queryFromInternet(w http.ResponseWriter, r *http.Request){
   //vars := mux.Vars(r)
   //key := vars["id"]

   api := gomdb.Init("38966891")
   query := &gomdb.QueryData{Title: "Macbeth", SearchType: gomdb.MovieSearch}
   res, err := api.Search(query)
   if err != nil {
      fmt.Println(err)
      return
   }

   print(res)

    json.NewEncoder(w).Encode(res)

   query = &gomdb.QueryData{Title: "Macbeth", Year: "2015"}
   res2, err := api.MovieByTitle(query)
   if err != nil {
      fmt.Println(err)
      return
   }
   print(res2)

   res3, err := api.MovieByImdbID("tt2884018")
   if err != nil {
      fmt.Println(err)
      return
   }
   print(res3)
}

func main() {
   fmt.Println("Rest API v2.0 - Mux Routers")
   handleRequests()
}

