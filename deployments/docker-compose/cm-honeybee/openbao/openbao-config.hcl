# OpenBao server configuration for cm-honeybee's secrets backend.
#
# Stores SSH access info and CSP credentials for cm-honeybee (KV v2 engine at
# secret/). TLS is disabled for local/dev; enable it (and a KMS auto-unseal
# stanza) for production. See README.md.
#
# Reference: https://openbao.org/docs/configuration/

# Persistent storage — data survives container restarts.
storage "file" {
  path = "/openbao/data"
}

# TCP listener (TLS disabled for local development).
listener "tcp" {
  address     = "0.0.0.0:8200"
  tls_disable = true
}

api_addr = "http://0.0.0.0:8200"

# Disable mlock for container compatibility (IPC_LOCK cap handles memory locking).
disable_mlock = true

# Web UI at http://<host>:8201/ui
ui = true
