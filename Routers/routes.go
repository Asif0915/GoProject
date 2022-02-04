package Routers

import (
	"github.com/AsifIITR/mongodb-go1/PersonControllers"
	"github.com/gorilla/mux"
)

var RegisterBookandPersonStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/person", PersonControllers.CreatePerson).Methods("POST")
	router.HandleFunc("/people", PersonControllers.GetPerson).Methods("GET")
	router.HandleFunc("/person/{id}", PersonControllers.UpdatePerson).Methods("PUT")
	router.HandleFunc("/person/{id}", PersonControllers.DeletePerson).Methods("DELETE")
	router.HandleFunc("/people/{id}", PersonControllers.GetPeopleEndPoint).Methods("GET")
	//router.HandleFunc("/aggrt",PersonControllers.GetPeopleEndpoint).Methods("GET") // its not a specific method what should be hers
	// router.HandleFunc("/book", BookControllers.CreateBook).Methods("POST")
	// router.HandleFunc("/book/{id}", BookControllers.UpdateBook).Methods("PUT")
	// router.HandleFunc("/book1", BookControllers.GetBook).Methods("GET")
	// router.HandleFunc("/book/{id}", BookControllers.DeleteBook).Methods("DELETE")
	//for login pages and Authentication using JWT
	// router.HandleFunc("/login", Login_reqhandellers.LoginPageHandler).Methods("POST")
	// router.HandleFunc("/logout", Login_reqhandellers.LoginPageHandler).Methods("POST")
	// router.HandleFunc("/dashboard", Login_reqhandellers.LoginPageHandler)
}

// client = c
// 	fmt.Println(client)
//router := mux.NewRouter()
