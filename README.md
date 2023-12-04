# go-congkit
Chinese input method "Congkit" engine written in Go

This project aims to create an input engine for Congkit input method using Go.



### Features

- Typing : Entering Congkit radicals will list the corresponding Chinese words with the subsequence radicals.
- Associated Phrase and Vocabularies: Finished typing a Chinese word will show a list of associated phrases.



### Data Source

This project needs a data source for mapping Chinese characters and Congkit radicals.

The source could be borrowed from an earlier project which placed the table under Public Domain.

Congkit Data Source: https://gitlab.freedesktop.org/cangjie/libcangjie/-/raw/master/data/table.txt



### Description

This project provides a binary which takes the input Congkit radicals as English characters and outputs the corresponding matching Chinese word list.



### Use as package

`go get github.com/antonyho/go-congkit`

```
import (
    congkit "github.com/antonyho/go-congkit/engine"
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
./congkit [congkit_radicals]

Usage of ./congkit:
  -d string
    	Custom database file path (default "congkit.db")
  -database string
    	Custom database file path (default "congkit.db")
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
    	Congkit version(3/5) (default 5)
  -version int
    	Congkit version(3/5) (default 5)
```

#### Usage Example #1
```
❯ ./congkit hqi
[我 牫 𥫻]
```

#### Usage Example #2
```
❯ ./congkit -v=5 ykmhm
[產]
```

#### Usage Example #3
```
❯ ./congkit -v=3 yhhqm
[產 産]
```

#### Usage Example #4
```
❯ ./congkit -s oiar
[仓]
```


### To-Do Plan

- [ ] Benchmarks
- [ ] Support wildcard for uncertain radical
- [ ] Type frequency (consent needed)