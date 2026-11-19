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

## One directory per piece of work
Each piece of work lives in its own directory at the repository root: a program as
`package main`, run with `go run ./<name>`, or a library package with its tests.
Directory names follow Go package names: lowercase, one word, no hyphens or
underscores.

Why: a directory per piece of work keeps each one isolated and easy to find again,
and an unfinished one does not get in the way of the next. Package-style names build
one naming habit across programs and packages. Go's module layout guide suggests
`cmd/` for repositories that mix commands and packages; this repository keeps each
piece of work whole at the root instead.

Source: Go, "Organizing a Go module" (go.dev/doc/modules/layout); Effective Go,
"Package names".
