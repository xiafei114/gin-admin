package demo

// CommonFile Product对象
type CommonFile struct {
	RecordID    string `json:"record_id" swaggo:"false,记录ID"`
	ContentType string `json:"content_type" binding:"required" swaggo:"true,内容类型"`
	FileName    string `json:"file_name" binding:"required" swaggo:"true,文件名称"`
	FilePath    string `json:"file_path"  swaggo:"false,文件位置"` // 文件位置
	FileType    int    `json:"file_type"  swaggo:"false,文件类型"` // 文件类型
	FileURL     string `json:"file_url"  swaggo:"false,文件url"` // 文件位置
}
