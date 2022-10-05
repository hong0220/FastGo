package base

// 公共请求参数
type Page struct {
	// todo title

	// 第几页
	PageNum int `json:"pageNum" p:"pageNum" d:"1" v:"bail|required|integer" dc:"第几页"`
	// 每页大小
	PageSize int `json:"pageSize" p:"pageSize" d:"10" v:"bail|required|integer" dc:"每页大小"`
	// 排序方式
	OrderBy string `json:"orderBy" p:"orderBy" dc:"排序方式，如果是多个排序是id asc, name desc"`
}

// 列表公共返回
type ListRes struct {
	PageNum int `json:"pageNum" p:"pageNum" dc:"当前页"`
	Total   int `json:"total" p:"total" dc:"总数"`
}
