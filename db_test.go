package gormManager

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestNewConnectionManagerForMysql(t *testing.T) {
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

	dataSourceName = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root3",
		"root3",
		"199.199.199.199",
		"3306",
		"test2",
	)
	m.Add("test2", &ConnectionConfig{
		DatabaseDriverName: DRIVER_MY_SQL,
		DataSourceName: dataSourceName,
	})

	conn, err := m.Get("test1").GetGormDB()
	errIsNil := true
	if err != nil {
		errIsNil = false
	}
	assert.Equal(t, true, errIsNil, err)
	sql := `insert into test(nickname) values(?)`
	db := conn.Exec(sql, fmt.Sprintf("test1.field.value.%s", time.Now().Format("2006-01-02 15:04:05")))
	assert.Equal(t, int64(1), db.RowsAffected, "test1.insertData.error", db.Error)


	conn, err = m.Get("test2").GetGormDB()
	errIsNil = true
	if err != nil {
		errIsNil = false
	}
	assert.Equal(t, true, errIsNil, err)
	sql = `insert into test(nickname) values(?)`
	db = conn.Exec(sql, fmt.Sprintf("test2.field.value.%s", time.Now().Format("2006-01-02 15:04:05")))
	assert.Equal(t, int64(1), db.RowsAffected, "test2.insertData.error", db.Error)
}

func TestNewConnectionManagerForSqlite(t *testing.T) {
	dir, _ := os.Getwd()
	fmt.Println(dir)

	m := NewConnectionManager()
	dataSourceName := dir + "/tmp/sqlite3.1.db"
	m.Add("sqlite1", &ConnectionConfig{
		DatabaseDriverName: DRIVER_SQLITE3,
		DataSourceName: dataSourceName,
	})

	dataSourceName = dir + "/tmp/sqlite3.2.db"
	m.Add("sqlite2", &ConnectionConfig{
		DatabaseDriverName: DRIVER_SQLITE3,
		DataSourceName: dataSourceName,
	})

	assert.Equal(t, 2, m.Length(), "driver.length.func.error")

	conn, err := m.Get("sqlite1").GetGormDB()
	errIsNil := true
	if err != nil {
		errIsNil = false
	}
	assert.Equal(t, true, errIsNil, err)
	sql := `insert into test(nickname) values(?)`
	db := conn.Exec(sql, fmt.Sprintf("sqlite1.field.value.%s", time.Now().Format("2006-01-02 15:04:05")))
	assert.Equal(t, int64(1), db.RowsAffected, "sqlite1.insertData.error", db.Error)

	conn, err = m.Get("sqlite2").GetGormDB()
	errIsNil = true
	if err != nil {
		errIsNil = false
	}
	assert.Equal(t, true, errIsNil, err)
	sql = `insert into test(nickname) values(?)`
	db = conn.Exec(sql, fmt.Sprintf("sqlite2.field.value.%s", time.Now().Format("2006-01-02 15:04:05")))
	assert.Equal(t, int64(1), db.RowsAffected, "sqlite2.insertData.error", db.Error)

	assert.Less(t, 1, len(m.String()))
}