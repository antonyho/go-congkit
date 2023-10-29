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