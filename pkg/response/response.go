package response

type Response struct {
	Solution []Solution `json:"solution"`
}

type Solution []int
