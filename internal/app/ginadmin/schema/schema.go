package schema

// HTTPError HTTP响应错误
type HTTPError struct {
	Error HTTPErrorItem `json:"error" swaggo:"true,错误项"`
}

// HTTPErrorItem HTTP响应错误项
type HTTPErrorItem struct {
	Code    int    `json:"code" swaggo:"true,错误码"`
	Message string `json:"message" swaggo:"true,错误信息"`
}

// HTTPStatus HTTP响应状态
type HTTPStatus struct {
	Status string `json:"status" swaggo:"true,状态(OK)"`
}

// HTTPResponse HTTP响应列表数据
type HTTPResponse struct {
	Message   string      `json:"message" swaggo:"true,返回消息"`
	Result    interface{} `json:"result" swaggo:"true,返回结果"`
	Status    int         `json:"status" swaggo:"true,返回状态码"`
	Timestamp int64       `json:"timestamp" swaggo:"true,返回时间戳"`
}

// HTTPList HTTP响应列表数据
type HTTPPage struct {
	Data       interface{} `json:"data"`
	PageNo     int         `json:"pageNo"`
	PageSize   int         `json:"pageSize"`
	TotalPage  int         `json:"totalPage"`
	TotalCount int         `json:"totalCount"`
}

// HTTPList HTTP响应列表数据
type HTTPData struct {
	Data interface{} `json:"data"`
	// TotalCount int         `json:"totalCount"`
}

// PaginationParam 分页查询条件
type PaginationParam struct {
	PageIndex int // 页索引
	PageSize  int // 页大小
}

// PaginationResult 分页查询结果
type PaginationResult struct {
	Total int // 总数据条数
}

// Init 初始化
func Init() {
}
