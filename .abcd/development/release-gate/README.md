# Release gate — the pre-tag procedure

Every `v*` release must clear this gate before the tag is pushed. The gate spans
**two enforcement planes**, by design: deterministic checks run in CI (they need
no model); semantic checks run **host-side in the agent harness** (they need a
model, which CI does not have). This runbook is the single human enumeration of
both planes; it owns the semantic half and defers to
[`release.yml`](../../../.github/workflows/release.yml) as the source of truth
for the deterministic half.

Design of record: [`../plans/2026-07-11-iss35-semantic-release-gate.md`](../plans/2026-07-11-iss35-semantic-release-gate.md).

## Deterministic gates (CI-enforced)

The [`release.yml`](../../../.github/workflows/release.yml) `verify` job runs
these, in order, on macOS + Linux. `release.yml` is authoritative; this list is
the human-readable mirror.

1. Format (gofmt)
2. Build
3. Vet
4. Test
5. Test (race, internal)
6. Record-lint (design-record drift gate)
7. Docs-lint (docs-currency gate)
8. Reviews-charter discipline (RD001-RD003)
9. Smoke every command (self-discovering harness)

> **Design target — not yet armed.** A `record-lint` rule keeping this list in
> lockstep with `release.yml`'s `verify` steps lands in phase 3 (see the design
> doc). Until it arms, this mirror is kept honest by hand; per
> [`../principles/loud-staging.md`](../principles/loud-staging.md) the gap is
> stated, not hidden.

## Semantic gates (host-run, before the tag)

CI cannot run these — they spawn LLM agents. Run them in the agent harness
against the exact commit to be tagged:

1. **`docs-currency-reviewer`** — verifies every user-facing claim still matches
   the code (the semantic complement of `docs lint`; see
   [`../brief/04-surfaces/10-docs.md`](../brief/04-surfaces/10-docs.md)).
2. **Brief↔surface cross-check** — [`brief-surface-crosscheck.js`](brief-surface-crosscheck.js),
   the Direction-A semantic half of the iss-35 graduation: the brief's surface
   *prose* (flags, sub-verbs, exit codes, schema fields, counts) vs. the shipped
   binary's actual behaviour. The deterministic Direction-B half is the
   `surface_coverage` `record-lint` rule and already runs in CI.

## Recording the semantic verdict

Each semantic pass records its outcome as a **signed, commit-sha-keyed receipt**
— a Verification Summary Attestation (VSA) shape carrying `verificationResult`
(PROMOTE / HOLD), the pinned judge-model snapshot, the detector version, and the
failing categories. A `record-lint` rule then verifies, before a tag, that a
PROMOTE receipt exists for `HEAD`'s commit; a missing or HOLD receipt **blocks**
the release (fail-closed — an un-run semantic pass is never a silent pass).

> **Design target — not yet armed.** The receipt schema + fail-closed verify
> rule land in phase 2, the signing + in-`release.yml` verification in phase 4
> (see the design doc). Until then the semantic gates are enforced by this
> runbook as maintainer discipline, and that provisional status is stated here
> rather than implied away.

## Procedure

1. Land all work on the release commit; open the release the normal way (branch
   → PR → merge). The `verify` job gates the merge.
2. On the merged commit, run the two semantic gates above in the harness.
3. Record each verdict as a receipt keyed to that commit's sha.
4. Tag `vX.Y.Z` on the commit. Once the fail-closed verify rule is armed, the
   tag is rejected unless every semantic receipt is present and PROMOTE.
