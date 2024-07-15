package api

type PageReq struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

func (p *PageReq) setPage() {
	if p.Page <= 0 {
		p.Page = 1
	}
}

func (p *PageReq) setSize() {
	if p.Size <= 0 {
		p.Size = 10
	}
}

type PageRes struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
}
