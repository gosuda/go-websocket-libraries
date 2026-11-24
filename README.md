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
**Last Updated:** Tue, 24 Nov 2026 12:55:23 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 326.88 | 23.77 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.5.0 | 341.63 | 22.72 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 620.38 | 12.50 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.2 | 430.00 | 18.04 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 737.12 | 10.57 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)