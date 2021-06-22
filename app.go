package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"io/ioutil"

	userdb "github.com/anandraj-coder/UserAPI/DB"
	"github.com/gorilla/mux"
)

/*vars usage string := "Welcome to Literature User API\n
user json : {ID: '"1"', FName: '"Anand"', LName: '"Rajagopalan"', Email: '"anandraj@yahoo.com"'}
Usage:\n
Return all users (GET) : /users
Return a user (GET) : /user/{id}
Add a user (POST) : /user
Update a user (PUT): /USER/{id}
Delete a user (DELETE): /user/{id}"*/

type userProfile struct {
	ID    string `json:"Id"`
	FName string `json:"fname"`
	LName string `json:"lname"`
	Email string `json:"email"`
}

const addr = ":8080"

var userProfiles []userProfile

func main() {
	userProfiles = []userProfile{
		{ID: "1", FName: "Anand", LName: "Rajagopalan", Email: "anandraj@yahoo.com"},
		{ID: "2", FName: "Jayashree", LName: "Anand", Email: "jayanalak@hotmail.com"},
		{ID: "3", FName: "Gavin", LName: "Leo Rhynie", Email: "jayanalak@hotmail.com"},
		{ID: "4", FName: "Brian", LName: "Thomson", Email: "jayanalak@hotmail.com"},
		{ID: "5", FName: "Gary", LName: "Mccormick", Email: "blah@hotmail.com"},
	}
	handleRequests()
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/users", returnalAllUsers)
	myRouter.HandleFunc("/user", createNewUser).Methods("POST")
	myRouter.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{id}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/user/{id}", returnUser)
	log.Fatal(http.ListenAndServe(addr, myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	var output = userdb.InitUser()
	fmt.Fprintf(w, output)
}

func returnalAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllUsers")
	//fmt.Println(userProfiles)
	json.NewEncoder(w).Encode(userProfiles)
}

func returnUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnUser")
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprintf(w, "Key: "+key)
	for _, user := range userProfiles {
		if user.ID == key {
			json.NewEncoder(w).Encode(user)
		}
	}
}
func createNewUser(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	fmt.Println("Endpoint Hit: createNewUser")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newUser userProfile
	json.Unmarshal(reqBody, &newUser)
	// update our global Articles array to include
	// our new Article
	userProfiles = append(userProfiles, newUser)
	json.NewEncoder(w).Encode(newUser)
	//fmt.Fprintf(w, "%+v", string(reqBody))
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteUser")
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	id := vars["id"]

	// we then need to loop through all our articles
	for index, user := range userProfiles {
		// if our id path parameter matches one of our
		// articles
		if user.ID == id {
			// updates our Articles array to remove the
			// article
			userProfiles = append(userProfiles[:index], userProfiles[index+1:]...)
		}
	}

}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateUser")
	vars := mux.Vars(r)
	id := vars["id"]
	for index, user := range userProfiles {
		user = userProfiles[index]
		if user.ID == id {
			reqBody, _ := ioutil.ReadAll(r.Body)
			var newUser userProfile
			json.Unmarshal(reqBody, &newUser)
			fmt.Println("Current User ID:" + userProfiles[index].ID)
			fmt.Println("New User ID:" + newUser.ID)
			userProfiles[index].ID = newUser.ID
			fmt.Println("Current User ID after update:" + userProfiles[index].ID)
			userProfiles[index].FName = newUser.FName
			userProfiles[index].LName = newUser.LName
			userProfiles[index].Email = newUser.Email
		}
	}
}
