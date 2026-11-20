# Highlights

Moments in the learning that caught the learner's interest, newest last.

- 2026-11-20 · go-basics ([#2](https://github.com/AhmetSIRIM/learn-go/pull/2), [66c0235](https://github.com/AhmetSIRIM/learn-go/commit/66c02351e26cdeabae07ecdd6604bd1c552f8e12), [spec](https://go.dev/ref/spec#Defer_statements)): `defer` takes a snapshot of its arguments on the defer line: a deferred print showed 1, the count at that moment, although the count was set to 2 before the print ran.
- 2026-11-20 · go-basics ([#2](https://github.com/AhmetSIRIM/learn-go/pull/2), [66c0235](https://github.com/AhmetSIRIM/learn-go/commit/66c02351e26cdeabae07ecdd6604bd1c552f8e12), [spec](https://go.dev/ref/spec#Defer_statements)): Deferred calls run last-in, first-out: two defers in a row printed "deferred second" before "deferred first". That order is what cleanup needs: open a connection, then a transaction on it, and defer each close right after its open; the transaction closes first and the connection last, with no thought about the order.
