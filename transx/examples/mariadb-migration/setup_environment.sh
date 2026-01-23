#!/bin/bash

# Environment setup script for MariaDB migration testing
# This script sets up MariaDB containers with SSH access for testing migrations

# Color definitions
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Script directory (for storing SSH keys)
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SSH_KEY_DIR="${SCRIPT_DIR}/ssh_keys"

echo -e "${BLUE}=========================================${NC}"
echo -e "${BLUE}MariaDB Migration Environment Setup Script${NC}"
echo -e "${BLUE}=========================================${NC}"

# Check for required tools
echo -e "${YELLOW}Checking for required tools...${NC}"

# Check for Docker
if ! command -v docker &> /dev/null; then
    echo -e "${RED}Docker is not installed. Please install Docker:${NC}"
    echo -e "    curl -sSL get.docker.com | sh"
    echo -e "    sudo usermod -aG docker \${USER}"
    exit 1
else
    echo -e "${GREEN}✓ Docker is installed${NC}"
fi

# Check if Docker daemon is running
if ! docker info &> /dev/null; then
    echo -e "${RED}Docker daemon is not running. Please start Docker:${NC}"
    echo -e "    sudo systemctl start docker"
    exit 1
else
    echo -e "${GREEN}✓ Docker daemon is running${NC}"
fi

# Function to generate SSH keypair for container access
generate_ssh_keys() {
    echo -e "${YELLOW}Generating SSH keypair for container access...${NC}"
    
    # Create SSH key directory
    mkdir -p "${SSH_KEY_DIR}"
    
    # Generate keypair if not exists
    if [ ! -f "${SSH_KEY_DIR}/id_rsa" ]; then
        ssh-keygen -t rsa -b 4096 -f "${SSH_KEY_DIR}/id_rsa" -N "" -C "mariadb-migration-test"
        chmod 600 "${SSH_KEY_DIR}/id_rsa"
        chmod 644 "${SSH_KEY_DIR}/id_rsa.pub"
        echo -e "${GREEN}✓ SSH keypair generated${NC}"
    else
        echo -e "${GREEN}✓ SSH keypair already exists${NC}"
    fi
}

# Function to stop and remove existing containers
cleanup_containers() {
    echo -e "${YELLOW}Cleaning up existing containers...${NC}"
    docker stop mariadb_source mariadb_target 2>/dev/null || true
    docker rm mariadb_source mariadb_target 2>/dev/null || true
    
    # Clean up backup directories
    rm -rf ~/mariadb_source_backup/* 2>/dev/null || true
    rm -rf ~/mariadb_target_backup/* 2>/dev/null || true
    rm -rf ~/mariadb_backup/* 2>/dev/null || true
    
    # Clean up SSH keys
    if [ -d "${SSH_KEY_DIR}" ]; then
        rm -rf "${SSH_KEY_DIR}"
        echo -e "${GREEN}✓ SSH keys cleaned up${NC}"
    fi
    
    echo -e "${GREEN}✓ Cleaned up existing containers and data${NC}"
}

# Function to create Dockerfile for MariaDB with SSH
create_dockerfile() {
    echo -e "${YELLOW}Creating Dockerfile for MariaDB with SSH...${NC}"
    
    cat > /tmp/mariadb-ssh.Dockerfile << 'EOF'
FROM mariadb:latest

# Install SSH server and required packages
RUN apt-get update && apt-get install -y \
    openssh-server \
    rsync \
    && rm -rf /var/lib/apt/lists/* \
    && mkdir -p /var/run/sshd \
    && mkdir -p /root/.ssh \
    && chmod 700 /root/.ssh

# Configure SSH
RUN sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config \
    && sed -i 's/#PubkeyAuthentication yes/PubkeyAuthentication yes/' /etc/ssh/sshd_config \
    && sed -i 's/#PasswordAuthentication yes/PasswordAuthentication no/' /etc/ssh/sshd_config \
    && echo "AllowAgentForwarding yes" >> /etc/ssh/sshd_config

# Create entrypoint script that starts both SSH and MariaDB
RUN echo '#!/bin/bash\n\
/usr/sbin/sshd\n\
exec docker-entrypoint.sh "$@"' > /docker-entrypoint-ssh.sh \
    && chmod +x /docker-entrypoint-ssh.sh

EXPOSE 22 3306

ENTRYPOINT ["/docker-entrypoint-ssh.sh"]
CMD ["mariadbd", "--ssl=0"]
EOF

    echo -e "${GREEN}✓ Dockerfile created${NC}"
}

# Function to build custom MariaDB image with SSH
build_mariadb_ssh_image() {
    echo -e "${YELLOW}Building MariaDB with SSH image...${NC}"
    
    create_dockerfile
    docker build -t mariadb-ssh:latest -f /tmp/mariadb-ssh.Dockerfile /tmp
    rm /tmp/mariadb-ssh.Dockerfile
    
    echo -e "${GREEN}✓ MariaDB with SSH image built${NC}"
}

# Create backup directories
create_directories() {
    echo -e "${YELLOW}Creating backup directories...${NC}"
    mkdir -p ~/mariadb_source_backup
    mkdir -p ~/mariadb_target_backup
    mkdir -p ~/mariadb_backup  # Local backup for direct mode
    echo -e "${GREEN}✓ Backup directories created${NC}"
}

# Function to start MariaDB source container with SSH
start_source_container() {
    echo -e "${YELLOW}Starting MariaDB source container with SSH...${NC}"
    
    docker run -d --name mariadb_source \
        -e MARIADB_ROOT_PASSWORD=source_password \
        -e MARIADB_DATABASE=testdb \
        -e MARIADB_USER=testuser \
        -e MARIADB_PASSWORD=testpass \
        -p 3306:3306 \
        -p 2222:22 \
        -v ~/mariadb_source_backup:/backup \
        mariadb-ssh:latest
    
    echo -e "${YELLOW}Waiting for source container to be ready...${NC}"
    sleep 10
    
    # Copy SSH public key to container
    docker cp "${SSH_KEY_DIR}/id_rsa.pub" mariadb_source:/root/.ssh/authorized_keys
    docker exec mariadb_source chmod 600 /root/.ssh/authorized_keys
    docker exec mariadb_source chown root:root /root/.ssh/authorized_keys
    
    # For Agent-Forward: copy private key to source container (to SSH to target)
    docker cp "${SSH_KEY_DIR}/id_rsa" mariadb_source:/root/.ssh/id_rsa
    docker exec mariadb_source chmod 600 /root/.ssh/id_rsa
    
    # Add target container to known_hosts (will be updated after target starts)
    docker exec mariadb_source bash -c "mkdir -p /root/.ssh && touch /root/.ssh/known_hosts"
    
    # Wait for MariaDB to be ready
    while ! docker exec mariadb_source mariadb -uroot -psource_password --ssl=0 -e "SELECT 1" &>/dev/null; do
        echo -e "${YELLOW}Waiting for source MariaDB to start...${NC}"
        sleep 3
    done
    
    echo -e "${GREEN}✓ MariaDB source container is running${NC}"
    echo -e "${GREEN}  - MariaDB port: 3306${NC}"
    echo -e "${GREEN}  - SSH port: 2222${NC}"
}

# Function to start MariaDB target container with SSH
start_target_container() {
    echo -e "${YELLOW}Starting MariaDB target container with SSH...${NC}"
    
    docker run -d --name mariadb_target \
        -e MARIADB_ROOT_PASSWORD=target_password \
        -e MARIADB_DATABASE=testdb \
        -e MARIADB_USER=testuser \
        -e MARIADB_PASSWORD=testpass \
        -p 3307:3306 \
        -p 2223:22 \
        -v ~/mariadb_target_backup:/backup \
        mariadb-ssh:latest
    
    echo -e "${YELLOW}Waiting for target container to be ready...${NC}"
    sleep 10
    
    # Copy SSH public key to container
    docker cp "${SSH_KEY_DIR}/id_rsa.pub" mariadb_target:/root/.ssh/authorized_keys
    docker exec mariadb_target chmod 600 /root/.ssh/authorized_keys
    docker exec mariadb_target chown root:root /root/.ssh/authorized_keys
    
    # Wait for MariaDB to be ready
    while ! docker exec mariadb_target mariadb -uroot -ptarget_password --ssl=0 -e "SELECT 1" &>/dev/null; do
        echo -e "${YELLOW}Waiting for target MariaDB to start...${NC}"
        sleep 3
    done
    
    echo -e "${GREEN}✓ MariaDB target container is running${NC}"
    echo -e "${GREEN}  - MariaDB port: 3307${NC}"
    echo -e "${GREEN}  - SSH port: 2223${NC}"
}

# Function to configure SSH connectivity between containers
configure_ssh_connectivity() {
    echo -e "${YELLOW}Configuring SSH connectivity between containers...${NC}"
    
    # Get host IP (for container-to-container communication via host)
    HOST_IP=$(hostname -I | awk '{print $1}')
    
    # Add target to source's known_hosts (for Agent-Forward mode)
    docker exec mariadb_source bash -c "ssh-keyscan -p 2223 ${HOST_IP} >> /root/.ssh/known_hosts 2>/dev/null"
    docker exec mariadb_source bash -c "ssh-keyscan -p 2223 host.docker.internal >> /root/.ssh/known_hosts 2>/dev/null" || true
    
    # Also add localhost entries for testing
    docker exec mariadb_source bash -c "ssh-keyscan -p 2223 localhost >> /root/.ssh/known_hosts 2>/dev/null" || true
    
    echo -e "${GREEN}✓ SSH connectivity configured${NC}"
    echo -e "${GREEN}  Host IP: ${HOST_IP}${NC}"
}

# Function to populate source database with test data
populate_test_data() {
    echo -e "${YELLOW}Populating source database with test data...${NC}"
    docker exec mariadb_source mariadb -uroot -psource_password --ssl=0 -e "
        USE testdb;
        CREATE TABLE IF NOT EXISTS users (
            id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(100) NOT NULL,
            email VARCHAR(100) UNIQUE NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
        INSERT INTO users (name, email) VALUES 
            ('John Doe', 'john@example.com'),
            ('Jane Smith', 'jane@example.com'),
            ('Bob Johnson', 'bob@example.com'),
            ('Alice Brown', 'alice@example.com'),
            ('Charlie Wilson', 'charlie@example.com');
        
        CREATE TABLE IF NOT EXISTS orders (
            id INT AUTO_INCREMENT PRIMARY KEY,
            user_id INT,
            product_name VARCHAR(100),
            quantity INT,
            price DECIMAL(10,2),
            order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (user_id) REFERENCES users(id)
        );
        INSERT INTO orders (user_id, product_name, quantity, price) VALUES 
            (1, 'Laptop', 1, 1299.99),
            (2, 'Mouse', 2, 29.99),
            (3, 'Keyboard', 1, 89.99),
            (1, 'Monitor', 1, 249.99),
            (4, 'Headphones', 1, 159.99);
    "
    echo -e "${GREEN}✓ Test data populated in source database${NC}"
}

# Function to test SSH connectivity
test_ssh_connectivity() {
    echo -e "${YELLOW}Testing SSH connectivity...${NC}"
    
    HOST_IP=$(hostname -I | awk '{print $1}')
    
    # Test SSH to source container
    echo -e "${YELLOW}Testing SSH to source container (port 2222)...${NC}"
    if ssh -i "${SSH_KEY_DIR}/id_rsa" -o StrictHostKeyChecking=no -o ConnectTimeout=5 -p 2222 root@localhost echo "SSH OK" 2>/dev/null; then
        echo -e "${GREEN}✓ SSH to source container works${NC}"
    else
        echo -e "${RED}✗ SSH to source container failed${NC}"
    fi
    
    # Test SSH to target container
    echo -e "${YELLOW}Testing SSH to target container (port 2223)...${NC}"
    if ssh -i "${SSH_KEY_DIR}/id_rsa" -o StrictHostKeyChecking=no -o ConnectTimeout=5 -p 2223 root@localhost echo "SSH OK" 2>/dev/null; then
        echo -e "${GREEN}✓ SSH to target container works${NC}"
    else
        echo -e "${RED}✗ SSH to target container failed${NC}"
    fi
    
    # Test SSH from source to target (Agent-Forward scenario)
    echo -e "${YELLOW}Testing SSH from source to target (Agent-Forward)...${NC}"
    if docker exec mariadb_source ssh -i /root/.ssh/id_rsa -o StrictHostKeyChecking=no -o ConnectTimeout=5 -p 2223 root@${HOST_IP} echo "SSH OK" 2>/dev/null; then
        echo -e "${GREEN}✓ SSH from source to target works (Agent-Forward ready)${NC}"
    else
        echo -e "${RED}✗ SSH from source to target failed${NC}"
        echo -e "${YELLOW}  Note: This is needed for relay mode (Agent-Forward)${NC}"
    fi
}

# Display final status
display_status() {
    HOST_IP=$(hostname -I | awk '{print $1}')
    
    echo ""
    echo -e "${BLUE}=========================================${NC}"
    echo -e "${GREEN}Environment Setup Completed!${NC}"
    echo -e "${BLUE}=========================================${NC}"
    
    echo ""
    echo -e "${YELLOW}Container Status:${NC}"
    docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}" --filter "name=mariadb_"
    
    echo ""
    echo -e "${YELLOW}SSH Key Location:${NC}"
    echo -e "  Private Key: ${SSH_KEY_DIR}/id_rsa"
    echo -e "  Public Key:  ${SSH_KEY_DIR}/id_rsa.pub"
    
    echo ""
    echo -e "${YELLOW}Source Container:${NC}"
    echo -e "  MariaDB: localhost:3306 (root/source_password)"
    echo -e "  SSH:     localhost:2222 (root, key auth)"
    
    echo ""
    echo -e "${YELLOW}Target Container:${NC}"
    echo -e "  MariaDB: localhost:3307 (root/target_password)"
    echo -e "  SSH:     localhost:2223 (root, key auth)"
    
    echo ""
    echo -e "${YELLOW}Host IP (for container-to-container):${NC} ${HOST_IP}"
    
    echo ""
    echo -e "${YELLOW}Test SSH connections:${NC}"
    echo -e "  ssh -i ${SSH_KEY_DIR}/id_rsa -p 2222 root@localhost"
    echo -e "  ssh -i ${SSH_KEY_DIR}/id_rsa -p 2223 root@localhost"
    
    echo ""
    echo -e "${YELLOW}Run migration:${NC}"
    echo -e "  ${GREEN}./mariadb-migration --config=direct-mode-config.json --verbose${NC}"
    echo -e "  ${GREEN}./mariadb-migration --config=relay-mode-config.json --verbose${NC}"
}

# Main setup process
case "${1:-all}" in
    "cleanup")
        cleanup_containers
        ;;
    "source")
        generate_ssh_keys
        create_directories
        build_mariadb_ssh_image
        start_source_container
        populate_test_data
        test_ssh_connectivity
        ;;
    "target")
        generate_ssh_keys
        create_directories
        build_mariadb_ssh_image
        start_target_container
        test_ssh_connectivity
        ;;
    "all")
        cleanup_containers
        generate_ssh_keys
        create_directories
        build_mariadb_ssh_image
        start_source_container
        start_target_container
        configure_ssh_connectivity
        populate_test_data
        test_ssh_connectivity
        display_status
        ;;
    "test-ssh")
        test_ssh_connectivity
        ;;
    *)
        echo -e "${RED}Usage: $0 [all|cleanup|source|target|test-ssh]${NC}"
        echo -e "${YELLOW}  all      - Setup both containers with SSH and test data (default)${NC}"
        echo -e "${YELLOW}  cleanup  - Stop and remove containers, clean up SSH keys${NC}"
        echo -e "${YELLOW}  source   - Setup only source container${NC}"
        echo -e "${YELLOW}  target   - Setup only target container${NC}"
        echo -e "${YELLOW}  test-ssh - Test SSH connectivity${NC}"
        exit 1
        ;;
esac
