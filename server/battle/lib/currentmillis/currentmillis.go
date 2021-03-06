package currentmillis

import "time"

// StubNow For Test.
var StubNow func() int64

// Second.
const Second int64 = 1000

// Minute.
const Minute int64 = 60 * Second

// Now returns the number of the current time milliseconds elapsed since UNIX epoch time.
func Now() int64 {
	if StubNow != nil {
		return StubNow()
	}
	return time.Now().UnixNano() / 1000000
}

// Millis returns the number of the given time milliseconds elapsed since UNIX epoch time.
func Millis(t time.Time) int64 {
	return t.UnixNano() / 1000000
}

// Time returns the local Time corresponding to the given milliseconds,
func Time(ms int64) time.Time {
	sec := ms / 1000
	nsec := (ms % 1000) * 1000000
	return time.Unix(sec, nsec)
}

// Duration converts the given milliseconds to time.Duration,
func Duration(ms int64) time.Duration {
	return time.Duration(ms * 1000000)
}

// DurationMillis converts the given time.Duration to milliseconds,
func DurationMillis(d time.Duration) int64 {
	return d.Nanoseconds() / 1000000
}
