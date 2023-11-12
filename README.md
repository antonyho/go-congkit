# go-cangjie
Chinese input method "Cangjie" engine written in Go

This project aims to create an input engine for Cangjie input method using Go.



### Features

- Typing : Entering Cangjie radicals will list the corresponding Chinese words with the subsequence radicals.
- Associated Phrase and Vocabularies: Finished typing a Chinese word will show a list of associated phrases.



### Data Source

This project needs a data source for mapping Chinese characters and Cangjie radicals.

The source could be borrowed from an earlier project which placed the table under Public Domain.

Cangjie Data Source: https://gitlab.freedesktop.org/cangjie/libcangjie/-/raw/master/data/table.txt



### Description

This project provides a binary which takes the input Cangjie radicals as English characters and outputs the corresponding matching Chinese word list.



### Use as package

`go get github.com/antonyho/go-cangjie`

```
import (
    cangjie "github.com/antonyho/go-cangjie/engine"
)
```



### Build the binary

#### On Linux
##### Only building the binary
```
make
```

##### Build the binary and generate the database file
```
make generate
```



### Use as executable process
```
./cangjie [cangjie_radicals]

Usage of ./cangjie:
  -d string
    	Custom database file path (default "cangjie.db")
  -database string
    	Custom database file path (default "cangjie.db")
  -e	Use 'Easy' input method
  -easy
    	Use 'Easy' input method
  -h	Print usages
  -help
    	Print usages
  -p	Predict the possible typing word
  -prediction
    	Predict the possible typing word
  -s	Output simplified Chinese word
  -simplified
    	Output simplified Chinese word
  -v int
    	Cangjie version(3/5) (default 5)
  -version int
    	Cangjie version(3/5) (default 5)
```

#### Usage Example #1
```
❯ ./cangjie hqi
[我 牫 𥫻]
```

#### Usage Example #2
```
❯ ./cangjie -v=5 ykmhm
[產]
```

#### Usage Example #3
```
❯ ./cangjie -v=3 yhhqm
[產 産]
```

#### Usage Example #4
```
❯ ./cangjie -s oiar
[仓]
```


### To-Do Plan

- [ ] Benchmarks
- [ ] Support wildcard for uncertain radical
- [ ] Type frequency (consent needed)