package request

type Request struct {
	Numbers []int `json:"numbers" binding:"required,unique,min=1"`
	Target  int   `json:"target" binding:"required"`
}
