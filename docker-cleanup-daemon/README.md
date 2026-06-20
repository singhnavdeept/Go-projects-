# Docker Cleanup Daemon

A high-performance, lightweight background service written in Go for Windows to periodically clean up stopped Docker containers, dangling images, and unused volumes from **Docker Desktop**.

---

## 🚀 Key Features

* **Safety Thresholds**: Filter and protect recently stopped development containers by setting an age threshold (e.g., only delete containers stopped for more than 24 hours).
* **CLI Inspection Mode**: Run a quick, safe check using command-line arguments without modifying the config file or running a background service.
* **Dry Run Mode**: Simulates resource pruning, logging what *would* have been deleted (including programmatically checking mounts for unused volumes).
* **Graceful Signal Handling**: Captures interrupt signals (`Ctrl+C`, `SIGTERM`) to shutdown cleanly, completing any active jobs first.
* **Startup Run**: Automatically performs a scan and cleanup cycle immediately on boot so you don't have to wait for the scheduler to trigger.

---

## ⚡ System Footprint

Go compiles directly to a native Windows machine code executable with zero runtime dependencies. Its footprint is extremely minimal:

* **RAM (Memory)**: **~10.4 MB**
* **CPU Usage**: **0%** (suspended by the OS kernel, waking up only at schedule intervals)
* **Binary Size**: **~10 MB** (self-contained executable)

---

## 🛠️ Configuration (`config.yaml`)

Manage the daemon's behavior in the `config.yaml` file located in the same directory:

```yaml
schedule: "0 */6 * * *"   # Cron schedule (runs every 6 hours)
cleanup:
  stopped_containers: true
  dangling_images: true
  unused_volumes: false
thresholds:
  containers_older_than: "24h" # Only prune containers stopped for longer than 24 hours
dry_run: false
```

* **schedule**: A standard 5-field cron spec. E.g., `*/10 * * * *` for every 10 minutes.
* **cleanup**: Toggles for which resource categories to scan.
* **thresholds**: Supports standard Go duration values (e.g., `2h`, `24h`, `30m`). Use `0s` to prune stopped resources instantly.
* **dry_run**: If `true`, logs operations without deleting any data.

---

## 🏃 How to Run

Make sure **Docker Desktop for Windows** is running.

### 1. Run as a Continuous Daemon
Uses settings defined in `config.yaml` and executes on the schedule in the background:
```powershell
.\docker-cleanup-daemon.exe
```

### 2. Check-only Dry Run (One-time)
Quickly inspect what would be deleted and exit immediately:
```powershell
.\docker-cleanup-daemon.exe --dry-run --once
```

### 3. Run Live Cleanup Once
Executes a single cleanup run immediately and exits:
```powershell
.\docker-cleanup-daemon.exe --once
```

---

## 🛠️ Build and Compilation

To rebuild the binary from source:

1. Ensure Go (v1.20+) is installed.
2. Build using the Go toolchain:
   ```powershell
   go build -o docker-cleanup-daemon.exe
   ```
