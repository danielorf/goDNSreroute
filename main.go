// Borrowed structure from https://jameshfisher.com/2017/08/04/golang-dns-server.html

package main

import (
	"log"
	"net"
	"strconv"

	"github.com/miekg/dns"
)

var domainsToAddresses = map[string]string{
	//Real Addresses
	"jameshfisher.com.": "104.198.14.53",
	"jpmorgan.net.":     "209.132.183.105",
	"duckduckgo.com.":   "50.18.200.106",

	//Fake addresses
	"google.com.":   "1.2.3.4",
	"suse.com.":     "5.6.7.8",
	"ldaptest.com.": "127.0.0.1",
	"oidctest.com.": "127.0.0.1",
}

type handler struct{}

func dnsLookup(domain string, server string, m dns.Msg, dnsClient dns.Client) string {
	address := ""
	m.SetQuestion(domain, dns.TypeA)
	reply, _, err := dnsClient.Exchange(&m, server+":53")
	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("Took %v", rtt)
	if len(reply.Answer) == 0 {
		log.Fatal("No results")
	} else {
		address = reply.Answer[0].(*dns.A).A.String()
	}
	return address
}

func (this *handler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {

	// dn
	msg := dns.Msg{}
	msg.SetReply(r)

	dnsClient := dns.Client{}
	server := "192.168.100.3"
	m := dns.Msg{}

	switch r.Question[0].Qtype {
	case dns.TypeA:
		msg.Authoritative = true
		domain := msg.Question[0].Name
		address, ok := domainsToAddresses[domain]
		if ok {
			log.Printf("DNSReroute rerouted %s to %s", domain, address)
		} else {
			address = dnsLookup(domain, server, m, dnsClient)
			log.Printf("DNSLookup found %s at %s", domain, address)
		}
		msg.Answer = append(msg.Answer, &dns.A{
			Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60},
			A:   net.ParseIP(address),
		})
	}
	w.WriteMsg(&msg)
}

func main() {
	srv := &dns.Server{Addr: ":" + strconv.Itoa(5354), Net: "udp"}
	srv.Handler = &handler{}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to set udp listener %s\n", err.Error())
	}
}
