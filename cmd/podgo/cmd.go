package main

import (
	"context"
	"os"

	"github.com/vcokltfre/podgo"
)

func main() {
	if len(os.Args) != 2 {
		println("Usage: podgo <domain>")
		return
	}

	pronouns, err := podgo.GetPronounsResolved(context.Background(), os.Args[1], false)
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

	if pref := pronouns.Preferred(); pref != nil {
		println("Preferred pronouns are:", pref.Subject, pref.Object, pref.PossessiveDeterminer, pref.PossessivePronoun, pref.Reflexive)
	}
}
