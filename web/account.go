package web

import (
	"fmt"
	"net/http"
	"strings"
	"thyself/data"
)

// The regex to validate Email is pretty messy, so just check length
// and if there's a @ and .
func ValidateEmail(emailAddr string) bool {
	return strings.Count(emailAddr, "@") == 1 && strings.Count(emailAddr, ".") >= 1 &&
		len(emailAddr) > 4 && len(emailAddr) < 32
}

func ValidatePassword(password string) bool {
	return len(password) >= 6 && len(password) < 32
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	repass := r.FormValue("repass")
	password := r.FormValue("password")

	session, _ := cookieStore.Get(r, defaultSessionName)

	valid := true

	fmt.Printf("Email is %s (%s) , pass is %s - re %s \n", email, ValidateEmail(email), password, repass)
	if !ValidatePassword(password) {
		session.AddFlash("alert : Password must be atleast 6 letters long")
		valid = false
	} else if repass != password {
		session.AddFlash("alert : Passwords do not match")
		valid = false
	}
	if !ValidateEmail(email) {
		session.AddFlash("alert : That doesn't look like a valid email address")
		valid = false
	}

	if valid {
		user, err := data.Registeruser(email, password, 0)
		if err == nil {
			// Log the user in
			session.Values["user_id"] = user
			session.AddFlash("success : Registration successful. Welcome to Thyself.io!")
			session.Save(r, w)

			HomepageHandler(w, r)
		} else {
			session.AddFlash("error : User exists already")
			valid = false
		}
	}

	// intentionally not else if. valid is set to false on failed registration above
	if !valid {
		errors := BuildMessages(w, r)
		renderedPage := string(TemplateMessage.Render(map[string]string{
			"message": errors + PartialRegisterForm}, nil))
		fmt.Fprintln(w, renderedPage)
	}

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")
		session, _ := cookieStore.Get(r, defaultSessionName)

		if ValidateEmail(email) && ValidatePassword(password) {
			user, err := data.AuthUser(email, password)
			if err == nil {
				session.Values["user_id"] = user.User_ID
				session.AddFlash("success : Login Successful!")
				session.AddFlash("error : Could not Login. Username or Password does not match.")
			}
		} else {
			session.AddFlash("error : Username and password doesn't seem right")
		}
		session.Save(r, w)
	}

	if isAuth(r) {
		HomepageHandler(w, r) // Homepage will redirect you to the journal page
	} else {
		errors := BuildMessages(w, r)
		renderedPage := string(TemplateMessage.Render(map[string]string{
			"message": errors + PartialLoginForm}, nil))
		fmt.Fprintln(w, renderedPage)
	}
	//http.Redirect(w, r, "/", 302)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := cookieStore.Get(r, defaultSessionName)
	session.Values["user_id"] = nil
	session.Save(r, w)
	HomepageHandler(w, r)
}

func isAuth(r *http.Request) bool {
	return GetLoggedInUser(r) != ""
}

func GetLoggedInUser(r *http.Request) string {
	session, _ := cookieStore.Get(r, defaultSessionName)
	if session != nil && session.Values["user_id"] != nil {
		return session.Values["user_id"].(string)
	} else {
		return ""
	}

}
