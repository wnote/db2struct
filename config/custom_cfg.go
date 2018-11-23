package config

var CustomCfg = map[string]string{
	"db_type":   "mysql",
	"file_mode": "0777",
	// Add table's prefix if exist
	"table_prefix": "",
	"db_dn":        "root:123456@tcp(127.0.0.1:3306)/antman_project?charset=utf8&parseTime=true&loc=Local",
	// path for go model directory
	"out_put_dir": "./model",
}
