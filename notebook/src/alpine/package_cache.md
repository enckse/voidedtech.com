Package Cache
===

Utilizing `nginx` to provide an Alpine Linux package cache for apk files. The idea
is, on a network, to use nginx to provide a common repository of cached packages
that _any_ machine has requested that can then be re-used by subsequent requests.

Below is an example configuration file (that would then be "included" into an `nginx.conf`)

```
/etc/nginx/apk.conf
---
server_names_hash_bucket_size 128;

server
{
    listen      9999;
    root        /srv/http/apk/;
    autoindex   on;

    # Requests for package db, signature files and files db should redirect upstream without caching
    location ~ \.tar\.gz$ {
        proxy_pass http://mirrors$request_uri;
    }

    # Requests for actual packages should be served directly from cache if available.
    #   If not available, retrieve and save the package from an upstream mirror.
    location ~ \.apk$ {
        try_files $uri @pkg_mirror;
    }

    # Retrieve package from upstream mirrors and cache for future requests
    location @pkg_mirror {
        proxy_store    on;
        proxy_redirect off;
        proxy_store_access  user:rw group:rw all:r;
        proxy_next_upstream error timeout http_404;
        proxy_pass          http://mirrors$request_uri;
    }
}

# Upstream mirrors
# - Configure as many backend mirrors as you want in the blocks below
# - Servers are used in a round-robin fashion by nginx
# - Add "backup" if you want to only use the mirror upon failure of the other mirrors
# - Use separate mirror server blocks to be able to use mirrors that have different paths to the package repos
upstream mirrors {
    server 127.0.0.1:8001;
}

# the proxy_pass directive should look like this
# proxy_pass http://mirror.domain.example/path/to/repo$request_uri;
#
# Notice that $request_uri replaces the /$repo/os/$arch part of
# the mirror address. See more examples below.

server
{
    listen      127.0.0.1:8001;

    location / {
        proxy_pass       http://dl-cdn.alpinelinux.org$request_uri;
    }
}
```

Which can then be referenced by an Alpine install

```
vim /etc/apk/repositories
---
http://<host>:9999/alpine/v3.14/main
http://<host>:9999/alpine/v3.14/community
```
