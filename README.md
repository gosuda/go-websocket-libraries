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
**Last Updated:** Wed, 07 Jan 2026 03:53:18 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 239.46 | 32.33 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 289.32 | 26.75 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 548.23 | 14.12 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 373.22 | 20.76 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 633.24 | 12.23 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)