package apis

type Example struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ExampleList struct {
	List  []ExampleList `json:"list"`
	Total int64         `json:"total"`
}

type ExampleListParams struct {
	Username *string `json:"username" form:"username"`
	Password *string `json:"password" form:"password"`
	Limit    *int    `json:"limit" form:"pageSize"`
	Offset   *int    `json:"offset" form:"current"`
}
