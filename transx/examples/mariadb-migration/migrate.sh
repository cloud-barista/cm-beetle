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
  echo -e "${BLUE}Simple MariaDB Migration Tool${NC}"
  echo -e "${YELLOW}Usage:${NC}"
  echo -e "  $0 [direct|relay] [flags]"
  echo -e ""
  echo -e "${YELLOW}Modes:${NC}"
  echo -e "  direct    - Direct mode migration (default)"
  echo -e "  relay     - Relay mode migration (source → relay node → destination)"
  echo -e ""
  echo -e "${YELLOW}Flags:${NC}"
  echo -e "  --backup     Run only the backup step"
  echo -e "  --transfer   Run only the transfer step"
  echo -e "  --restore    Run only the restore step"
  echo -e "  --verbose    Enable verbose logging"
  echo -e ""
  echo -e "${YELLOW}Examples:${NC}"
  echo -e "  $0 direct --verbose      Run direct mode migration with verbose logging"
  echo -e "  $0 relay --backup        Run only the backup step in relay mode"
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
if [ "$1" == "direct" ] || [ "$1" == "relay" ]; then
  MODE="$1"
  shift
fi

# Set configuration file based on mode
if [ "$MODE" == "relay" ]; then
  CONFIG="relay-mode-config.json"
else
  CONFIG="direct-mode-config.json"
fi

# Check prerequisites
check_prerequisites

# Show pre-migration status
show_pre_migration_status

# Construct command to run
CMD_ARGS="--config=$CONFIG $@"

echo -e "${BLUE}=========================================${NC}"
echo -e "${BLUE}Running MariaDB Migration (${MODE} mode)${NC}"
echo -e "${BLUE}=========================================${NC}"

echo -e "${GREEN}Executing: ./mariadb-migration $CMD_ARGS${NC}"
./mariadb-migration $CMD_ARGS

# Check result and show post-migration status
if [ $? -eq 0 ]; then
  echo -e "${GREEN}Migration completed successfully!${NC}"
  show_post_migration_status
else
  echo -e "${RED}Migration failed!${NC}"
  exit 1
fi
