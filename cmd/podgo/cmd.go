package main

import (
	"os"

	"github.com/vcokltfre/podgo"
)

func main() {
	if len(os.Args) != 2 {
		println("Usage: podgo <domain>")
		return
	}

	pronouns, err := podgo.GetPronouns(os.Args[1], false)
	if err != nil {
		println("Error:", err.Error())
		return
	}

	if pronouns.Any {
		println("Accepts any pronouns")
	}

	for _, p := range pronouns.Accept {
		println(p.Subject, p.Object, p.PossessiveDeterminer, p.PossessivePronoun, p.Reflexive)
	}
}
