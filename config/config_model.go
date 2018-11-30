package config

type OutputConf struct {
	DbDn         string   `json:"db_dn"`
	OutputDir    string   `json:"output_dir"`
	TablePrefix  string   `json:"table_prefix"`
	SelectTables []string `json:"select_tables"`
}

type Config struct {
	DbType     string       `json:"db_type"`
	FileMode   string       `json:"file_mode"`
	OutputList []OutputConf `json:"output_list"`
}
