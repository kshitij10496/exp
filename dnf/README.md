# dnf - DNS Forwarder

This is an attempt at the [Coding Challenge #44 - DNS Forwarder](https://codingchallenges.substack.com/p/coding-challenge-44-dns-forwarder).

# Step Zero

Using Go 1.21 to learn UDP network programming.

# Step 1

**Objective**: Create a UDP server that will listen on a specified port for incoming requests.

**Approach**: Add a `-port` flag using the `flag` package from the standard library with a default value of `53`.

# Step 2

**Objective**: Parse the request packet that has been sent to your server.

**Approach**: Using the [github.com/miekg/dns](https://github.com/miekg/dns) package, I did not implement the actual byte parsing logic. My goal is to get to the meat of the problem this tool solves. I will probably try to revisit this once I am through with all the steps.