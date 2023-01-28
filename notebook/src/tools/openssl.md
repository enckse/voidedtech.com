openssl
===

## imap

`openssl` can connect to an imap server and run commands (e.g. fastmail)

```
#!/usr/bin/env bash
{
  sleep 3
  echo "a1 LOGIN myemail@example.org myapppassword"
  sleep 3
  echo 'a1 LIST "" "*"'
  sleep 3
} | openssl s_client -connect imap.fastmail.com:993 -crlf >/dev/null 2>&1
```
