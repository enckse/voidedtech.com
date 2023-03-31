Git
===

## one-liners

### line counts

get line counts for all git-tracked files

```
for f in $(git ls-files); do printf "%-10s %s\n" $(wc -l $f); done | sort -n
```
