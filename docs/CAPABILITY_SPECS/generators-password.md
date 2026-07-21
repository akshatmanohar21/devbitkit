# Capability Specification: Password (Generators domain)

> See `docs/DOMAIN_MAP.md` for why this capability's domain is in scope, and
> `docs/ADR/0001-Generators-Domain-Inclusion.md` for the accepted decision.
> This document describes *what the capability does*, not why it exists.

## Command

```
dbk generate password [flags]
```

## Flags

| Flag | Type | Default | Description |
|------|------|---------|--------------|
| `--length` | int | `16` | Password length, in characters. |
| `--no-letters` | bool | `false` | Exclude letter characters (A–Z, a–z) from the charset. |
| `--no-numbers` | bool | `false` | Exclude number characters (0–9) from the charset. |
| `--no-symbols` | bool | `false` | Exclude symbol characters (`!@#$%^&*()-_=+`) from the charset. |
| `--count` | int | `1` | Number of passwords to generate. Each is printed on its own line. |

## Behavior

- Each character is drawn independently from the active charset using
  `crypto/rand`, Go's cryptographically secure random source, avoiding
  modulo bias (see `internal/generators/password.go` for the implementation
  detail — this document tracks behavior, not implementation).
- The active charset is the union of whichever of letters/numbers/symbols
  are not excluded via flags. At least one class must remain included.
- With `--count` greater than 1, each generated password is fully
  independent — no relationship or shared entropy between them.
- Unrecognized positional arguments (anything not matching a defined flag)
  are rejected with an error; they are not silently ignored.

## Error Conditions

| Condition | Behavior |
|-----------|----------|
| `--length` ≤ 0 | Error: `"password length must be positive, got <n>"`. Exit code 1. |
| `--count` ≤ 0 | Error: `"count must be positive, got <n>"`. Exit code 1. |
| All of `--no-letters`, `--no-numbers`, `--no-symbols` set simultaneously | Error: `"no characters available: cannot exclude all character sets"`. Exit code 1. |
| Unrecognized argument passed | Error: `"unexpected argument passed: <args>"`. Exit code 1. |
| Underlying `crypto/rand` call fails | Error wraps the underlying cause via `%w`. Exit code 1. This is expected to be rare (OS-level RNG failure). |

## Examples

```
dbk generate password
dbk generate password --length 24
dbk generate password --no-symbols
dbk generate password --no-letters --no-symbols
dbk generate password --count 5
```

## Status

Implemented. Covered by 7 automated tests in
`internal/generators/password_test.go`: output length, basic randomness
sanity, exclusion correctness for each of letters/numbers/symbols
individually, and both error conditions (invalid length, all classes
excluded).