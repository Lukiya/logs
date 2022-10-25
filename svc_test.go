package main

import (
	"context"
	"testing"
	"time"

	"github.com/Lukiya/logs/model"
	"github.com/Lukiya/logs/svc"
	"github.com/syncfuture/go/u"
)

func TestWriteLog(t *testing.T) {
	logService := new(svc.LogService)
	logService.Write(context.Background(), &model.WriteLogCommand{
		ClientID: "DL",
		LogEntry: &model.LogEntry{
			Level:        model.LogLevel_Debug,
			TraceNo:      "xxxxx",
			Message:      "AAA",
			Error:        "BBB",
			CreatedOnUtc: time.Now().UTC().UnixMilli(),
			Payload:      u.StrToBytes(`{"name":"test","score":3.98}`),
		},
	})

	time.Sleep(1 * time.Second)
}
