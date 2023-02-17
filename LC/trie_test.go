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
	c := '-'
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

// DFS Walk
func (o *Trie) Walk() {
	var walk func(*Trie, string)
	walk = func(n *Trie, w string) {
		log.Print(n, " ", w)

		for i, child := range n.Children {
			if child != nil {
				walk(child, w+string(byte(i)+'a'))
			}
		}
	}

	walk(o, "")
}

func TestInsert(t *testing.T) {
	T := NewTrie()
	for _, w := range []string{"a", "b", "c", "ab", "ac", "bcz", "abc", "abcdefghi", "z"} {
		T.Insert(w)
	}

	T.Walk()
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

func TestWildcard(t *testing.T) {
	T := &Trie{}
	T.Insert("a")
	T.Insert("app")
	T.Insert("apple")
	T.Insert("approve")
	T.Insert("application")

	T.Walk()

	log.Print(T.findWildcard(".", '.'), " ? .")
	log.Print(T.findWildcard("..", '.'), " ? ..")
	log.Print(T.findWildcard("app", '.'), " ? app")
	log.Print(T.findWildcard("a.p.e", '.'), " ? a.p.e")
	log.Print(T.findWildcard("*******", '*'), " ? *******")
}

// Find a word in Trie with wildcards: word or w.rd/w*rd
func (o *Trie) findWildcard(word string, wildcard byte) *Trie {
	var dfs func(*Trie, int) *Trie
	dfs = func(n *Trie, idx int) *Trie {
		if idx == len(word) {
			return n
		}

		if word[idx] == wildcard {
			for _, child := range n.Children {
				if child != nil {
					if n := dfs(child, idx+1); n != nil && n.IsNode {
						return n
					}
				}
			}
		} else {
			child := n.Children[word[idx]-'a']
			if child == nil {
				return nil
			}
			return dfs(child, idx+1)
		}

		return nil
	}

	return dfs(o, 0)
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

	/* f, err := os.Open("passwords.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	rdr := bufio.NewReader(f) */

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

		var dfs func(P, string) bool
		dfs = func(p P, suffix string) bool {
			if len(suffix) == 0 {
				log.Printf("+ %q", board)
				return true
			}

			b := board[p.i][p.j]
			board[p.i][p.j] = '*'

			log.Printf("> %s %q", suffix, board)

			dir := []int{0, 1, 0, -1, 0}
			for i := range dir[:4] {
				q := P{p.i + dir[i], p.j + dir[i+1]}
				if pValid(q) && suffix[0] == board[q.i][q.j] {
					if dfs(q, suffix[1:]) {
						return true
					}
				}
			}

			board[p.i][p.j] = b

			log.Printf("- %q", board)
			return false
		}

		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if board[i][j] != word[0] {
					continue
				}
				if dfs(P{i, j}, word[1:]) {
					return true
				}
			}
		}
		return false
	}

	boardSet := func() [][]byte { return [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}} }

	board := boardSet()
	for i := range board {
		for j := range board[i] {
			fmt.Printf("| %c ", board[i][j])
		}
		fmt.Println("|")
	}

	log.Print("true ?= ", exist(boardSet(), "ABCCED"))
	log.Print("true ?= ", exist(boardSet(), "SEE"))
	log.Print("false ?= ", exist(boardSet(), "ABCB"))
	log.Print("true ?= ", exist(boardSet(), "ABCCFSADEESE"))
	log.Print("true ?= ", exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCEFSADEESE"))
}

// 139m Word Break
func Test139(t *testing.T) {
	wordBreak := func(s string, wordDict []string) bool {
		Visited := map[int]struct{}{}

		var possible func(int) bool
		possible = func(start int) bool {
			if start == len(s) {
				return true
			}

			if _, ok := Visited[start]; ok {
				return false
			}

			subs := s[start:]
			for _, w := range wordDict {
				log.Printf("%d %s %s", start, subs, w)

				if strings.HasPrefix(subs, w) {
					if possible(start + len(w)) {
						return true
					} else {
						Visited[start+len(w)] = struct{}{}
					}
				}
			}

			return false
		}

		return possible(0)
	}

	// DP: Bottom-Up
	bottomUp := func(s string, wordDict []string) bool {
		D := make([]bool, len(s)+1)
		D[0] = true
		for i := 1; i <= len(s); i++ {
			for _, w := range wordDict {
				if i-len(w) >= 0 && s[i-len(w):i] == w && D[i-len(w)] {
					D[i] = true
				}
			}
		}

		return D[len(s)]
	}

	for _, f := range []func(string, []string) bool{wordBreak, bottomUp} {
		log.Print("(TLE) false ?= ", f("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}))
		log.Print("true ?= ", f("applepenapple", []string{"apple", "pen"}))
		log.Print("false ?= ", f("catsandogs", []string{"cats", "dog", "sand", "and", "cat"}))
	}
}

// 472h Concatenated Words
func Test472(t *testing.T) {
	findAllConcatenatedWordsInADict := func(words []string) []string {
		W := []string{}

		for _, w := range words {
			// DP: w[0..n] -> Y/N segmented?
			D := make([]bool, len(w)+1)
			D[0] = true

			for i := 1; i <= len(w); i++ {
				for j := 0; j < len(words) && !D[i]; j++ {
					if w == words[j] {
						continue
					}
					if i-len(words[j]) >= 0 && D[i-len(words[j])] && w[i-len(words[j]):i] == words[j] {
						D[i] = true
					}
				}
			}

			if D[len(w)] {
				W = append(W, w)
			}
		}

		return W
	}

	trieSearcher := func(words []string) []string {
		type Trie struct {
			isNode   bool
			children [26]*Trie
		}
		insert := func(n *Trie, word string) {
			for i := 0; i < len(word); i++ {
				child := n.children[word[i]-'a']
				if child == nil {
					child = &Trie{}
					n.children[word[i]-'a'] = child
				}
				n = child
			}
			n.isNode = true
		}
		search := func(n *Trie, word string) bool {
			for i := 0; i < len(word); i++ {
				child := n.children[word[i]-'a']
				if child == nil {
					return false
				}
				n = child
			}
			return n.isNode
		}

		trie := &Trie{}
		for _, w := range words {
			insert(trie, w)
		}

		W := []string{}
		for _, w := range words {
			D := make([]bool, len(w)+1)
			D[0] = true

			for i := 1; i <= len(w); i++ {
				for j := 0; j < i && !D[i]; j++ {
					if i == len(w) && j == 0 {
						continue
					}
					if D[j] && search(trie, w[j:i]) {
						D[i] = true
					}
				}
			}

			if D[len(w)] {
				W = append(W, w)
			}
		}
		return W
	}

	hashSearcher := func(words []string) []string {
		W := []string{}

		m := map[string]bool{}
		for _, w := range words {
			m[w] = true
		}

		for _, w := range words {
			D := make([]bool, len(w)+1)
			D[0] = true

			for i := 1; i <= len(w); i++ {
				for j := 0; j < i && !D[i]; j++ {
					if i == len(w) && j == 0 {
						continue
					}

					if D[j] && m[w[j:i]] {
						D[i] = true
					}
				}
			}

			if D[len(w)] {
				W = append(W, w)
			}
		}

		return W
	}

	for _, f := range []func([]string) []string{findAllConcatenatedWordsInADict, trieSearcher, hashSearcher} {
		log.Print(f([]string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}))
	}
}
