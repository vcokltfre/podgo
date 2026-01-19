package main

import (
	"context"
	"os"

	"github.com/vcokltfre/podgo"
)

func main() {
	var pronouns *podgo.Pronouns
	var err error

	if os.Getenv("DRY") != "1" {
		pronouns, err = podgo.GetPronounsResolved(context.Background(), os.Args[1], false)
	} else {
		pronouns, err = podgo.GetPronouns([]string{
			"they/them/their/theirs/themself;preferred",
			"it/its",
			"invalid;record",
		}, false)
	}
	if err != nil {
		panic(err)
	}

	if pronouns.None {
		println("Does not accept pronouns, use name instead")
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
