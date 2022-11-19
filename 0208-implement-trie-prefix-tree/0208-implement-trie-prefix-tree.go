type Trie struct {
    children map[rune]*Trie
    endOfWord bool
}


func Constructor() Trie {
    return Trie{children: make(map[rune]*Trie)}
}


func (this *Trie) Insert(word string)  {
    currRoot := this
    for _, char := range word {
        if _, ok := currRoot.children[char]; !ok {
            newNode := Constructor()
            currRoot.children[char] = &newNode
        }
        currRoot = currRoot.children[char]
    }
    currRoot.endOfWord = true
}


func (this *Trie) Search(word string) bool {
    currRoot := this
    for _, char := range word {
        if _, ok := currRoot.children[char]; !ok {
            return false
        }
        currRoot = currRoot.children[char]
    }
    return currRoot.endOfWord
}


func (this *Trie) StartsWith(prefix string) bool {
    currRoot := this
    for _, char := range prefix {
        if _, ok := currRoot.children[char]; !ok {
            return false
        }
        currRoot = currRoot.children[char]
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