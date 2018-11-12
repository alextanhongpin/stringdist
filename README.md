# stringdist

[![](https://godoc.org/github.com/alextanhongpin/stringdist?status.svg)](http://godoc.org/github.com/alextanhongpin/stringdist)

`stringdist` package contains several string metrics for calculating edit distance between two different strings. This includes the _Levenshtein Distance_, _Damerau Levenshtein_ (both _Optimal String Alignment_, OSA and _true_ damerau levenshtein), Jaro, Jaro Winkler and additionally a _BK-Tree_ that can be used for autocorrect.

## Algorithms

- __Levenshtein__: A string metric for measuring the difference between two sequence. Done by computing the _minimum_ number of single-edit character edit (`insertion`, `substitution` and `deletion`) required to change from one word to another.
- __Damerau-Levenshteim__: similar to Levenshtein, but allows transposition of two adjacent characters. Can be computed with two different algorithm - _Optimal String Alignment_, (OSA) and _true damerau-levenshtein_. The assumption for ASA is taht no substring is edited more than once.
- __Jaro__: Jaro distance between two words is the minimum number of single-character transpositions required to change one word into the other.
- __Jaro-Winkler__: Similar to Jaro, but uses a prefix scale which gives more favourable ratings to strings that match from the beginning for a set prefix length.
- __BK-Tree__: A tree data structure specialized to index data in a metric space. Can be used for approximate string matching in a dictionary.

Other algorithms to explore:
- Sift3/4 algorithm
- Soundex
- Metaphone
- Hamming Distance
- Symspell
- Linspell

## Thoughts

- Autocorrect can be implemented using any of the distance metrics (such as levenshtein) with BK-Tree
- Distance metric can be supplied to bk-tree through an interface.
- Dictionary words can first be supplied to the tree, and subsequent words can be added later through other means (syncing, streaming, pub-sub)
- The tree can be snapshotted periodically to avoid rebuild (e.g. using `gob`), test should be conducted to see if rebuilding the tree is faster than reloading the whole tree.
- Build tree through prefix (A-Z) would result in better performance (?). How to avoid hotspots (more characters in A than Z)?
- Can part of the tree be transmitted through the network?
- How to blacklist words that are not supposed to be searchable? (profanity words)
- 


## References
- https://en.wikibooks.org/wiki/Algorithm_Implementation/Strings/Dice%27s_coefficient#Javascript
- https://en.wikipedia.org/wiki/Wikipedia:AutoWikiBrowser/Typos#C
- https://ii.nlm.nih.gov/MTI/Details/trigram.shtml
- https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance
- https://en.wikipedia.org/wiki/Bitap_algorithm
- https://lingpipe-blog.com/2006/12/13/code-spelunking-jaro-winkler-string-comparison/
- Adjustment for longer string http://citeseerx.ist.psu.edu/viewdoc/download;jsessionid=7DCFAEBBA89D749D9D901DFA621FCA31?doi=10.1.1.64.7405&rep=rep1&type=pdf
- Table 6 shows the test cases https://www.census.gov/srd/papers/pdf/rrs2006-02.pdf
- http://alias-i.com/lingpipe/demos/tutorial/stringCompare/read-me.html

