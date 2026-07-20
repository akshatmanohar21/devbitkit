# Domain Map

> **Purpose.** This document defines how Devbitkit organizes its capability
> space.
>
> Capabilities are grouped into logical domains based on the developer's intent
> rather than the underlying implementation or technology.
>
> Every capability belongs to exactly one domain. A capability may expose
> multiple operations, but ownership of the capability remains exclusive to its
> domain.
>
> A domain's presence in this document does not, by itself, justify its
> existence. Every domain listed here has independently passed the Inclusion
> Test defined in `DESIGN_NARRATIVE.md` (Section 2). This document assumes
> that test and does not restate it — see the narrative document for the
> reasoning behind why a domain is or isn't in scope.
>
> The Domain Map provides architectural organization only. Behaviour,
> operations, command syntax, and implementation details are documented
> elsewhere.

---

# Domain Hierarchy

Devbitkit organizes its capability space using three hierarchical concepts.

```text
Domain
    ↓
Capability
    ↓
Operation
```

A **Domain** represents a broad category of developer intent that has passed
the Inclusion Test.

A **Capability** represents a specific developer workflow within that domain.

An **Operation** represents an action that can be performed on a capability.

For example:

```text
Encryption
    ↓
Envelope Encryption
    ↓
wrap
unwrap
rotate
```

---

# Domain Principles

Every domain should satisfy the following principles.

- A domain enters this document only after passing the Inclusion Test in
  `DESIGN_NARRATIVE.md` — either secret-adjacent absence of any adequate
  offline tool, or genuine cross-platform behavioural divergence in an
  existing tool. Neither "it's easy to build" nor "it fits alongside existing
  domains" is sufficient on its own.
- Domains represent developer intent rather than implementation details.
- A capability belongs to exactly one domain.
- Domains should remain cohesive and narrowly focused.
- New domains should only be introduced when existing domains no longer
  describe the capability naturally, **and** the new domain independently
  passes the Inclusion Test.
- Domain names should remain stable over time.

---

# Current Domains

The following domains constitute the current capability space of Devbitkit.
Each entry notes which clause of the Inclusion Test it satisfies.

---

## Generators

### Inclusion basis

Test A (secret-adjacent absence) for password and API-key generation — no
offline baseline tool for these exists on either major OS, and generating
credentials via a browser tool means trusting an unverifiable third party with
material that may end up in real use. Test B (fragmentation) partially, for
UUID generation, where GNU and BSD `uuidgen` implementations diverge in
default case, supported flags, and available UUID versions.

### Capabilities

- Password
- API Key
- UUID

### Non-goals

This domain does not perform cryptographic key generation for use in
encryption workflows (key pairs, data-encryption keys). Those belong to the
Encryption domain, where generation is tied to a specific cryptographic
scheme rather than a standalone random value.

---

## Encryption

### Inclusion basis

Test A. No coherent, single-command tool exists for envelope encryption
(data-key generation, wrapping under a key-encryption key, unwrapping,
rotation), and this is the domain with the highest correctness and security
stakes in the project. See `DESIGN_NARRATIVE.md` Sections 4 and 9 — this
domain requires a written, reviewed ADR before implementation.

### Capabilities

- Envelope Encryption (DEK/KEK wrap, unwrap, rotate)

### Non-goals

This domain does not expose general-purpose, algorithm-agnostic
encrypt/decrypt operations with user-selectable cipher modes. It implements
one opinionated, safe-by-default envelope-encryption workflow rather than a
broad cryptography toolkit. It does not own JWT signing/verification unless a
future ADR explicitly brings JWT into scope under the Inclusion Test — JWT is
not currently in scope (see Rejected Domains below).

---

## Hashing

### Inclusion basis

Test B. `sha256sum` (Linux, GNU coreutils) and `shasum`/`md5` (macOS/BSD) are
different binaries with different names, different flags, and inconsistent
algorithm coverage under one command.

### Capabilities

- Hash (algorithm-selectable digest generation)

### Non-goals

This domain does not perform cryptographic key derivation or password
hashing with a security-sensitive purpose (e.g. bcrypt/argon2 for credential
storage) unless a future ADR brings that in as a distinct capability. As
initially scoped, this covers general-purpose digest generation only.

---

## Encoding

### Inclusion basis

Test B for Base64 specifically — GNU and BSD `base64` diverge on the decode
flag (`-d` vs `-D`) and on default line-wrap behaviour. Hex and URL encoding
have no coherent, consistently-available baseline tool on either major OS at
all, which is the same absence-of-tool reasoning that underlies Test A,
applied here to non-secret data.

### Capabilities

- Base64
- URL Encoding
- Hex

### Non-goals

This domain does not perform formatting, validation, or cryptographic
operations.

---

## Time

### Inclusion basis

Test B. GNU and BSD `date` diverge in flag syntax for parsing dates and
converting between timezones, to the point that a working command on one
platform produces a parse error on the other.

### Capabilities

- Unix Timestamp
- Timezone
- RFC3339
- ISO8601
- Duration

### Non-goals

This domain does not perform scheduling or task execution.

---

## Find/Replace

### Inclusion basis

Test B. BSD `sed -i` requires a backup-suffix argument that GNU `sed -i`
treats as optional; omitting it on BSD does not error cleanly, it silently
misinterprets the next argument as the suffix. This is a sharper and more
dangerous divergence than a simple flag mismatch.

### Capabilities

- Find/Replace (literal or simple-pattern substitution in a file, with
  optional in-place editing)

### Non-goals

This domain does not implement a general-purpose regular-expression or
stream-editing language. It is deliberately not a `sed` clone — see
`DESIGN_NARRATIVE.md` Section 4.

---

## Archives

### Inclusion basis

Test B. GNU and BSD `tar` diverge on extended-attribute handling, archive
format defaults, and supported flags (e.g. `--delete`, `--sort` are
GNU-only), to the point that an archive created on one platform can produce
warnings or fail to extract cleanly on the other.

### Capabilities

- Create
- Extract
- List

### Non-goals

This domain does not attempt flag-for-flag parity with either GNU or BSD
tar's complete feature set. It covers common create/extract/list operations
only. This is the second-highest implementation-risk domain in the project
after Encryption, owing to archive-format and compression correctness — see
`DESIGN_NARRATIVE.md` Section 4.

---

# Rejected Domains

The following were evaluated against the Inclusion Test and did not pass.
They are recorded here, not to be silently reconsidered later without a new
ADR explicitly re-applying the test:

- **JSON formatting/querying** — `jq` is already consistent, cross-platform,
  and mature. No fragmentation, no absence.
- **HTTP requests** — `curl` is already consistent, cross-platform, and
  mature. Same reasoning.
- **JWT decode/verify/mint** — dedicated cross-platform tools (`jwt-cli`,
  `compiledpanda/jwt`, and others) already solve this consistently. No
  fragmentation, no absence.
- **Case conversion, colour-format conversion, QR code generation,
  line-ending conversion** — no fragmentation, no secret-adjacency; the only
  argument for inclusion was implementation cost being low, which the
  Inclusion Test explicitly does not accept on its own.

A domain listed here may be reconsidered only if new evidence changes which
clause of the Inclusion Test applies — for example, a genuine cross-platform
divergence in a tool listed above that wasn't previously known. That
evidence, and the resulting decision, belongs in a new ADR, not a silent edit
to this list.

---

# Evolution

The Domain Map is expected to evolve incrementally, but only in the direction
the Inclusion Test permits. A new domain is added here only after:

1. It has been evaluated in writing against `DESIGN_NARRATIVE.md` Section 2,
   and
2. That evaluation is recorded as an ADR.

This document deliberately does not maintain a "future domains" backlog.
Pre-listing domains as inevitable additions was tried in an earlier version
of this project's planning and is the mechanism that let scope drift in the
first place — a listed "future domain" tends to get treated as already
decided by the time anyone gets to it, which defeats the purpose of the
Inclusion Test. Any new domain, whenever it's proposed, starts from zero and
is evaluated on its own merits at that time.