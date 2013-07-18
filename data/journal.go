package data

import (
	"time"
	//	"strings"
	"errors"
	"thyself/log"
	"thyself/util"
)

type JournalEntry struct {
	User_ID string `json:"user_id"`
	Je_Time int64  `json:"time"`
	Je_Text string `json:"text"`
}

func UpsertJournalEntry(user_id string, t time.Time, je_text string) {
	log.Info("data : Add Journal Entry : ", user_id, " : time : ", t)
	existingJe, err := GetJournalEntry(user_id, t)

	if err == nil { // a journal entry exists already. update it if it's different
		start_date, end_date := util.GetDayEndpoints(t)
		if je_text != existingJe.Je_Text { // text is different
			_, err := SQL_UPDATE_JE.Exec(t.Unix(), je_text, user_id, start_date, end_date)
			log.Debug(err, "DATA: JournalEntry : Update : FAILED : ")
		}
	} else {
		_, err := SQL_ADD_JE.Exec(user_id, t.Unix(), je_text)
		log.Debug(err, "DATA: JournalEntry : Insert : FAILED : ")
	}
}

func GetJournalEntry(user_id string, t time.Time) (JournalEntry, error) {
	start_date, end_date := util.GetDayEndpoints(t)
	row := SQL_RETRIEVE_JE.QueryRow(user_id, start_date, end_date)
	je := JournalEntry{User_ID: user_id, Je_Text: ""}
	if err := row.Scan(&je.Je_Time, &je.Je_Text); err != nil {
		return je, errors.New("Could not find journal entry")
	} else {
		return je, nil
	}
}
