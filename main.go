package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type User struct {
	Email   string
	Name    string
	Address string
	Job     string
	Image 	string
}

var users = []User{
	{
		Email:   "maxblack@gmail.com",
		Name:    "Max Black",
		Address: "Brooklyn, New York City",
		Job:     "Waitress and Baker",
		Image:   "https://i.pinimg.com/736x/d3/81/cb/d381cb0a2d103be6b0a9208071d744ef.jpg",
	},
	{
		Email:   "carolinechanning@gmail.com",
		Name:    "Caroline Channing",
		Address: "Manhattan, New York City",
		Job:     "Waitress and Enterpreneur",
		Image: "https://westweekever.com/wp-content/uploads/2013/01/channing.jpg",
	},
	{
		Email:   "earle@gmail.com",
		Name:    "Earl",
		Address: "Brooklyn, New York City",
		Job:     "Cashier",
		Image: "https://qph.cf2.quoracdn.net/main-qimg-81584b0118b197a8d46cccc8f52d2a16-lq",
	},
	{
		Email:   "oleg@gmail.com",
		Name:    "Oleg",
		Address: "Brooklyn, New York City",
		Job:     "Chef",
		Image: "https://artworks.thetvdb.com/banners/person/461027/61517658.jpg",
	},
	{
		Email:   "sophie@gmail.com",
		Name:    "Sophie Kachinsky",
		Address: "Brooklyn, New York City",
		Job:     "Owner of Sophie’s Cleaning Company",
		Image: "https://www.hallofseries.com/wp-content/uploads/2015/11/2-broke-girls-sophie.jpg",
	},
	{
		Email:   "hanlee@gmail.com",
		Name:    "Han Lee",
		Address: "Brooklyn, New York City",
		Job:     "Owner of Williamsburg Diner",
		Image: "https://ew.com/thmb/d2Om5GG3rMRcorUr98MtuaK-lXk=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/matthew-moy-2-broke-girls_510x380-a76c8df08e3b4e188c660bd05bc664d1.jpg",
	},
}

var PORT = ":2020"

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/profile", profile)
	http.HandleFunc("/", getUsers)
	http.HandleFunc("/user-list", getUsers)

	http.ListenAndServe(PORT, nil)
}

func profile(w http.ResponseWriter, r *http.Request) {
	// check method
	if r.Method != "GET" {
		http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}

	// get email from query params
	email := r.URL.Query().Get("email")
	
	// get user by email
	var user User
	for _, u := range users {
		if u.Email == email {
			user = u
			break
		}
	}

	// render profile page
	tpl, err := template.ParseFiles("profile.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, user)
}

func login(w http.ResponseWriter, r *http.Request) {
	// check method
	if r.Method == "POST" {
		email := r.FormValue("email")
		// get user by email
		var user User
		for _, u := range users {
			if u.Email == email {
				user = u
				break
			} 
		} 
		// if user not found, render login page with error message
		if user.Email == "" {
			tpl, err := template.ParseFiles("notfound.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			tpl.Execute(w, "User not found")
			return
		}
		
		// redirect to profile page
		redirectURL := fmt.Sprintf("/profile?email=%s", user.Email)
		http.Redirect(w, r, redirectURL, http.StatusSeeOther)
		return
	}

	// render login page for GET request
	if r.Method == "GET" {
		tpl, err := template.ParseFiles("login.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tpl.Execute(w, nil)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	// parsing to user-list.html
	tpl, err := template.ParseFiles("user-list.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, users)
}