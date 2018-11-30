# db2struct
A tool for change database(mysql) struct to go struct

## Get Start
```bash
go get -u github.com/wnote/db2struct
cd /path/to/wnote/db2struct
# Edit config.json
{
  "db_type":   "mysql",
  "file_mode": "0777",
  "output_list":[
    {
      "table_prefix": "pjt_",
      "db_dn":        "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true&loc=Local",
      "output_dir":  "./model/test",
      "select_tables":["project.*"] // regexp,default is all table in this database
    }
  ]
}
// default is ./config.json
go run main.go
// Ts has default template in config/default_template.go
go run main.go -c ./config.json 
go run main.go -c ./config.json -t example_model.tpl
```