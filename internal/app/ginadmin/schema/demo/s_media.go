package demo

// Media Product对象
type Media struct {
	RecordID     string `json:"record_id" swaggo:"false,记录ID"`
	CommonFileID string `json:"common_file_id" binding:"required" swaggo:"true,内容类型"`
	InfoNo       string `json:"info_no" binding:"required" swaggo:"true,信息编号"`
	InfoDesc     string `json:"info_desc"  swaggo:"false,信息说明"`
}
