# Design Narrative

> **Milestone snapshot.** Initial design narrative for Devbitkit.
>
> This document defines the long-lived philosophy, architectural intent, and design reasoning that
> establish the identity of the project. It intentionally operates above implementation details.
>
> This document is **narrative**, not **specification**. Capability behaviour, command contracts,
> implementation details, architectural decisions, and development policies are maintained within
> their respective reference artifacts. Where this document differs from an Architecture Decision
> Record (ADR) or a capability specification, the reference artifact takes precedence.

---

# 1. What this project is

Devbitkit is a command-line utility that provides a small set of developer workflows that have no
adequate existing tool — either because no coherent tool exists at all, or because the tool that
exists behaves inconsistently across operating systems.

Devbitkit is not an attempt to unify or replace mature, already-consistent developer tools. Utilities
like jq, curl, and dedicated single-purpose CLIs solve their problems well, are already cross-platform,
and already carry years of documentation and muscle memory behind them. A tool that only offers a new
syntax for something already solved does not reduce friction — it adds a fourth convention to
remember alongside three that already work.

Devbitkit exists specifically where one of two conditions holds:

1. **No coherent tool exists for the workflow**, and developers currently rely on throwaway scripts,
   ad hoc shell one-liners, or third-party websites — including, in some cases, pasting sensitive
   material such as private keys, credentials, or plaintext secrets into a browser-based tool with no
   way to verify what happens to that data afterward.
2. **A tool exists, but its behaviour genuinely diverges across operating systems** — commonly between
   GNU/Linux and BSD/macOS — such that the same command produces different output, requires different
   flags, or fails silently depending on the platform it runs on.

Every capability considered for Devbitkit is evaluated against the Inclusion Test defined in Section
2 before it is accepted into scope. Ease of implementation is deliberately not part of that test —
see Section 2 for why.

---

# 2. The Inclusion Test

This is the central filter of the project and takes precedence over every other principle in this
document. A capability may be added to Devbitkit only if it satisfies **at least one** of the
following:

## A. Secret-Adjacent Absence

The workflow handles sensitive material — private keys, credentials, tokens, data-encryption keys —
and no adequate offline tool exists for it. The differentiator here is trust boundary, not
convenience: a local, offline binary is categorically safer than a browser-based tool for anything
touching secret material, because there is no way to verify what a third-party website does with
data submitted to it.

*Examples: envelope encryption / data-key wrapping, JWT signing with a private key, generation of
passwords or API keys using a cryptographically secure random source.*

## B. Cross-Platform Fragmentation

A baseline tool exists, but its behaviour genuinely diverges across operating systems — different
flags, different defaults, or silent failure modes — to the point that a script or habit built on one
platform breaks or misbehaves on another.

*Examples: `sed -i` (a mandatory vs. optional backup-suffix argument between BSD and GNU sed, where
omitting it on BSD silently misinterprets the rest of the command), `base64` (different decode flags
and line-wrap behaviour between implementations), hashing utilities (`sha256sum` on Linux vs.
`shasum`/`md5` on macOS being entirely different binaries with different interfaces), `tar`
(incompatible extended-attribute and archive-format handling between BSD and GNU implementations),
`date` (incompatible flag syntax for parsing dates and converting timezones).*

## What does not pass

**"No tool does exactly this, but building one would be cheap and low-risk" is not sufficient on its
own.** Nearly any small utility clears that bar — it describes the default state of a well-scoped
feature, not a reason to build it. A capability that is easy to build, does not touch secret
material, and does not face genuine platform divergence does not belong in Devbitkit, no matter how
naturally it seems to fit alongside the domains that do. The cost of including something is not the
time it takes to write — it is the long-term maintenance surface, the dilution of `--help` output,
and the drift away from a clear, describable product identity. This test exists specifically to keep
that drift from happening one convenient addition at a time.

Any proposal to add a new capability must be evaluated against Test A and Test B explicitly, in
writing, before implementation begins — see Section 9 on ADRs.

---

# 3. The Command Language

The fundamental abstraction within Devbitkit is **the command language**.

Capabilities that pass the Inclusion Test are presented through a shared command structure rather
than independent interfaces. Developers should not need to learn a different mental model for each
capability. Once a developer understands how one capability behaves, every other capability should
feel immediately familiar.

This consistency governs *how* an included capability is exposed. It is not, on its own, a reason to
include a capability — that question is answered entirely by Section 2.

The command language is considered part of the product itself. Changes affecting command structure,
argument conventions, naming, or behavioural expectations are architectural decisions and must be
treated accordingly.

---

# 4. The Capability Model

Every piece of functionality exposed by Devbitkit is modeled as a **Capability**, grouped into
functional **Domains**. The domains currently in scope, and the specific reasoning that qualifies
each one under Section 2, are:

- **Generators** — secure password and API-key generation qualifies under Test A (offline,
  cryptographically sound, no browser round-trip for anything sensitive). UUID generation qualifies
  partially under Test B, given divergence between GNU and BSD `uuidgen` implementations.

- **Encryption (envelope encryption / DEK handling)** — qualifies under Test A. This is the domain
  with the highest correctness and security stakes in the project. It requires a written design
  review — algorithm choice, mode, nonce strategy, key-wrap and rotation scheme — before any
  implementation begins, not after. See Section 9.

- **Hashing** — qualifies under Test B (`sha256sum` vs. `shasum`/`md5` divergence between Linux and
  macOS).

- **Encoding** (base64, hex, URL-encoding) — base64 qualifies under Test B directly; hex and
  URL-encoding qualify under the absence clause of Test A's reasoning extended to non-secret data —
  no coherent, consistently-available baseline tool exists for either on a fresh install of either
  major OS.

- **Time** — qualifies under Test B (`date` flag and timezone-handling divergence between GNU and
  BSD).

- **Find/Replace** — a deliberately narrow capability (literal or simple pattern find-and-replace in
  a file, not a general-purpose stream-editing language), qualifying under Test B via `sed -i`'s
  divergent and occasionally silently-dangerous behaviour across platforms.

- **Archives** — qualifies under Test B via `tar`'s incompatible flag and format handling between GNU
  and BSD. Second only to Encryption in implementation risk, owing to the correctness demands of
  archive-format and compression handling; scoped to common create/extract/list operations rather
  than full parity with either GNU or BSD tar's complete flag set.

A capability owns only one responsibility. A capability may expose multiple operations where
appropriate, but those operations must remain cohesive under the same conceptual problem space, and
must not evolve into a collection of loosely related commands. Where a domain's real-world
counterpart implements a large and complex feature set of its own (`sed`'s regex language, `openssl`'s
full cipher-suite surface), Devbitkit implements only the common-path workflows within that domain —
not a clone of the underlying tool's full capability. This constraint applies most strictly to
Find/Replace (not a general regex engine) and Encryption (a scoped, opinionated envelope-encryption
workflow, not a general-purpose cryptography toolkit with exposed algorithm and mode selection).

---

# 5. The Interface Model

Devbitkit separates **capability** from **interface**.

Capabilities define behaviour. Interfaces expose behaviour.

The command-line interface is the sole interface of the project. No graphical interface is planned or
designed for at this stage. If one is proposed in the future, it will be evaluated on its own merits
at that time — including whether it is justified at all — rather than assumed as a natural extension
of the CLI.

---

# 6. Design Philosophy

Several principles govern architectural and implementation decisions within Devbitkit, listed in
priority order. Where principles conflict, the higher-listed principle takes precedence.

## Justified Inclusion (highest priority)

Nothing is added to this project without passing the Inclusion Test in Section 2. This principle
supersedes all others below, including consistency and simplicity, because a consistent and simple
interface to the wrong set of capabilities is still the wrong product.

## CLI First

The command-line interface is the only interface of the project (see Section 5).

## Offline First

Every capability must function without requiring network connectivity. This is a load-bearing
property, not a soft default: it is part of what makes the trust-boundary argument in Test A hold. A
tool that phones home undermines the exact guarantee that justifies its existence for secret-adjacent
workflows.

## Consistency Over Convenience

Among capabilities that have already passed the Inclusion Test, a coherent, predictable command
structure is preferred over shorter or cleverer syntax for any individual command.

## Correctness Before Breadth, Especially for Encryption

For the Encryption domain specifically, a narrow set of correct, safe-by-default operations is
strongly preferred over broad algorithm or mode flexibility. Where exposing a choice to the user (for
example, cipher mode) would introduce a failure mode that fails silently rather than loudly, that
choice should not be exposed by default.

## Scriptability

Capabilities should integrate naturally with shell scripting, pipelines, automation systems, and
continuous integration environments.

## Predictability

Equivalent inputs should produce equivalent outputs. Capabilities should avoid hidden state, implicit
behaviour, or surprising side effects unless explicitly requested by the user.

---

# 7. Boundaries

Devbitkit intentionally maintains a narrow scope. It does not attempt to become:

- A shell
- A terminal emulator
- A package manager
- A programming language
- An integrated development environment
- A cloud platform
- An AI assistant
- A hosted service
- An authentication platform
- A general-purpose replacement for any tool that is already consistent and already cross-platform

Devbitkit's purpose is to fill specific, evidenced gaps — not to consolidate ownership of the
developer tooling ecosystem. The number of domains in the project is bounded entirely by how many
capabilities actually pass the Inclusion Test. A proposed domain that cannot be justified under
Section 2, in writing, does not enter scope, regardless of how naturally it seems to fit alongside
domains that already have.

---

# 8. Architectural Direction

Devbitkit is a single static binary with domains implemented as subcommands. The architecture is
intentionally flat: one command-line entrypoint, one internal package per domain, and shared internal
utilities only where genuinely common across domains (I/O handling, error formatting, and similar).
No plugin system or capability-engine abstraction is being built ahead of need. If the project's scope
grows enough in the future to justify that complexity, it will be introduced then, under its own ADR
— not speculatively now.

Implementation favours Go's standard library first. Third-party packages are used only where the
standard library is genuinely inadequate for a common-path workflow, and any such dependency is
treated as a supply-chain decision worth recording, not a default choice.

---

# 9. Decision Philosophy

Architectural decisions within Devbitkit are intentionally conservative. New capabilities,
interfaces, or architectural patterns are accepted only when they pass the Inclusion Test (Section 2)
and strengthen the project's actual, current scope rather than expand it speculatively.

Every significant decision — including every new domain, and certainly any design affecting the
Encryption domain — is documented through an Architecture Decision Record (ADR) before implementation
begins, not after. For the Encryption domain specifically, the ADR is reviewed before any code is
written, given the correctness and security stakes involved.

The project values deliberate evolution over rapid expansion.

---

# 10. Status and Roadmap at this Milestone

At this milestone, Devbitkit's foundational identity is defined: a small set of command-line
workflows, gated by the Inclusion Test, prioritizing offline execution and a consistent command
language among the capabilities that qualify.

**Planned build order**, sequenced for learning curve as much as domain priority:

1. **Go fundamentals** — language basics, module structure, the `testing` package, before any domain
   work begins.
2. **Generators** — first domain implementation. Chosen as a starting point for its low correctness
   risk (secure randomness via `crypto/rand` is a solved, narrow problem) relative to its real
   differentiator under Test A.
3. **One fragmentation domain** (Hashing or Encoding) — reinforces the CLI and subcommand pattern
   established in Generators without introducing new security surface.
4. **Envelope encryption, conceptual study** — no code at this stage; understanding data-key and
   key-encryption-key relationships, key wrapping, rotation, and why authenticated modes are
   preferred over unauthenticated ones.
5. **Encryption domain design review, then implementation** — a written ADR covering algorithm, mode,
   nonce strategy, and key-wrap/rotation scheme, reviewed before implementation begins.

The project's governing philosophy is:

- Include only what passes the Inclusion Test — secret-adjacent absence, or genuine cross-platform
  fragmentation.
- Treat offline execution as load-bearing, not a convenience.
- Keep the architecture flat and matched to actual current scope.
- Treat the Encryption domain's correctness as the project's highest-stakes concern, and design
  accordingly.

Future work will define individual capability specifications and ADRs for each domain listed in
Section 4, in the build order above.