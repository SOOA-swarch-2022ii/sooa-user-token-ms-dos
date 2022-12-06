package routes

import (
	control "github.com/SOOA-swarch-2022ii/sooa_user_token_ms_dos/controller"
	"github.com/gorilla/mux"
)

// Routes -> define endpoints
func Routes() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/reg-ur/", control.NewUser).Methods("POST")
	router.HandleFunc("/ur={username}", control.GetUserByUN).Methods("GET")
	//login route
	router.HandleFunc("/lg", control.Login).Methods("POST")


	return router
}
