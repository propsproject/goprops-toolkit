package rediscache

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
)

type CommandBuilder struct {
	Commands    []Command
}

type Command struct {
	CMD  string
	ARGS []interface{}
}

func (s *CommandBuilder) newCMD(cmd string, args []interface{}) Command {
	return Command{CMD: cmd, ARGS: args}
}

func (s *CommandBuilder) CMD(cmd string, args []interface{}) *CommandBuilder {
	s.Commands = append(s.Commands, s.newCMD(cmd, args))
	return s
}

func (s *CommandBuilder) DO(conn redis.Conn) ([]interface{}, []error) {
	defer conn.Close()

	var (
		results []interface{}
		errs []error
	)
	for _, cmd := range s.Commands {
		err := conn.Send(cmd.CMD, cmd.ARGS...)
		if err != nil {
			errs = append(errs, fmt.Errorf("error sending pipeline command to redis (CMD: %s) (ARGS: %v) (%v)", cmd.CMD, cmd.ARGS, err))
		}

		return nil, errs
	}

	err := conn.Flush()
	if err != nil {
		errs = append(errs, fmt.Errorf("error flushing pipeline closing connection (%v)", err))
		return nil, errs
	}

	for range s.Commands {
		v, err := conn.Receive()
		if err != nil {
			errs = append(errs, fmt.Errorf("error receiving results from pipeline command to redis (%v)", err))
		}
		results = append(results, v)
	}
	return results, errs
}

func Commander() *CommandBuilder {
	return &CommandBuilder{make([]Command, 0)}
}

