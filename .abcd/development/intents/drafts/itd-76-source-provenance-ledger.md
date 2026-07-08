---
id: itd-76
slug: source-provenance-ledger
spec_id: null
kind: standalone
suggested_kind: null
reclassification_history: []
---

# abcd Lets You Consult Sources You Cannot Cite — and Remembers Every Debt

## Press Release

> **Consult any source freely; cite only by deliberate human choice — abcd keeps the ledger in between.** A developer-researcher often works from material they are not free to name in public: working papers under review, a collaborator's private repository, notes shared under an NDA. The agent should be able to *read* that material when it bears on a decision — and must never *cite* it in anything published. abcd manages this as three coupled pieces. A **local-only corpus** (a user-level directory outside every repo, itself a no-remote git repository) holds the documents and a machine-readable bibliography (CSL-JSON; a `custom` block carries `confidential`, `permission_status`, retrieval keywords, and banned aliases). An **append-only provenance ledger** (JSONL, one file per consuming repo) records every meaningful influence — which source, which decision, what claim, what kind of influence — with `cited_publicly: false` until the human flips it. And the **guardrails are mechanical**, not aspirational: the confidential entries *generate* the pattern block in each repo's untracked private-names banlist, so the existing pre-commit guard blocks the leak even when discipline fails; a `cite-guard` scan clears any text before it is shared. Later, the same machine-readable trail — sources, decisions, claims — reconstructs an academic paper, rendered publicly only through the `confidential != true` filter.
>
> "I could never let an agent near my working papers before, because one helpful footnote could burn a collaborator's trust," said Alice, a researcher-developer. "Now it reads everything, records what influenced what, and cites nothing. When Bob's paper is finally published, I flip one flag and the citation appears in the next render — with the whole influence trail already written."

## Why This Matters

Automatic citation is a virtue that becomes a breach the moment a source is confidential: an agent that helpfully names "the working paper this design follows" in a commit message has leaked something no history rewrite fully recalls. The naive fix — keep the material away from the agent — throws away exactly the context that makes its design work good. The resolution is to split *consultation* from *citation* and put a durable, machine-readable record between them: influence is captured eagerly and automatically (cheap, local, append-only), citation happens lazily and manually (when permission exists). The ledger is also the seed of something bigger: a paper whose claims trace to decisions and whose decisions trace to sources is *reconstructable* rather than rewritten from memory.

This composes three existing abcd designs rather than inventing new machinery: the two-layer name banlist ([itd-74](itd-74-name-banlist.md)) supplies the leak guard; the append-only audit chain ([itd-16](itd-16-hash-chain-merkle-audit.md)) supplies the tamper-evidence model the ledger's no-remote git history approximates; and the provenance substrate ([`09-provenance-substrate.md`](../../brief/05-internals/09-provenance-substrate.md)) already defines citation blocks, a source registry, and an NDA-aware publish gate for *ingested* content — this intent extends the same stance to *consulted* content.

## What It Looks Like

- **`abcd source add`** registers a document: CSL-JSON entry (confidentiality, permission status, keywords, aliases), original stored, text extracted into a grep-friendly corpus. Consultation is plain search over that corpus — no index, no service.
- **`abcd source ledger`** appends an influence record — `{decision_ref, claim, source_key, locator, influence, cited_publicly}` — and commits it; corrections are new lines, never edits.
- **`abcd source sync-banlist`** projects every confidential entry's identifying strings into the repo's untracked private-names banlist (the itd-74 private layer), so the pre-commit guard enforces what the convention promises. **`abcd source cite-check`** scans any text and reports offending entries by key only — its output is safe to relay.
- **Paper reconstruction** walks the ledger: claims grouped by decision, citations resolved from the bibliography, rendered to PDF and HTML from one markdown source; the public render draws only on the `confidential != true` filter and the human-flipped `cited_publicly` flags.

## Dogfood (already running)

The convention-first prototype is live for this repo's development: the corpus, ledger, guard scripts, and an agent-side consultation skill exist as a user-level scaffold, and the generated banlist block feeds the same `.githooks/pre-commit` guard that itd-74 generalises. The feature is to lift that scaffold into abcd verbs so any managed repo inherits it. See the plan ([`2026-07-08-confidential-sources-scaffold.md`](../../plans/2026-07-08-confidential-sources-scaffold.md)) and the SOTA survey ([`2026-07-08-confidential-sources-provenance-sota.md`](../../research/notes/2026-07-08-confidential-sources-provenance-sota.md)).

## Open Questions

- Verb shape: a `source` domain under the CLI with plugin-surface commands, or fold `ledger` into the itd-16 audit chain as a sub-verb application?
- Ledger ownership once work spans machines: per-repo files in the user-level corpus (current), or per-repo with a sync discipline?
- Whether `sync-banlist` should run automatically from the pre-commit guard itself (freshness) or stay an explicit step (surprise-free).
- How the paper pipeline's public render *proves* it consulted only permitted citations — a lint over the rendered bibliography against the ledger, or a build that structurally cannot see confidential entries.
