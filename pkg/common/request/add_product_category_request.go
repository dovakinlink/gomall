package request

type AddProductCategoryRequest struct {
	ParentId    int    `json:"parent_id"`
	Name        string `json:"name"`
	Level       int    `json:"level"`
	ShowStatus  int    `json:"show_status"`
	Sort        int    `json:"sort"`
	Keywords    string `json:"keywords"`
	Description string `json:"description"`
}
