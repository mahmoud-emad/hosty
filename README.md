# Hosty

Hosty is a small CLI tool for managing and connecting to remote machines using memorable names instead of IP addresses.

Instead of remembering:

```bash
ssh root@185.206.122.30
```

You can save the machine:

```bash
hosty set studio 185.206.122.30
```

And connect to it later:

```bash
hosty connect studio
```

## Commands

```bash
hosty set <name> <ip>
hosty get <name>
hosty list
hosty forget <name>
hosty connect <name>
```

Hosty supports both IPv4 and IPv6 addresses, with support for usernames and ports.

## Configuration

Hosts are stored locally in:

```text
~/.config/hosty/hosts.json
```

Hosty does not manage SSH keys or implement the SSH protocol itself. It uses the system `ssh` command, so your existing SSH configuration, keys, and agent continue to work.

## Status

Hosty is currently a small learning project written in Go.
