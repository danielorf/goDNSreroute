# goDNSreroute

goDNSreroute is a tool for developers to capture and re-route DNS queries for network testing

#### Background
The concept of a redirecting DNS server for the purpose of testing is derived from the [dns-reroute](https://github.com/nanoscopic/dns-reroute) project by [nanoscopic](https://github.com/nanoscopic).
The DNS server portion of this project borrows heavily from code located [here](https://jameshfisher.com/2017/08/04/golang-dns-server.html).


#### Introduction
goDNSreroute captures and re-routes DNS queries for specified domains.  For domains not specified, goDNSreroute hands off the query to a typical DNS server (8.8.8.8 by default).  This allows the user to re-route DNS queries for network testing while not affecting network connectivity otherwise.


#### Usage
At the top of main.go are three variables used to configure goDNSreroute server:
- dnsPort: Specifies the port which goDNSreroute listens.  DNS servers typically listen on port 53.  Considering this is well-within the reserved port range (0-1024), it is often helpful to be able to run this server on an unreserved port (>1024).
- dnsLookupServer: Specifies the upstream DNS server for domians not listed in domainsToAddresses.  This allows DNS quries not listed in domainsToAddresses to be resolved correctly.
- domainsToAddresses: List of domain names and IP addresses to re-route.

DNS resolution can be tested with `dig`.  Examples below:
- `dig @{DNS server IP} -p {DNS server port} {domain to look up}`
- `dig @127.0.0.1 -p 5354 google.com`
