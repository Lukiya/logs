package dal

import (
	"github.com/Lukiya/logs/core"
	"github.com/Lukiya/logs/dal/ch"
	"github.com/Lukiya/logs/dal/mongodb"
	"github.com/Lukiya/logs/dal/mysql"
	"github.com/Lukiya/logs/dal/redis"
	"github.com/Lukiya/logs/model"
	"github.com/syncfuture/go/slog"
)

type ILogDAL interface {
	InsertLogEntry(dbName, tableName string, logEntry *model.LogEntry) error
	GetLogEntry(query *model.LogEntryQuery) (*model.LogEntry, error)
	GetLogEntries(query *model.LogEntriesQuery) ([]*model.LogEntry, int64, error)
	GetDatabases(clientID string) ([]string, error)
	GetTables(database string) ([]string, error)
}

type IClientDAL interface {
	InsertClient(*model.LogClient) error
	GetClient(id string) (*model.LogClient, error)
	UpdateClient(*model.LogClient) error
	DeleteClient(id string) error
	GetClients(in *model.LogClientsQuery) ([]*model.LogClient, error)
}

func NewLogDAL() ILogDAL {
	provider := core.GrpcCP.GetStringDefault("DataAccess.Provider", "clickhouse")
	if provider == "clickhouse" {
		ch.Init()
		r := new(ch.ClickHouseDAL)
		return r
	} else if provider == "mongodb" {
		mongodb.Init()
		r := new(mongodb.MongoDAL)
		return r
	} else if provider == "mysql" {
		mysql.Init()
		r := new(mysql.MySqlDAL)
		return r
	}

	slog.Fatalf("Provider '%s' is not supported", provider)
	return nil
}

func NewClientDAL() IClientDAL {
	r := new(redis.RedisDAL)
	r.Init(core.GrpcCP)
	return r
}
