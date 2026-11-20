# Workflow

## The next piece of work is asked as a choice
The agent asks for the next piece of work with `AskUserQuestion`: two or three
options, its recommendation first. Each option says in one sentence what sets it
apart, on the axes where the options actually differ:

- depth or breadth: building on finished work, or opening new ground;
- the area: the Go language, server-side development, or computer science
  fundamentals.

Within a piece of work the agent sets the tasks in order without asking.

Why: the learner wants the direction visible without carrying the decision. The
options make it visible, and the recommendation carries the decision unless the
learner shifts it. Asking only between pieces of work, with few options, keeps the
choice light; a question at every task, or a long menu, can overwhelm the learner.

## Plans stay short and loose
A look ahead (a roadmap, the next few pieces of work) is offered rarely, and held as a
direction rather than a commitment: the learner's focus and interests change from
session to session.

Why: a detailed roadmap tires the learner and goes stale as the direction changes.

## One piece of work per pull request, by default
A piece of work is done on its own branch, in atomic commits. A pull request closes
it; the learner approves opening and merging it, and its description follows
`.github/pull_request_template.md`. It is merged with a merge commit,
never squashed. The learner may choose another way for a piece of work; the agent
follows it.

Why: the commit history is the record of the learning; a squash collapses it into one
commit and the record is gone. A default rather than a fixed rule leaves room to try
other ways of working.

## Directories follow depth and breadth
A piece of work that builds on an earlier one extends that piece's directory; a piece
of work that opens new ground starts a new directory at the repository root. A
directory is named for what it covers, whether that is a level, a domain, or
something else. Directory names follow Go's package naming convention.

Why: a directory grows as its line of learning deepens, and a new one marks new
ground, so each line stays in one place. The split follows the same depth-or-breadth
axis the next piece of work is chosen on. Package-style names build one naming habit
across programs and packages. This repository keeps each directory at the root rather
than under `cmd/`.

Source: Go, "Organizing a Go module" (go.dev/doc/modules/layout); Effective Go,
"Package names".
