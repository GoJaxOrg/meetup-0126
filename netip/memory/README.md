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