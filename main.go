package main

import (
	"context"
	"example/Go-API/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)

	s := http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	signChan := make(chan os.Signal)
	signal.Notify(signChan, os.Interrupt)
	signal.Notify(signChan, os.Kill)

	sig := <-signChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}

// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:admin@friendbook.ac2qv.mongodb.net/?retryWrites=true&w=majority"))
// if err != nil {
// 	log.Fatal(err)
// }
// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// err = client.Connect(ctx)
// if err != nil {
// 	log.Fatal(err)
// }
// defer client.Disconnect(ctx)

// database := client.Database("friendbook-friends")
// friendsCollection := database.Collection("friends")

// cursor, err := friendsCollection.Find(ctx, bson.M{})
// if err != nil {
// 	log.Fatal(err)
// }

// // var friends []bson.M
// // if err = cursor.All(ctx, &friends); err != nil {
// // 	log.Fatal(err)
// // }
// // for _, friend := range friends {
// // 	fmt.Println(friend["username"])
// // }
// defer cursor.Close(ctx)

// for cursor.Next(ctx) {
// 	var friend bson.M
// 	if err = cursor.Decode(&friend); err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(friend)
// }
