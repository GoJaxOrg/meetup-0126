# Why the new netip Package?

Here are some reasons identified by Brad Fitz:

- The contents are mutable. The underlying type of net.IP is []byte, which means that net.IP is passed by reference, and
any function that handles it can change its contents.

- Cannot be compared directly. Because slice cannot be compared directly, net.IP cannot be used directly to determine if
two addresses are equal using ==, nor can it be used as the key of a map.

- There are two types of addresses in the standard library, net.IP and net.IPAddr. Common IPv4 and IPv6 addresses are
stored using net.IP. IPv6 link-local addresses need to be stored using net.IPAddr (because of the additional storage
of the link’s NIC). Since there are two types of addresses, it’s a matter of determining which one to use or both.

- Takes up a lot of memory. A single slice header message takes 24 bytes (64-bit platforms, see Russ’s article for
details). So the memory footprint of net.IP contains 24 bytes of header information and 4 bytes (IPv4) or 6 bytes
(IPv6) of address data. If the local link NIC (zone) needs to be stored, then net.IPAddr also needs a 16-byte string
header and the specific NIC name.

- Memory needs to be allocated from the heap. Each time memory is allocated from the heap, it puts extra pressure on the GC.

- Cannot distinguish between IPv4 addresses and IPv4-mapped IPv6 addresses (in the form of ::ffff:192.168.1.1) when
parsing IP addresses from strings.

- Expose implementation details to the outside world. The definition of net.IP is type IP []byte, and the underlying
[]byte is part of the API and cannot be changed.

# Profiling

The `TestProfileProcess10KAddresses` func found in each example (net.IP and netip) allocate a large number of objects and save the heap before exiting. They don't do any processing.

The output is shown below in case you don't want to run `pprof` yourself. Only the data for the allocation func is shown.

_net.IP Heap Usage_

```shell
Total: 6.37MB
ROUTINE ======================== github.com/runlevl4/meetup-0126/netip/memory/net-ip.processIPAddresses in /Users/carybeuershausen/source/golang/meetup-0126/netip/memory/net-ip/main.go
    2.76MB     3.26MB (flat, cum) 51.18% of Total
         .          .     33:	var (
         .          .     34:		ipAddresses []net.IP
         .          .     35:	)
         .          .     36:	for _, ip := range ipList {
         .          .     37:		tip := net.ParseIP(ip)
    1.26MB     1.26MB     38:		ipAddresses = append(ipAddresses, tip)
         .          .     39:	}
         .          .     40:
         .          .     41:	if print {
         .          .     42:		for i, nip := range ipAddresses {
    1.50MB        2MB     43:			fmt.Printf("%d. %s\n", i+1, nip)
         .          .     44:		}
         .          .     45:	}
         .          .     46:}
         .          .     47:
         .          .     48:func loadIPList(file string) []string {
```

_netip Heap Usage_

```shell
Total: 3.21MB
ROUTINE ======================== github.com/runlevl4/meetup-0126/netip/memory/netip.processIPAddresses in /Users/carybeuershausen/source/golang/meetup-0126/netip/memory/netip/main.go
    1.10MB     1.10MB (flat, cum) 34.31% of Total
         .          .     27:	var (
         .          .     28:		ipAddresses []netip.Addr
         .          .     29:	)
         .          .     30:	for _, ip := range ipList {
         .          .     31:		tip, _ := netip.ParseAddr(ip)
    1.10MB     1.10MB     32:		ipAddresses = append(ipAddresses, tip)
         .          .     33:	}
         .          .     34:
         .          .     35:	if print {
         .          .     36:		for i, nip := range ipAddresses {
         .          .     37:			fmt.Printf("%d. %s\n", i+1, nip)
```

You can see that total memory dropped and the func dropped from 51.18% for `net.IP` to 34.31% for `netip`. Remember that the less you put on the heap, the less work the garbage collector has to do while also reducing the opportunity for memory leaks.

