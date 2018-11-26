# db2struct
A tool for change database(mysql) struct to go struct

## Get Start
```bash
go get -u github.com/wnote/db2struct
cd /path/to/wnote/db2struct
# Edit config/custom_cfg.go. Change the db_dn,out_put_dir 
# Add table_prefix if has.
vim config/custom_cfg.go
# Edit config/struct_file_tpl.go if neccessary
# Edit config/db_struct_map.go if neccessory
go run main.go
```