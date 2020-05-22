# histy

Histy is a histogram renderer.

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
k1 (4)
[0,	7]	[                                        ]	(0/4)
[8,	15]	[========================================]	(2/4)
[16,	23]	[====================                    ]	(1/4)
[24,	31]	[                                        ]	(0/4)
[32,	39]	[                                        ]	(0/4)
[40,	47]	[                                        ]	(0/4)
[48,	55]	[                                        ]	(0/4)
[56,	63]	[                                        ]	(0/4)
[64,	71]	[                                        ]	(0/4)
[72,	79]	[                                        ]	(0/4)
[80,	87]	[                                        ]	(0/4)
[88,	95]	[                                        ]	(0/4)
[96,	103]	[====================                    ]	(1/4)

k2 (3)
[0,	7]	[========================================]	(1/3)
[8,	15]	[========================================]	(1/3)
[16,	23]	[                                        ]	(0/3)
[24,	31]	[                                        ]	(0/3)
[32,	39]	[                                        ]	(0/3)
[40,	47]	[                                        ]	(0/3)
[48,	55]	[                                        ]	(0/3)
[56,	63]	[                                        ]	(0/3)
[64,	71]	[                                        ]	(0/3)
[72,	79]	[                                        ]	(0/3)
[80,	87]	[                                        ]	(0/3)
[88,	95]	[                                        ]	(0/3)
[96,	103]	[========================================]	(1/3)
```

## LICENSE

MIT.

## TODO

* Come up with a better name, maybe log2hist.
* Explore alternate histogram implementations: [prometheus](https://github.com/prometheus/client_golang/blob/master/prometheus/histogram.go), [tally](https://github.com/uber-go/tally/blob/master/histogram.go), [llhist](https://github.com/circonus-labs/circonusllhist), [bcc](https://github.com/iovisor/bcc/blob/master/src/python/bcc/table.py) ([`helpers.h`](https://github.com/iovisor/bcc/blob/74e66b4f6730e0708f97150ac23d5951c5684ff8/src/cc/export/helpers.h#L765)), [bpftrace](https://github.com/iovisor/bpftrace/blob/1ece0d0b1441aa70d4a6b324fb852954a5989eab/src/output.cpp#L166).
