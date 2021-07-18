package models

type ResponseData struct {
	SheetNameList []string     `json:"sheet_name_list"`
	Sheet         *Sheet       `json:"sheet"`
	SheetList     []*SheetList `json:"sheet_list"`
}

type Sheet struct {
	TableHeader []string   `json:"table_header"`
	TableData   [][]string `json:"table_data"`
}

type SheetList struct {
	SheetIndex int    `json:"sheet_index"`
	SheetName  string `json:"sheet_name"`
}
