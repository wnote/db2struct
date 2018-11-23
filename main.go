package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/wnote/db2struct/config"
	"github.com/wnote/db2struct/utils"
)

func main() {
	outDir := utils.CheckAndMakeDir(config.CustomCfg["out_put_dir"]+"/", config.CustomCfg["file_mode"])

	allTables := utils.GetAllTables()
	for _, tableName := range allTables {
		// Replace the tag {struct_name},{table_name}.
		structStr, parseTableName := utils.ReplaceStruct(config.MODEL_FILE_TEMPLATE, tableName)

		// Replace the tag {struct_field},{column_name},{package}
		packageName := utils.GetPackageNameFromOutPutDir(outDir)
		structStr = utils.ReplaceFields(structStr, packageName, tableName)

		// Write structStr to model file
		utils.CreateAndWriteFile(outDir+parseTableName+".go", structStr, config.CustomCfg["file_mode"])
	}
}
