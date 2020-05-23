# log2hist

log2hist chews numbers, spits out ascii histograms.

## Example

```
$ shuf -i 2000-65000 -n 200 | log2hist

    [1024, 2048)        1 |                                                    |
    [2048, 4096)       10 |@@@@@                                               |
    [4096, 8192)       14 |@@@@@@@@                                            |
   [8192, 16384)       27 |@@@@@@@@@@@@@@@                                     |
  [16384, 32768)       57 |@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@                    |
  [32768, 65536)       91 |@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@|
```

## LICENSE

MIT.

## Credits

The code and idea are heavily influenced by the log2 histograms from [bpftrace](https://github.com/iovisor/bpftrace/blob/1ece0d0b1441aa70d4a6b324fb852954a5989eab/src/output.cpp#L166).
