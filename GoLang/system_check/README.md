# SystemCheck ([REMOVED])

---
This is a proof of concept of integrating System Check functions (as requested...) into [REMOVED]

<b>Note:</b> This version will leave out any company information and/or allusions to it.  With that being said, the code also has to be modified some to fulfill said criteria

---
## Getting Started

---
The compressed file should contain the GO code, as well as teh information for the Docker containers used.  Ideally, you'll just have to spin up the docker container, enter the container, navigate to the [REMOVED] directory, and execute the code

---
## Prerequisites

---
You may or may not need to match the software versions, but for your safe knowledge, these are what I am currently running:
```
$ docker --version
Docker version 27.5.1, build 9f9e405

docker-compose --version
docker-compose version 1.29.2, build unknown
```
---
## Setting Up

---
<b><u>Spinning up Docker Container</u></b>
- From the terminal, navigate to the corresponding "system_check" directory
- Bring up the container with the command `docker-compose up -d [--build]`
  - note: build may or may not be necessary

<b><u>Entering the Container</u></b>
- From the terminal, you can enter a shell session within the container with the command `docker exec -it $CONTAINER_NAME /bin/bash` substituting $CONTAINER_NAME with the name of the container.
  - If you kept everything default, the name of the container will be "ubuntu_22".  This is determined by the "container_name" field in the docker-compose.yml
---
## Running Tests

---
After the container has been spun up and you've entered the container, you'll want to navigate to the proper directory: **/home/golang/system_check**

Execute the GO code by running the command `go run main.go`...

```
root@508aecd34cd4:/home/golang/system_check# go run main.go 
------------------
IO STATS
------------------
Linux 6.8.0-52-generic (508aecd34cd4)   02/08/25        _x86_64_        (2 CPU)

avg-cpu:  %user   %nice %system %iowait  %steal   %idle
           8.11    0.05    3.26    0.15    0.00   88.43

Device             tps    kB_read/s    kB_wrtn/s    kB_dscd/s    kB_read    kB_wrtn    kB_dscd
loop0             0.00         0.00         0.00         0.00         17          0          0
loop1             0.10         1.81         0.00         0.00      10918          0          0
loop10            0.01         0.12         0.00         0.00        706          0          0
loop11            0.13         2.29         0.00         0.00      13825          0          0
loop12            0.88        17.58         0.00         0.00     106136          0          0
loop13            0.31        13.04         0.00         0.00      78742          0          0
loop14            0.00         0.00         0.00         0.00         10          0          0
loop2             0.01         0.19         0.00         0.00       1132          0          0
loop3             1.08         5.06         0.00         0.00      30538          0          0
loop4             0.56         9.95         0.00         0.00      60042          0          0
loop5             4.00       100.62         0.00         0.00     607402          0          0
loop6             0.10         1.82         0.00         0.00      11013          0          0
loop7             0.01         0.07         0.00         0.00        441          0          0
loop8             0.01         0.05         0.00         0.00        324          0          0
loop9             0.26        13.15         0.00         0.00      79401          0          0
sda              59.26      1203.03      1392.19         0.00    7262328    8404197          0
sr0               0.05         1.54         0.00         0.00       9288          0          0


avg-cpu:  %user   %nice %system %iowait  %steal   %idle
           7.04    0.00    2.51    0.00    0.00   90.45

Device             tps    kB_read/s    kB_wrtn/s    kB_dscd/s    kB_read    kB_wrtn    kB_dscd
loop0             0.00         0.00         0.00         0.00          0          0          0
loop1             0.00         0.00         0.00         0.00          0          0          0
loop10            0.00         0.00         0.00         0.00          0          0          0
loop11            0.00         0.00         0.00         0.00          0          0          0
loop12            0.00         0.00         0.00         0.00          0          0          0
loop13            0.00         0.00         0.00         0.00          0          0          0
loop14            0.00         0.00         0.00         0.00          0          0          0
loop2             0.00         0.00         0.00         0.00          0          0          0
loop3             0.00         0.00         0.00         0.00          0          0          0
loop4             0.00         0.00         0.00         0.00          0          0          0
loop5             0.00         0.00         0.00         0.00          0          0          0
loop6             0.00         0.00         0.00         0.00          0          0          0
loop7             0.00         0.00         0.00         0.00          0          0          0
loop8             0.00         0.00         0.00         0.00          0          0          0
loop9             0.00         0.00         0.00         0.00          0          0          0
sda              14.00        76.00         8.00         0.00         76          8          0
sr0               0.00         0.00         0.00         0.00          0          0          0


avg-cpu:  %user   %nice %system %iowait  %steal   %idle
           4.08    0.00    1.53    0.00    0.00   94.39

Device             tps    kB_read/s    kB_wrtn/s    kB_dscd/s    kB_read    kB_wrtn    kB_dscd
loop0             0.00         0.00         0.00         0.00          0          0          0
loop1             0.00         0.00         0.00         0.00          0          0          0
loop10            0.00         0.00         0.00         0.00          0          0          0
loop11            0.00         0.00         0.00         0.00          0          0          0
loop12            0.00         0.00         0.00         0.00          0          0          0
loop13            0.00         0.00         0.00         0.00          0          0          0
loop14            0.00         0.00         0.00         0.00          0          0          0
loop2             0.00         0.00         0.00         0.00          0          0          0
loop3             0.00         0.00         0.00         0.00          0          0          0
loop4             0.00         0.00         0.00         0.00          0          0          0
loop5             0.00         0.00         0.00         0.00          0          0          0
loop6             0.00         0.00         0.00         0.00          0          0          0
loop7             0.00         0.00         0.00         0.00          0          0          0
loop8             0.00         0.00         0.00         0.00          0          0          0
loop9             0.00         0.00         0.00         0.00          0          0          0
sda             153.00         0.00       776.00         0.00          0        776          0
sr0               0.00         0.00         0.00         0.00          0          0          0


avg-cpu:  %user   %nice %system %iowait  %steal   %idle
           3.03    0.00    1.52    0.00    0.00   95.45

Device             tps    kB_read/s    kB_wrtn/s    kB_dscd/s    kB_read    kB_wrtn    kB_dscd
loop0             0.00         0.00         0.00         0.00          0          0          0
loop1             0.00         0.00         0.00         0.00          0          0          0
loop10            0.00         0.00         0.00         0.00          0          0          0
loop11            0.00         0.00         0.00         0.00          0          0          0
loop12            0.00         0.00         0.00         0.00          0          0          0
loop13            0.00         0.00         0.00         0.00          0          0          0
loop14            0.00         0.00         0.00         0.00          0          0          0
loop2             0.00         0.00         0.00         0.00          0          0          0
loop3             0.00         0.00         0.00         0.00          0          0          0
loop4             0.00         0.00         0.00         0.00          0          0          0
loop5             0.00         0.00         0.00         0.00          0          0          0
loop6             0.00         0.00         0.00         0.00          0          0          0
loop7             0.00         0.00         0.00         0.00          0          0          0
loop8             0.00         0.00         0.00         0.00          0          0          0
loop9             0.00         0.00         0.00         0.00          0          0          0
sda               0.00         0.00         0.00         0.00          0          0          0
sr0               0.00         0.00         0.00         0.00          0          0          0


avg-cpu:  %user   %nice %system %iowait  %steal   %idle
           2.53    0.00    2.02    0.00    0.00   95.45

Device             tps    kB_read/s    kB_wrtn/s    kB_dscd/s    kB_read    kB_wrtn    kB_dscd
loop0             0.00         0.00         0.00         0.00          0          0          0
loop1             0.00         0.00         0.00         0.00          0          0          0
loop10            0.00         0.00         0.00         0.00          0          0          0
loop11            0.00         0.00         0.00         0.00          0          0          0
loop12            0.00         0.00         0.00         0.00          0          0          0
loop13            0.00         0.00         0.00         0.00          0          0          0
loop14            0.00         0.00         0.00         0.00          0          0          0
loop2             0.00         0.00         0.00         0.00          0          0          0
loop3             0.00         0.00         0.00         0.00          0          0          0
loop4             0.00         0.00         0.00         0.00          0          0          0
loop5             0.00         0.00         0.00         0.00          0          0          0
loop6             0.00         0.00         0.00         0.00          0          0          0
loop7             0.00         0.00         0.00         0.00          0          0          0
loop8             0.00         0.00         0.00         0.00          0          0          0
loop9             0.00         0.00         0.00         0.00          0          0          0
sda             110.00         0.00      5824.00         0.00          0       5824          0
sr0               0.00         0.00         0.00         0.00          0          0          0


avg-cpu:  %user   %nice %system %iowait  %steal   %idle
           3.54    0.00    2.02    0.00    0.00   94.44

Device             tps    kB_read/s    kB_wrtn/s    kB_dscd/s    kB_read    kB_wrtn    kB_dscd
loop0             0.00         0.00         0.00         0.00          0          0          0
loop1             0.00         0.00         0.00         0.00          0          0          0
loop10            0.00         0.00         0.00         0.00          0          0          0
loop11            0.00         0.00         0.00         0.00          0          0          0
loop12            0.00         0.00         0.00         0.00          0          0          0
loop13            0.00         0.00         0.00         0.00          0          0          0
loop14            0.00         0.00         0.00         0.00          0          0          0
loop2             0.00         0.00         0.00         0.00          0          0          0
loop3             0.00         0.00         0.00         0.00          0          0          0
loop4             0.00         0.00         0.00         0.00          0          0          0
loop5             0.00         0.00         0.00         0.00          0          0          0
loop6             0.00         0.00         0.00         0.00          0          0          0
loop7             0.00         0.00         0.00         0.00          0          0          0
loop8             0.00         0.00         0.00         0.00          0          0          0
loop9             0.00         0.00         0.00         0.00          0          0          0
sda              13.00         4.00       660.00         0.00          4        660          0
sr0               0.00         0.00         0.00         0.00          0          0          0


avg-cpu:  %user   %nice %system %iowait  %steal   %idle
           3.03    0.00    1.01    0.00    0.00   95.96

Device             tps    kB_read/s    kB_wrtn/s    kB_dscd/s    kB_read    kB_wrtn    kB_dscd
loop0             0.00         0.00         0.00         0.00          0          0          0
loop1             0.00         0.00         0.00         0.00          0          0          0
loop10            0.00         0.00         0.00         0.00          0          0          0
loop11            0.00         0.00         0.00         0.00          0          0          0
loop12            0.00         0.00         0.00         0.00          0          0          0
loop13            0.00         0.00         0.00         0.00          0          0          0
loop14            0.00         0.00         0.00         0.00          0          0          0
loop2             0.00         0.00         0.00         0.00          0          0          0
loop3             0.00         0.00         0.00         0.00          0          0          0
loop4             0.00         0.00         0.00         0.00          0          0          0
loop5             0.00         0.00         0.00         0.00          0          0          0
loop6             0.00         0.00         0.00         0.00          0          0          0
loop7             0.00         0.00         0.00         0.00          0          0          0
loop8             0.00         0.00         0.00         0.00          0          0          0
loop9             0.00         0.00         0.00         0.00          0          0          0
sda               0.00         0.00         0.00         0.00          0          0          0
sr0               0.00         0.00         0.00         0.00          0          0          0


avg-cpu:  %user   %nice %system %iowait  %steal   %idle
           3.05    0.00    2.03    0.00    0.00   94.92

Device             tps    kB_read/s    kB_wrtn/s    kB_dscd/s    kB_read    kB_wrtn    kB_dscd
loop0             0.00         0.00         0.00         0.00          0          0          0
loop1             0.00         0.00         0.00         0.00          0          0          0
loop10            0.00         0.00         0.00         0.00          0          0          0
loop11            0.00         0.00         0.00         0.00          0          0          0
loop12            0.00         0.00         0.00         0.00          0          0          0
loop13            0.00         0.00         0.00         0.00          0          0          0
loop14            0.00         0.00         0.00         0.00          0          0          0
loop2             0.00         0.00         0.00         0.00          0          0          0
loop3             0.00         0.00         0.00         0.00          0          0          0
loop4             0.00         0.00         0.00         0.00          0          0          0
loop5             0.00         0.00         0.00         0.00          0          0          0
loop6             0.00         0.00         0.00         0.00          0          0          0
loop7             0.00         0.00         0.00         0.00          0          0          0
loop8             0.00         0.00         0.00         0.00          0          0          0
loop9             0.00         0.00         0.00         0.00          0          0          0
sda               0.00         0.00         0.00         0.00          0          0          0
sr0               0.00         0.00         0.00         0.00          0          0          0


avg-cpu:  %user   %nice %system %iowait  %steal   %idle
          12.18    0.00    4.06    0.00    0.00   83.76

Device             tps    kB_read/s    kB_wrtn/s    kB_dscd/s    kB_read    kB_wrtn    kB_dscd
loop0             0.00         0.00         0.00         0.00          0          0          0
loop1             0.00         0.00         0.00         0.00          0          0          0
loop10            0.00         0.00         0.00         0.00          0          0          0
loop11            0.00         0.00         0.00         0.00          0          0          0
loop12            0.00         0.00         0.00         0.00          0          0          0
loop13            0.00         0.00         0.00         0.00          0          0          0
loop14            0.00         0.00         0.00         0.00          0          0          0
loop2             0.00         0.00         0.00         0.00          0          0          0
loop3             0.00         0.00         0.00         0.00          0          0          0
loop4             0.00         0.00         0.00         0.00          0          0          0
loop5             0.00         0.00         0.00         0.00          0          0          0
loop6             0.00         0.00         0.00         0.00          0          0          0
loop7             0.00         0.00         0.00         0.00          0          0          0
loop8             0.00         0.00         0.00         0.00          0          0          0
loop9             0.00         0.00         0.00         0.00          0          0          0
sda               6.00         0.00      1644.00         0.00          0       1644          0
sr0               0.00         0.00         0.00         0.00          0          0          0


avg-cpu:  %user   %nice %system %iowait  %steal   %idle
          42.13    0.00   11.17    0.00    0.00   46.70

Device             tps    kB_read/s    kB_wrtn/s    kB_dscd/s    kB_read    kB_wrtn    kB_dscd
loop0             0.00         0.00         0.00         0.00          0          0          0
loop1             0.00         0.00         0.00         0.00          0          0          0
loop10            0.00         0.00         0.00         0.00          0          0          0
loop11            0.00         0.00         0.00         0.00          0          0          0
loop12            0.00         0.00         0.00         0.00          0          0          0
loop13            0.00         0.00         0.00         0.00          0          0          0
loop14            0.00         0.00         0.00         0.00          0          0          0
loop2             0.00         0.00         0.00         0.00          0          0          0
loop3             0.00         0.00         0.00         0.00          0          0          0
loop4             0.00         0.00         0.00         0.00          0          0          0
loop5             0.00         0.00         0.00         0.00          0          0          0
loop6             0.00         0.00         0.00         0.00          0          0          0
loop7             0.00         0.00         0.00         0.00          0          0          0
loop8             0.00         0.00         0.00         0.00          0          0          0
loop9             0.00         0.00         0.00         0.00          0          0          0
sda              17.00         8.00      7804.00         0.00          8       7804          0
sr0               0.00         0.00         0.00         0.00          0          0          0



------------------
PROCESSES STUCK IN 'D' STATE
------------------
-
-
-
-
-
-
-
-
-
-
-
-
-
-
-
-
-
-
-
-
-
-
-
-
-
-
-
-
-
-
------------------
CPU USAGE
------------------
top - 18:12:42 up  1:40,  0 users,  load average: 0.35, 0.53, 0.49
Tasks:   5 total,   1 running,   4 sleeping,   0 stopped,   0 zombie
%Cpu(s): 46.7 us, 13.3 sy,  0.0 ni, 36.7 id,  0.0 wa,  0.0 hi,  3.3 si,  0.0 st
MiB Mem :   3868.6 total,    160.7 free,   2445.5 used,   1262.4 buff/cache
MiB Swap:   2140.0 total,   1792.2 free,    347.8 used.   1128.6 avail Mem 

    PID USER      PR  NI    VIRT    RES    SHR S  %CPU  %MEM     TIME+ COMMAND
      1 root      20   0    4628   3584   3072 S   0.0   0.1   0:00.01 bash
      9 root      20   0    4628   3712   3200 S   0.0   0.1   0:00.01 bash
    480 root      20   0 1237420  17332   9600 S   0.0   0.4   0:00.03 go
    515 root      20   0 1225856   2560   1408 S   0.0   0.1   0:00.00 main
    550 root      20   0    7184   3200   2816 R   0.0   0.1   0:00.00 top

------------------
IP INFORMATION
------------------
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
25: eth0@if26: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:ac:12:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0


ERROR: File '/etc/systemd/network/05-eth0.network' not found!
```
---
## Next Steps

---

---
## Misc

---

---
