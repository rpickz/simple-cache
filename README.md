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
BenchmarkHashCache_Get/Byte16x10-12             55476240                22.6 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte16x100-12            53066190                24.4 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte16x1000-12           51842634                22.3 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte16x10000-12          53608676                22.3 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte128x10-12            52284294                22.3 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte128x100-12           48275804                21.4 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte128x1000-12          52628275                21.5 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte128x10000-12         55039910                21.5 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte1024x10-12           52595434                21.4 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte1024x100-12          52467171                22.7 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte1024x1000-12         53088578                21.3 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte1024x10000-12        47793980                21.7 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte8192x10-12           53019108                22.8 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte8192x100-12          53827578                22.2 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte8192x1000-12         55411922                20.8 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte8192x10000-12        52094169                21.7 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte65536x10-12          48434702                21.5 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte65536x100-12         53755449                22.2 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte65536x1000-12        55333750                20.8 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte65536x10000-12       47729581                24.7 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte524288x10-12         54539494                22.7 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte524288x100-12        52138134                21.6 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte524288x1000-12       48840132                21.3 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte524288x10000-12      52410690                22.1 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte4194304x10-12        50724450                21.5 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte4194304x100-12       49020811                22.1 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte4194304x1000-12      53594049                22.2 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte4194304x10000-12     53532528                21.7 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte33554432x10-12       51295472                22.7 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte33554432x100-12      55162215                21.6 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte33554432x1000-12     54156126                21.6 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte33554432x10000-12            47652477                22.5 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte268435456x10-12              54292471                22.1 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte268435456x100-12             55740627                21.3 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte268435456x1000-12            53665546                21.7 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte268435456x10000-12           46311897                22.0 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte2147483648x10-12             55116741                22.4 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte2147483648x100-12            54401917                22.9 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte2147483648x1000-12           59320562                21.8 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Get/Byte2147483648x10000-12          56376190                21.4 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Set/Byte16x10-12                      2226417               579 ns/op             220 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte16x100-12                     2214525               585 ns/op             221 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte16x1000-12                    2228475               574 ns/op             220 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte16x10000-12                   2235366               577 ns/op             219 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte128x10-12                     2202376               580 ns/op             222 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte128x100-12                    2240728               568 ns/op             218 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte128x1000-12                   2256168               567 ns/op             217 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte128x10000-12                  2238990               574 ns/op             219 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte1024x10-12                    2246172               568 ns/op             218 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte1024x100-12                   2244782               572 ns/op             218 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte1024x1000-12                  2212513               572 ns/op             221 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte1024x10000-12                 2242311               572 ns/op             218 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte8192x10-12                    2218930               571 ns/op             220 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte8192x100-12                   2239107               563 ns/op             219 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte8192x1000-12                  2259296               565 ns/op             217 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte8192x10000-12                 2213745               573 ns/op             221 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte65536x10-12                   2226502               577 ns/op             220 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte65536x100-12                  2241740               572 ns/op             218 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte65536x1000-12                 2122578               583 ns/op             229 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte65536x10000-12                2228971               572 ns/op             219 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte524288x10-12                  2218023               571 ns/op             220 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte524288x100-12                 2222146               577 ns/op             220 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte524288x1000-12                2222866               571 ns/op             220 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte524288x10000-12               2259063               560 ns/op             217 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte4194304x10-12                 2258872               568 ns/op             217 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte4194304x100-12                2221099               573 ns/op             220 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte4194304x1000-12               2239100               577 ns/op             219 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte4194304x10000-12              2194965               582 ns/op             222 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte33554432x10-12                2215353               617 ns/op             221 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte33554432x100-12               1916194               696 ns/op             249 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte33554432x1000-12              1869169               629 ns/op             254 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte33554432x10000-12             2173021               586 ns/op             224 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte268435456x10-12               2468854               541 ns/op             202 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte268435456x100-12              2533442               518 ns/op             198 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte268435456x1000-12             2536617               522 ns/op             198 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte268435456x10000-12            2484727               524 ns/op             201 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte2147483648x10-12              2617057               520 ns/op             193 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte2147483648x100-12             2618032               494 ns/op             193 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte2147483648x1000-12            2609324               514 ns/op             194 B/op          2 allocs/op
BenchmarkHashCache_Set/Byte2147483648x10000-12           2475064               561 ns/op             202 B/op          2 allocs/op
BenchmarkHashCache_Delete-12                            45795818                25.1 ns/op             0 B/op          0 allocs/op
BenchmarkHashCache_Len-12                               80880376                13.5 ns/op             0 B/op          0 allocs/op

pkg: simplecache/web
BenchmarkTransactionalHandlerFunc/Get-12                 4661180               270 ns/op             131 B/op          3 allocs/op
BenchmarkTransactionalHandlerFunc/Put-12                 2650690               402 ns/op             560 B/op          3 allocs/op
BenchmarkTransactionalHandlerFunc/Delete-12             41117755                29.9 ns/op             0 B/op          0 allocs/op
```
