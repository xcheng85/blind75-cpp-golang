package main

// stable layer
// function as first-class citizen
// the async will return string assumably

type Resolvor func(string)
type Rejector func(error)
type Executor func() (string, error)

type Promisify struct {
	resolver Resolvor
	rejector Rejector
}

// function as first-class citizen
// chain pattern, builder pattern, 
func (p *Promisify) RegisterResolver(r Resolvor) *Promisify{
	p.resolver = r
	return p
}

func (p *Promisify) RegisterRejector(r Rejector) *Promisify{
	p.rejector = r
	return p
}

// .resolve.reject node.js

func (p *Promisify) Execute(ex Executor) {
	// non-block, anonymous
	go func(p *Promisify) {
		res, err := ex()
		if err != nil {
			p.rejector(err)
		} else {
			p.resolver(res)
		}
	}(p)
}