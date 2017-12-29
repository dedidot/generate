## Getting Started

Stringer is tools for generate random string for golang. Stringer is have type: 

1. alnum : 0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ
2. alpha : abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ
3. lowalnum : 0123456789abcdefghijklmnopqrstuvwxy
4. numeric : 0123456789
5. nozero : 123456789

### Installing

```sh
$ go get github.com/dedidot/generate/stringer
```

### Usage

```
import(
    "github.com/dedidot/generate/stringer"
)
```

```sh
strRandom := stringer.RandomStr(18, "lowalnum") 
```

18 = length
lowalnum = type