package data

import (
	"fmt"
	"time"
)

type RateLimit struct {
	reset   int64
	key     string
	limit   int64
	per     int64
	current int64
}

/*
Free = 18 entries per day.
Basic = 100 entries per day. No roll-over
Premium = 310 entries per day. Rolls-over up to a total of 5000 entries.
Elite = 17,000 entries per month. No roll-over.
*/

const FREE_LIMIT = 18     // 24 entries
const FREE_EXPIRE = 86400 // 1 day
const FREE_ROLL_OVER_LIMIT = 0

const BASIC_LIMIT = 100
const BASIC_EXPIRES = 86400
const BASIC_ROLL_OVER_LIMIT = 0

const PREMIUM_LIMIT = 310
const PREMIUM_EXPIRES = 86400
const PREMIUM_ROLL_OVER_LIMIT = 5000

const ELITE_LIMIT = 17000     // It's huge because it's for a whole month.
const ELITE_EXPIRES = 2628000 // 1 month
const ELITE_ROLL_OVER_LIMIT = 0

func getExpireAt() time.Time {
	// Get and insert the user's plan's expire time here.
	rightNow := time.Now() // Save it so that it doesn't change while we're calling each field
	secondsToday := (rightNow.Hour() * 3600) + (rightNow.Minute() * 60) + rightNow.Second()
	// Assert that secondsToday < 86,400
	expireDuration := time.Duration(ELITE_EXPIRES - secondsToday) // param = seconds
	// Should return the time at midnight on that day
	expireDate := time.Now().Add(expireDuration)
	return expireDate
}

// Key = userid.
func initLimit(key_prefix string) {
	reset := time.Now().Unix() // per) * per + per
	key := key_prefix + fmt.Sprintf("%s", reset)
	/*
			limit = limit
			per = per
		   p = redis.pipeline()
		   p.incr(self.key)
		   p.expireat(self.key, self.reset + self.expiration_window)
		   self.current = min(p.execute()[0], limit)
		   remaining = property(lambda x: x.limit - x.current)
		   over_limit = property(lambda x: x.current >= x.limit)
	*/
	fmt.Println("Key is ", key)
}
