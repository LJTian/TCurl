package cmd

import (
	"fmt"
	"github.com/ljtian/tcurl/tcurl-cmd/pkg/curl"
	"github.com/ljtian/tcurl/tcurl-cmd/pkg/db"
	"github.com/ljtian/tcurl/tcurl-cmd/pkg/define"
	"github.com/ljtian/tcurl/tcurl-cmd/pkg/envVar"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var AllArgs define.TCurl
var DbArgs define.DBInfo

func init() {

	// 设置默认环境变量，无值时才会设置
	envVar.SetDefaultEnv([]envVar.DefaultEnv{
		{envVar.Uri, ""},
		{envVar.ClientName, ""},
		{envVar.Times, "100"},
		{envVar.TimeInterval, "5"},
		{envVar.TimeOut, "5"},
		{envVar.CoroutineNum, "5"},
		{envVar.SaveDB, "false"},
		{envVar.DbConnectUri, ""},
	})
	rootCmd.PersistentFlags().StringVarP(&AllArgs.Uri, "uri", "U", envVar.GetEnvString(envVar.Uri), "web service addr")
	rootCmd.PersistentFlags().StringVarP(&AllArgs.ClientName, "clientName", "N", envVar.GetEnvString(envVar.ClientName), "client name")
	rootCmd.PersistentFlags().IntVarP(&AllArgs.Times, "times", "T", envVar.GetEnvInt(envVar.Times), "Number of cycles")
	rootCmd.PersistentFlags().IntVarP(&AllArgs.Intervals, "intervals", "I", envVar.GetEnvInt(envVar.TimeInterval), "Intervals")
	rootCmd.PersistentFlags().IntVarP(&AllArgs.TimeOut, "timeout", "t", envVar.GetEnvInt(envVar.TimeOut), "timeout period")
	rootCmd.PersistentFlags().IntVarP(&AllArgs.CoroutineNum, "coroutineNum", "c", envVar.GetEnvInt(envVar.CoroutineNum), "coroutine number")
	rootCmd.PersistentFlags().BoolVarP(&AllArgs.SaveDB, "saveDB", "S", envVar.GetEnvBool(envVar.SaveDB), "Whether the data is saved in the database")
	rootCmd.PersistentFlags().StringVarP(&DbArgs.DbConnectUri, "dbUri", "D", envVar.GetEnvString(envVar.DbConnectUri), "Database connect address")
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// 生成随机字符串
func generateRandomString(length int) string {

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

var rootCmd = &cobra.Command{
	Use:   "TCurl",
	Short: "TCurl is an http access command client",
	Long:  `TCurl is an http client mainly used for web service access and recording command line programs`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(AllArgs)
		if AllArgs.Uri == "" {
			fmt.Println("URI 为空， 请使用 -h 查看使用说明")
			os.Exit(1)
		}
		if AllArgs.SaveDB {
			if DbArgs.DbConnectUri == "" {
				fmt.Println("数据库相关内容为空")
				os.Exit(1)
			}
			if err := db.StartDB(DbArgs.DbConnectUri); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		fmt.Println("开始执行")
		if AllArgs.ClientName == "" {
			AllArgs.ClientName = fmt.Sprintf("clientName_%s", generateRandomString(5))
			fmt.Printf(" ClientName is [%s]\n", AllArgs.ClientName)
		}

		ch1 := make(chan int, AllArgs.CoroutineNum)
		for i := 0; i < AllArgs.CoroutineNum; i++ {
			go func(i int, ch chan int) {
				if err := curl.Run(AllArgs, i); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				ch1 <- 1
			}(i, ch1)
		}
		for i := 0; i < AllArgs.CoroutineNum; i++ {
			<-ch1
		}
		fmt.Println("循环结束")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
