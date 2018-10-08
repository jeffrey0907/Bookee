package comm

type PageRequest struct {
	PageNo   int `json:"pn" form:"pn,default=0"`
	PageSize int `json:"ps" form:"ps,default=10"`
}
