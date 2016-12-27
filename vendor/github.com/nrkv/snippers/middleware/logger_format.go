package middleware

// LoggerFormat is a struct that specifies format of printed log unit
type LoggerFormat struct {
	Format string // classic printf-style format
	Value  string // predefined type or hardcoded value
}

const (
	// LoggerTime will be replaced with time of request creation
	LoggerTime = "{time}"
	// LoggerMethod will be replaced with request method
	LoggerMethod = "{method}"
	// LoggerURI will be replaced with request URI
	LoggerURI = "{uri}"
	// LoggerStatus will be replaced with response status
	LoggerStatus = "{status}"
	// LoggerLatency will be replaced with duration between request and response
	LoggerLatency = "{latency}"
)

// DefaultLoggerConfig is a config that setups log to following format:
// TIME RFC3339 | METHOD | STATUS | LATENCY | URI
// Example:
// 2016/12/24 17:05:57 | GET | 200 | 3.903567ms | /api/accounts/42
var DefaultLoggerConfig = []LoggerFormat{
	{"%s", LoggerTime},
	{"%s", " | "},
	{"%-6s", LoggerMethod},
	{"%s", " | "},
	{"%d", LoggerStatus},
	{"%s", " | "},
	{"%15s", LoggerLatency},
	{"%s", " | "},
	{"%s", LoggerURI},
	{"%s", "\n"},
}
