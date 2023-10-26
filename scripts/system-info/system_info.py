import os
import platform
import socket
import subprocess
from collections import defaultdict
from datetime import datetime

import distro
import netifaces
import psutil
import yaml  # Required for YAML file operations
from prettytable import PrettyTable

# Optional packages for GPU and Graphic Card information
try:
    import GPUtil

    GPU_AVAILABLE = True
except ImportError:
    GPU_AVAILABLE = False


def bytes_to_gb(bytes_value):
    """Convert bytes to gigabytes."""
    return round(bytes_value / (1024**3), 2)


def get_os_info():
    """Fetch Operating System information."""
    boot_time = datetime.fromtimestamp(psutil.boot_time()).strftime(
        "%Y-%m-%d %H:%M:%S"
    )
    try:
        logged_user = os.getlogin()
    except OSError:
        logged_user = os.getenv("USER") or os.getenv("USERNAME")
    uname_info = platform.uname()

    try:
        dist = platform.linux_distribution()
    except AttributeError:
        name = distro.name()  # e.g.: 'Ubuntu'
        version = distro.version()  # e.g.: '20.04'
        codename = distro.codename()  # e.g.: 'focal'
        dist = f"{name} {version} {codename.capitalize()}"

    return "OS", [
        ("Name", platform.system()),
        ("Version", platform.version()),
        ("Architecture", platform.architecture()),
        ("Platform", platform.platform()),
        ("Boot Time", boot_time),
        ("Logged User", logged_user),
        ("Kernel Version", uname_info.release),
        ("Distribution", dist),
    ]


def get_cpu_info():
    """Fetch CPU information."""
    cpu_info = psutil.cpu_freq()
    cpu_stats = psutil.cpu_stats()
    cpu_times_percent = psutil.cpu_times_percent()
    return "CPU", [
        ("Brand", platform.processor()),
        ("Core Count (Logical)", psutil.cpu_count(logical=True)),
        ("Core Count (Physical)", psutil.cpu_count(logical=False)),
        ("Current Frequency", f"{cpu_info.current} MHz"),
        ("Context Switches", cpu_stats.ctx_switches),
        ("Interrupts", cpu_stats.interrupts),
        ("Architecture", platform.machine()),
        ("User Time (%)", cpu_times_percent.user),
        ("System Time (%)", cpu_times_percent.system),
        ("Idle Time (%)", cpu_times_percent.idle),
    ]


def get_memory_info():
    """Fetch Memory information."""
    virtual_memory = psutil.virtual_memory()
    swap_memory = psutil.swap_memory()
    return "Memory", [
        ("Total (GB)", bytes_to_gb(virtual_memory.total)),
        ("Available (GB)", bytes_to_gb(virtual_memory.available)),
        ("Used (GB)", bytes_to_gb(virtual_memory.used)),
        ("Cached (GB)", bytes_to_gb(virtual_memory.cached)),
        ("Swap Total (GB)", bytes_to_gb(swap_memory.total)),
        ("Swap Free (GB)", bytes_to_gb(swap_memory.free)),
        ("Swap Used (GB)", bytes_to_gb(swap_memory.used)),
    ]


def get_disk_info():
    """Fetch Disk information."""
    disk_info = psutil.disk_usage("/")
    disk_io = psutil.disk_io_counters()
    # partitions = psutil.disk_partitions()
    # partition_info = ", ".join([p.device for p in partitions])
    return "Disk", [
        ("Total (GB)", bytes_to_gb(disk_info.total)),
        ("Free (GB)", bytes_to_gb(disk_info.free)),
        ("Used (GB)", bytes_to_gb(disk_info.used)),
        ("Read Count", disk_io.read_count),
        ("Write Count", disk_io.write_count)
        # ("Partitions", partition_info),
        # ("File System Types", ', '.join([p.fstype for p in partitions]))
    ]


def get_network_info():
    """Fetch Network information."""
    net_info = psutil.net_if_addrs()
    net_stats = psutil.net_if_stats()
    active_interfaces = [k for k, v in net_stats.items() if v.isup]
    mac_address = "Unknown"
    for interface, addrs in net_info.items():
        for addr in addrs:
            if addr.family == psutil.AF_LINK:
                mac_address = addr.address
                break
    return "Network", [
        ("Hostname", socket.gethostname()),
        ("IP Address", socket.gethostbyname(socket.gethostname())),
        ("MAC Address", mac_address),
        ("Active Interfaces", ", ".join(active_interfaces)),
        ("Gateway", socket.gethostbyname(socket.gethostname())),
        ("DNS", socket.gethostbyname(socket.getfqdn())),
    ]


# IP and MAC addresses using netifaces
def get_advanced_network_info():
    """Fetch advanced network information using netifaces."""
    # Add a function using netifaces
    interfaces = netifaces.interfaces()
    iface_info = []
    for iface in interfaces:
        addresses = netifaces.ifaddresses(iface)
        if netifaces.AF_INET in addresses:
            ip_info = addresses[netifaces.AF_INET][0]
            ip = ip_info.get("addr", "N/A")
            netmask = ip_info.get("netmask", "N/A")
            broadcast = ip_info.get("broadcast", "N/A")
            if netifaces.AF_LINK in addresses:
                mac_info = addresses[netifaces.AF_LINK][0]
                mac = mac_info.get("addr", "N/A")
                iface_info.append((iface, ip, netmask, broadcast, mac))
            else:
                iface_info.append((iface, ip, netmask, broadcast, "N/A"))
    return "Advanced Network", iface_info


def get_gpu_info():
    """Fetch GPU information if available."""
    info_list = []
    if GPU_AVAILABLE:
        gpus = GPUtil.getGPUs()
        for gpu in gpus:
            info_list.append(
                (f"{gpu.name}", f"Driver Version: {gpu.driver_version}")
            )
    else:
        info_list.append(("Status", "Not Available"))
    return "GPU", info_list


def get_graphic_card_info():
    try:
        lspci_output = subprocess.check_output(
            "lspci | grep VGA", shell=True, text=True
        )
        info_list = [
            ("VGA Compatible Controller", line.split(":")[-1].strip())
            for line in lspci_output.split("\n")
            if line
        ]
    except Exception:
        info_list = [("Status", "Not Available")]

    return "Graphic Card", info_list


def get_process_info():
    """Fetch Process information."""
    process_dict = defaultdict(list)

    for proc in psutil.process_iter(["name", "username"]):
        process_dict[proc.info["username"]].append(proc.info["name"])

    # Sorting and grouping by username
    sorted_process = defaultdict(list)
    for username, names in sorted(process_dict.items()):
        sorted_names = sorted(set(names))
        sorted_process[username].extend(sorted_names)

    # Group by prefix and limit to 80 characters per line
    grouped_process = defaultdict(list)
    for username, processes in sorted_process.items():
        temp = defaultdict(list)
        for proc in processes:
            prefix = proc[:4]
            temp[prefix].append(proc)

        for prefix, names in temp.items():
            line = ""
            for name in names:
                if len(line + name + ", ") > 80:
                    grouped_process[username].append(line[:-2])
                    line = ""
                line += name + ", "

            if line:
                grouped_process[username].append(line[:-2])

    # Convert to list of tuples
    process_list = []
    for username, processes in grouped_process.items():
        for proc in processes:
            process_list.append((username, proc))

    return "Process", sorted(process_list)


# New function to save system information to a YAML file
def save_to_yaml(data):
    """Save dictionary data to YAML file."""
    with open("system_info.yaml", "w") as file:
        yaml.dump(data, file, default_flow_style=False)


def main():
    """
    Main function to collect and print all system information
    and save it to YAML.
    """

    system_table = PrettyTable()
    system_table.field_names = ["Index", "Group", "Category", "Info"]

    index = 1
    system_info = {}  # Dictionary to store all system info

    info_functions = [
        get_os_info,
        get_cpu_info,
        get_memory_info,
        get_disk_info,
        get_network_info,
        get_gpu_info,
        get_graphic_card_info,
    ]

    for func in info_functions:
        group, info_func = func()
        system_info[group] = {category: info for category, info in info_func}
        for category, info in info_func:
            system_table.add_row([index, group, category, info])
            index += 1

    print("System Information:")
    print(system_table)

    # Fetch advanced network information using netifaces
    advanced_network_table = PrettyTable()
    advanced_network_table.field_names = [
        "Interface",
        "IP Address",
        "Netmask",
        "Broadcast",
        "MAC Address",
    ]

    _, advanced_network_info = get_advanced_network_info()
    advanced_network_data = []
    for row in advanced_network_info:
        advanced_network_table.add_row(row)
        advanced_network_data.append(
            {
                "Interface": row[0],
                "IP Address": row[1],
                "Netmask": row[2],
                "Broadcast": row[3],
                "MAC Address": row[4],
            }
        )

    print("\nAdvanced Network Information:")
    print(advanced_network_table)
    system_info[
        "Advanced Network"
    ] = advanced_network_data  # Add to system info dictionary

    # Fetch process information
    process_table = PrettyTable()
    process_table.field_names = ["User", "Process Names"]

    _, process_info = get_process_info()
    process_data = []
    for username, grouped_name in process_info:
        process_table.add_row([username, grouped_name])
        process_data.append({username: grouped_name})

    print("\nProcess Information:")
    print(process_table)
    system_info["Process"] = process_data  # Add to system info dictionary

    save_to_yaml(system_info)  # Save all gathered info to a YAML file


if __name__ == "__main__":
    main()
