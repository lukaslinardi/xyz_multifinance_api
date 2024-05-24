package general

import "time"

// message db connection.
const (
	ConnectDBSuccess    string = "Connected to DB"

	ConnectDBFail    string = "Could not connect database, error"

	ClosingDBSuccess string = "Database conn gracefully close"
	ClosingDBFailed  string = "Error closing DB connection"

	Success string = "success"
	Fail    string = "fail"

	DataNotFound string = "no data found"

	DBTimeLayout       string = "2006-01-02 15:04:05"
	ResponseTimeLayout string = "2006-01-02T15:04:05-0700"
)

//URL type
const (
	URLPublic  = "public"  //without expired time
	URLLimited = "limited" //with expired time
)

const (
	LogRotationTime = time.Duration(24) * time.Hour
	MaxRotationFile = 4
)

const (
	SessionContextKey = "session"
)
