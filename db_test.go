package gorm_manager

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAll(t *testing.T) {
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
	sql := `insert test(nickname) values(?)`
	db := conn.Exec(sql, fmt.Sprintf("test1.field.value.%s", time.Now().Format("2006-01-02 15:04:05")))
	assert.Equal(t, int64(1), db.RowsAffected, "test1.insertData.error", db.Error)


	conn, err = m.Get("test2").GetGormDB()
	errIsNil = true
	if err != nil {
		errIsNil = false
	}
	assert.Equal(t, true, errIsNil, err)
	sql = `insert test(nickname) values(?)`
	db = conn.Exec(sql, fmt.Sprintf("test2.field.value.%s", time.Now().Format("2006-01-02 15:04:05")))
	assert.Equal(t, int64(1), db.RowsAffected, "test2.insertData.error", db.Error)
}