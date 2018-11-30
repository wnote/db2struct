package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
	"github.com/urfave/cli"
	"github.com/wnote/db2struct/config"
	"github.com/wnote/db2struct/core"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	app := cli.NewApp()
	app.Name = "db2struct"
	app.Usage = "change table to golang model struct"
	app.Version = "v2.0"

	currentDir, err := filepath.Abs("./")
	if err != nil {
		fmt.Println(err)
		return
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: currentDir + "/config.json",
			Usage: "The file path for config.json",
		},
		cli.StringFlag{
			Name:  "tpl, t",
			Usage: "The file path for model template (It has default value in framework)",
		},
	}

	app.Action = func(c *cli.Context) error {
		configFile := c.String("config")
		configJson := core.GetFileContent(configFile)
		conf := config.Config{}
		err := json.Unmarshal(configJson, &conf)
		if err != nil {
			panic(err)
		}

		templatePath := c.String("tpl")
		var template string
		if templatePath != "" {
			templateByte, err := ioutil.ReadFile(templatePath)
			if err != nil {
				fmt.Println(err)
				fmt.Println("It will use framework template.")
			}
			template = string(templateByte)
		}
		if template == "" {
			template = config.MODEL_STRUCT_TEMPLATE
		}

		for _, out := range conf.OutputList {
			outDir := core.CheckAndMakeDir(out.OutputDir+"/", conf.FileMode)
			db := core.GetDb(conf.DbType, out.DbDn)
			allTables := core.GetAllTables(db)
			regTables := out.SelectTables
			for _, tableName := range allTables {
				if len(regTables) > 0 {
					matched := false
					for _, regStr := range regTables {
						reg, err := regexp.Compile(regStr)
						if err != nil {
							panic(err)
						}
						matched = reg.MatchString(tableName)
						if matched {
							break
						}
					}
					if !matched {
						continue
					}
				}
				// Replace the tag {struct_name},{table_name}.
				structStr, parseTableName := core.ReplaceStruct(template, out.TablePrefix, tableName)

				// Replace the tag {struct_field},{column_name},{package}
				packageName := core.GetPackageNameFromOutPutDir(outDir)
				structStr = core.ReplaceFields(db, structStr, packageName, tableName)

				// Write structStr to model file
				core.CreateAndWriteFile(outDir+parseTableName+".go", structStr, conf.FileMode)
			}
		}

		return nil
	}
	app.Run(os.Args)

	/*for _, customCfg := range config.CustomCfgList {
		outDir := utils.CheckAndMakeDir(customCfg["out_put_dir"]+"/", config.GlobalCfg["file_mode"])
		db := utils.GetDb(config.GlobalCfg["db_type"], customCfg["db_dn"])
		allTables := utils.GetAllTables(db)
		for _, tableName := range allTables {
			// Replace the tag {struct_name},{table_name}.
			structStr, parseTableName := utils.ReplaceStruct(config.MODEL_STRUCT_TEMPLATE, customCfg["table_prefix"], tableName)

			// Replace the tag {struct_field},{column_name},{package}
			packageName := utils.GetPackageNameFromOutPutDir(outDir)
			structStr = utils.ReplaceFields(db, structStr, packageName, tableName)

			// Write structStr to model file
			utils.CreateAndWriteFile(outDir+parseTableName+".go", structStr, config.GlobalCfg["file_mode"])
		}
	}*/
}
