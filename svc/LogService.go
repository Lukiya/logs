package svc

import (
	"context"
	"fmt"
	"time"

	"github.com/Lukiya/logs/core"
	"github.com/Lukiya/logs/model"
	"github.com/syncfuture/go/serr"
	"github.com/syncfuture/go/u"
	"github.com/syncfuture/host"
)

type LogService struct{}

func (self *LogService) WriteLogEntry(_ context.Context, in *model.WriteLogCommand) (*model.LogEntryResult, error) {
	r := new(model.LogEntryResult)

	if _asyncWriting {
		go func() {
			err := write(in)
			u.LogError(err)
		}()
	} else {
		err := write(in)
		if u.LogError(err) {
			r.Message = err.Error()
		}
	}

	return r, nil
}

func write(in *model.WriteLogCommand) error {
	client, err := _clientDAL.GetClient(in.ClientID)
	if err != nil {
		return err
	} else if client == nil {
		return serr.Errorf("Client '%s' not found", in.ClientID)
	} else {
		if in.LogEntry.Level < client.Level {
			// skip
			return nil
		}

		// determine database and table
		createdOnUtc := time.UnixMilli(in.LogEntry.CreatedOnUtc)

		var dbName, tableName string
		switch client.DBPolicy {
		case 1: // By Year
			dbName = fmt.Sprintf("%s%s_%04d", core.LOG_DB_PREFIX, client.ID, createdOnUtc.Year())
			// Use month as table name
			tableName = fmt.Sprintf("%02d", createdOnUtc.Month())
			break
		case 2: // By Month
			dbName = fmt.Sprintf("%s%s_%04d%02d", core.LOG_DB_PREFIX, client.ID, createdOnUtc.Year(), createdOnUtc.Month())
			// Use day as table name
			tableName = fmt.Sprintf("%02d", createdOnUtc.Day())
			break
		case 3: // By Day
			dbName = fmt.Sprintf("%s%s_%04d%02d%02d", core.LOG_DB_PREFIX, client.ID, createdOnUtc.Year(), createdOnUtc.Month(), createdOnUtc.Day())
			// Use hour as table name
			tableName = fmt.Sprintf("%02d", createdOnUtc.Hour())
			break
		default:
			dbName = fmt.Sprintf("%s%s", core.LOG_DB_PREFIX, client.ID)
			// Use year as table name
			tableName = fmt.Sprintf("%02d", createdOnUtc.Year())
		}
		// generate id
		in.LogEntry.ID = host.GenerateID()

		err = _logDAL.InsertLogEntry(dbName, tableName, in.LogEntry)
		return err
	}
}

func (self *LogService) GetLogEntry(_ context.Context, query *model.LogEntryQuery) (*model.LogEntryResult, error) {
	r := new(model.LogEntryResult)

	_logDAL.GetLogEntry(query)

	return r, nil
}
func (self *LogService) GetLogEntries(_ context.Context, query *model.LogEntriesQuery) (*model.LogEntriesResult, error) {
	r := new(model.LogEntriesResult)
	var err error
	r.LogEntries, r.TotalCount, err = _logDAL.GetLogEntries(query)
	if u.LogError(err) {
		r.Message = err.Error()
	}

	return r, nil
}
