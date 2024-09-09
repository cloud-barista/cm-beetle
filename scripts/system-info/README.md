# System Information Extractor

This script, written in Python, collects various hardware and software information from the host system.
It can be useful for extracting and referencing the configuration of the source computing environment in cloud migration projects.
This script helps to quickly understand the current state of the system and diagnose issues.
It is useful for checking and referencing the configuration of the source computing environment for the cloud migrator project.

## Prerequisites

Ensure the apt packages are intstalled.

```bash
sudo apt-get install python3-dev
```

## Tested on

- `Python 3.10.12`
- `pip 20.0.2`

## Dependencies

- `psutil`
- `prettytable`
- `netifaces`
- `pyyaml`
- `GPUtil`
- `distro`

### Setup venv (optional)

```bash
cd [PROJECT_ROOT]/scripts/system-info
python3 -m venv .venv
source .venv/bin/activate
```

Note - deactivate venv

```bash
deactivate
```

### Install dependencies

```bash
pip3 install -r requirements.txt
```

## Execution

```bash
python3 system_info.py
```

## Features

### 1. OS Information Collection

Collects OS name, version, architecture, platform, boot time, currently logged-in user, kernel version, and distribution information.

```python
def get_os_info():
    """Fetch Operating System information."""
    ...
    return "OS", [
        ("Name", platform.system()),
        ...
    ]
```

### 2. CPU Information Collection

Collects CPU brand, core count, current frequency, context switches, interrupts, architecture, user time, system time, and idle time.

```python
def get_cpu_info():
    """Fetch CPU information."""
    ...
    return "CPU", [
        ("Brand", platform.processor()),
        ...
    ]
```

### 3. Memory Information Collection

Collects total memory, available memory, used memory, cached memory, and swap memory information.

```python
def get_memory_info():
    """Fetch Memory information."""
    ...
    return "Memory", [
        ("Total (GB)", bytes_to_gb(virtual_memory.total)),
        ...
    ]
```

### 4. Disk Information Collection

Collects total disk capacity, available capacity, used capacity, read and write count.

```python
def get_disk_info():
    """Fetch Disk information."""
    ...
    return "Disk", [
        ("Total (GB)", bytes_to_gb(disk_info.total)),
        ...
    ]
```

### 5. Network Information Collection

Collects host name, IP address, MAC address, active network interfaces, gateway, and DNS information.

```python
def get_network_info():
    """Fetch Network information."""
    ...
    return "Network", [
        ("Hostname", socket.gethostname()),
        ...
    ]
```

### 6. Advanced Network Information Collection

Collects IP address, netmask, broadcast, and MAC address of each network interface.

```python
def get_advanced_network_info():
    """Fetch advanced network information using netifaces."""
    ...
    return "Advanced Network", iface_info
```

### 7. GPU Information Collection

Collects GPU name and driver version (if GPUtil is installed).

```python
def get_gpu_info():
    """Fetch GPU information if available."""
    ...
    return "GPU", info_list
```

### 8. Graphic Card Information Collection

Collects VGA compatible controller information.

```python
def get_graphic_card_info():
    ...
    return "Graphic Card", info_list
```

### 9. Process Information Collection

Collects information on all running processes and users on the system.

```python
def get_process_info():
    """Fetch Process information."""
    ...
    return "Process", sorted(process_list)
```

### 10. Data Saving

Saves all collected system information to a YAML file.

```python
def save_to_yaml(data):
    """Save dictionary data to YAML file."""
    ...
```

## Output

The collected information is printed to the console in a neat table format, and is also saved to a `system_info.yaml` file.

Example

```
System Information:
+-------+--------------+-----------------------+----------------------------------------------------------+
| Index |    Group     |        Category       |                           Info                           |
+-------+--------------+-----------------------+----------------------------------------------------------+
|   1   |      OS      |          Name         |                          Linux                           |
|   2   |      OS      |        Version        |       #224-Ubuntu SMP Mon Jun 19 13:30:12 UTC 2023       |
|   3   |      OS      |      Architecture     |                     ('64bit', 'ELF')                     |
|   4   |      OS      |        Platform       | Linux-4.15.0-213-generic-x86_64-with-Ubuntu-18.04-bionic |
|   5   |      OS      |       Boot Time       |                   2023-09-18 11:11:51                    |
|   6   |      OS      |      Logged User      |                           son                            |
|   7   |      OS      |     Kernel Version    |                    4.15.0-213-generic                    |
|   8   |      OS      |      Distribution     |                   Ubuntu 18.04 bionic                    |
|   9   |     CPU      |         Brand         |                          x86_64                          |
|   10  |     CPU      |  Core Count (Logical) |                            4                             |
|   11  |     CPU      | Core Count (Physical) |                            4                             |
|   12  |     CPU      |   Current Frequency   |                        3600.0 MHz                        |
|   13  |     CPU      |    Context Switches   |                        3132898179                        |
|   14  |     CPU      |       Interrupts      |                        1738814431                        |
|   15  |     CPU      |      Architecture     |                          x86_64                          |
|   16  |     CPU      |     User Time (%)     |                           17.0                           |
|   17  |     CPU      |    System Time (%)    |                           0.0                            |
|   18  |     CPU      |     Idle Time (%)     |                           21.0                           |
|   19  |    Memory    |       Total (GB)      |                           7.79                           |
|   20  |    Memory    |     Available (GB)    |                           2.22                           |
|   21  |    Memory    |       Used (GB)       |                           4.85                           |
|   22  |    Memory    |      Cached (GB)      |                           1.95                           |
|   23  |    Memory    |    Swap Total (GB)    |                           2.0                            |
|   24  |    Memory    |     Swap Free (GB)    |                           0.86                           |
|   25  |    Memory    |     Swap Used (GB)    |                           1.14                           |
|   26  |     Disk     |       Total (GB)      |                          97.87                           |
|   27  |     Disk     |       Free (GB)       |                           3.95                           |
|   28  |     Disk     |       Used (GB)       |                          88.91                           |
|   29  |     Disk     |       Read Count      |                         5742877                          |
|   30  |     Disk     |      Write Count      |                         23171054                         |
|   31  |   Network    |        Hostname       |                           son                            |
|   32  |   Network    |       IP Address      |                        127.0.1.1                         |
|   33  |   Network    |      MAC Address      |                    aa:83:85:16:60:d5                     |
|   34  |   Network    |   Active Interfaces   |             lo, docker0, enp0s3, veth6c9eb12             |
|   35  |   Network    |        Gateway        |                        127.0.1.1                         |
|   36  |   Network    |          DNS          |                        127.0.1.1                         |
|   37  | Graphic Card |         Status        |                      Not Available                       |
+-------+--------------+-----------------------+----------------------------------------------------------+

Advanced Network Information:
+-----------------+------------+---------------+----------------+-------------------+
|    Interface    | IP Address |    Netmask    |   Broadcast    |    MAC Address    |
+-----------------+------------+---------------+----------------+-------------------+
|        lo       | 127.0.0.1  |   255.0.0.0   |      N/A       | 00:00:00:00:00:00 |
|      enp0s3     | 10.0.2.15  | 255.255.255.0 |   10.0.2.255   | 08:00:27:bd:9c:3e |
| br-3ede32bcf02b | 172.18.0.1 |  255.255.0.0  | 172.18.255.255 | 02:42:ca:27:eb:e0 |
|     docker0     | 172.17.0.1 |  255.255.0.0  | 172.17.255.255 | 02:42:a6:17:95:87 |
+-----------------+------------+---------------+----------------+-------------------+

Process Information:
+-----------------+--------------------------------------------------------------------------------+
|       User      |                                 Process Names                                  |
+-----------------+--------------------------------------------------------------------------------+
|      avahi      |                                  avahi-daemon                                  |
| aws-replication |                                      java                                      |
| aws-replication |                    run_linux_migration_scripts_periodically                    |
| aws-replication |                                  tail, tailer                                  |
| aws-replication |                             update_onprem_volumes                              |
|      colord     |                                     colord                                     |
|       gdm       |                                    (sd-pam)                                    |
|       gdm       |                                    Xwayland                                    |
|       gdm       |                     at-spi-bus-launcher, at-spi2-registryd                     |
|       gdm       |                                  dbus-daemon                                   |
|       gdm       |                              gdm-wayland-session                               |
|       gdm       |                       gnome-session-binary, gnome-shell                        |
|       gdm       |  gsd-a11y-settings, gsd-clipboard, gsd-color, gsd-datetime, gsd-housekeeping   |
|       gdm       |  gsd-keyboard, gsd-media-keys, gsd-mouse, gsd-power, gsd-print-notifications   |
|       gdm       |    gsd-rfkill, gsd-screensaver-proxy, gsd-sharing, gsd-smartcard, gsd-sound    |
...
|       root      |                                      ksmd                                      |
|       root      |               ksoftirqd/0, ksoftirqd/1, ksoftirqd/2, ksoftirqd/3               |
|       root      |                                     kstrp                                      |
|       root      |                                    kswapd0                                     |
|       root      |                               kthreadd, kthrotld                               |
|       root      | kworker/0:0, kworker/0:0H, kworker/0:1, kworker/0:1H, kworker/0:2, kworker/0:3 |
|       root      | kworker/1:0, kworker/1:0H, kworker/1:1, kworker/1:1H, kworker/1:2, kworker/1:3 |
|       root      | kworker/2:0, kworker/2:0H, kworker/2:1, kworker/2:1H, kworker/2:2, kworker/3:0 |
|       root      |       kworker/3:0H, kworker/3:1, kworker/3:1H, kworker/3:2, kworker/u8:0       |
|       root      |             kworker/u8:1, kworker/u8:3, kworker/u8:4, kworker/u9:0             |
|       root      |  loop0, loop1, loop10, loop11, loop12, loop13, loop14, loop15, loop16, loop17  |
|       root      | loop18, loop19, loop2, loop20, loop21, loop22, loop23, loop24, loop25, loop26  |
|       root      | loop27, loop28, loop29, loop3, loop30, loop31, loop32, loop33, loop34, loop35  |
|       root      |    loop36, loop37, loop38, loop39, loop4, loop5, loop6, loop7, loop8, loop9    |
|       root      |                                       md                                       |
|       root      |               migration/0, migration/1, migration/2, migration/3               |
|       root      |                                  mm_percpu_wq                                  |
|       root      |                                     netns                                      |
|       root      |                                networkd-dispat                                 |
|       root      |                                   oom_reaper                                   |
|       root      |                                  packagekitd                                   |
|       root      |                                    polkitd                                     |
|       root      |                       rcu_bh, rcu_sched, rcu_tasks_kthre                       |
|       root      |      scsi_eh_0, scsi_eh_1, scsi_eh_2, scsi_tmf_0, scsi_tmf_1, scsi_tmf_2       |
|       root      |                                     snapd                                      |
|       root      |                                      sshd                                      |
|       root      |                                      sudo                                      |
|       root      |            systemd, systemd-journald, systemd-logind, systemd-udevd            |
...
|       son       |       evolution-calendar-factory, evolution-calendar-factory-subprocess        |
|       son       |                           evolution-source-registry                            |
|       son       |                                 gdm-x-session                                  |
|       son       |            gnome-keyring-daemon, gnome-session-binary, gnome-shell             |
|       son       |       gnome-shell-calendar-server, gnome-software, gnome-terminal-server       |
|       son       |                        goa-daemon, goa-identity-service                        |
...
|       son       |                                    systemd                                     |
|       son       |                        update-manager, update-notifier                         |
|       son       |        xdg-desktop-portal, xdg-desktop-portal-gtk, xdg-document-portal         |
|       son       |                              xdg-permission-store                              |
|      syslog     |                                    rsyslogd                                    |
| systemd-resolve |                                systemd-resolved                                |
|     whoopsie    |                                    whoopsie                                    |
+-----------------+--------------------------------------------------------------------------------+
```
