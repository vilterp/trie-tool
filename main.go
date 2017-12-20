package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type trie map[string]trie

func NewTrie() trie {
	return make(trie)
}

func (t trie) add(path []string) {
	if len(path) == 0 {
		return
	}
	child, ok := t[path[0]]
	if !ok {
		child = NewTrie()
		t[path[0]] = child
	}
	child.add(path[1:])
}

func (t trie) Print(w io.Writer) {
	t.visit(func(key string, depth int) {
		for i := 0; i < depth; i++ {
			fmt.Fprint(w, "\t")
		}
		fmt.Fprintln(w, key)
	})
}

func (t trie) visit(f func(key string, depth int)) {
	visitRecur(f, t, 0)
}

func visitRecur(f func(key string, depth int), t trie, depth int) {
	for key, child := range t {
		f(key, depth)
		visitRecur(f, child, depth+1)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	t := NewTrie()
	for {
		text, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}
		newlineStripped := text[:len(text)-1]
		path := strings.Split(newlineStripped, ".")
		t.add(path)
	}
	t.Print(os.Stdout)
}
