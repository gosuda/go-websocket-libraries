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
**Last Updated:** Fri, 18 Sep 2026 07:37:14 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 337.88 | 22.96 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 364.90 | 21.26 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 632.27 | 12.29 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 444.44 | 17.47 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 728.04 | 10.69 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)