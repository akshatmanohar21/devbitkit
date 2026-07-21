# Feature Catalog

> **Purpose.** This document is the canonical inventory of every capability in
> Devbitkit and its current implementation status. Where `DOMAIN_MAP.md`
> answers *does this belong, and under which clause of the Inclusion Test*,
> this document answers *is it built yet*.
>
> This document does not restate qualification reasoning — see
> `DOMAIN_MAP.md` and `DESIGN_NARRATIVE.md` Section 2 for that. It tracks
> status only.

---

# Status Definitions

| Status | Meaning |
|--------|---------|
| Not Started | Passed the Inclusion Test and is listed in the Domain Map. No design or code work has begun. |
| Design | An ADR and/or capability specification is in progress. |
| In Progress | Implementation has started but the capability is not yet usable end-to-end. |
| Implemented | Usable, tested, and documented. |
| Deprecated | Previously implemented, no longer maintained. Reason recorded in an ADR. |

A capability's row should link to its ADR and capability specification once
those exist. Until then, the row simply tracks status against the build order
in `DESIGN_NARRATIVE.md` Section 10.

---

# Catalog

| Domain | Capability | Status | ADR | Capability Spec | Notes |
|--------|-----------|--------|-----|------------------|-------|
| Generators | Password | Implemented | [ADR-0001](ADR/0001-Generators-Domain-Inclusion.md) | [Spec](CAPABILITY_SPECS/generators-password.md) | Flags (`--length`, `--no-letters`, `--no-numbers`, `--no-symbols`, `--count`) implemented, manually verified, and covered by 7 automated tests (length, randomness sanity, per-flag exclusion x3, both error conditions). |
| Generators | API Key | Not Started | — | — | |
| Generators | UUID | Not Started | — | — | |
| Hashing | Hash | Not Started | — | — | Candidate for build order step 3. |
| Encoding | Base64 | Not Started | — | — | Candidate for build order step 3. |
| Encoding | URL Encoding | Not Started | — | — | |
| Encoding | Hex | Not Started | — | — | |
| Time | Unix Timestamp | Not Started | — | — | |
| Time | Timezone | Not Started | — | — | |
| Time | RFC3339 | Not Started | — | — | |
| Time | ISO8601 | Not Started | — | — | |
| Time | Duration | Not Started | — | — | |
| Find/Replace | Find/Replace | Not Started | — | — | |
| Archives | Create | Not Started | — | — | |
| Archives | Extract | Not Started | — | — | |
| Archives | List | Not Started | — | — | |
| Encryption | Envelope Encryption | Not Started | — | — | Requires reviewed ADR before implementation begins (`DESIGN_NARRATIVE.md` Section 9). Build order step 5, after conceptual study in step 4. |

---

# Maintenance

This table should be updated as part of the Capability Lifecycle
(`WORKING_AGREEMENT.md`, Section 5) — specifically at the "Documentation
Update" stage, and again whenever a capability changes status. A capability
should never sit at "Implemented" without a linked capability specification,
and the Encryption domain should never move past "Design" without a linked,
accepted ADR.