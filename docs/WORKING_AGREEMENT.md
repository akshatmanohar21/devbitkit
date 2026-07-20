# Working Agreement

> **Purpose.** This document defines how Devbitkit evolves. It establishes the
> principles, processes, and expectations that govern architectural decisions,
> capability additions, implementation, and documentation.
>
> This document intentionally focuses on **how decisions are made**, not **how
> the software is implemented**.

---

# 1. Philosophy

Devbitkit is developed with the belief that software quality is determined long
before code is written.

Implementation is expected to follow thoughtful design rather than drive it.
Architectural decisions should be deliberate, documented, and understandable
months or years after they were made.

The objective is not to build features as quickly as possible, but to build a
tool whose behaviour remains predictable, consistent, and maintainable
throughout its lifetime — and whose scope stays bounded to what it actually
solves, rather than what would merely be convenient to add.

---

# 2. Development Principles

The following principles govern all work within the project.

- Documentation precedes implementation.
- Every capability must pass the Inclusion Test (`DESIGN_NARRATIVE.md`,
  Section 2) before it is discussed as a design, not after.
- Every significant architectural decision is recorded.
- Simplicity is preferred over unnecessary flexibility.
- Consistency is valued over feature count, but consistency is never, on its
  own, a reason to add a capability that does not independently pass the
  Inclusion Test.
- Long-term maintainability is preferred over short-term convenience.

---

# 3. Decision Hierarchy

When disagreements arise, project artifacts take precedence in the following
order.

1. Design Narrative
2. Architecture Decision Records (ADRs)
3. Capability Specification
4. Architecture Documentation
5. Source Code

If implementation contradicts an accepted ADR, the ADR remains authoritative
until it is explicitly superseded.

If any document in this repository — including this one — appears to permit a
capability that would not pass the Inclusion Test in `DESIGN_NARRATIVE.md`
Section 2, the Design Narrative governs. No other document is authorized to
define or loosen the inclusion criteria.

---

# 4. Documentation Responsibilities

Every document within the repository exists to answer one specific question.

| Document | Purpose |
|----------|---------|
| Design Narrative | Why the project exists, and what qualifies for inclusion |
| Project Identity | What the project is |
| Working Agreement | How the project evolves |
| Domain Map | Where capabilities belong, and why each one qualified |
| Capability Specification | What the project can do |
| Architecture | How the project is implemented |
| ADR | Why a specific decision was made |

Documentation should not duplicate responsibilities already owned by another
document. In particular: **no document other than the Design Narrative
defines or restates the criteria for what may be added to the project.**
Other documents may reference the Inclusion Test; none may redefine it.

---

# 5. Capability Lifecycle

Every capability introduced into Devbitkit follows the same lifecycle.

```text
Problem
    ↓
Inclusion Test evaluation (DESIGN_NARRATIVE.md, Section 2)
    ↓
Discussion
    ↓
Design Review
    ↓
ADR (required for any new domain; required before implementation for the
     Encryption domain specifically)
    ↓
Capability Specification
    ↓
Implementation
    ↓
Testing
    ↓
Documentation Update
```

The Inclusion Test evaluation happens first and in writing. A capability that
has not been checked against Section 2 of the Design Narrative should not
proceed to Discussion or Design Review — those stages assume the capability
has already earned a place in the project, not that it's still being argued
for.

No capability should skip later stages unless the change is purely corrective
and does not affect project behaviour.

---

# 6. Architecture Decision Records

Architectural decisions are documented through ADRs.

An ADR should capture:

- The problem being solved.
- When the problem was solved.
- The context surrounding the decision.
- Which clause of the Inclusion Test the decision satisfies, if the ADR
  concerns a new capability or domain.
- The accepted solution.
- Alternatives considered.
- Consequences of the decision.

ADRs are immutable historical records.

Existing ADRs must never be modified to reflect new thinking. Instead, they are
superseded by newer ADRs that explain why the previous decision changed.

---

# 7. Implementation Expectations

Implementation should remain aligned with the project's documented philosophy.

Code should be:

- Predictable.
- Readable.
- Testable.
- Modular.
- Consistent.

Implementation should not introduce behaviour that contradicts established
architectural decisions without first updating the relevant design artifacts.

For the Encryption domain specifically, implementation must not begin until
the corresponding ADR — covering algorithm, mode, nonce strategy, and
key-wrap/rotation scheme — has been reviewed and accepted. This is a stricter
requirement than the general lifecycle in Section 5, given the domain's
correctness and security stakes.

---

# 8. Documentation Standards

Documentation is considered part of the product.

Every document should:

- Explain one concept.
- Avoid implementation details unless explicitly intended.
- Prefer explaining *why* before *how*.
- Use consistent terminology throughout the repository.
- Remain concise without sacrificing clarity.

Documentation should evolve alongside the software rather than after it.

---

# 9. Review Principles

Every proposed capability or architectural change should be evaluated using
the following questions, in this order:

1. **Does this pass the Inclusion Test** (`DESIGN_NARRATIVE.md`, Section 2)?
   If no, the review ends here — none of the following questions are
   evaluated, because they only matter for something already in scope.
2. Does this preserve the command language's consistency (Section 3 of the
   Design Narrative)?
3. Does this introduce complexity disproportionate to the problem it solves?
4. Will a future contributor understand, from the record, why this decision
   was made?

A capability is never accepted solely because it would be useful, well-built,
or easy to add. It is accepted because it passes the Inclusion Test — usefulness
and ease are properties nearly everything has, and neither is evidence of
belonging in this project.

---

# 10. Evolution

Devbitkit is expected to evolve incrementally.

Architectural improvements are encouraged, but should remain compatible with
the project's core philosophy, and specifically must not loosen or bypass the
Inclusion Test.

The project intentionally prefers deliberate evolution over rapid expansion.

Growth should strengthen the command language among capabilities that have
already qualified for inclusion — not increase the number of capabilities for
its own sake. A larger command surface is not, on its own, a sign of a
healthier project.

---

# 11. Status at this Milestone

At this milestone, Devbitkit establishes its development philosophy.

The project adopts documentation-first development, architecture-driven
decision making, and ADR-backed evolution as its governing process — with the
Inclusion Test in `DESIGN_NARRATIVE.md` Section 2 as the single, non-duplicated
gate for what enters the project's scope.

Future contributors are expected to follow these agreements unless they are
explicitly superseded through the project's established decision process.