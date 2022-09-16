Memory
===

## Dirty Memory

View status of dirty memory (e.g. `sync` is running before `umount`)
```
watch -n 1 grep -e Dirty: -e Writeback: /proc/meminfo 
```
