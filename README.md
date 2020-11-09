# Simple Cache

Simple Cache is a simple cache implementation with a focus on readability, clarity and performance.

Store values in a key-value store which are computationally expensive to generate or retrieve, and
have them subsequently returned to you with lightning speed from the cache.

The Simple Cache has few features, intentionally.  Values are stored with a lifetime
duration.  Once this duration has elapsed, they are cleaned up and removed from the cache.  You can
store values indefinitely too.

You can deliberately remove entries from the cache using the `Delete` method.

## Examples

```go
package main

import (
    "simplecache"
    "time"
)

func main() {
    c := simplecache.New(time.Minute)
    defer c.Close()

    // Store a value against 'a-key' for 5 minutes...
    c.Set("a-key", "some-complex-value", time.Minute * 5)

    // Set a value indefinitely
    c.SetIndefinite("a-key", "some-complex-value")

    // Retrieve a value
    value, ok := c.Get("a-key")
    if !ok {
        // Not found the value
    }

    // Delete a value
    c.Delete("a-key")
}
```

## Benchmarks

Simple Cache has been benchmarked against each of its functions.

Simple Cache was benchmarked against a machine with the following specification:

* __Make/model__: MacBook Pro (15-inch, 2018) 
* __CPU__: 2.9 GHz 6-Core Intel Core i9
* __RAM__: 32 GB 2400 MHz DDR4

The following are the results:

```
Test                                Iterations              Nanoseconds per Op      Bytes alloc     Allocs
pkg: simplecache
BenchmarkHashCache_Get/Byte16x10-12             45408501                22.1 ns/op
BenchmarkHashCache_Get/Byte16x100-12            50895264                22.0 ns/op
BenchmarkHashCache_Get/Byte16x1000-12           49851592                21.8 ns/op
BenchmarkHashCache_Get/Byte16x10000-12          55497226                20.9 ns/op
BenchmarkHashCache_Get/Byte128x10-12            53074515                22.9 ns/op
BenchmarkHashCache_Get/Byte128x100-12           51671214                23.0 ns/op
BenchmarkHashCache_Get/Byte128x1000-12          55143558                21.7 ns/op
BenchmarkHashCache_Get/Byte128x10000-12         51909048                24.8 ns/op
BenchmarkHashCache_Get/Byte1024x10-12           54565129                22.7 ns/op
BenchmarkHashCache_Get/Byte1024x100-12          46009594                22.8 ns/op
BenchmarkHashCache_Get/Byte1024x1000-12         55138800                20.8 ns/op
BenchmarkHashCache_Get/Byte1024x10000-12        53092317                21.1 ns/op
BenchmarkHashCache_Get/Byte8192x10-12           53481157                21.4 ns/op
BenchmarkHashCache_Get/Byte8192x100-12          49936424                22.0 ns/op
BenchmarkHashCache_Get/Byte8192x1000-12         53615840                21.8 ns/op
BenchmarkHashCache_Get/Byte8192x10000-12        53922898                20.9 ns/op
BenchmarkHashCache_Get/Byte65536x10-12          50914202                22.5 ns/op
BenchmarkHashCache_Get/Byte65536x100-12         54566184                22.3 ns/op
BenchmarkHashCache_Get/Byte65536x1000-12        52880586                21.5 ns/op
BenchmarkHashCache_Get/Byte65536x10000-12       51452203                21.5 ns/op
BenchmarkHashCache_Get/Byte524288x10-12         54579868                23.6 ns/op
BenchmarkHashCache_Get/Byte524288x100-12        47658624                22.3 ns/op
BenchmarkHashCache_Get/Byte524288x1000-12       54127789                21.4 ns/op
BenchmarkHashCache_Get/Byte524288x10000-12      49856299                22.9 ns/op
BenchmarkHashCache_Get/Byte4194304x10-12        55792357                22.8 ns/op
BenchmarkHashCache_Get/Byte4194304x100-12       48027753                22.4 ns/op
BenchmarkHashCache_Get/Byte4194304x1000-12      53771000                21.6 ns/op
BenchmarkHashCache_Get/Byte4194304x10000-12     49844733                23.8 ns/op
BenchmarkHashCache_Get/Byte33554432x10-12       54154126                22.0 ns/op
BenchmarkHashCache_Get/Byte33554432x100-12      50052015                21.4 ns/op
BenchmarkHashCache_Get/Byte33554432x1000-12     53546043                22.9 ns/op
BenchmarkHashCache_Get/Byte33554432x10000-12            55401746                22.3 ns/op
BenchmarkHashCache_Get/Byte268435456x10-12              54343042                21.9 ns/op
BenchmarkHashCache_Get/Byte268435456x100-12             54233667                23.7 ns/op
BenchmarkHashCache_Get/Byte268435456x1000-12            54720808                20.9 ns/op
BenchmarkHashCache_Get/Byte268435456x10000-12           52620936                24.7 ns/op
BenchmarkHashCache_Get/Byte2147483648x10-12             54573882                21.6 ns/op
BenchmarkHashCache_Get/Byte2147483648x100-12            48646939                22.6 ns/op
BenchmarkHashCache_Get/Byte2147483648x1000-12           56910170                22.8 ns/op
BenchmarkHashCache_Get/Byte2147483648x10000-12          55692628                23.8 ns/op
BenchmarkHashCache_Set/Byte16x10-12                      2258113               569 ns/op
BenchmarkHashCache_Set/Byte16x100-12                     2238192               564 ns/op
BenchmarkHashCache_Set/Byte16x1000-12                    2255827               564 ns/op
BenchmarkHashCache_Set/Byte16x10000-12                   2225949               569 ns/op
BenchmarkHashCache_Set/Byte128x10-12                     2237157               570 ns/op
BenchmarkHashCache_Set/Byte128x100-12                    2256166               571 ns/op
BenchmarkHashCache_Set/Byte128x1000-12                   2254378               560 ns/op
BenchmarkHashCache_Set/Byte128x10000-12                  2237065               560 ns/op
BenchmarkHashCache_Set/Byte1024x10-12                    2244127               562 ns/op
BenchmarkHashCache_Set/Byte1024x100-12                   2246329               561 ns/op
BenchmarkHashCache_Set/Byte1024x1000-12                  2289274               552 ns/op
BenchmarkHashCache_Set/Byte1024x10000-12                 2299333               551 ns/op
BenchmarkHashCache_Set/Byte8192x10-12                    2281873               554 ns/op
BenchmarkHashCache_Set/Byte8192x100-12                   2282673               556 ns/op
BenchmarkHashCache_Set/Byte8192x1000-12                  2251870               557 ns/op
BenchmarkHashCache_Set/Byte8192x10000-12                 2276329               557 ns/op
BenchmarkHashCache_Set/Byte65536x10-12                   2275825               558 ns/op
BenchmarkHashCache_Set/Byte65536x100-12                  2290562               555 ns/op
BenchmarkHashCache_Set/Byte65536x1000-12                 2276390               555 ns/op
BenchmarkHashCache_Set/Byte65536x10000-12                2224436               565 ns/op
BenchmarkHashCache_Set/Byte524288x10-12                  2251174               559 ns/op
BenchmarkHashCache_Set/Byte524288x100-12                 2243061               569 ns/op
BenchmarkHashCache_Set/Byte524288x1000-12                2267726               558 ns/op
BenchmarkHashCache_Set/Byte524288x10000-12               2288089               561 ns/op
BenchmarkHashCache_Set/Byte4194304x10-12                 2246990               570 ns/op
BenchmarkHashCache_Set/Byte4194304x100-12                2228106               574 ns/op
BenchmarkHashCache_Set/Byte4194304x1000-12               2224502               563 ns/op
BenchmarkHashCache_Set/Byte4194304x10000-12              2269345               561 ns/op
BenchmarkHashCache_Set/Byte33554432x10-12                2292748               557 ns/op
BenchmarkHashCache_Set/Byte33554432x100-12               2157140               566 ns/op
BenchmarkHashCache_Set/Byte33554432x1000-12              2307084               556 ns/op
BenchmarkHashCache_Set/Byte33554432x10000-12             2265688               559 ns/op
BenchmarkHashCache_Set/Byte268435456x10-12               2561427               515 ns/op
BenchmarkHashCache_Set/Byte268435456x100-12              2479664               527 ns/op
BenchmarkHashCache_Set/Byte268435456x1000-12             2491854               524 ns/op
BenchmarkHashCache_Set/Byte268435456x10000-12            2506264               517 ns/op
BenchmarkHashCache_Set/Byte2147483648x10-12              2611123               530 ns/op
BenchmarkHashCache_Set/Byte2147483648x100-12             2624196               510 ns/op
BenchmarkHashCache_Set/Byte2147483648x1000-12            2589937               502 ns/op
BenchmarkHashCache_Set/Byte2147483648x10000-12           2593828               506 ns/op
BenchmarkHashCache_Delete-12                            50107819                24.1 ns/op
BenchmarkHashCache_Len-12                               92155755                12.6 ns/op

pkg: simplecache/web
BenchmarkTransactionalHandlerFunc/Get-12                 4693284               252 ns/op
BenchmarkTransactionalHandlerFunc/Put-12                 3195512               368 ns/op
BenchmarkTransactionalHandlerFunc/Delete-12             41423528                27.9 ns/op
```
