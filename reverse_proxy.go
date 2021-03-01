package trie_tree

//ip address 0-9&.
const CharCount = 11

//Trie树
type TrieNode struct {
	IsLeaf bool
	Url    string
	Child  []*TrieNode
}

func NewTrieNode(count int) *TrieNode {
	return &TrieNode{IsLeaf: false, Url: "", Child: make([]*TrieNode, count)}
}

type DNSCache struct {
	root *TrieNode
}

func NewDNSCache() *DNSCache {
	return &DNSCache{root: NewTrieNode(CharCount)}
}

func (d *DNSCache) getIndexFromRune(char rune) int {
	if char == '.' {
		return 10
	} else {
		return int(char) - '0'
	}
}

func (d *DNSCache) getRuneFromIndex(i int) rune {
	if i == 10 {
		return '.'
	} else {
		return rune('0' + i)
	}
}

// 把一个IP地址和相应的URL添加到trie树中，最后一个节点是URL
func (d *DNSCache) Insert(ip, url string) {
	pCrawl := d.root
	for _, v := range []rune(ip) {
		// 根据当前遍历到的IP中的字符，找到子节点的索引
		index := d.getIndexFromRune(v)
		// 如果子节点不存在，创建一个
		if pCrawl.Child[index] == nil {
			pCrawl.Child[index] = NewTrieNode(CharCount)
		}
		// 移动到子节点
		pCrawl = pCrawl.Child[index]
	}
	// 在叶子节点中存储IP地址对应的URL
	pCrawl.IsLeaf = true
	pCrawl.Url = url
}

// 通过IP地址找到对应的URL
func (d *DNSCache) SearchDNSCache(ip string) string {
	pCrawl := d.root
	for _, v := range []rune(ip) {
		index := d.getIndexFromRune(v)
		if pCrawl.Child[index] == nil {
			return ""
		}
		pCrawl = pCrawl.Child[index]
	}
	// 返回找到的URL
	if pCrawl != nil && pCrawl.IsLeaf {
		return pCrawl.Url
	}
	return ""
}
