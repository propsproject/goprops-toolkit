package rediscache

const (
	SET = "SET"
	GET = "GET"
	RPUSH = "RPUSH"
	LPOP = "LPOP"
	LLEN = "LLEN"
	LPUSH = "LPUSH"
	RPOPLPUSH = "RPOPLPUSH"
	PING = "PING"
	LREM = "LREM"
	SADD = "SADD"

	ExpireInSeconds = "EX"
	ExpireInMilliseconds = "PX"
	OnlyIfKeyNotExist = "NX"
	OnlyIfKeyExists = "XX"

	errOnOpenConnection = "could not open redis connection: %v"
)