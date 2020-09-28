package dbUtil

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

//数据库配置
const (
	userName = "mayantao"
	password = "Mayantao110"
	ip       = "120.79.253.180"
	port     = "3306"
	dbName   = "mayantao"
)

type Book struct {
	Bid          string
	BName        string
	BAuth        string
	Desc         string
	Cat          string
	CatId        int
	Cnt          string
	RankCnt      string
	RankNum      int
	YearMonth    string
	YearMonthDay string
	Date         time.Time
}

//Db数据库连接池
var DB *sql.DB

//注意方法名大写，就是public
func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("connnect success")
}

func InsertBook(book *Book) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return false
	}
	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO go_books (`bid`, `bName`, `bAuth`, `desc`, `cat`, `catId`, `cnt`,`rankCnt`,`rankNum`,`YearMonth`,`YearMonthDay`,`createTime`) " +
		"VALUES (?, ?,?, ?,?, ?,?, ?,?,?, ?,?)")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}

	//将参数传递到sql语句中并且执行
	res, err := stmt.Exec(book.Bid, book.BName, book.BAuth, book.Desc, book.Cat, book.CatId, book.Cnt, book.RankCnt, book.RankNum, book.YearMonth, book.YearMonthDay, book.Date)
	if err != nil {
		fmt.Println("Exec fail", err)
		return false
	}
	//将事务提交
	_ = tx.Commit()
	//获得上一个插入自增的id
	fmt.Println(res.LastInsertId())
	return true
}
