package timewheel

type JobContext struct {
	Event interface{}
}

// Job 到达时间执行的Job
type Job func(key interface{}, jobContext JobContext)
