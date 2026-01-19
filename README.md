# podgo

Pronouns-Over-DNS implementation for Go

Spec-compliant parsing for https://github.com/CutieZone/pronouns-over-dns. Primary pronouns are prioritised, per spec, and appear as the first element(s) of the `Accept` field.

## Usage

For the CLI:

```bash
go install github.com/vcokltfre/podgo/cmd/podgo@latest
podgo domain.tld
```

For the library, example usage can be found in `cmd/podgo/cmd.go` and `example/main.go`.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
