# ADR-0001: Generators Domain Inclusion

## Status

Accepted (recorded retroactively — see Process Note below).

## Problem

Developers regularly need to generate passwords, API keys, and UUIDs for
testing and setup purposes. No single offline tool covers all three:

- Secure password/API-key generation has no standard baseline tool on either
  major OS. `pwgen` is not preinstalled and must be separately installed; its
  default mode (pronounceable/phonetic generation) trades randomness for
  memorability, which is not a safe default for credential generation.
- `uuidgen` exists on both Linux (util-linux) and macOS (BSD), but the two
  implementations diverge in default case, supported flags, and available
  UUID versions.

## When

Decided during initial project scoping, prior to any implementation.
Recorded retroactively in this ADR after the Password capability was already
implemented — see Process Note.

## Context

This domain was evaluated against the Inclusion Test defined in
`DESIGN_NARRATIVE.md`, Section 2, which requires at least one of:

- **Test A (secret-adjacent absence):** the workflow handles sensitive
  material and no adequate offline tool exists.
- **Test B (cross-platform fragmentation):** a baseline tool exists but
  diverges genuinely across operating systems.

## Decision

Generators is accepted as a domain.

- **Password and API-key generation** qualify under Test A. Generating
  credentials via a browser-based tool means trusting an unverifiable third
  party with material that may end up in real use; an offline, local binary
  removes that trust dependency entirely. No adequate CLI tool exists that
  defaults to full-entropy random generation without also defaulting to (or
  only offering) phonetic/pronounceable generation.
- **UUID generation** qualifies under Test B, given the confirmed divergence
  between GNU (util-linux) and BSD `uuidgen`.

## Alternatives Considered

- **Rely on existing tools per-platform** (`openssl rand -base64`,
  `/dev/urandom` one-liners, `pwgen`, platform `uuidgen`) — rejected as the
  status quo the domain exists to improve; requires remembering multiple
  tools and installing at least one (`pwgen`) that isn't preinstalled
  anywhere.
- **Wrap `pwgen` rather than reimplement** — rejected. `pwgen`'s default
  behavior (phonetic generation) is a security trade-off Devbitkit does not
  want as a default, and wrapping an external binary reintroduces an install
  dependency the Offline First principle is meant to avoid.

## Consequences

- Devbitkit owns correctness of its own randomness source
  (`crypto/rand`, not `math/rand`) for every capability in this domain.
  Getting this wrong would be a security defect, not a cosmetic bug.
- Future capabilities proposed under this domain (e.g. any additional
  generator) must still independently satisfy the Inclusion Test — domain
  membership does not grant automatic inclusion to new capabilities within
  it.
- Phonetic/memorable password generation, if ever requested, must be an
  explicit opt-in flag, never a default, per `DESIGN_NARRATIVE.md`'s
  Correctness Before Breadth principle applied here.

## Process Note

This ADR was written after the Password capability was already implemented,
not before, which is a deviation from the Capability Lifecycle defined in
`WORKING_AGREEMENT.md` Section 5 (ADR precedes Capability Specification,
which precedes Implementation). The reasoning captured here already existed
in project discussion prior to implementation — it was simply never recorded
as a formal ADR at the time. Going forward, any new domain's ADR is written
and accepted before its Domain Map entry is added, not after.