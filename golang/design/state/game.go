package state

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// unstable, tied to the app
type GameContext struct {
	SecretNumber int
	RetriesLeft  int
	Won          bool
	Next         IGameState
}

type IGameState interface {
	Execute(*GameContext) bool
}

// real state machine
type StartState struct {
}

type PromptState struct {
}

type FinishState struct {
}

func (s *StartState) Execute(ctx *GameContext) bool {
	// what is the mutation
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	// Generate a random number between 0 and 10
	ctx.SecretNumber = rng.Intn(10)
	fmt.Println("max retries allowed: ")
	fmt.Fscanf(os.Stdin, "%d\n", &ctx.RetriesLeft)

	// state change
	ctx.Next = &PromptState{}

	return true
}

func (s *PromptState) Execute(ctx *GameContext) bool {
	fmt.Printf("Guess a number between [0, 10], %d tries left\n", ctx.RetriesLeft)
	var guessNum int
	fmt.Fscanf(os.Stdin, "%d\n", &guessNum)
	// no operator-- in go
	ctx.RetriesLeft = ctx.RetriesLeft - 1

	if guessNum == ctx.SecretNumber {
		ctx.Won = true
		ctx.Next = &FinishState{}
	}

	if ctx.RetriesLeft == 0 {
		ctx.Next = &FinishState{}
	}

	// other case, stay in this state, self-loop
	return true
}

func (s *FinishState) Execute(ctx *GameContext) bool {
	if ctx.Won {
		fmt.Println("You Won")
	} else {
		fmt.Println("You Lose")
	}
	// exit state machine
	return false
}
