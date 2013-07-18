package data

import (
	"errors"
	"thyself/log"
	"thyself/util"
)

type User struct {
	User_ID string
	Email   string
}

// Don't bother with time created and last logged in. They provide little value to us.
func Registeruser(email, rawPass string) (string, error) {
	log.Info("Registering user " + email)
	user_id := util.GenID(5)
	pass_hash := util.HashPass(rawPass)
	_, err := SQL_CREATE_USER.Exec(email, user_id, pass_hash)
	log.Debug(err, "Error registering user")
	return user_id, err
}

// Returns the user object with userid and email
func AuthUser(email, rawpass string) (User, error) {
	row := SQL_RETRIEVE_PASS.QueryRow(email)
	var pass_hash, user_id string
	if err := row.Scan(&pass_hash, &user_id); err != nil {
		log.Debug(err, "Error scanning username and password")
		return User{Email: email}, errors.New("User does not exist")
	} else {
		if util.AuthPass(pass_hash, rawpass) {
			return User{Email: email, User_ID: user_id}, nil
		} else {
			return User{Email: email}, errors.New("Username or password does not match")
		}
	} 
}
