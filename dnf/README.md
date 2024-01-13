# dnf - DNS Forwarder

This is an attempt at the [Coding Challenge #44 - DNS Forwarder](https://codingchallenges.substack.com/p/coding-challenge-44-dns-forwarder).

# Step Zero

Using Go 1.21 to learn UDP network programming.

# Step 1

**Objective**: Create a UDP server that will listen on a specified port for incoming requests.

**Approach**: Add a `-port` flag using the `flag` package from the standard library with a default value of `53`.
