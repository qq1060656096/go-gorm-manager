package gormManager

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pkg/errors"
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
	if !ok {
		return nil
	}
	return con
}

func (m *ConnectionManager) Exist(name string) bool {
	con := m.Get(name)
	if con == nil {
		return false
	}
	return true
}

func (m *ConnectionManager) length() int {
	return len(m.connList)
}

func (m *ConnectionManager) Reconnection(name string) (*Connection, error) {
	con := m.Get(name)
	if con == nil {
		return nil, errors.New("gorm_manager.connectionNameNotExist")
	}
	db, err := gorm.Open(con.conf.DatabaseDriverName, con.conf.DataSourceName)
	if err != nil {
		return nil, err
	}
	con.db = db
	return con, nil
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
