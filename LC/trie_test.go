package lc

import (
	"bufio"
	"fmt"
	"log"
	"strings"
	"testing"
	"time"
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
		c = '*'
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
	return n != nil && n.IsNode
}

func (o *Trie) StartsWith(prefix string) bool { return o.find(prefix) != nil }

func TestTrie(t *testing.T) {
	T := NewTrie()
	T.Insert("trie")
	T.Insert("is")
	T.Insert("a")
	T.Insert("prefix")
	T.Insert("tree")

	log.Print(T)
	log.Print(T.Children['i'-'a'])
	log.Print(T.Children['i'-'a'].Children['s'-'a'])

	log.Print("the? tree? -> ", T.Search("the"), T.Search("tree"))
	log.Print("pre? trie? the? -> ", T.StartsWith("pre"), T.StartsWith("trie"), T.StartsWith("the"))
}

type eTrie struct {
	Child  [26 + 26 + 10 + 1]*eTrie // A..Za..z0..9*'
	IsNode bool
}

func (o eTrie) String() string {
	child := make([]byte, len(o.Child))
	isNode := '-'
	if o.IsNode {
		isNode = '*'
	}
	for i, c := range o.Child {
		if c != nil {
			switch {
			case 0 <= i && i < 26:
				child[i] = 'A' + byte(i)
			case 26 <= i && i < 52:
				child[i] = 'a' + byte(i-26)
			case 52 <= i && i < 62:
				child[i] = '0' + byte(i-52)
			default:
				child[i] = '*'
			}
		} else {
			child[i] = '-'
		}
	}
	return fmt.Sprintf("[%s %c]", child, isNode)
}

// Trie
func TestTrieSearch(t *testing.T) {
	bToI := func(b byte) int {
		switch {
		case 'A' <= b && b < 'Z':
			return int(b) - 'A'
		case 'a' <= b && b < 'z':
			return 26 + int(b) - 'a'
		case '0' <= b && b < '9':
			return 52 + int(b) - '0'
		}
		return 62
	}

	trieSearch := func(n *eTrie, prefix string) *eTrie {
		for i := 0; i < len(prefix); i++ {
			child := n.Child[bToI(prefix[i])]
			if child == nil {
				return nil
			}
			n = child
		}
		return n
	}

	rdr := strings.NewReader("testing\nprefix\ntree\nTrie\nsearch\n1234\n-sign\n+sign")
	/*
		  f, err := os.Open("passwords.txt")
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()
			rdr := bufio.NewReader(f)
	*/

	T := &eTrie{}

	wc, ts := 0, time.Now()
	scr := bufio.NewScanner(rdr)
	for scr.Scan() {
		wrd := scr.Text()

		n := T
		for i := 0; i < len(wrd); i++ {
			child := n.Child[bToI(wrd[i])]
			if child == nil {
				child = &eTrie{}
			}
			n.Child[bToI(wrd[i])] = child
			n = child
		}
		n.IsNode = true

		wc++
	}
	if err := scr.Err(); err != nil {
		t.Fatal(err)
	}
	log.Printf("Trie Load: #%d %v", wc, time.Since(ts))
	log.Print(T)

	log.Print("Trie Search...")
	for _, wrd := range []string{"computer", "Trie", "pre", "*sign"} {
		ts := time.Now()
		n := trieSearch(T, wrd)
		log.Printf("? %-11s [%5t %5t]  %v", wrd, n != nil && n.IsNode, n != nil, time.Since(ts))
	}
}

// 79m Word Search
func Test79(t *testing.T) {
	exist := func(board [][]byte, word string) bool {
		type P struct{ i, j int }

		m, n := len(board), len(board[0])
		pValid := func(p P) bool { return p.i >= 0 && p.j >= 0 && m > p.i && n > p.j }

		var dfs func(P, string, [][]bool) bool
		dfs = func(p P, suffix string, Vis [][]bool) bool {
			if suffix == "" {
				return true
			}

			Vis[p.i][p.j] = true
			for _, dir := range []P{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				q := P{p.i + dir.i, p.j + dir.j}
				if pValid(q) && !Vis[q.i][q.j] && suffix[0] == board[q.i][q.j] {
					if dfs(q, suffix[1:], Vis) {
						return true
					}
				}
			}

			return false
		}

		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if board[i][j] != word[0] {
					continue
				}

				Vis := make([][]bool, m)
				for i := range Vis {
					Vis[i] = make([]bool, n)
				}
				if dfs(P{i, j}, word[1:], Vis) {
					return true
				}
			}
		}
		return false
	}

	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	for i := range board {
		for j := range board[i] {
			fmt.Printf("| %c ", board[i][j])
		}
		fmt.Println("|")
	}

	log.Print("true ?= ", exist(board, "ABCCED"))
	log.Print("true ?= ", exist(board, "SEE"))
	log.Print("false ?= ", exist(board, "ABCB"))
	log.Print("true ?= ", exist(board, "FBCCS"))
}
