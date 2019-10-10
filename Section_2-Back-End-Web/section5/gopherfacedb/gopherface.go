package main

import (
	"log"

	"github.com/EngineerKamesh/gofullstack/volume2/section5/gopherfacedb/common"
	"github.com/EngineerKamesh/gofullstack/volume2/section5/gopherfacedb/common/datastore"
	"github.com/EngineerKamesh/gofullstack/volume2/section5/gopherfacedb/handlers"
	"github.com/EngineerKamesh/gofullstack/volume2/section5/gopherfacedb/middleware"

	"net/http"
	"os"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	WEBSERVERPORT = ":8080"
)

func main() {

	db, err := datastore.NewDatastore(datastore.MYSQL, "gopherface:gopherface@/gopherfacedb")
	//db, err := datastore.NewDatastore(datastore.MONGODB, "localhost:27017")
	//db, err := datastore.NewDatastore(datastore.REDIS, "localhost:6379")

	if err != nil {
		log.Print(err)
	}

	defer db.Close()

	env := common.Env{DB: db}

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/register", handlers.RegisterHandler).Methods("GET,POST")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST")
	r.HandleFunc("/feed", handlers.FeedHandler).Methods("GET")
	r.HandleFunc("/friends", handlers.FriendsHandler).Methods("GET")
	r.HandleFunc("/find", handlers.FindHandler).Methods("GET,POST")
	r.HandleFunc("/profile", handlers.MyProfileHandler).Methods("GET")
	r.HandleFunc("/profile/{username}", handlers.ProfileHandler)
	r.HandleFunc("/triggerpanic", handlers.TriggerPanicHandler)
	r.HandleFunc("/foo", handlers.FooHandler)
	r.Handle("/signup", handlers.SignUpHandler(&env)).Methods("GET", "POST")
	r.HandleFunc("/postpreview", handlers.PostPreviewHandler).Methods("GET", "POST")
	r.HandleFunc("/upload-image", handlers.UploadImageHandler).Methods("GET", "POST")
	r.HandleFunc("/upload-video", handlers.UploadVideoHandler).Methods("GET", "POST")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.Handle("/", middleware.PanicRecoveryHandler(ghandlers.LoggingHandler(os.Stdout, r)))
	http.ListenAndServe(WEBSERVERPORT, nil)

}
