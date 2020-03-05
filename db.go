package gormManager

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	DRIVER_MY_SQL = "mysql"
	DRIVER_POSTGRE_SQL = "postgres"
	DRIVER_SQLITE3 = "sqlite3"
	DRIVER_SQL_SERVER = "mssql"
)

type ConnectionManager struct {
	connList map[string]*Connection
}

type Connection struct {
	db   *gorm.DB
	conf *ConnectionConfig
}

type ConnectionConfig struct {
	DatabaseDriverName   string
	DataSourceName string
}

func NewConnectionManager() *ConnectionManager {
	m := ConnectionManager{
		connList: make(map[string]*Connection),
	}
	return &m
}

func (m *ConnectionManager) Add(name string, conf *ConnectionConfig) {
	m.connList[name] = &Connection{
		conf: conf,
	}
}

func (m *ConnectionManager) Remove(name string) {
	delete(m.connList, name)
}

func (m *ConnectionManager) Get(name string) *Connection {
	con, ok := m.connList[name]
	fmt.Println(name, con, ok, m.connList)
	if !ok {
		return nil
	}
	fmt.Println(name, con, ok, m.connList)

	return con
}

func (m *ConnectionManager) Exist(name string) bool {
	con := m.Get(name)
	if con == nil {
		return false
	}
	return true
}

func (m *ConnectionManager) Length() int {
	return len(m.connList)
}

func (m ConnectionManager) String() string {
	type tmpT struct {
		HasDB bool
		ConnConfig ConnectionConfig
	}
	list := make(map[string]tmpT)
	for k, v := range m.connList {
		tmp := tmpT{}
		if v.db != nil {
			tmp.HasDB = true
		}
		tmp.ConnConfig = *v.conf
		list[k] = tmp
	}
	bytes, _ := json.Marshal(list)
	return string(bytes)
}



func (c *Connection) GetGormDB() (*gorm.DB, error) {
	if c.db == nil {
		err := c.ReconnectGormDB()
		return c.db, err
	}
	return c.db, nil
}

func (c *Connection) ReconnectGormDB() error {
	db, err := gorm.Open(c.conf.DatabaseDriverName, c.conf.DataSourceName)
	if err != nil {
		return err
	}
	c.db = db
	return nil
}

func (c *Connection) DisconnectGormDB() bool {
	if c.db != nil {
		c.db.Close()
		c.db = nil
	}
	return true
}
