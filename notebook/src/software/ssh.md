ssh
===

# Profile

## noprofile

Connect without profile

```
ssh -t <host> bash --norc --noprofile
```

# Host Keys

## Scanning

scan keys for git forwarding/relaying/mirroring/etc.

```
ssh-keyscan -t rsa github.com >> ~/.ssh/known_hosts
```

<sub><sup>Updated: 2021-12-13</sup></sub>
