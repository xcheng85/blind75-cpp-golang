package command

import (
	"errors"
	"fmt"
)

// 0: idle
// 1: up
// 2: down

// multi-thread access queue

type IElevatorCommand interface {
	// like ngrx reducer
	Execute(level int, state int) (int, error)
}

type ElevatorUpCommand struct {
	level int
}

func (c *ElevatorUpCommand) Execute(level int, up bool) (int, error) {
	if up && level <= c.level {
		fmt.Printf("Moving up from %d to %d\n", level, c.level)
		return c.level, nil
	} else {
		return level, errors.New("pending up")
	}
}

type ElevatorDownCommand struct {
	level int
}

func (c *ElevatorDownCommand) Execute(level int, up bool) (int, error) {
	if !up && level >= c.level {
		fmt.Printf("Moving down from %d to %d\n", level, c.level)
		return c.level, nil
	} else {
		return level, errors.New("pending down")
	}
}
