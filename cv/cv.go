package cv

type CV struct {
	Info BaseInfo
	Exps []Experience
}

type BaseInfo struct {
	Name     string
	Age      string
	Position string
}

type Experience struct {
	Company  string
	Start    string
	End      string
	Position string
	Duty     []string
	Awards   []string
}

type Projects struct {
	PrjName string
	Techs   []string
	Details []string
}
