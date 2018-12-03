package config

type OutputConf struct {
	DbDn         string   `json:"db_dn"`
	OutputDir    string   `json:"output_dir"`
	TablePrefix  string   `json:"table_prefix"`
	SelectTables []string `json:"select_tables"`
	ModelPackage string   `json:"model_package"`
}

type Config struct {
	DbType      string       `json:"db_type"`
	FileMode    string       `json:"file_mode"`
	OutConfList []OutputConf `json:"out_conf_list"`
}
