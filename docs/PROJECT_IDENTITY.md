# Project Identity

This document defines the canonical identity of the project. These values are
considered stable and should remain consistent across repositories, releases,
documentation, and implementations.

Where this document would need to describe infrastructure not yet justified
by an actual domain or an ADR — configuration systems, extension mechanisms,
and similar — it deliberately does not, consistent with
`DESIGN_NARRATIVE.md` Section 8: architecture is built to match current
scope, not ahead of it.

---

## Project Name

**Devbitkit**

The official name of the project.

---

## CLI Executable

`dbk`

The primary executable exposed to users.

---

## Canonical Repository

`devbitkit`

The canonical repository name.

---

## Naming Conventions

The following naming conventions are considered canonical throughout the
project.

| Entity | Convention |
|---------|------------|
| Executable | `dbk` |
| Capability | `generate`, `encrypt`, `hash`, `encode`, `time`, `replace`, `archive` |
| Operations | Verb-oriented (`generate`, `wrap`, `hash`, `decode`, `convert`) |
| Environment Variables | `DBK_*` |

Capability names above reflect the current domain list in `DOMAIN_MAP.md` and
should be updated there first if a domain is added, renamed, or removed —
this table follows that document, not the reverse.

---

## Reserved Identifiers

The following prefix is reserved by the project:

- `dbk`

No plugin or extension namespace is reserved at this time. Devbitkit does not
currently have a plugin system or third-party extension mechanism — see
`DESIGN_NARRATIVE.md` Section 8. If one is introduced in the future under its
own ADR, a reserved prefix for it belongs in that ADR and should be added
here only once that architecture actually exists, not in advance of it.

---

## Configuration

Devbitkit does not currently have a persistent user configuration system.
Every domain in scope as of this document (`DOMAIN_MAP.md`) is a stateless,
single-invocation operation with no need for stored state, defaults, or
per-user settings. A configuration mechanism — including its location, format,
and the `DBK_*` environment variables that might complement it — should be
designed only once a specific domain's capability spec identifies an actual
need for it, and recorded as its own ADR at that time.

---

## Naming Philosophy

The project is named **Devbitkit** because it represents a small, deliberately
bounded collection of developer capabilities ("bits") that fill specific gaps
— either the absence of any adequate offline tool for handling sensitive
material, or genuine cross-platform inconsistency in a tool that otherwise
already exists. It is not a general umbrella for unrelated utilities and does
not aim to grow indefinitely; see `DESIGN_NARRATIVE.md` Section 2 for the test
that governs what does and doesn't belong under this name.