package rediscache

import (
	"fmt"
	"github.com/mediocregopher/radix.v2/redis"
	"encoding/base64"
	"time"
	"strings"
)

// command is a non exported struct that putcmd and getcmd derive from respectively
type command struct {
	cmd string
	data string
	args []string
	instance *redis.Client
}

func (c *command) LPop(list string, value *interface{}, d interface{}) error {
	c.cmd = fmt.Sprintf("%s %s", LPOP, list)
	resp, err := c.toExecution().Execute()
	if err != nil {
		return err
	}

	return c.decode(resp, value, d)
}

func (c *command) Get(key string, value *interface{}, d interface{}) error {
	c.cmd = fmt.Sprintf("%s %s", GET, key)
	resp, err := c.toExecution().Execute()
	if err != nil {
		return err
	}

	return c.decode(resp, value, d)
}

//RPOPLPUSH used to maintain a FIFO job queue
func (c *command) RPOPLPUSH(src, dst string, value *interface{}, d interface{}) error {
	c.cmd = fmt.Sprintf("%s %s %s", RPOPLPUSH, src, dst)
	resp, err := c.toExecution().Execute()
	if err != nil {
		return err
	}

	return c.decode(resp, value, d)
}

//LLEN length of list
func (c *command) LLEN(key string) (int, error) {
	c.cmd = fmt.Sprintf("%s %s", LLEN, key)
	resp, err := c.toExecution().Execute()
	if err != nil {
		return 0, err
	}

	return resp.Int()
}

func (c *command) decode(resp *redis.Resp, value *interface{}, d interface{}) error {
	b, err := resp.Bytes()
	if err != nil {
		return err
	}

	err = NewDecoder(d).Decode(b, value)
	if err != nil {
		return err
	}

	return nil
}

func (c *command) toExecution() *execution {
	return &execution{instance: c.instance, cmd: c.cmd}
}

// WithData will create a new encoder from the data argument you provide that satisfies the Encoder interface.
// It will return an initialized putcmd struct with commands only for adding data to redis
func (c *command) WithData(data interface{}) (*putcmd, error) {
	switch v := data.(type) {
	case Encoder:
		b, err := NewEncoder(data).Encode()
		if err != nil {
			return nil, err
		}
		c.data = base64.StdEncoding.EncodeToString(b)
		return c.toPutCMD(), nil
	case string:
		c.data = base64.StdEncoding.EncodeToString([]byte(data.(string)))
		return c.toPutCMD(), nil
	case []byte:
		c.data = base64.StdEncoding.EncodeToString(data.([]byte))
		return c.toPutCMD(), nil
	default:
		return nil, fmt.Errorf("unsupported type %v, please use Encoder, string, or bytes", v)
	}
}

// transforms a command to a putcmd essentially stripping away any get functions in an attempt to protect user from invalid cmds
func (c *command) toPutCMD() *putcmd {
	return &putcmd{data: c.data, instance: c.instance, cmd: c.cmd}
}

// struct with one method receiver Execute, this essentially ends any cmd manipulation after certain sequences
type execution command

// Execute generic function executes redis cmd and returns the resp or an error
func (e *execution) Execute() (*redis.Resp, error) {
	if resp := e.instance.Cmd(e.cmd, e.args); resp.Err != nil {
		return nil, fmt.Errorf("could not execute redis command %s: %v", e.cmd, resp.Err)
	} else {
		return resp, nil
	}
}

// struct that provides method receivers for getting data out of redis
type putcmd command

type CMDOptions struct {
	args []string
	builder strings.Builder
}

func (o *CMDOptions) Build() []string {
	return o.args
}

func NewCMDOptions() *CMDOptions {
	return &CMDOptions{}
}

func (o *CMDOptions) SetExpire(duration time.Duration, EX bool) *CMDOptions {
	if inSeconds := EX; inSeconds {
		o.args = append(o.args, ExpireInSeconds, fmt.Sprintf("%.0f", duration.Seconds()))
		return o
	}

	millis := duration * time.Millisecond
	o.args = append(o.args, ExpireInMilliseconds, fmt.Sprintf("%.0f", millis.Seconds()))
	return o
}

func (o *CMDOptions) SetNX() *CMDOptions {
	o.args = append(o.args, OnlyIfKeyNotExist)
	return o
}

func (o *CMDOptions) SetXX() *CMDOptions {
	o.args = append(o.args, OnlyIfKeyExists)
	return o
}

// Set builds the redis set command string. Can be followed up with .Expire(...) which will append the EX nd expiration in seconds to the command and return an execution
func (pc *putcmd) Set(key string, opts *CMDOptions) *execution {
	pc.cmd = SET
	pc.args = append([]string{key, pc.data}, opts.Build()...)
	return pc.toExecution()
}

func (pc *putcmd) RPush(list string) *execution {
	pc.cmd = RPUSH
	pc.args = []string{list, pc.data}
	return pc.toExecution()
}

func (pc *putcmd) LPush(list string) *execution {
	pc.cmd = LPUSH
	pc.args = []string{list, pc.data}
	return pc.toExecution()
}

func (pc *putcmd) SADD(set string) *execution {
	pc.cmd = SADD
	pc.args = []string{set, pc.data}
	return pc.toExecution()
}

func (pc *putcmd) toExecution() *execution {
	return &execution{instance: pc.instance, cmd: pc.cmd, args: pc.args}
}

