package poke

type Result struct {
	Name     string
	Response []string
	Metrics  map[string]interface{}
	Err      error
}

type results []*Result

func NewResult() *Result {
	r := &Result{}
	r.Metrics = make(map[string]interface{})
	return r
}
