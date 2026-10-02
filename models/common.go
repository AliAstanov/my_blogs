package models

type Common struct {
	TableName  string `json:"table_name"`
	ColumnName string `json:"column_name"`
	ExpValue   any    `json:"exp_value"`
}

type CheckisExistsReq struct {
	IsExists bool
	Status   string
}

type GetListReq struct {
	Limit  int32  `json:"limit"`
	Page   int32  `json:"page"`
	Search string `json:"search"`
}
