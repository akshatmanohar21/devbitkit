# Open Threads

> **Purpose.** This document tracks unresolved decisions, corrections owed to
> existing documents, and designed-but-not-yet-implemented work. It exists so
> a session can end mid-thought without losing the reasoning behind where
> things were left.
>
> This is different from `FEATURE_CATALOG.md`, which tracks implementation
> status per capability. This document tracks *everything else that needs
> attention* — including things that aren't capabilities at all (a correction
> to an ADR, a licensing question, an undecided design trade-off).
>
> When an item here is resolved, move the reasoning into its proper home
> (an ADR, a capability spec, the Feature Catalog) and remove it from this
> file. This file should never be the permanent record of a decision — only
> the holding area before one is made.

---

# Corrections Owed

## ADR-0001 overstates Password's justification

`docs/ADR/0001-Generators-Domain-Inclusion.md` currently claims "no adequate
CLI tool exists that defaults to full-entropy random generation without also
defaulting to phonetic/pronounceable generation." This is misleading: `pwgen
-s` already produces fully random, non-phonetic output — the claim is true of
pwgen's *default* behavior, not its *capability*.

**What still holds:** pwgen's default (no flags) is the weaker phonetic mode,
which is a real footgun if someone forgets `-s`. Devbitkit's password
generator has no such footgun — full entropy is unconditional. There's also a
legitimate "marginal cost is near-zero once you're installing `dbk` for
Encryption anyway" argument. Neither of these is "no tool exists," and the
ADR should be corrected to say so honestly rather than overstate the case.

**Action needed:** Rewrite the Decision section of ADR-0001 to state the
narrower, accurate justification (safer default + marginal-cost bundling),
not "no adequate tool exists." Do this before the ADR is treated as settled
history — it's still early enough that a direct edit is reasonable; once
more decisions cite it, it would need a superseding ADR instead per
`WORKING_AGREEMENT.md` Section 6 ("ADRs are immutable historical records...
superseded by newer ADRs").

---

# Design Decisions Made, Not Yet Implemented

## Passphrase mode (merged into `password`, not a separate capability)

After extended discussion, the design landed on: passphrase generation is a
**mode of the existing `password` capability**, not a separate `passphrase`
subcommand. Rationale: to the consumer of the output, a password and a
passphrase are the same kind of thing (a string); the generation strategy is
an internal detail. This mirrors 1Password's own generator UI (one generator,
a mode toggle).

**Flag design agreed:**

- `--passphrase` (bool) — mode switch. Words instead of characters.
- `--length` — shared with character mode. In passphrase mode, means total
  output length in characters (words + separators + any appended digit/symbol),
  not word count. `--word-count` as a separate flag was considered and
  rejected in favor of this unification.
- `--no-numbers` / `--no-symbols` — shared meaning across both modes ("output
  should not contain this character class"), achieved differently
  underneath: excluded from charset in character mode, not appended in
  passphrase mode.
- `--no-letters` — **not valid in passphrase mode** (words are inherently
  letters) and must be explicitly rejected, not silently ignored:
  ```go
  if *passphrase && *noLetters {
      return fmt.Errorf("--no-letters is not valid with --passphrase (words are inherently letters)")
  }
  ```
- `--add-number` / `--add-symbol` — new flags, passphrase-mode only, append a
  single digit/symbol to satisfy password-policy requirements (e.g.
  `elephant-lantern-92`). Reserve 1 character from the `--length` budget when
  set.
- `--strict` (bool) — when set, exact `--length` match is mandatory; errors
  if unreachable, rather than falling back to the closest achievable length.
- Default (non-`--strict`) behavior: attempt exact match; if genuinely
  unreachable (only realistic case: `--length` below 3, the shortest
  possible single word), fall back to the closest achievable length and
  print an explicit note — never silently substitute without saying so.

**Not yet decided:** exact wording of the fallback note; whether
`--capitalize` (capitalize each word) is worth adding — was mentioned once,
never resolved either way.

## Exact-length word-combination algorithm

Verified against the actual embedded wordlist
(`internal/generators/wordlist.txt`, EFF large wordlist, 7,776 words): word
lengths run 3–9 with no gaps (counts: 82 words at length 3, up to 1,779 at
length 8). This gapless property means the sum of `n` word lengths can hit
every integer between `3n` and `9n` continuously — which makes exact-length
matching solvable by direct computation, not search.

**Algorithm agreed (not yet coded):**

1. Precompute `map[int][]string` (word length → words of that length) once,
   from the embedded wordlist.
2. For target total length `L`: try word counts `n = 1, 2, 3, ...`. For each
   `n`, compute `targetSum = L - (n-1)×separatorLength - reservedChars` (the
   latter only if `--add-number`/`--add-symbol` is set). Check whether
   `targetSum` falls within `[3n, 9n]`. The first `n` that passes is
   guaranteed solvable — no trial and error.
3. Distribute `targetSum` across the `n` words: start each at minimum length
   3, randomly hand out the remaining `targetSum - 3n` surplus (each word
   can absorb up to +6) until exhausted.
4. For each word's decided length, pick one random word from that length's
   precomputed bucket via `crypto/rand` (same bias-free approach as
   Password's character selection).

**Why this replaces the earlier "randomized backtracking with retries" idea**
(briefly proposed, then retracted in the same session): retries have
unbounded worst-case cost and are unnecessary once the gapless-range property
is recognized. This algorithm is `O(n)` in the number of words, with no
possibility of retry blowup. Worth remembering as a general pattern: check
whether a generation problem's structure allows direct computation before
reaching for retry-based search.

**Not yet implemented.** This is the most algorithmically complex piece of
code in the project so far — deliberately deferred to a session with full
attention, not squeezed in at the end of a long one.

---

# Design Decisions Not Yet Made

## `--copy` (clipboard) mechanism

Agreed the feature is worth building (avoids leaving a generated password in
terminal scrollback/shell history — ties back to the domain's core
secret-handling justification). **Not decided:** implementation approach.

- **Option A:** shell out via `os/exec` to the OS-native clipboard tool
  (`pbcopy` on macOS; try `wl-copy`, then `xclip`/`xsel` as fallbacks on
  Linux). Stdlib-only, but Linux has no guaranteed clipboard tool
  preinstalled — this is a real, not hypothetical, cross-platform
  fragmentation problem, and Devbitkit would be depending on an external tool
  it can't guarantee is present. Needs a clear, honest error message when
  nothing is available, not a silent no-op.
- **Option B:** a third-party Go clipboard package (e.g. `atotto/clipboard`,
  `golang.design/x/clipboard`) — handles cross-platform detection
  internally, at the cost of exactly the kind of dependency
  `DESIGN_NARRATIVE.md` Section 8 says should be "a supply-chain decision
  worth recording, not a default reach."

No implementation started. Decide before building — this is a real trade-off,
not a default to reach for without thought.

---

# Housekeeping

- **`internal/generators/wordlist.txt` needs to actually be committed to the
  repo.** It was generated and verified this session (7,776 words, EFF large
  wordlist, dice-index column stripped) but needs to be saved into the repo
  by hand.
- **Licensing note owed before open-sourcing.** The EFF wordlist is published
  under a Creative Commons license (likely CC BY, unconfirmed exact variant
  as of this writing). Before the repo goes public, confirm the exact license
  terms directly from EFF's site and add attribution — a `LICENSES.md` or a
  note in the relevant `CAPABILITY_SPEC` is probably the right home for it,
  not yet decided which.
- **`go:embed` not yet introduced.** Needed to bundle `wordlist.txt` into the
  compiled binary. First real use of Go's `embed` package in this project —
  worth introducing deliberately, same as pointers and `flag.FlagSet` were.
- **Empty placeholder doc files** (`ANTI_GOALS.md`, `ARCHITECTURE.md`,
  `CONTRIBUTING.md`, `PRODUCT_PRINCIPLES.md`, `ROADMAP.md`, `START_HERE.md`)
  are sitting in the repo tree. Earlier discussion concluded three duplicate
  `DESIGN_NARRATIVE.md` content and two are premature for a solo
  pre-contributor project — but they were never actually deleted or
  otherwise resolved. Decide: delete, or leave as intentional stubs with a
  one-line note explaining why they're empty.
- **API Key and UUID capabilities** — not started. Both covered by
  ADR-0001 already; no new ADR needed when work begins on them.
- **Password's capability spec** (`docs/CAPABILITY_SPECS/generators-password.md`)
  will need a substantial rewrite once `--passphrase` mode ships — it
  currently only documents character-mode flags and behavior.

---

# How to Use This Document

At the start of a session, check here first, before `FEATURE_CATALOG.md`, if
picking up mid-thread rather than starting a new capability from scratch.
When an item is resolved, move it to its permanent home and delete it from
here — an item lingering here after being resolved elsewhere is a sign this
document has drifted out of sync with reality, the same failure mode the
Domain Map's old "Future Domains" section had.