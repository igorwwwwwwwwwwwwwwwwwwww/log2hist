# log2hist

log2hist is a histogram renderer based on the log2 histograms from [bpftrace](https://github.com/iovisor/bpftrace/blob/1ece0d0b1441aa70d4a6b324fb852954a5989eab/src/output.cpp#L166).

It transforms input in the shape of:

```
k1 10
k1 15
k1 20
k1 100
k2 1
k2 10
k2 100
```

Into a set of histograms such as:

```
k1
        [16, 32)        2 |@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@|
        [32, 64)        1 |@@@@@@@@@@@@@@@@@@@@@@@@@@                          |
       [64, 128)        0 |                                                    |
      [128, 256)        1 |@@@@@@@@@@@@@@@@@@@@@@@@@@                          |

k2
             [1]        1 |@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@|
          [2, 4)        0 |                                                    |
          [4, 8)        0 |                                                    |
         [8, 16)        0 |                                                    |
        [16, 32)        1 |@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@|
        [32, 64)        0 |                                                    |
       [64, 128)        0 |                                                    |
      [128, 256)        1 |@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@|
```

## LICENSE

MIT.
