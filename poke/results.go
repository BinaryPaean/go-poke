package poke

type Result struct {
	Name     string
	Target   string
	Response []string `json:"-"`
	Metrics  map[string]interface{}
	Err      error
}

type results []*Result

func NewResult(name, target string) *Result {
	r := &Result{
		Name:   name,
		Target: target}
	r.Metrics = make(map[string]interface{})
	return r
}
