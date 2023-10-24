package main

import (
	"fmt"
	"net"
	"regexp"

	"github.com/miekg/dns"
)

func resolve(domain, dnsServer string) (string, error) {
	// Validate and transform the domain to FQDN.
	domain = dns.Fqdn(domain)

	// Validate the dns server.
	ns := net.ParseIP(dnsServer)
	if ns == nil {
		return "", fmt.Errorf("invalid dns server provided: %s", dnsServer)
	}

	for {
		resp, err := queryDNS(domain, ns)
		if err != nil {
			return "", fmt.Errorf("querying DNS server %s: %w", ns.String(), err)
		}
		addr, err := parseReply(resp)
		if err != nil {
			return "", fmt.Errorf("parsing DNS response: %w", err)
		}
		return addr, nil
	}
}

func parseReply(resp *dns.Msg) (string, error) {
	for _, record := range resp.Answer {
		if record.Header().Rrtype == dns.TypeA {
			if ip := parseRecord(record.String()); ip != "" {
				return ip, nil
			}
		}
	}
	return "", fmt.Errorf("no IP address record found")
}

func parseRecord(input string) string {
	// Define a regular expression to match an IP address
	// Here, we're assuming that the IP address is in the format of 4 groups of digits separated by periods.
	re := regexp.MustCompile(`(\d+\.\d+\.\d+\.\d+)$`)

	// Find the last matching IP address in the input string
	return re.FindString(input)
}

func queryDNS(address string, ns net.IP) (*dns.Msg, error) {
	msg := &dns.Msg{
		Question: []dns.Question{
			{
				Name:   address,
				Qtype:  dns.TypeA,
				Qclass: dns.ClassINET,
			},
		},
	}

	// Use UDP network layer for DNS query.
	client := dns.Client{
		Net: "udp",
	}

	// Use port 53 by default.
	nsAddr := net.UDPAddr{
		IP:   ns,
		Port: 53,
	}

	// TODO: Explore the duration for DNS query.
	resp, _, err := client.Exchange(msg, nsAddr.String())
	if err != nil {
		return nil, err
	}
	return resp, nil
}
