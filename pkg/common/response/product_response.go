package response

type CategoryResponse struct {
	Id       int    `json:"id"`
	ParentId int    `json:"parent_id"`
	Name     string `json:"name"`
	LeafNode int    `json:"leaf_node"`
}
