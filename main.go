package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/wnote/db2struct/config"
	"github.com/wnote/db2struct/utils"
)

func main() {
	for _, customCfg := range config.CustomCfgList {
		outDir := utils.CheckAndMakeDir(customCfg["out_put_dir"]+"/", config.GlobalCfg["file_mode"])
		db := utils.GetDb(config.GlobalCfg["db_type"], customCfg["db_dn"])
		allTables := utils.GetAllTables(db)
		for _, tableName := range allTables {
			// Replace the tag {struct_name},{table_name}.
			structStr, parseTableName := utils.ReplaceStruct(config.MODEL_FILE_TEMPLATE, customCfg["table_prefix"], tableName)

			// Replace the tag {struct_field},{column_name},{package}
			packageName := utils.GetPackageNameFromOutPutDir(outDir)
			structStr = utils.ReplaceFields(db, structStr, packageName, tableName)

			// Write structStr to model file
			utils.CreateAndWriteFile(outDir+parseTableName+".go", structStr, config.GlobalCfg["file_mode"])
		}
	}
}
