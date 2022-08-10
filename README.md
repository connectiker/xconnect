# xConnect
`xConnect` provides the abstraction layers & managers for pluggable `xAddons` used in a microservice architecture, 
the basic interfaces & managers to define expectations and vocabulary.

```
+-----------+           +---------+
| xConnect  | ------>   | xAddons |
+----+------+           +-----+---+
     |                        |
     |                        |
     +------------+-----------+
                  |
                  |
                  V
             +-----------+ 
             | xService  | 
             +-----------+

```
