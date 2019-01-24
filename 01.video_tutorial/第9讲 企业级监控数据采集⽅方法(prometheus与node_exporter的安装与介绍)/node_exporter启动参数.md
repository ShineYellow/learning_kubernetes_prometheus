
## Enabled by default
-----------------------

Name Description OS

arp

Exposes ARP statistics from /proc/net/arp.

Linux bcache

Exposes bcache statistics from /sys/fs/bcache/.

Linux conntrack

Shows conntrack statistics (does nothing if no /proc/ sys/net/netfilter/ present).

Linux

cpu

Exposes CPU statistics

Darwin, Dragonfly, FreeBSD, Linux diskstats

Exposes disk I/O statistics.

Darwin, Linux edac

Exposes error detection and correction statistics.

Linux entropy

Exposes available entropy.

Linux exec

Exposes execution statistics.

Dragonfly, FreeBSD filefd

Exposes file descriptor statistics from /proc/sys/fs/ file-nr.

Linux filesystem

Exposes filesystem statistics, such as disk space used.

Darwin, Dragonfly, FreeBSD, Linux, OpenBSD hwmon

Expose hardware monitoring and sensor data from /sys/ class/hwmon/.

Linux

infiniband

Exposes network statistics specific to InfiniBand and Intel OmniPath configurations.

Linux ipvs

Exposes IPVS status from /proc/net/ip_vs and stats from /proc/net/ip_vs_stats.

Linux

loadavg

Exposes load average.

Darwin, Dragonfly, FreeBSD, Linux, NetBSD, OpenBSD,

Solaris mdadm

Exposes statistics about devices in /proc/mdstat (does nothing if no /proc/mdstat present).

Linux

meminfo

Exposes memory statistics.

Darwin, Dragonfly, FreeBSD, Linux, OpenBSD netdev

Exposes network interface statistics such as bytes transferred.

Darwin, Dragonfly, FreeBSD, Linux, OpenBSD netstat

Exposes network statistics from /proc/net/netstat. This is the same information as netstat -s.

Linux

nfs

Exposes NFS client statistics from /proc/net/rpc/nfs. This is the same information as nfsstat -c.

Linux

nfsd

Exposes NFS kernel server statistics from /proc/net/ rpc/nfsd. This is the same information as nfsstat -s.

Linux

sockstat

Exposes various statistics from /proc/net/sockstat.

Linux stat

Exposes various statistics from /proc/stat. This includes boot time, forks and interrupts.

Linux textfile

Exposes statistics read from local disk. The -- collector.textfile.directory flag must be set. any

time

Exposes the current system time.

any

timex

Exposes selected adjtimex(2) system call stats.

Linux uname

Exposes system information as provided by the uname system call.

Linux

vmstat

Exposes statistics from /proc/vmstat.

Linux wifi

Exposes WiFi device and station statistics.

Linux xfs

Exposes XFS runtime statistics. Linux (kernel 4.4+)

zfs

Exposes ZFS performance statistics. Linux

Disabled by default


Name Description OS

bonding

Exposes the number of configured and active slaves of Linux bonding interfaces.

Linux buddyinfo

Exposes statistics of memory fragments as reported by / proc/buddyinfo.

Linux devstat

Exposes device statistics Dragonfly, FreeBSD

drbd

Exposes Distributed Replicated Block Device statistics (to version 8.4)

Linux interrupts

Exposes detailed interrupts statistics.

Linux, OpenBSD ksmd

Exposes kernel and system statistics from /sys/kernel/ mm/ksm.

Linux

logind

Exposes session counts from logind.

Linux meminfo_numa

Exposes memory statistics from /proc/meminfo_numa.

Linux mountstats

Exposes filesystem statistics from /proc/self/ mountstats. Exposes detailed NFS client statistics.

Linux ntp

Exposes local NTP daemon health to check time

any

qdisc

Exposes queuing discipline statistics Linux

runit

Exposes service status from runit.

any

supervisord

Exposes service status from supervisord.

any

systemd

Exposes service and system status from systemd.

Linux tcpstat

Exposes TCP connection status information from /proc/ net/tcp and /proc/net/tcp6. (Warning: the current version has potential performance issues in high load situations.)


## 另外 这些项⽬ 可以通过 node_exporter的⾃定义 启动参数 来 实现 开启
------------------------------------------------------------

./node_exporter —help


--collector.arp Enable the arp collector (default: enabled).

--collector.bcache Enable the bcache collector (default: enabled).

--collector.bonding Enable the bonding collector (default: disabled).

--collector.buddyinfo Enable the buddyinfo collector (default: disabled).

--collector.conntrack Enable the conntrack collector (default: enabled).

--collector.cpu Enable the cpu collector (default: enabled).

--collector.diskstats Enable the diskstats collector (default: enabled).

--collector.drbd Enable the drbd collector (default: disabled).

--collector.edac Enable the edac collector (default: enabled).

--collector.entropy Enable the entropy collector (default: enabled).

--collector.filefd Enable the filefd collector (default: enabled).

--collector.filesystem Enable the filesystem collector (default: enabled).

--collector.gmond Enable the gmond collector (default: disabled).

--collector.hwmon Enable the hwmon collector (default: enabled).

--collector.infiniband Enable the infiniband collector (default: enabled).

--collector.interrupts Enable the interrupts collector (default: disabled).

--collector.ipvs Enable the ipvs collector (default: enabled).

--collector.ksmd Enable the ksmd collector (default: disabled).

--collector.loadavg Enable the loadavg collector (default: enabled).

--collector.logind Enable the logind collector (default: disabled).

--collector.mdadm Enable the mdadm collector (default: enabled).
