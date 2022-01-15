package model

type ProductCategory struct {
	Id          int    `gorm:"id" json:"id"`
	ParentId    int    `gorm:"parent_id" json:"parent_id"`     // 上级分类ID，0表示一级分类
	Name        string `gorm:"name" json:"name"`               // 分类名称
	Level       int    `gorm:"level" json:"level"`             // 分类级别： 0->1级 1->2级
	ShowStatus  int    `gorm:"show_status" json:"show_status"` // 是否展示 0:不展示，1:展示
	Sort        int    `gorm:"sort" json:"sort"`               // 展示排序
	Keywords    string `gorm:"keywords" json:"keywords"`       // 关键字
	Description string `gorm:"description" json:"description"` // 分类描述
	LeafNode    int    `gorm:"leaf_node" json:"leaf_node"`     // 是否为叶子节点分类，0:是，1:否
}

func (*ProductCategory) TableName() string {
	return "product_category"
}

type Brand struct {
	Id   int    `gorm:"id" json:"id"`
	Name string `gorm:"name" json:"name"` // 品牌名称
	Sort int    `gorm:"sort" json:"sort"` // 展示排序
	Logo string `gorm:"logo" json:"logo"` // logo图片
}

func (*Brand) TableName() string {
	return "brand"
}

type Member struct {
	Id         int    `gorm:"id" json:"id"`
	Username   string `gorm:"username" json:"username"` // 用户名（唯一标识）
	Password   string `gorm:"password" json:"password"` // 密码
	Nickname   string `gorm:"nickname" json:"nickname"` // 昵称
	Phone      string `gorm:"phone" json:"phone"`       // 手机号
	Status     int    `gorm:"status" json:"status"`     // 账号状态：0-启用 1-禁用
	CreateTime string `gorm:"create_time" json:"create_time"`
	UpdateTime string `gorm:"update_time" json:"update_time"`
}

func (*Member) TableName() string {
	return "member"
}

type Product struct {
	Id                  int     `gorm:"id" json:"id"`
	BrandId             int     `gorm:"brand_id" json:"brand_id"`                           // 品牌ID
	BrandName           string  `gorm:"brand_name" json:"brand_name"`                       // 品牌名称
	ProductCategoryId   int     `gorm:"product_category_id" json:"product_category_id"`     // 产品分类ID
	ProductCategoryName string  `gorm:"product_category_name" json:"product_category_name"` // 商品分类名称
	Name                string  `gorm:"name" json:"name"`                                   // 商品名称
	ProductSn           string  `gorm:"product_sn" json:"product_sn"`                       // 商品编号
	DeleteStatus        int     `gorm:"delete_status" json:"delete_status"`                 // 是否软删除，0:未删除，1:已删除
	PublishStatus       int     `gorm:"publish_status" json:"publish_status"`               // 上架状态，0:下架，1:上架
	Sort                int     `gorm:"sort" json:"sort"`                                   // 商品展示排序
	Sales               int     `gorm:"sales" json:"sales"`                                 // 销量
	Price               float64 `gorm:"price" json:"price"`                                 // 商品价格
	OriginalPrice       float64 `gorm:"original_price" json:"original_price"`               // 市场价格/展示价格
	Stock               int     `gorm:"stock" json:"stock"`                                 // 库存
	DetailHtml          string  `gorm:"detail_html" json:"detail_html"`                     // 商品详情HTML
	DetailMobileHtml    string  `gorm:"detail_mobile_html" json:"detail_mobile_html"`       // 商品详情手机端HTML
	CreateTime          string  `gorm:"create_time" json:"create_time"`
	UpdateTime          string  `gorm:"update_time" json:"update_time"`
}

func (*Product) TableName() string {
	return "product"
}
