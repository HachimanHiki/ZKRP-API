package selftype

type Deagnosis struct {
	DeagnosisCode string `json:"deagnosisCode" example:"J0190"`
	DuringTime int `json:"duringTime" example:"28"`
}

type DeagnosisDate struct {
	Deagnosis
	Date string `json:"date" example:"20200505"`
}

type Procedure struct {
	ProcedureCode string `json:"procedureCode" example:"3950"`
	DuringTime int `json:"duringTime" example:"28"`
}

type ProcedureDate struct {
	Procedure
	Date string `json:"date" example:"20200505"`
}
/*
type ProveRequired struct {
	RequestTime string `json:"requestTime" example:"20200531"`
	Deagnosis []DeagnosisDate `json:"deagnosisDate"`
	Procedure []ProcedureDate `json:"procedureDate"`
}
*/
type ProvePackage struct {
	Type string `json:"type" example:"Deagnosis"`
	Code string `json:"code" example:"J0390"`
	Prove string `json:"prove"`
	Lowerbound int64 `json:"lowerbound" example:"20200503"`
	Upperbound int64 `json:"upperbound" example:"20200531"`
	Commitment string `json:"commitment" example:"0x1af2e641e834f6c503e0e6b9b593323342e7d45ad0aa299bc934110a4e13ae2f33c3a0fc695f397d51e8b4699b4d9f9397b3bc3e20fa42cb1a540ff8f19dc118"`
}

type Verify struct {
	ProvePackages []ProvePackage `json:"provePackages"`
	//UserName string `json:"userName" example:"\u738b\u6625\u5b0c"`
}

type Event struct {
	EventName string `json:"eventName" example:"marathon"`
	Deagnosis []Deagnosis `json:"deagnosis"`
	Procedure []Procedure `json:"procedure"`
	VarifyURL string `json:"varifyURL" example:"http://localhost:8080/verify"`
}