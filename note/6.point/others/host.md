# host

在某些情况下，比如vpn不给力等，可能会遇到对于某些网站的dns解析非常慢。

在这种情况下我们可以通过 `dig` 命令得到需要访问的目标域名的ip， 然后直接写入本地的映射表中，跳过dns这一步，从而提升网络体验

```shell
# dig @'dns服务器的地址' '所需要访问的域名'
>>> $ dig @10.32.0.2 github.com
; <<>> DiG 9.10.6 <<>> @10.32.0.2 github.com
; (1 server found)
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 31521
;; flags: qr rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 4096
;; QUESTION SECTION:
;github.com.            IN    A

;; ANSWER SECTION:
github.com.        1    IN    A    140.82.113.3

;; Query time: 267 msec
;; SERVER: 10.32.0.2#53(10.32.0.2)
;; WHEN: Thu Feb 29 16:15:09 CST 2024
;; MSG SIZE  rcvd: 55
```

这里的 `ANSWER SECTION` 就提供了目标域名的ip。
