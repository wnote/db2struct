package config

// table_prefix (Add table's prefix if exist.)
// db_dn (database dn config)
// out_put_dir (Path for go model directory.)
var CustomCfgList = []map[string]string{
	{
		"table_prefix": "geo_",
		"db_dn":        "root:123456@tcp(127.0.0.1:3306)/antman_geo?charset=utf8&parseTime=true&loc=Local",
		"out_put_dir":  "./model/antman_geo",
	},
	{
		"table_prefix": "pjt_",
		"db_dn":        "root:123456@tcp(127.0.0.1:3306)/antman_project?charset=utf8&parseTime=true&loc=Local",
		"out_put_dir":  "./model/antman_project",
	},
}

var GlobalCfg = map[string]string{
	"db_type":   "mysql",
	"file_mode": "0777",
}
