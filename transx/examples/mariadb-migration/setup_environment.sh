#!/bin/bash

# Environment setup script for MariaDB migration testing
# This script sets up local MariaDB containers for testing migrations

# Color definitions
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

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
    echo -e "${RED}Docker is required for running MariaDB containers${NC}"
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
# Create backup directories
echo -e "${YELLOW}Creating backup directories...${NC}"
mkdir -p ~/mariadb_source_backup
mkdir -p ~/mariadb_target_backup
mkdir -p ~/mariadb_relay_temp
echo -e "${GREEN}✓ Backup directories created${NC}"

# Function to stop and remove existing containers
cleanup_containers() {
    echo -e "${YELLOW}Cleaning up existing containers...${NC}"
    docker stop mariadb_source mariadb_target 2>/dev/null || true
    docker rm mariadb_source mariadb_target 2>/dev/null || true
    echo -e "${GREEN}✓ Cleaned up existing containers${NC}"
}

# Function to start MariaDB source container
start_source_container() {
    echo -e "${YELLOW}Starting MariaDB source container...${NC}"
    docker run -d --name mariadb_source \
        -e MARIADB_ROOT_PASSWORD=source_password \
        -e MARIADB_DATABASE=testdb \
        -e MARIADB_USER=testuser \
        -e MARIADB_PASSWORD=testpass \
        -p 3306:3306 \
        -v ~/mariadb_source_backup:/backup \
        mariadb:latest --ssl=0
    
    echo -e "${YELLOW}Waiting for source container to be ready...${NC}"
    sleep 15
    
    # Wait for MariaDB to be ready
    while ! docker exec mariadb_source mariadb -uroot -psource_password --ssl=0 -e "SELECT 1" &>/dev/null; do
        echo -e "${YELLOW}Waiting for source MariaDB to start...${NC}"
        sleep 3
    done
    
    echo -e "${GREEN}✓ MariaDB source container is running on port 3306${NC}"
}

# Function to start MariaDB target container
start_target_container() {
    echo -e "${YELLOW}Starting MariaDB target container...${NC}"
    docker run -d --name mariadb_target \
        -e MARIADB_ROOT_PASSWORD=target_password \
        -e MARIADB_DATABASE=testdb \
        -e MARIADB_USER=testuser \
        -e MARIADB_PASSWORD=testpass \
        -p 3307:3306 \
        -v ~/mariadb_target_backup:/backup \
        mariadb:latest --ssl=0
    
    echo -e "${YELLOW}Waiting for target container to be ready...${NC}"
    sleep 15
    
    # Wait for MariaDB to be ready
    while ! docker exec mariadb_target mariadb -uroot -ptarget_password --ssl=0 -e "SELECT 1" &>/dev/null; do
        echo -e "${YELLOW}Waiting for target MariaDB to start...${NC}"
        sleep 3
    done
    
    echo -e "${GREEN}✓ MariaDB target container is running on port 3307${NC}"
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

# Main setup process
case "${1:-all}" in
    "cleanup")
        cleanup_containers
        ;;
    "source")
        start_source_container
        populate_test_data
        ;;
    "target")
        start_target_container
        ;;
    "all")
        cleanup_containers
        start_source_container
        start_target_container
        populate_test_data
        ;;
    *)
        echo -e "${RED}Usage: $0 [all|cleanup|source|target]${NC}"
        echo -e "${YELLOW}  all     - Setup both containers with test data (default)${NC}"
        echo -e "${YELLOW}  cleanup - Stop and remove existing containers${NC}"
        echo -e "${YELLOW}  source  - Setup only source container${NC}"
        echo -e "${YELLOW}  target  - Setup only target container${NC}"
        exit 1
        ;;
esac

echo -e "${BLUE}=========================================${NC}"
echo -e "${GREEN}Container setup completed!${NC}"
echo -e "${BLUE}=========================================${NC}"

# Display container status
if [ "${1:-all}" = "all" ] || [ "${1:-all}" = "source" ] || [ "${1:-all}" = "target" ]; then
    echo ""
    echo -e "${YELLOW}Container Status:${NC}"
    docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}" --filter "name=mariadb_"
    
    echo ""
    echo -e "${YELLOW}Database Connection Information:${NC}"
    echo -e "${GREEN}Source MariaDB:${NC}"
    echo -e "  Host: localhost"
    echo -e "  Port: 3306"
    echo -e "  Root Password: source_password"
    echo -e "  Database: testdb"
    echo -e "  User: testuser / Password: testpass"
    echo ""
    echo -e "${GREEN}Target MariaDB:${NC}"
    echo -e "  Host: localhost"
    echo -e "  Port: 3307"
    echo -e "  Root Password: target_password"
    echo -e "  Database: testdb"
    echo -e "  User: testuser / Password: testpass"
    echo ""
    
    echo -e "${YELLOW}Test the connections:${NC}"
    echo -e "${GREEN}docker exec mariadb_source mysql -uroot -psource_password -e 'SELECT COUNT(*) FROM testdb.users;'${NC}"
    echo -e "${GREEN}docker exec mariadb_target mysql -uroot -ptarget_password -e 'SHOW DATABASES;'${NC}"
    echo ""
fi

echo -e "${YELLOW}Available migration configuration files:${NC}"
echo -e "${GREEN}direct-mode-config.json${NC} - Direct mode migration (local-to-local)"
echo -e "${GREEN}relay-mode-config.json${NC} - Relay mode migration"
echo ""
echo -e "${YELLOW}How to run migration:${NC}"
echo -e "${GREEN}./mariadb-migration --config=direct-mode-config.json --verbose${NC}"
echo -e "${GREEN}./mariadb-migration --config=relay-mode-config.json --verbose${NC}"
echo ""
echo -e "${YELLOW}Container Management:${NC}"
echo -e "${GREEN}./setup_environment.sh all${NC}      - Setup both containers with test data"
echo -e "${GREEN}./setup_environment.sh cleanup${NC}  - Stop and remove containers"
echo -e "${GREEN}./setup_environment.sh source${NC}   - Setup only source container"
echo -e "${GREEN}./setup_environment.sh target${NC}   - Setup only target container"
echo ""
