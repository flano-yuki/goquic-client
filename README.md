# goquic-client
goquic client cli

## sample
```
$ ./goquic-client https://www.google.co.jp
 HTTP/2.0 200 200
Connection: close
:status: 200
accept-ranges: none
alt-svc: quic=":443"; ma=2592000; v="34,33,32,31,30,29,28,27,26,25"
alternate-protocol: 443:quic
cache-control: private, max-age=0
content-type: text/html; charset=Shift_JIS
date: Fri, 01 Jul 2016 18:07:27 GMT
```

## option

```
$ ./goquic-client -h header1=hoge -h header2=fuga -q https://www.google.co.jp
```

```
  --liblog, -l: libquic Log level (defaults to: 2)
   --sring, -s: POST body string
 --headers, -h: HTTP Request headers (-h key=value)
    --body, -b: Show body
   --quiet, -q: no output
```
