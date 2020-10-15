# что необходимо собирать

- Средняя загрузка системы (load average).

**Linux**
```
uptime
 08:26:47 up 52 days,  4:28,  1 user,  load average: 0.04, 0.01, 0.00
```

**Darwin**
```
$ uptime
11:25  up 5 days, 12:27, 3 users, load averages: 2.19 2.65 2.99
```

- Средняя загрузка CPU (%user_mode, %system_mode, %idle)

**Linux**
```
# top -i -b -n1 | egrep '^%Cpu'
%Cpu(s):  0.2 us,  0.1 sy,  0.0 ni, 99.6 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st
```
**Darwin**
```
$ top -l 2 -n 0 | egrep '^CPU usage' | tail -n1
CPU usage: 9.0% user, 11.20% sys, 79.80% idle
```
- Загрузка дисков:

    - tps (transfers per second);
    - KB/s (kilobytes (read+write) per second);
    
    **Linux** `iostat -d -o JSON`
          
    ```json
  {"sysstat": {
        "hosts": [
            {
                "nodename": "web-coder.dev",
                "sysname": "Linux",
                "release": "4.15.0-112-generic",
                "machine": "x86_64",
                "number-of-cpus": 1,
                "date": "10/02/2020",
                "statistics": [
                    {
                        "disk": [
                            {"disk_device": "loop0", "tps": 0.00, "kB_read/s": 0.00, "kB_wrtn/s": 0.00, "kB_read": 1, "kB_wrtn": 0},
                            {"disk_device": "sda", "tps": 0.42, "kB_read/s": 1.49, "kB_wrtn/s": 11.54, "kB_read": 6723037, "kB_wrtn": 51952924},
                            {"disk_device": "sdb", "tps": 0.00, "kB_read/s": 0.00, "kB_wrtn/s": 0.01, "kB_read": 4224, "kB_wrtn": 27748}
                        ]
                    }
                ]
            }
        ]
      }}
  ```
    **Darwin** 
    ```
  $ iostat
                 disk0       cpu    load average
       KB/t  tps  MB/s  us sy id   1m   5m   15m
      49.11  122  5.85  10  9 80  1.97 2.89 3.12
  ```
  
    - CPU (%user_mode, %system_mode, %idle)
    
    **Linux** `iostat -c -o JSON`
    ```json
  {"sysstat": {
    	"hosts": [
    		{
    			"nodename": "web-coder.dev",
    			"sysname": "Linux",
    			"release": "4.15.0-112-generic",
    			"machine": "x86_64",
    			"number-of-cpus": 1,
    			"date": "10/02/2020",
    			"statistics": [
    				{
    					"avg-cpu":  {
                            "user": 0.25, 
                            "nice": 0.00, 
                            "system": 0.11, 
                            "iowait": 0.00, 
                            "steal": 0.00, 
                            "idle": 99.63
                         }
    				}
    			]
    		}
    	]
    }}
  ```

- Информация о дисках по каждой файловой системе:

    - использовано мегабайт, % от доступного количества;

**Linux** 
```
  df -BM
       Filesystem     1M-blocks   Used Available Use% Mounted on
       udev                462M     0M      462M   0% /dev
       tmpfs                99M     7M       93M   7% /run
       /dev/sda          22583M 11555M     9863M  54% /
       tmpfs               493M     0M      493M   0% /dev/shm
       tmpfs                 5M     0M        5M   0% /run/lock
       tmpfs               493M     0M      493M   0% /sys/fs/cgroup
       tmpfs                99M     0M       99M   0% /run/user/0
  ```
**Darwin**
```
df -m
Filesystem    1M-blocks   Used Available Capacity iused      ifree %iused  Mounted on
/dev/disk1s5     429118  10729    256070     5%  488318 4393683282    0%   /
devfs                 0      0         0   100%     674          0  100%   /dev
/dev/disk1s1     429118 152369    256070    38% 2432282 4391739318    0%   /System/Volumes/Data
/dev/disk1s4     429118   9217    256070     4%      10 4394171590    0%   /private/var/vm
/dev/disk0s3      47555  44989      2566    95%  578752 4294388527    0%   /Volumes/backups
map auto_home         0      0         0   100%       0          0  100%   /System/Volumes/Data/home
```
    - использовано inode, % от доступного количества.
**Linux**
```
df -i
Filesystem      Inodes  IUsed   IFree IUse% Mounted on
udev            118182    382  117800    1% /dev
tmpfs           126088   1842  124246    2% /run
/dev/sda       1440000 396831 1043169   28% /
tmpfs           126088      1  126087    1% /dev/shm
tmpfs           126088      3  126085    1% /run/lock
tmpfs           126088     18  126070    1% /sys/fs/cgroup
tmpfs           126088     11  126077    1% /run/user/0
```

- Top talkers по сети:

    - по протоколам: protocol (TCP, UDP, ICMP, etc), bytes, % от sum(bytes) за последние M), сортируем по убыванию процента;
**Linux**
```
cat /proc/net/dev
Inter-|   Receive                                                |  Transmit
 face |bytes    packets errs drop fifo frame compressed multicast|bytes    packets errs drop fifo colls carrier compressed
br-c48cbfa55133:       0       0    0    0    0     0          0         0        0       0    0    0    0     0       0          0
br-42e2a7b3bd01:       0       0    0    0    0     0          0         0        0       0    0    0    0     0       0          0
br-8ed58366592d:       0       0    0    0    0     0          0         0        0       0    0    0    0     0       0          0
docker0:       0       0    0    0    0     0          0         0        0       0    0    0    0     0       0          0
  eth0: 1656622526 5501805    0    1    0     0          0         0 1740053148 5674047    0    0    0     0       0          0
    lo: 1457381   12488    0    0    0     0          0         0  1457381   12488    0    0    0     0       0          0
```
    - по трафику: source ip:port, destination ip:port, protocol, bytes per second (bps), сортируем по убыванию bps.

- Статистика по сетевым соединениям:

    - слушающие TCP & UDP сокеты: command, pid, user, protocol, port;
    - количество TCP соединений, находящихся в разных состояниях (ESTAB, FIN_WAIT, SYN_RCV и пр.).
