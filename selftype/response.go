package selftype

type JSONResponse struct {
    Status string	`json:"status" example:"success"`
    Message string	`json:"message"`
}

type ProvePackagesResponse struct {
	JSONResponse
	ProvePackage []ProvePackage `json:"data "`
}

type EventResponse struct {
	JSONResponse
	Event Event `json:"data"`
}