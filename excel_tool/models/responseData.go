package models

type ResponseData struct {
	SheetNameList []string `json:"sheet_name_list"`
	Sheet         *Sheet   `json:"sheet"`
}

type Sheet struct {
	TableHeader []string   `json:"table_header"`
	TableData   [][]string `json:"table_data"`
}
