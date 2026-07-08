# DECISIONS

Append-only, one line per decision, newest last. Date-prefixed. Architecture-shaping
decisions graduate to an ADR under [`../development/decisions/adrs/`](../development/decisions/adrs/).
Graduate this file to per-file `decisions/<date>--<slug>.md` if size or
parallel-agent merge contention bites.

- 2026-07-06 — Rebuild abcd from scratch in Go, no external tools (specstory,
  RepoPrompt, flow-next, Ralph, codex); ship an MVP, extend via the companion harness then
  Claude Code.
- 2026-07-06 — Transport-agnostic Go core; CLI is the reliable default front door;
  MCP is an additive front door on the same core, added later.
- 2026-07-06 — Peer with the companion harness via conventions + MCP; no Go dependency either way.
- 2026-07-06 — LLM work host-delegated by default; native/CLI/API/MCP oracles are
  opt-in adapters.
- 2026-07-06 — Spec/task layer native-minimal; the companion harness `ccpm` the primary deeper
  backend; flow-next dropped. Autonomous run not a Ralph port (host orchestrators).
- 2026-07-06 — Single repo, curated release (no dev→public mirror); the repo is the
  marketplace. Private companion repo deferred (trigger: shared transcripts).
- 2026-07-06 — Three-tier `.abcd/` layout: development (durable) / work (shared) /
  .work.local (local). `docs/` user-facing only.
- 2026-07-06 — Module path `github.com/REPPL/abcd-cli`; Cobra approved as the CLI
  framework (matches ferry and the companion harness).
- 2026-07-08 — Confidential sources: global user-level corpus (CSL-JSON + grep
  corpus, local no-remote git), append-only JSONL influence ledger per repo,
  banlist patterns generated from confidential entries into the itd-74 private
  guard; convention + skill first, `abcd source` verbs deferred (itd-76). Quarto
  chosen for eventual paper reconstruction; RAG rejected at this scale.
- 2026-07-08 — Personas in any scenario are always Alice, Bob, Carol (in that
  order); the user is they/them. Recorded as a principle.
- 2026-07-08 — itd-76 grilled: leak guard promises literal strings only
  (paraphrase risk stated, handled behaviourally + review); citation is a
  two-level AND (source permission_status AND per-line cited_publicly); author
  bans default on with per-source ban_authors opt-out; standalone `source`
  domain (itd-16 a possible backend, not a dependency); pre-commit auto-
  refreshes the generated banlist; public render proven by structural filter
  AND post-render lint; team share of citation data via committed
  `.abcd/work/references.json` (share/ingest); durability = machine backup +
  git bundle, multi-machine deferred.
- 2026-07-08 — `~/.abcd/` blessed as abcd's user-level home (fourth tier,
  additive to repo `.abcd/`), path configurable; relocation wizard recorded as
  itd-77.
