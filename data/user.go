package data

import (
	"errors"
	"thyself/log"
	"thyself/util"
)

type User struct {
	User_ID string
	Email   string
	Tier    int
}

// Don't bother with time created and last logged in. They provide little value to us.
func Registeruser(email, rawPass string, tier int32) (string, error) {
	log.Info("Registering user " + email)
	user_id := util.GenID(5)
	pass_hash := util.HashPass(rawPass)
	_, err := SQL_CREATE_USER.Exec(email, user_id, pass_hash, tier)
	log.Debug(err, "Error registering user")
	return user_id, err
}

// Returns the user object with userid, email and tier.
func AuthUser(email, rawpass string) (User, error) {
	row := SQL_RETRIEVE_PASS.QueryRow(email)
	var pass_hash, user_id string
	var tier int
	if err := row.Scan(&pass_hash, &user_id, &tier); err != nil {
		log.Debug(err, "Error scanning username and password")
	} else if util.AuthPass(pass_hash, rawpass) {
		authUser := User{Email: email, User_ID: user_id, Tier: tier}
		return authUser, nil
	}
	return User{Email: email}, errors.New("Could not find user")
}
