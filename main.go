package main

import pokeapi "github.com/Asfand-Yar-Hassan/PokeDex/internal/poke-api"

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.Client{},
	}
	startRepl(&cfg)
}
