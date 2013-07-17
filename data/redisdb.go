package data

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"thyself/log"
	"time"
)

var rdb, conErr = redis.DialTimeout("tcp", ":6379", 0, 1*time.Second, 1*time.Second)

type Word struct {
	Value         string
	ValueType     int
	IsVerb        string  `redis:"v"`
	IsNoun        string  `redis:"n"`
	Plural        string  `redis:"p"`
	Infinitive    string  `redis:"i"`
	CategoryQuery string  `redis:"c"`
	VerbCount     float64 `redis:"vCt"`
	NounCount     float64 `redis:"nCt"`
	IsFiltered    string  `redis:"f"`
}

var nilWord = Word{}

func (w *Word) String() string {
	return fmt.Sprintf("(Verb %s, Noun %s, Plurla %s, Infinitive %s, CategoryQuery %s, VerbCount %f, NounCount %f)",
		w.IsVerb, w.IsNoun, w.Plural, w.Infinitive, w.CategoryQuery, w.VerbCount, w.NounCount)
}

// Establish the database connection.
func RedisInit() {
	// Don't panic
	log.Debug(conErr, "  REDIS connection error ")
}

// TODO: Limit how many often this gets called. Don't attempt to reconnect more than once a second
func RedisReInit() bool {
	log.Info("REDIS : Attempting Reconnection ")
	rdb, conErr = redis.DialTimeout("tcp", ":6379", 0, 1*time.Second, 1*time.Second)
	if rdb != nil {
		log.Info("REDIS : ReConnected Successfully ")
		return true
	} else {
		log.Info("REDIS : Reconnection Failed ")
		return false
	}
}

func GetWord(query string) *Word {
	if rdb == nil {
		log.Info("REDIS : DB Down")
		if !RedisReInit() {
			return nil
		}
	}
	// println("Get Word : " + query)
	qWord := Word{}
	v, err := redis.Values(rdb.Do("HGETALL", "w:"+query))
	if err != nil {
		log.Debug(err, "NLP : Word : "+query+" : Not Found : ")
		return nil
	}
	redis.ScanStruct(v, &qWord)
	if qWord == nilWord {
		return nil
	} else {
		qWord.Value = query
	}
	return &qWord
}

func GetCategories(categoryQuery string) []string {
	if rdb == nil {
		log.Info("REDIS : DB Down")
		if !RedisReInit() {
			return nil
		}
	}
	// Return the top ten categories.
	v, err := redis.Strings(rdb.Do("ZREVRANGE", "c:"+categoryQuery, 0, 10))
	log.Debug(err, "NLP : Category : "+categoryQuery+" : Not Found ")
	return v
}
