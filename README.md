# Elgluth TCP workload

Program to make a workload test for your TCP-servers. ⭐pls
Linux|Termux|Mac OS

Windows is tend to scold when it sees the fact of spamming with TCP-requests

❗️ BE SURE TO AVOID USING IT ON WEAK SERVERS, THEY MAY SUFFER ❗️

<!--Installation-->
## Installation

0. Go Lang installation
   
 ```rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.1.linux-amd64.tar.gz```
 
 ```export PATH=$PATH:/usr/local/go/bin```

1. Clone repo

```cd /usr/local/go/src```

```git clone github.com/LLlE0/ELGLUTH_TCP```


2. Build and run

```go build .```

```./elgluth | ip:port / domen:port | amount of goroutines (200-1000 recommended)```

Example:

```./elgluth 128.128.128.128:55689 1000```

```./elgluth google.com:443 800```

##
<h2>Made by LLlE0, 11130 (c)</h2>
