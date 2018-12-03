
# db2struct  Documentation
[点我查看中文版](./docs/README_cn.md)

## 1, About
 - A tool for change database(mysql) struct to go struct

## 2, Overview
* Change table struct to golang struct
* Support multiple output config

## 3, Getting Started

+ go get -u github.com/wnote/db2struct


+ Edit config.json

   cd /path/to/wnote/db2struct

   ### vi config.json
     ```json
     {
       "db_type":   "mysql",
       "file_mode": "0777",
       "out_conf_list":[
         {
           "table_prefix": "pjt_",
           "db_dn":        "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true&loc=Local",
           "output_dir":  "./model/test",
           "select_tables":["project.*"],
           "model_package":"test_package"
         }
       ]
     }
    ```
    ```
    db_type: current only support mysql
    file_mode: output file permission
    out_conf_list: model output config list
       table_prefix: If all table has prefix ,then set it,otherwise fill blank
       db_dn: mysql connection db
       output_dir: model output dir
       select_tables: null is default,  all table will output the models
                      // regexp,default is all table in this database
    ```
    ### Simple config.json
         ```json
         {
           "out_conf_list":[
             {
               "db_dn":        "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true&loc=Local",
               "output_dir":  "./model/test",
             }
           ]
         }
        ```
+ Configure example
    [config.json Example](./config.json)
    [Template Example](./docs/example_model.tpl)

+ Run tools
    ```bash
    // default is ./config.json
    go run main.go
    // Ts has default template in config/default_template.go
    go run main.go -c ./config.json 
    go run main.go -c ./config.json -t example_model.tpl
    ```

## Contributing

[You can send pull request to deliver a better db2struct](https://github.com/wnote/db2struct)

### Contributor
   * aengineer@foxmail.com

## License

© wnote, 2018~time.Now

Released under the [MIT License](https://github.com/wnote/db2struct/blob/master/License)

