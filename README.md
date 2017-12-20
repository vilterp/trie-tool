# trie-tool

Print out dot-delimited paths as a tree. (Output indented with tabs)

E.g. Given `file.txt`:
```
a.b
b.c.d
```

Run:

```sh
$ cat file.txt | trie-tool
a
	b
b
	c
		d
```
