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
**Last Updated:** Tue, 28 Apr 2026 05:46:26 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 236.40 | 32.63 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 295.00 | 26.30 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 523.54 | 14.80 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 339.88 | 22.76 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 596.86 | 13.02 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)