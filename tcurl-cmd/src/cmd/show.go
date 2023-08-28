package cmd

import (
	"fmt"
	"github.com/ljtian/tcurl/tcurl-cmd/pkg/db"
	"github.com/ljtian/tcurl/tcurl-cmd/pkg/echarts"
	"github.com/ljtian/tcurl/tcurl-cmd/pkg/envVar"
	"github.com/spf13/cobra"
	"os"
)

var ClientName string

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.PersistentFlags().StringVarP(&ClientName, "ClientName", "C", envVar.GetEnvString(envVar.ClientName), "clientName")
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show all logs by clientName",
	Long:  `show all logs by clientName`,
	Run: func(cmd *cobra.Command, args []string) {
		if DbArgs.DbConnectUri == "" {
			fmt.Println("数据库相关内容为空")
			os.Exit(1)
		}
		if ClientName == "" {
			fmt.Println("ClientName 内容为空")
			os.Exit(1)
		}
		if err := db.StartDB(DbArgs.DbConnectUri); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		//logs, err := db.SelectLogsByClientName(ClientName)
		//if err != nil {
		//	fmt.Println(err)
		//	os.Exit(1)
		//}

		echarts.ShowWeb(ClientName)
	},
}
