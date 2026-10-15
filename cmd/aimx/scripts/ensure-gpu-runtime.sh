#!/bin/bash
# GPU Driver, Docker Compose & NVIDIA Container Toolkit Initializer for AIMX
# Adapted from CB-Tumblebug (scripts/usecases/llm/installGpuDriver.sh) with idempotent check-then-install flow.

set -e
set -o pipefail

export DEBIAN_FRONTEND=noninteractive
export NEEDRESTART_MODE=a
export NEEDRESTART_SUSPEND=1

if [ -d /etc/needrestart/conf.d ]; then
    echo "\$nrconf{restart} = 'a';" | sudo tee /etc/needrestart/conf.d/99-autorestart.conf > /dev/null 2>&1 || true
fi

APT_OPTS=(-o Dpkg::Options::="--force-confdef" -o Dpkg::Options::="--force-confold" -o DPkg::Lock::Timeout=600)
APT_INSTALL=(sudo DEBIAN_FRONTEND=noninteractive apt-get install -y -qq "${APT_OPTS[@]}")

wait_for_apt_locks() {
    for i in $(seq 1 12); do
        if ! sudo fuser /var/lib/dpkg/lock-frontend >/dev/null 2>&1 && \
           ! sudo fuser /var/lib/apt/lists/lock >/dev/null 2>&1; then
            break
        fi
        echo "  Waiting for apt/dpkg lock... (${i}/12)"
        sleep 5
    done
}

# ============================================================
# 1. GPU Driver Check -> Install if Missing (CB-Tumblebug logic)
# ============================================================
if command -v nvidia-smi >/dev/null 2>&1 && nvidia-smi >/dev/null 2>&1; then
    DRV_VER=$(nvidia-smi --query-gpu=driver_version --format=csv,noheader 2>/dev/null | head -1 || echo "unknown")
    GPU_NAME=$(nvidia-smi --query-gpu=name --format=csv,noheader 2>/dev/null | head -1 || echo "NVIDIA GPU")
    echo "✓ NVIDIA GPU driver already active: ${GPU_NAME} (Driver ${DRV_VER})"
else
    echo "NVIDIA driver not active; running GPU driver installation flow..."
    wait_for_apt_locks

    if ! command -v lspci >/dev/null 2>&1; then
        sudo apt-get -o DPkg::Lock::Timeout=600 update -qq
        "${APT_INSTALL[@]}" pciutils
    fi

    GPU_INFO=$(sudo lspci | grep -i nvidia || true)
    if [ -z "$GPU_INFO" ]; then
        echo "ERROR: No NVIDIA GPU detected via lspci!"
        sudo lspci | grep -i -E "vga|3d|display" || true
        exit 1
    fi
    echo "  Detected PCI GPU: $GPU_INFO"

    BROKEN_PKGS=$(dpkg -l 2>/dev/null | grep -E "nvidia|cuda|libnvidia" | grep -v "^ii " | grep -v "^un " | awk '{print $2}' || true)
    if [ -n "$BROKEN_PKGS" ]; then
        sudo dpkg --force-all --purge $BROKEN_PKGS 2>&1 | tail -5 || true
        sudo rm -rf /var/lib/dkms/nvidia 2>/dev/null || true
        sudo dpkg --configure -a 2>/dev/null || true
    fi

    sudo apt-get -o DPkg::Lock::Timeout=600 update -qq

    if ! grep -q "blacklist nouveau" /etc/modprobe.d/blacklist-nouveau.conf 2>/dev/null; then
        sudo tee /etc/modprobe.d/blacklist-nouveau.conf > /dev/null << 'EOF'
blacklist nouveau
options nouveau modeset=0
EOF
        sudo update-initramfs -u 2>/dev/null || true
    fi

    KERNEL_VERSION=$(uname -r)
    "${APT_INSTALL[@]}" "linux-headers-${KERNEL_VERSION}" 2>&1 | tail -3 || "${APT_INSTALL[@]}" linux-headers-generic 2>&1 | tail -3 || true
    "${APT_INSTALL[@]}" "linux-modules-extra-${KERNEL_VERSION}" build-essential dkms 2>&1 | tail -3 || true

    BUILD_DIR="/lib/modules/${KERNEL_VERSION}/build"
    if [ -f "$BUILD_DIR/.config" ]; then
        KERNEL_GCC_VER=$(grep 'CONFIG_CC_VERSION_TEXT=' "$BUILD_DIR/.config" 2>/dev/null | sed 's/.*gcc.*[ (]\([0-9]\+\)\..*/\1/' || true)
        SYSTEM_GCC_VER=$(gcc -dumpversion 2>/dev/null | cut -d. -f1 || true)
        if [ -n "$KERNEL_GCC_VER" ] && [ -n "$SYSTEM_GCC_VER" ] && [ "$KERNEL_GCC_VER" -gt "$SYSTEM_GCC_VER" ] 2>/dev/null; then
            "${APT_INSTALL[@]}" "gcc-${KERNEL_GCC_VER}" 2>&1 | tail -3
            sudo mkdir -p /etc/dkms/framework.conf.d
            echo "export CC=/usr/bin/gcc-${KERNEL_GCC_VER}" | sudo tee /etc/dkms/framework.conf.d/cc.conf > /dev/null
        fi
    fi

    UBUNTU_VERSION=$(. /etc/os-release 2>/dev/null && echo "${VERSION_ID}" | tr -d '.' || echo "2204")
    ARCH=$(dpkg --print-architecture 2>/dev/null || echo "amd64")
    [ "$ARCH" = "arm64" ] && ARCH_PATH="sbsa" || ARCH_PATH="x86_64"
    [[ "$UBUNTU_VERSION" == "2404" ]] && REPO_DIST="ubuntu2404" || REPO_DIST="ubuntu2204"
    KEYRING_URL="https://developer.download.nvidia.com/compute/cuda/repos/${REPO_DIST}/${ARCH_PATH}/cuda-keyring_1.1-1_all.deb"

    KEYRING_FILE=$(mktemp)
    curl -fsSL "$KEYRING_URL" -o "$KEYRING_FILE" 2>/dev/null || wget -q "$KEYRING_URL" -O "$KEYRING_FILE"
    sudo dpkg -i --force-confdef --force-confold "$KEYRING_FILE" 2>/dev/null || sudo dpkg -i "$KEYRING_FILE"
    rm -f "$KEYRING_FILE"
    sudo apt-get -o DPkg::Lock::Timeout=600 update -qq

    IS_VGPU=false
    if sudo lspci -nn 2>/dev/null | grep -i nvidia | grep -Eqi "GRID|vGPU"; then
        IS_VGPU=true
    elif [ -d /sys/bus/mdev/devices ] && ls -1 /sys/bus/mdev/devices 2>/dev/null | grep -q .; then
        IS_VGPU=true
    fi

    if [ "$IS_VGPU" = true ]; then
        DRIVER_CANDIDATES=("cuda-drivers-580" "nvidia-driver-580" "cuda-drivers")
    else
        DRIVER_CANDIDATES=("cuda-drivers-580-open" "cuda-drivers-580" "nvidia-driver-580-open" "nvidia-driver-580" "cuda-drivers")
    fi

    DRIVER_INSTALLED=false
    for CANDIDATE in "${DRIVER_CANDIDATES[@]}"; do
        echo "  Attempting NVIDIA driver package: ${CANDIDATE}..."
        if "${APT_INSTALL[@]}" -o Dpkg::Options::="--force-overwrite" "$CANDIDATE" 2>&1 | tail -10; then
            DRIVER_INSTALLED=true
            echo "  ✓ Installed ${CANDIDATE}"
            break
        fi
    done

    if [ "$DRIVER_INSTALLED" != true ]; then
        echo "ERROR: Failed to install NVIDIA driver from candidates: ${DRIVER_CANDIDATES[*]}"
        exit 1
    fi

    if ! dpkg -l nvidia-modprobe 2>/dev/null | grep -q "^ii"; then
        "${APT_INSTALL[@]}" nvidia-modprobe 2>&1 | tail -3 || true
    fi
    sudo modprobe nvidia 2>/dev/null || true
    sudo modprobe nvidia-uvm 2>/dev/null || true
fi

# ============================================================
# 2. Docker Engine & Docker Compose v2 Check -> Install if Missing
# ============================================================
if command -v docker >/dev/null 2>&1; then
    echo "✓ Docker Engine already installed: $(docker --version)"
else
    echo "Installing Docker Engine..."
    wait_for_apt_locks
    curl -fsSL https://get.docker.com | sudo sh
fi

if sudo docker compose version >/dev/null 2>&1; then
    echo "✓ Docker Compose v2 already installed: $(sudo docker compose version)"
else
    echo "Installing Docker Compose v2 plugin..."
    wait_for_apt_locks
    sudo apt-get -o DPkg::Lock::Timeout=600 update -qq
    "${APT_INSTALL[@]}" docker-compose-v2 || "${APT_INSTALL[@]}" docker-compose-plugin
fi

# ============================================================
# 3. NVIDIA Container Toolkit Check -> Install & Configure if Missing
# ============================================================
if command -v nvidia-ctk >/dev/null 2>&1; then
    echo "✓ NVIDIA Container Toolkit already installed: $(nvidia-ctk --version 2>/dev/null | head -1)"
else
    echo "Installing NVIDIA Container Toolkit (nvidia-ctk)..."
    wait_for_apt_locks
    curl -fsSL https://nvidia.github.io/libnvidia-container/gpgkey | \
        sudo gpg --dearmor --yes -o /usr/share/keyrings/nvidia-container-toolkit-keyring.gpg
    curl -s -L https://nvidia.github.io/libnvidia-container/stable/deb/nvidia-container-toolkit.list | \
        sed 's#deb https://#deb [signed-by=/usr/share/keyrings/nvidia-container-toolkit-keyring.gpg] https://#g' | \
        sudo tee /etc/apt/sources.list.d/nvidia-container-toolkit.list > /dev/null
    sudo apt-get -o DPkg::Lock::Timeout=600 update -qq
    "${APT_INSTALL[@]}" nvidia-container-toolkit 2>&1 | tail -5
fi

if ! grep -q "nvidia" /etc/docker/daemon.json 2>/dev/null; then
    echo "Configuring Docker daemon for NVIDIA GPU runtime..."
    sudo nvidia-ctk runtime configure --runtime=docker >/dev/null 2>&1 || true
    sudo systemctl restart docker
    echo "✓ Docker NVIDIA runtime configured and restarted"
else
    echo "✓ Docker NVIDIA runtime already configured in /etc/docker/daemon.json"
fi

echo "RUNTIME_READY"
