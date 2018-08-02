package rediscache

import (
	"encoding/base64"
	"fmt"
	"github.com/kc1116/builder"
	"github.com/mediocregopher/radix.v2/redis"
	"time"
)

type CommandBuilder builder.Builder

func (cb CommandBuilder) CMD(cmd string) CommandBuilder {
	return builder.Set(cb, "CMD", cmd).(CommandBuilder)
}

func (cb CommandBuilder) ARGS(args ...string) CommandBuilder {
	return builder.Extend(cb, "ARGS", args).(CommandBuilder)
}

func (cb CommandBuilder) VAL(data ...interface{}) CommandBuilder {
	return builder.Extend(cb, "Data", data).(CommandBuilder)
}

func (cb CommandBuilder) DO(conn *redis.Client) *Resp {
	cmd := builder.GetStruct(cb).(Command)
	return &Resp{conn.Cmd(cmd.Format())}
}

func (cb CommandBuilder) kv(key, value string) interface{} {
	return map[string]string{key: value}
}

// WithData will create a new encoder from the data argument you provide that satisfies the Encoder interface.
// It will return an initialized putcmd struct with commands only for adding data to redis
func (cb CommandBuilder) encodeB64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Commander() CommandBuilder {
	return builder.Register(CommandBuilder{}, Command{}).(CommandBuilder)
}

type Command struct {
	CMD  string
	ARGS []string
	Data []interface{}
}

func (c Command) Format() (string, []string, []interface{}) {
	return c.CMD, c.ARGS, c.Data
}

type Resp struct {
	RedisResp *redis.Resp
}

func (r *Resp) Decode(value *interface{}, d interface{}) error {
	if r.RedisResp.Err != nil {
		return fmt.Errorf("could not execute redis command: %v", r.RedisResp.Err.Error())
	}

	b, err := r.RedisResp.Bytes()
	if err != nil {
		return err
	}

	err = NewDecoder(d).Decode(b, value)
	if err != nil {
		return err
	}

	return nil
}

func DurationInSec(duration time.Duration) string {
	return fmt.Sprintf("%.0f", duration.Seconds())
}

func DurationInMilli(duration time.Duration) string {
	millis := duration * time.Millisecond
	return fmt.Sprintf("%.0f", millis.Seconds())
}
