package response

type OrderPropsSuccessResponse struct {
	Status string      `json:"status"`
	Value  []Requisite `json:"value"`
}
