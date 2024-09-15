package lodash

import "time"

func SetTimeZone(tz string) *time.Location {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		panic(err)
	}

	time.Local = loc
	return time.Local
}
