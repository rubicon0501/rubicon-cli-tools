package cmd

import (
	"log"
	"rubicon-cli-tools/internal/sql2struct"

	"github.com/spf13/cobra"
)

var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var tableName string

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql转换和处理",
	Long:  "sql转换和处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql转换为struct",
	Long:  "sql转换为struct",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:  dbType,
			Host:    host,
			User:    username,
			Passwd:  password,
			Charset: charset,
		}
		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect err: %v", err)
		}
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v", err)
		}

		template := sql2struct.NewStructTemplate()
		templateColumns := template.AssemblyColumns(columns)
		err = template.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("template.Generate err: %v", err)
		}
	},
}

func init() {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "u", "", "请输入数据库的账号")
	sql2structCmd.Flags().StringVarP(&password, "password", "p", "", "请输入数据库的密码")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "请输入数据库的IP和端口")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "c", "utf8mb4", "请输入数据库的编码")
	sql2structCmd.Flags().StringVarP(&dbType, "type", "t", "mysql", "请输入数据库实例类型")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "", "", "请输入数据库名称")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "", "", "请输入表名称")
}
