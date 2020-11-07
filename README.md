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
    "./src"
    "time"
)

func main() {
    c := cache.New(time.Minute)
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
BenchmarkGet/Byte16x10-12             51454510                22.7 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte16x100-12            49934908                21.8 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte16x1000-12           53226744                21.6 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte16x10000-12          56157685                24.4 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte128x10-12            51900474                25.1 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte128x100-12           43203547                24.8 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte128x1000-12          50396576                22.0 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte128x10000-12         53668971                23.9 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte1024x10-12           54387532                23.3 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte1024x100-12          45047198                25.0 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte1024x1000-12         53310292                24.9 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte1024x10000-12        49778852                23.7 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte8192x10-12           50307432                23.1 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte8192x100-12          53909013                27.0 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte8192x1000-12         56785495                22.1 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte8192x10000-12        54017059                22.8 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte65536x10-12          51067045                22.8 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte65536x100-12         51092523                23.8 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte65536x1000-12        53174803                23.3 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte65536x10000-12       52851549                23.6 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte524288x10-12         44416586                22.6 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte524288x100-12        51787684                25.6 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte524288x1000-12       51114158                22.5 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte524288x10000-12      46503898                23.2 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte4194304x10-12        52186716                25.7 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte4194304x100-12       52696802                26.4 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte4194304x1000-12      56605660                21.8 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte4194304x10000-12     51886293                22.1 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte33554432x10-12       48319400                24.2 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte33554432x100-12      52059420                23.1 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte33554432x1000-12     54425550                22.1 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte33554432x10000-12            54738646                22.6 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte268435456x10-12              52251931                23.2 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte268435456x100-12             49451070                25.5 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte268435456x1000-12            54135285                23.1 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte268435456x10000-12           49235263                23.3 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte2147483648x10-12             49466880                23.6 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte2147483648x100-12            52579755                25.1 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte2147483648x1000-12           54742369                22.8 ns/op             0 B/op          0 allocs/op
BenchmarkGet/Byte2147483648x10000-12          53958951                25.2 ns/op             0 B/op          0 allocs/op
BenchmarkSet/Byte16x10-12                      2242088               571 ns/op             218 B/op          2 allocs/op
BenchmarkSet/Byte16x100-12                     2192128               573 ns/op             222 B/op          2 allocs/op
BenchmarkSet/Byte16x1000-12                    2161458               578 ns/op             225 B/op          2 allocs/op
BenchmarkSet/Byte16x10000-12                   2233011               569 ns/op             219 B/op          2 allocs/op
BenchmarkSet/Byte128x10-12                     2228892               577 ns/op             220 B/op          2 allocs/op
BenchmarkSet/Byte128x100-12                    2190296               583 ns/op             223 B/op          2 allocs/op
BenchmarkSet/Byte128x1000-12                   2204491               571 ns/op             221 B/op          2 allocs/op
BenchmarkSet/Byte128x10000-12                  2202706               582 ns/op             222 B/op          2 allocs/op
BenchmarkSet/Byte1024x10-12                    2198372               586 ns/op             222 B/op          2 allocs/op
BenchmarkSet/Byte1024x100-12                   2228340               575 ns/op             220 B/op          2 allocs/op
BenchmarkSet/Byte1024x1000-12                  2216042               578 ns/op             220 B/op          2 allocs/op
BenchmarkSet/Byte1024x10000-12                 2242012               573 ns/op             218 B/op          2 allocs/op
BenchmarkSet/Byte8192x10-12                    2204413               580 ns/op             221 B/op          2 allocs/op
BenchmarkSet/Byte8192x100-12                   2205207               581 ns/op             221 B/op          2 allocs/op
BenchmarkSet/Byte8192x1000-12                  2189613               578 ns/op             223 B/op          2 allocs/op
BenchmarkSet/Byte8192x10000-12                 2211446               575 ns/op             221 B/op          2 allocs/op
BenchmarkSet/Byte65536x10-12                   2226477               585 ns/op             220 B/op          2 allocs/op
BenchmarkSet/Byte65536x100-12                  2213851               574 ns/op             221 B/op          2 allocs/op
BenchmarkSet/Byte65536x1000-12                 2168294               591 ns/op             224 B/op          2 allocs/op
BenchmarkSet/Byte65536x10000-12                2178804               579 ns/op             224 B/op          2 allocs/op
BenchmarkSet/Byte524288x10-12                  2156144               579 ns/op             225 B/op          2 allocs/op
BenchmarkSet/Byte524288x100-12                 2188779               584 ns/op             223 B/op          2 allocs/op
BenchmarkSet/Byte524288x1000-12                2192739               577 ns/op             222 B/op          2 allocs/op
BenchmarkSet/Byte524288x10000-12               2243992               574 ns/op             218 B/op          2 allocs/op
BenchmarkSet/Byte4194304x10-12                 2215375               577 ns/op             221 B/op          2 allocs/op
BenchmarkSet/Byte4194304x100-12                2192536               575 ns/op             223 B/op          2 allocs/op
BenchmarkSet/Byte4194304x1000-12               2278267               570 ns/op             215 B/op          2 allocs/op
BenchmarkSet/Byte4194304x10000-12              2226490               572 ns/op             220 B/op          2 allocs/op
BenchmarkSet/Byte33554432x10-12                2270582               562 ns/op             216 B/op          2 allocs/op
BenchmarkSet/Byte33554432x100-12               2294055               559 ns/op             214 B/op          2 allocs/op
BenchmarkSet/Byte33554432x1000-12              2266826               572 ns/op             216 B/op          2 allocs/op
BenchmarkSet/Byte33554432x10000-12             2232838               568 ns/op             219 B/op          2 allocs/op
BenchmarkSet/Byte268435456x10-12               2467710               531 ns/op             202 B/op          2 allocs/op
BenchmarkSet/Byte268435456x100-12              2479741               522 ns/op             201 B/op          2 allocs/op
BenchmarkSet/Byte268435456x1000-12             2527914               523 ns/op             198 B/op          2 allocs/op
BenchmarkSet/Byte268435456x10000-12            2473074               525 ns/op             202 B/op          2 allocs/op
BenchmarkSet/Byte2147483648x10-12              2572890               496 ns/op             195 B/op          2 allocs/op
BenchmarkSet/Byte2147483648x100-12             2561679               505 ns/op             196 B/op          2 allocs/op
BenchmarkSet/Byte2147483648x1000-12            2591049               499 ns/op             194 B/op          2 allocs/op
BenchmarkSet/Byte2147483648x10000-12           2552149               509 ns/op             197 B/op          2 allocs/op
BenchmarkDelete-12                                      46780657                27.1 ns/op             0 B/op          0 allocs/op
PASS
ok      _/Users/richardpickering/WebstormProjects/simple-cache/src/cache        132.513s
```

