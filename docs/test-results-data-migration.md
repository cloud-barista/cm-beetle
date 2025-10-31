# Data Migration API Test Results

> [!IMPORTANT]
>
> - Beetle provides a **relay-style data migration** API for security reasons.
> - `transx` package supports both **direct and relay-style data transfer**.

**Test Date:** October 29, 2025  
**Source Bucket Name:** `aws-ap-northeast-2-bucket-2uk0i5`  
**Target Bucket Name:** `aws-ap-northeast-2-bucket-2uk0i5-radcksl`  
**Target Cloud:** AWS ap-northeast-2  
**CM-Beetle Version:** v0.4.2

<img width="1651" height="575" alt="image" src="https://github.com/user-attachments/assets/d057a1a9-ecad-4ef0-9365-13477b82b93c" />

---

## Preparation

### Generate Dummy Data

**Status:** ✅ SUCCESS

**Run script:**

```bash
cd ~/dev/cloud-barista/cm-beetle/transx/examples/object-storage
./create-test-data.sh

tree /tmp/transx-test-data/
```

<details>
  <Summary>click to see the generated test data</Summary>

```
/tmp/transx-test-data/
├── README.md
├── backup
│   ├── config-backup.json
│   ├── data-backup.csv
│   └── old-backup.txt
├── config.json
├── data
│   ├── config.json
│   ├── dataset1.csv
│   ├── dataset2.csv
│   ├── process.log
│   ├── processed
│   │   ├── output1.csv
│   │   ├── output2.csv
│   │   ├── reports
│   │   │   ├── metrics.csv
│   │   │   ├── monthly.txt
│   │   │   └── quarterly.txt
│   │   └── summary.json
│   ├── raw
│   │   ├── 2025
│   │   │   ├── Q1
│   │   │   │   ├── feb-sales.csv
│   │   │   │   ├── jan-sales.csv
│   │   │   │   ├── process.log
│   │   │   │   └── summary.json
│   │   │   ├── february.csv
│   │   │   ├── january.csv
│   │   │   └── january.json
│   │   ├── metadata.json
│   │   ├── raw1.csv
│   │   └── raw2.csv
│   └── results.json
├── data.csv
├── debug.log
├── docs
│   ├── README.md
│   ├── api-reference.txt
│   ├── images
│   │   ├── diagram.png
│   │   ├── icon.svg
│   │   └── screenshot.jpg
│   ├── tutorial.pdf
│   └── user-guide.txt
├── logs
│   ├── access.log
│   ├── app.log
│   ├── archive
│   │   ├── app-2025-01.log
│   │   ├── app-2025-02.log
│   │   └── errors-2025-01.log
│   ├── debug.log
│   └── error.log
├── node_modules
│   ├── package1
│   │   ├── index.js
│   │   └── package.json
│   └── package2
│       └── package.json
├── root-file.txt
├── src
│   ├── build.log
│   ├── main.go
│   ├── models
│   │   ├── entities
│   │   │   ├── base.go
│   │   │   └── customer.go
│   │   ├── product.go
│   │   ├── schema.json
│   │   └── user.go
│   ├── package.json
│   ├── tests
│   │   ├── coverage.txt
│   │   ├── test.log
│   │   └── user_test.go
│   └── utils.go
├── temp
│   ├── cache.tmp
│   ├── file1.tmp
│   └── file2.tmp
└── temp.tmp

19 directories, 62 files
```

</details>

### Upload the test data to the source object storage

**Status:** ✅ SUCCESS

**Run script:**

```bash
# Change directory
cd ~/dev/cloud-barista/cm-beetle/transx/examples/object-storage

# Upload test data
./migrate.sh -c config-spider-upload.json -v
```

**(sample) config-spider-upload.json**

> [!NOTE]
>
> - masking secrets (e.g. xxxxxxxxxxxxxxxxx)

```json
{
  "source": {
    "endpoint": "",
    "dataPath": "/tmp/transx-test-data/"
  },
  "destination": {
    "endpoint": "http://localhost:1024/spider/s3",
    "dataPath": "aws-ap-northeast-2-bucket-2uk0i5/"
  },
  "destinationTransferOptions": {
    "method": "object-storage",
    "objectStorageOptions": {
      "client": "spider",
      "accessKeyId": "xxxxxxxxxx",
      "expiresIn": 3600,
      "timeout": 300,
      "maxRetries": 3,
      "useSSL": false
    }
  }
}
```

**Uploaded test data**

<img width="1651" height="1230" alt="image" src="https://github.com/user-attachments/assets/160c908e-0846-47dd-8c56-3d2bb8c5b3f8" />

## Test: Migrate Data from Source to Target Object Storage

**Status:** ✅ SUCCESS

**Request:**

> [!NOTE]
>
> - masking secrets (e.g. xxxxxxxxxxxxxxxxx)
> - Use `minio` clinet to access and download data from the source object storage
> - Use `spider` client to access and upload data to the target object storage
> - Set spider endpoint `http://cb-spider:1024/spider/s3` on Cloud-Migrator platform

```bash
curl -s -X POST http://localhost:8056/beetle/migration/data \
  -H "Content-Type: application/json" \
  -u "default:default" \
  -d '{
    "source": {
      "endpoint": "http://s3.ap-northeast-2.amazonaws.com",
      "dataPath": "	aws-ap-northeast-2-bucket-2uk0i5/"
    },
    "sourceTransferOptions": {
      "method": "object-storage",
      "objectStorageOptions": {
        "client": "minio",
        "accessKeyId": "xxxxxxxxxxxxxxxxx",
        "secretAccessKey": "xxxxxxxxxxxxxxxxx",
        "region": "ap-northeast-2",
        "expiresIn": 3600,
        "timeout": 300,
        "maxRetries": 3,
        "useSSL": true
      }
    },
    "destination": {
      "endpoint": "http://cb-spider:1024/spider/s3",
      "dataPath": "aws-ap-northeast-2-bucket-2uk0i5-radcksl/"
    },
    "destinationTransferOptions": {
      "method": "object-storage",
      "objectStorageOptions": {
        "client": "spider",
        "accessKeyId": "xxxxxxxxxxxxxxxxx",
        "expiresIn": 3600,
        "timeout": 300,
        "maxRetries": 3,
        "useSSL": false
      }
    }
  }'
```

**Response:**

```json
{
  "message": "Data migration completed successfully"
}
```

**HTTP Status:** 200 OK

**Migrated data**

<img width="1654" height="1227" alt="image" src="https://github.com/user-attachments/assets/99c46024-c732-40db-a25a-9e7849aeb6fd" />

---

## Summary

### Test Results

| Test | API         | Method | Status     | Response Time |
| ---- | ----------- | ------ | ---------- | ------------- |
| 1    | MigrateData | POST   | ✅ SUCCESS | ~4s           |

--

## Conclusion

- ✅ The Data Migration API has been successfully tested and validated:
