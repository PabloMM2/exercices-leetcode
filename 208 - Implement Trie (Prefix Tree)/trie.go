type Trie struct {
    Children map[rune]*Trie
    IsEnd bool
}

func Constructor() Trie {
    return Trie{Children: make(map[rune]*Trie)}
}

func (this *Trie) Insert(word string) {
    current := this

    for _, char := range word {
        if _, ok := current.Children[char]; !ok {
            newTrie := Constructor()
            current.Children[char] = &newTrie
        }
        current = current.Children[char]
    }
    
    current.IsEnd = true
}

func (this *Trie) Search(word string) bool {
    current := this

    for _, char := range word {
        if _, ok := current.Children[char]; !ok {
           return false
        }
        current = current.Children[char]
    }

    return current.IsEnd
}

func (this *Trie) StartsWith(prefix string) bool {
    current := this

    for _, char := range prefix {
        if _, ok := current.Children[char]; !ok {
            return false
        }
        current = current.Children[char]
    }
    
    return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */