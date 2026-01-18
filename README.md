# podgo

Pronouns-Over-DNS implementation for Go

Spec-compliant parsing for https://github.com/CutieZone/pronouns-over-dns. Primary pronouns are prioritised, per spec, and appear as the first element(s) of the `Accept` field.

## Usage

For the CLI:

```bash
go install github.com/vcokltfre/podgo/cmd/podgo@latest
podgo domain.tld
```

For the library:

```go
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
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
