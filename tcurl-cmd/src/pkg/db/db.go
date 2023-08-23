package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func SendDb(data []byte, clineName string, lastTime time.Time) (err error) {

	if lastTime.IsZero() {
		lastTime = time.Now()
	}
	//fmt.Printf("DBSend data is [%s]\n", data)
	var log Logs
	if len(data) != 0 {
		if err = json.Unmarshal(data, &log); err != nil {
			fmt.Println(err)
			return
		}
		NaTime := time.Since(lastTime).Milliseconds()
		log.TimeSinceLast = float64(NaTime) / 1000
	} else {
		log.Code = 01
	}

	log.ClientName = clineName
	db := db.Create(&log)
	if db.Error != nil {
		fmt.Println(db.Error)
		return errors.New("插入数据库失败")
	}

	return
}

func connect(dsn string) (err error) {

	fmt.Printf("进行数据库链接，链接地址[%s]\n", dsn)
	//dsn := "root:123456@tcp(81.70.17.60:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("链接数据库失败:[%s]\n", err)
		return
	}
	fmt.Printf("链接数据库成功~\n")
	return
}

func StartDB(dsn string) (err error) {

	if err = connect(dsn); err != nil {
		fmt.Println(err)
		return
	}
	db.AutoMigrate(&Logs{})
	return
}

func ListClientName() (clientNames []string, err error) {

	// 执行自定义的 SELECT 查询
	var names []string
	query := "SELECT client_name FROM logs.logs group by client_name"
	result := db.Raw(query).Scan(&names)

	if result.Error != nil {
		panic("查询失败")
	}
	clientNames = names
	return
}

func SelectLogsByClientName(name string) (logs []Logs, err error) {

	// 执行自定义的 SELECT 查询
	query := "SELECT * FROM logs.logs where client_name = ? "
	result := db.Raw(query, name).Scan(&logs)

	if result.Error != nil {
		panic("查询失败")
	}
	return
}

func ShowTimeLineLogsByClientName(name string) (times []float64, err error) {

	// 执行自定义的 SELECT 查询
	query := "SELECT time_since_last FROM logs.logs where client_name = ? "
	result := db.Raw(query, name).Scan(&times)

	if result.Error != nil {
		panic("查询失败")
	}
	return
}

func GetCoroutineNumByClientName(name string) (names []string, err error) {

	// 执行自定义的 SELECT 查询
	name = name + "%"
	query := "SELECT client_name FROM logs.logs WHERE client_name LIKE ? GROUP BY client_name"
	result := db.Raw(query, name).Scan(&names)

	if result.Error != nil {
		panic("查询失败")
	}
	return
}
