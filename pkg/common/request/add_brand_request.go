package request

type AddBrandRequest struct {
	Name string `json:"name"`
	Sort int    `json:"sort"`
	Logo string `json:"logo"`
}
