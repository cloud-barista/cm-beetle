# Beetle API Samples to Create and Delete Target Infra for Testing

Curl command samples for upper-level subsystems to create and delete test infrastructure via Beetle API.

## Prerequisites

- Beetle server running on `http://localhost:8056`
- Default authentication credentials: `default:default`
- CB-Tumblebug instance initialized and accessible
  - **Version matching required**:
    - Deployed Tumblebug version = Cloned Tumblebug version (for `make init`)
    - `cb-tumblebug/assets/assets/dump.gz` file version = Deployed Tumblebug version
  - **If version mismatch**: Clean up existing metadata and re-initialize with matching versions

## Table of Contents

1. [Infrastructure Recommendation](#infrastructure-recommendation)
2. [Infrastructure Creation](#infrastructure-creation)
3. [Infrastructure Deletion](#infrastructure-deletion)
4. [Troubleshooting: Individual Resource Deletion](#troubleshooting-individual-resource-deletion)

---

## Infrastructure Recommendation

Recommend optimal cloud infrastructure based on on-premise infrastructure information.

### Request

<details>
<summary>Click to expand full curl command</summary>

```bash
curl -u "default:default" -X 'POST' \
  'http://localhost:8056/beetle/recommendation/mci?option=register&connectionName=aws-ap-northeast-2' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "desiredCspAndRegionPair": {
    "csp": "aws",
    "region": "ap-northeast-2"
  },
  "nameSeed": "",
  "onpremiseInfraModel": {
    "network": {
      "ipv4Networks": {
        "defaultGateways": [
          {
            "interfaceName": "ens5",
            "ip": "10.0.1.1",
            "machineId": "ec268ed7-821e-9d73-e79f-961262161624"
          },
          {
            "interfaceName": "ens5",
            "ip": "10.0.1.1",
            "machineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
          },
          {
            "interfaceName": "ens5",
            "ip": "10.0.1.1",
            "machineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
          }
        ]
      },
      "ipv6Networks": {}
    },
    "nodes": [
      {
        "cpu": {
          "architecture": "x86_64",
          "cores": 1,
          "cpus": 1,
          "maxSpeed": 2.499,
          "model": "Intel(R) Xeon(R) Platinum 8259CL CPU @ 2.50GHz",
          "threads": 2,
          "vendor": "GenuineIntel"
        },
        "hostname": "ip-10-0-1-30",
        "machineId": "ec268ed7-821e-9d73-e79f-961262161624",
        "memory": {
          "available": 1,
          "totalSize": 2,
          "type": "DDR4"
        },
        "os": {
          "id": "ubuntu",
          "idLike": "debian",
          "name": "Ubuntu",
          "prettyName": "Ubuntu 22.04.3 LTS",
          "version": "22.04.3 LTS (Jammy Jellyfish)",
          "versionCodename": "jammy",
          "versionId": "22.04"
        }
      },
      {
        "cpu": {
          "architecture": "x86_64",
          "cores": 2,
          "cpus": 1,
          "maxSpeed": 2.499,
          "model": "Intel(R) Xeon(R) Platinum 8175M CPU @ 2.50GHz",
          "threads": 4,
          "vendor": "GenuineIntel"
        },
        "hostname": "ip-10-0-1-221",
        "machineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "memory": {
          "available": 15,
          "totalSize": 16,
          "type": "DDR4"
        },
        "os": {
          "id": "ubuntu",
          "idLike": "debian",
          "name": "Ubuntu",
          "prettyName": "Ubuntu 22.04.3 LTS",
          "version": "22.04.3 LTS (Jammy Jellyfish)",
          "versionCodename": "jammy",
          "versionId": "22.04"
        }
      },
      {
        "cpu": {
          "architecture": "x86_64",
          "cores": 1,
          "cpus": 1,
          "maxSpeed": 2.499,
          "model": "Intel(R) Xeon(R) Platinum 8259CL CPU @ 2.50GHz",
          "threads": 2,
          "vendor": "GenuineIntel"
        },
        "hostname": "ip-10-0-1-138",
        "machineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "memory": {
          "available": 7,
          "totalSize": 8,
          "type": "DDR4"
        },
        "os": {
          "id": "ubuntu",
          "idLike": "debian",
          "name": "Ubuntu",
          "prettyName": "Ubuntu 22.04.3 LTS",
          "version": "22.04.3 LTS (Jammy Jellyfish)",
          "versionCodename": "jammy",
          "versionId": "22.04"
        }
      }
    ]
  }
}'
```

</details>

### Notes

- **option=register**: Registers the recommendation result in CB-Tumblebug
- **connectionName**: Specifies the CB-Tumblebug connection to use (e.g., `aws-ap-northeast-2`)
- **nameSeed=""**: Using an empty string for `nameSeed` is recommended for testing as it simplifies resource naming and makes cleanup easier
- The request includes detailed on-premise infrastructure information for 3 nodes
- Firewall rules, network interfaces, and routing tables can be included for more accurate recommendations

### Alternative: Using JSON File

If you have the request data in a file (e.g., `recommendation-request.json`):

```bash
curl -u "default:default" -X 'POST' \
  'http://localhost:8056/beetle/recommendation/mci?option=register&connectionName=aws-ap-northeast-2' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d @recommendation-request.json
```

---

## Infrastructure Creation

Create multi-cloud infrastructure based on the recommendation result.

**Note**: Infrastructure creation is typically handled automatically when using `option=register` in the recommendation API. The following commands are for manual creation or verification purposes.

### Get All MCIs in a Namespace

Retrieve all migrated infrastructures in a namespace:

```bash
curl -u "default:default" -X 'GET' \
  'http://localhost:8056/beetle/migration/ns/mig01/mci' \
  -H 'accept: application/json'
```

### Get Specific MCI Details

Retrieve details of a specific MCI:

```bash
curl -u "default:default" -X 'GET' \
  'http://localhost:8056/beetle/migration/ns/mig01/mci/mmci01' \
  -H 'accept: application/json'
```

---

## Infrastructure Deletion

Delete the entire multi-cloud infrastructure including all associated resources (VMs, VNets, Security Groups, SSH Keys).

**This is the recommended way to clean up infrastructure.** Deleting the MCI will automatically remove all associated resources.

### Delete MCI

```bash
curl -u "default:default" -X 'DELETE' \
  'http://localhost:8056/beetle/migration/ns/mig01/mci/mmci01?option=terminate' \
  -H 'accept: application/json'
```

**Options:**

- `option=terminate`: Terminates and deletes all resources (recommended)
- `option=force`: Force deletion even if resources are in use

**Example:**

```bash
# Delete MCI with all resources
curl -u "default:default" -X 'DELETE' \
  'http://localhost:8056/beetle/migration/ns/mig01/mci/mmci01?option=terminate' \
  -H 'accept: application/json'
```

---

## Troubleshooting: Individual Resource Deletion

**⚠️ Important**: The following commands are for troubleshooting purposes only. Use them only when MCI deletion fails or leaves orphaned resources. In normal scenarios, deleting the MCI (as described in the [Infrastructure Deletion](#infrastructure-deletion) section) is sufficient.

### When to Use Individual Resource Deletion

Use these commands when:

- MCI deletion fails or hangs
- Orphaned resources remain after MCI deletion
- You need to clean up specific resources without affecting others
- Debugging resource dependency issues

### Deletion Order

If you need to delete resources individually, follow this order to avoid dependency issues:

1. MCI (Multi-Cloud Infrastructure)
2. Security Groups
3. SSH Keys
4. VNets and Subnets

### Delete Security Groups

Delete all security groups in the namespace:

```bash
curl -u "default:default" -X 'DELETE' \
  'http://localhost:8056/beetle/migration/ns/mig01/resources/securityGroup' \
  -H 'accept: application/json'
```

Delete a specific security group:

```bash
curl -u "default:default" -X 'DELETE' \
  'http://localhost:8056/beetle/migration/ns/mig01/resources/securityGroup/my-sg-01' \
  -H 'accept: application/json'
```

### Delete SSH Key

Delete the SSH key used for VM access:

```bash
curl -u "default:default" -X 'DELETE' \
  'http://localhost:8056/beetle/migration/ns/mig01/resources/sshKey/mig-sshkey-01' \
  -H 'accept: application/json'
```

### Delete VNet and Subnets

Delete the virtual network and all associated subnets:

```bash
curl -u "default:default" -X 'DELETE' \
  'http://localhost:8056/beetle/migration/ns/mig01/resources/vNet/mig-vnet-01?action=withsubnets' \
  -H 'accept: application/json'
```

**Actions:**

- `action=withsubnets`: Deletes the VNet along with all subnets

Delete only the VNet (without subnets):

```bash
curl -u "default:default" -X 'DELETE' \
  'http://localhost:8056/beetle/migration/ns/mig01/resources/vNet/mig-vnet-01' \
  -H 'accept: application/json'
```

### Complete Cleanup Script (For Troubleshooting)

**Note**: This script performs individual resource deletion and should only be used for troubleshooting. For normal cleanup, use the MCI deletion command described in the [Infrastructure Deletion](#infrastructure-deletion) section.

Here's a bash script to perform complete cleanup when troubleshooting:

<details>
<summary>Click to expand cleanup script</summary>

```bash
#!/bin/bash

BASE_URL="http://localhost:8056/beetle"
AUTH="default:default"
NAMESPACE="mig01"
MCI_ID="mmci01"
SSHKEY_ID="mig-sshkey-01"
VNET_ID="mig-vnet-01"

echo "🗑️  Starting cleanup process..."

# Step 1: Delete MCI
echo "1️⃣  Deleting MCI: $MCI_ID"
curl -u "$AUTH" -X 'DELETE' \
  "$BASE_URL/migration/ns/$NAMESPACE/mci/$MCI_ID?option=terminate" \
  -H 'accept: application/json'
echo -e "\n"

# Wait for MCI deletion to complete
echo "⏳ Waiting for MCI deletion to complete..."
sleep 10

# Step 2: Delete Security Groups
echo "2️⃣  Deleting Security Groups"
curl -u "$AUTH" -X 'DELETE' \
  "$BASE_URL/migration/ns/$NAMESPACE/resources/securityGroup" \
  -H 'accept: application/json'
echo -e "\n"

# Step 3: Delete SSH Key
echo "3️⃣  Deleting SSH Key: $SSHKEY_ID"
curl -u "$AUTH" -X 'DELETE' \
  "$BASE_URL/migration/ns/$NAMESPACE/resources/sshKey/$SSHKEY_ID" \
  -H 'accept: application/json'
echo -e "\n"

# Step 4: Delete VNet
echo "4️⃣  Deleting VNet: $VNET_ID"
curl -u "$AUTH" -X 'DELETE' \
  "$BASE_URL/migration/ns/$NAMESPACE/resources/vNet/$VNET_ID?action=withsubnets" \
  -H 'accept: application/json'
echo -e "\n"

echo "✅ Cleanup complete!"
```

</details>

Save this as `cleanup.sh`, make it executable (`chmod +x cleanup.sh`), and run it when troubleshooting resource cleanup issues.

### Check Resource Status

Before performing individual resource deletion, check the status of resources in your namespace:

```bash
# List all MCIs
curl -u "default:default" -X 'GET' \
  'http://localhost:8056/beetle/migration/ns/mig01/mci' \
  -H 'accept: application/json'

# List all VNets
curl -u "default:default" -X 'GET' \
  'http://localhost:8056/beetle/migration/ns/mig01/resources/vNet' \
  -H 'accept: application/json'

# List all Security Groups
curl -u "default:default" -X 'GET' \
  'http://localhost:8056/beetle/migration/ns/mig01/resources/securityGroup' \
  -H 'accept: application/json'

# List all SSH Keys
curl -u "default:default" -X 'GET' \
  'http://localhost:8056/beetle/migration/ns/mig01/resources/sshKey' \
  -H 'accept: application/json'
```

---

## Common Issues

1. **Authentication Failed**: Verify credentials are correct (`default:default`)
2. **Resource Not Found**: Check that the namespace, MCI ID, or resource ID exists
3. **MCI Deletion Failed**: If MCI deletion fails, refer to the [Troubleshooting: Individual Resource Deletion](#troubleshooting-individual-resource-deletion) section
4. **Timeout**: Some operations may take time; wait a few moments before retrying
5. **Orphaned Resources**: If resources remain after MCI deletion, use individual resource deletion commands from the troubleshooting section

---

## Additional Resources

- [Beetle API Documentation](../api/swagger.yaml)
- [Installation and Execution Guide](./installation-and-execution.md)
- [API Development Guide](./api-development-guide.md)

---

**Last Updated**: 2026-07-06
