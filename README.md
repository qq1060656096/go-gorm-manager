# go-gorm-driver

```go
import "github.com/qq1060656096/go-gorm-manager"

m := NewConnectionManager()
dataSourceName := "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
dataSourceName = fmt.Sprintf(
	"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	"root3",
	"root3",
	"199.199.199.199",
	"3306",
	"test1",
)
m.Add("test1", &ConnectionConfig{
	DatabaseDriverName: DRIVER_MY_SQL,
	DataSourceName: dataSourceName,
})

conn, err := m.Get("test1").GetGormDB()
if err == nil {
	sql := `insert test(nickname) values(?)`
    db := conn.Exec(sql, fmt.Sprintf("test1.field.value.%s", time.Now().Format("2006-01-02 15:04:05")))
}

```

```go
# go test sql
CREATE TABLE `test` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nickname` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;
```