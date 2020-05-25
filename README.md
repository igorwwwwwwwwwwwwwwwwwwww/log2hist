# log2hist

log2hist chews numbers, spits out ascii histograms.

## Example

```
$ shuf -i 2000-65000 -n 200 | log2hist

    [2048, 4096)        8  ∎∎∎
    [4096, 8192)       10  ∎∎∎∎
   [8192, 16384)       19  ∎∎∎∎∎∎∎∎∎
  [16384, 32768)       57  ∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  [32768, 65536)      106  ∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
```

## Credits

The code and idea are heavily influenced by the log2 histograms from [bpftrace](https://github.com/iovisor/bpftrace/blob/1ece0d0b1441aa70d4a6b324fb852954a5989eab/src/output.cpp#L166).

## License

MIT.
