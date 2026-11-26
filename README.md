# go-websocket-libraries

Go WebSocket Library Comparison

**Libraries Benchmarked:**

- https://github.com/gorilla/websocket
- https://github.com/lxzan/gws
- https://github.com/gobwas/ws
- https://github.com/lesismal/nbio
- https://github.com/coder/websocket

**Latest Benchmark Results:**

<!-- BENCHMARK_TABLE_START -->
**Last Updated:** Wed, 25 Nov 2026 11:41:22 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 386.69 | 20.07 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.5.0 | 489.08 | 15.90 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 784.41 | 9.90 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.2 | 532.99 | 14.57 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 1054.36 | 7.38 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)
## Benchmark environment

The Depot migration PR moves measurement from GitHub-hosted `ubuntu-latest` to `depot-ubuntu-24.04`. Results recorded before and after that migration use different machines. Do not interpret a change across that boundary as a library improvement or regression. The weekly/manual publisher checks out current `main`; queued refreshes coalesce, and an older rerun cannot cancel an active measurement. Depot runtime and billing savings have not been measured.
