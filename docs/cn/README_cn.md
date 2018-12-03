
# db2struct中文文档
[View English version](../../README.md)

## 1, 关于此工具
 - 快速将数据库表结构转换为golang模型的工具

## 2, 特性预览
* Change table struct to golang struct
* Support multiple output config

## 3, 安装使用

+ go get -u github.com/wnote/db2struct


+ 编辑 config.json

   cd /path/to/wnote/db2struct 进入目录

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
    db_type: 目前只支持mysql
    file_mode: 输出文件权限
    out_conf_list: 模型输出配置列表
       table_prefix: 定义的表名前辍，如果没有则不需要配置此项
       db_dn: 数据库连接字配置串
       output_dir: golang模型输出配置
       select_tables: 默认为空，空则表示数据库配置下的所有表都生成，可以使用正则表达式，如 ["project.*","test.*","table_name"]，只要匹配其中一个正则，则此表会输出
       model_package: 输出模型的package名字，默认为output_dir配置的目录名
    ```
    ### 简化版本 config.json
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
+ 配置文件示例

    [config.json 示例](./config.json)
    
    [model模板示例](./docs/example_model.tpl)

+ 如何运行此工具
    ```bash
    go run main.go
    go run main.go -c ./config.json   // -c没有配置，使用默认是 ./config.json
    go run main.go -c ./config.json -t example_model.tpl // -t没有设置时，使用默认配置 config/default_template.go定义好了
    ```

## 贡献

[如果你有什么bug修复及改进建议，请发送pull request](https://github.com/wnote/db2struct)
