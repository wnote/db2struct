package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

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
		panic(err)
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

	app.Action = AppAction
	app.Run(os.Args)
}

// Application run entry
func AppAction(c *cli.Context) error {
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
			fmt.Println(fmt.Sprintf("%v. It will use framework template.", err))
		}
		template = string(templateByte)
	}
	if template == "" {
		template = config.MODEL_STRUCT_TEMPLATE
	}

	confFileModel := conf.FileMode
	if confFileModel == "" {
		confFileModel = config.FileModel
	}

	for _, outConf := range conf.OutConfList {
		outDir := core.CheckAndMakeDir(outConf.OutputDir+"/", confFileModel)
		dbType := conf.DbType
		if dbType == "" {
			dbType = config.DbType
		}
		db := core.GetDb(dbType, outConf.DbDn)
		allTables := core.GetAllTables(db)
		regTables := outConf.SelectTables
		for _, tableName := range allTables {
			if matched := core.StringMatchSlice(regTables, tableName); !matched {
				continue
			}
			structStr, parseTableName := core.ReplaceStruct(template, outConf.TablePrefix, tableName)
			packageName := outConf.ModelPackage
			if packageName == "" {
				// Replace the tag {struct_field},{column_name},{package}
				packageName = core.GetPackageNameFromOutPutDir(outDir)
			}
			structStr = core.ReplaceFields(db, structStr, packageName, tableName)

			core.CreateAndWriteFile(outDir+parseTableName+".go", structStr, confFileModel)
		}
	}
	return nil
}
