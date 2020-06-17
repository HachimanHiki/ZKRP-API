package selftype

type Deagnosis struct {
	DeagnosisCode string
	DuringTime string
}

type Procedure struct {
	ProcedureCode string
	DuringTime string
}

type ProveRequired struct {
	RequestTime string `json:"requestTime"`
	Deagnosis []Deagnosis `json:"deagnosis"`
	Procedure []Procedure `json:"procedure"`
}

type ProvePackage struct {
	Type string `json:"type"`
	Code string `json:"code"`
	Prove string `json:"prove"`
	Lowerbound int64 `json:"lowerbound"`
	Upperbound int64 `json:"upperbound"`
	Commitment string `json:"commitment"`
}

type Verify struct {
	ProvePackages []ProvePackage `json:"provePackages"`
}


type EventName struct {
	Name string `json:"name"`
}

type Event struct {
	EventName string `json:"eventName"`
	Deagnosis []Deagnosis `json:"deagnosis"`
	Procedure []Procedure `json:"procedure"`
}