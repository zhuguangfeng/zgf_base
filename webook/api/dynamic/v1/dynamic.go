package v1

// 发布动态
type PublishDynamicReq struct {
	Title     string   `json:"title"`     //标题
	Content   string   `json:"content"`   //内容
	Category  int8     `json:"category"`  //分类
	Resources []string `json:"resources"` //资源(图片or视频)
}
type PublishDynamicRes struct {
}
