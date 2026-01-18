package main

import "github.com/vcokltfre/podgo"

func main() {
	pronouns, err := podgo.GetPronouns("domain.tld", false)
	if err != nil {
		panic(err)
	}

	if pronouns.Any {
		println("Accepts any pronouns")
	}

	for _, p := range pronouns.Accept {
		println(p.Subject, p.Object, p.PossessiveDeterminer, p.PossessivePronoun, p.Reflexive)
	}
}
