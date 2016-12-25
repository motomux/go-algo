package leetcode

// 359. Logger Rate Limiter
type Logger struct {
	logs map[string]int
}

/** Initialize your data structure here. */
func Constructor() Logger {
	return Logger{
		logs: make(map[string]int),
	}
}

/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
  If this method returns false, the message will not be printed.
  The timestamp is in seconds granularity. */
func (this *Logger) Shouldprintmessage(timestamp int, message string) bool {
	lastTs, ok := this.logs[message]
	if ok && timestamp-lastTs < 10 {
		return false
	}
	this.logs[message] = timestamp
	return true
}
