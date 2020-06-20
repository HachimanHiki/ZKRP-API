package selftype

type WesternMedicine struct {
	ID int	`json:"id"`
	DiseaseID string `json:"diseaseId"`
	DiseaseName string `json:"diseaseName"`
	MedicalDate	string `json:"medicalDate"`
	HospitalID int `json:"hospitalId"`
	HospitalName string `json:"hospitalName"`
}

type HospitalInfo struct {
	ID int	`json:"id"`
	DiseaseID string `json:"diseaseId"`
	DiseaseName string `json:"diseaseName"`
	OperationID	string `json:"operationId"`
	OperationName string `json:"operationName"`
	InHospitalDate string `json:"inHospitalDate"`
	OutHospitalDate string `json:"outHospitalDate"`
	HospitalID int `json:"hospitalId"`
	HospitalName string `json:"hospitalName"`
}

type ProveRequired struct {
	WesternMedicine []WesternMedicine `json:"westernMedicine"`
	HospitalInfo []HospitalInfo `json:"hospitalInfo"`
}