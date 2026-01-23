#!/bin/bash

# Simple migration wrapper script for running migrations with either direct or relay mode

# Color definitions
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# Default mode is direct
MODE="direct"

# Function to display usage
function show_usage {
  echo -e "${BLUE}MariaDB Migration Tool${NC}"
  echo -e "${YELLOW}Usage:${NC}"
  echo -e "  $0 [direct|relay|agent-fwd] [flags]"
  echo -e ""
  echo -e "${YELLOW}Modes:${NC}"
  echo -e "  direct     - Pull mode: SSH source → Local host (backup only)"
  echo -e "  relay      - Relay mode: SSH source → Local → SSH destination"
  echo -e "  agent-fwd  - Agent Forward: SSH source → SSH destination (direct)"
  echo -e ""
  echo -e "${YELLOW}Flags:${NC}"
  echo -e "  --backup     Run only the backup step"
  echo -e "  --transfer   Run only the transfer step"
  echo -e "  --restore    Run only the restore step"
  echo -e "  --verbose    Enable verbose logging"
  echo -e ""
  echo -e "${YELLOW}Examples:${NC}"
  echo -e "  $0 direct --verbose       Pull backup from source to local"
  echo -e "  $0 relay --verbose        Full migration via local relay"
  echo -e "  $0 agent-fwd --verbose    Direct transfer between containers"
  echo -e ""
}

# Function to check prerequisites
function check_prerequisites {
  echo -e "${CYAN}Checking prerequisites...${NC}"
  
  # Check if binary exists
  if [ ! -f "./mariadb-migration" ]; then
    echo -e "${YELLOW}Binary not found, building...${NC}"
    go build -o mariadb-migration main.go
    if [ $? -ne 0 ]; then
      echo -e "${RED}Failed to build binary${NC}"
      exit 1
    fi
    echo -e "${GREEN}✓ Binary built successfully${NC}"
  fi
  
  # Check if SSH keys exist
  if [ ! -f "./ssh_keys/id_rsa" ]; then
    echo -e "${RED}SSH keys not found. Run ./setup_environment.sh all first.${NC}"
    exit 1
  fi
  echo -e "${GREEN}✓ SSH keys found${NC}"
  
  # Check if config file exists
  if [ ! -f "$CONFIG" ]; then
    echo -e "${RED}Configuration file not found: $CONFIG${NC}"
    exit 1
  fi
  echo -e "${GREEN}✓ Configuration file found: $CONFIG${NC}"
  
  # Check if containers are running
  SOURCE_RUNNING=$(docker ps --filter "name=mariadb_source" --filter "status=running" -q)
  TARGET_RUNNING=$(docker ps --filter "name=mariadb_target" --filter "status=running" -q)
  
  if [ -z "$SOURCE_RUNNING" ] || [ -z "$TARGET_RUNNING" ]; then
    echo -e "${YELLOW}MariaDB containers not running. Starting environment...${NC}"
    ./setup_environment.sh all
    if [ $? -ne 0 ]; then
      echo -e "${RED}Failed to setup environment${NC}"
      exit 1
    fi
  else
    echo -e "${GREEN}✓ MariaDB containers are running${NC}"
  fi
  
  echo ""
}

# Function to show pre-migration status
function show_pre_migration_status {
  echo -e "${CYAN}Pre-migration database status:${NC}"
  echo -e "${YELLOW}Source database (testdb):${NC}"
  docker exec mariadb_source mariadb -uroot -psource_password --ssl=0 -e 'SELECT COUNT(*) as user_count FROM testdb.users; SELECT COUNT(*) as order_count FROM testdb.orders;' 2>/dev/null
  
  echo -e "${YELLOW}Target database (testdb):${NC}"
  TARGET_RESULT=$(docker exec mariadb_target mariadb -uroot -ptarget_password --ssl=0 -e 'SELECT COUNT(*) as user_count FROM testdb.users; SELECT COUNT(*) as order_count FROM testdb.orders;' 2>/dev/null)
  if [ $? -eq 0 ]; then
    echo "$TARGET_RESULT"
  else
    echo "No data (tables may not exist)"
  fi
  echo ""
}

# Function to show post-migration status
function show_post_migration_status {
  echo -e "${CYAN}Post-migration database status:${NC}"
  echo -e "${YELLOW}Target database (testdb):${NC}"
  docker exec mariadb_target mariadb -uroot -ptarget_password --ssl=0 -e 'SELECT COUNT(*) as user_count FROM testdb.users; SELECT COUNT(*) as order_count FROM testdb.orders;' 2>/dev/null
  echo ""
}

# Check if no arguments or help requested
if [ $# -eq 0 ] || [ "$1" == "--help" ] || [ "$1" == "-h" ]; then
  show_usage
  exit 0
fi

# Parse mode argument
if [ "$1" == "direct" ] || [ "$1" == "relay" ] || [ "$1" == "agent-fwd" ]; then
  MODE="$1"
  shift
fi

# Set configuration file based on mode
case "$MODE" in
  "relay")
    CONFIG="relay-mode-config.json"
    ;;
  "agent-fwd")
    CONFIG="agent-fwd-mode-config.json"
    ;;
  *)
    CONFIG="direct-mode-config.json"
    ;;
esac

# Check prerequisites
check_prerequisites

# For agent-fwd mode, create temporary config with actual host IP
ACTUAL_CONFIG="$CONFIG"
TEMP_CONFIG=""
if [ "$MODE" == "agent-fwd" ]; then
  HOST_IP=$(hostname -I | awk '{print $1}')
  echo -e "${YELLOW}Agent-Forward mode: Host IP = ${HOST_IP}${NC}"
  TEMP_CONFIG=$(mktemp)
  sed "s/HOST_IP_PLACEHOLDER/${HOST_IP}/g" "$CONFIG" > "$TEMP_CONFIG"
  ACTUAL_CONFIG="$TEMP_CONFIG"
fi

# Show pre-migration status
show_pre_migration_status

# Construct command to run
CMD_ARGS="--config=$ACTUAL_CONFIG $@"

echo -e "${BLUE}=========================================${NC}"
echo -e "${BLUE}Running MariaDB Migration (${MODE} mode)${NC}"
echo -e "${BLUE}=========================================${NC}"

echo -e "${GREEN}Executing: ./mariadb-migration --config=$CONFIG $@${NC}"
./mariadb-migration $CMD_ARGS
RESULT=$?

# Clean up temp config if created
if [ -n "$TEMP_CONFIG" ] && [ -f "$TEMP_CONFIG" ]; then
  rm -f "$TEMP_CONFIG"
fi

# Check result and show post-migration status
if [ $RESULT -eq 0 ]; then
  echo -e "${GREEN}Migration completed successfully!${NC}"
  show_post_migration_status
else
  echo -e "${RED}Migration failed!${NC}"
  exit 1
fi
