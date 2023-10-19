package lc

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func init() {
	log.Print("> Trie")
}

type Trie struct {
	Children [26]*Trie
	IsNode   bool
}

func (o Trie) String() string {
	sb := strings.Builder{}
	for i := byte(0); i < 26; i++ {
		c := byte('-')
		if o.Children[i] != nil {
			c = 'a' + i
		}
		sb.WriteByte(c)
	}
	c := ':'
	if o.IsNode {
		c = '+'
	}
	return fmt.Sprintf("[%s %c]", sb.String(), c)
}

func NewTrie() *Trie { return &Trie{} }

func (o *Trie) Insert(word string) {
	n := o
	for i := 0; i < len(word); i++ {
		c := word[i]
		if n.Children[c-'a'] == nil {
			n.Children[c-'a'] = &Trie{}
		}
		n = n.Children[c-'a']
	}
	n.IsNode = true
}

func (o *Trie) find(word string) *Trie {
	n := o
	for i := 0; i < len(word); i++ {
		c := word[i]
		if n.Children[c-'a'] != nil {
			n = n.Children[c-'a']
		} else {
			return nil
		}
	}
	return n
}

func (o *Trie) Search(word string) bool {
	n := o.find(word)
	if n != nil {
		return n.IsNode
	}
	return false
}

func (o *Trie) StartsWith(word string) bool { return o.find(word) != nil }

func TestTrie(t *testing.T) {
	T := NewTrie()
	T.Insert("trie")
	T.Insert("is")
	T.Insert("prefix")
	T.Insert("tree")

	log.Print(T)
	log.Print(T.Children['i'-'a'])
	log.Print(T.Children['i'-'a'].Children['s'-'a'])

	log.Print(T.Search("the"), T.Search("tree"))
	log.Print(T.StartsWith("pre"), T.StartsWith("tree"), T.StartsWith("the"))
}
