package SensitiveWordFilter

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

const (
	ReplaceString = "***"
	InvalidChar   = " ,~,!,@,#,$,%,^,&,*,(,),_,-,+,=,?,<,>,.,—,，,。,/,\\,|,《,》,？,;,:,：,',‘,；,“,¥,·"
)

var (
	InvalidChars map[rune]struct{}
	T            *Node
)

type Node struct {
	isEnd bool //是否为某个字符串的结尾
	Child map[rune]*Node
}

func (t *Node) setIsEnd(is bool) {
	t.isEnd = is
}

func (t *Node) addChild(char rune, node *Node) {
	t.Child[char] = node
}

func NewNode() *Node {
	return &Node{Child: make(map[rune]*Node)}
}

func NewTrie() *Node {
	node := NewNode()
	node.init()
	return node
}

func (t *Node) InsertWord(word string) {
	for _, v := range word {
		if node, ok := t.Child[v]; ok {
			t = node
		} else {
			node := NewNode()
			t.addChild(v, node)
			t = node
		}
	}
	t.setIsEnd(true)
}

func (t *Node) init() {
	path := "./Utils/SensitiveWordFilter/SensitiveWord.txt"
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		log.Println(err, "打开敏感词文件失败")
	}
	r := bufio.NewReader(f)
	for {
		word, _, err := r.ReadLine()
		t.InsertWord(string(word))
		if err == io.EOF {
			break
		}
	}
}

func (t *Node) Check(word string) (string, bool) {
	var (
		Has      = false //是否检测到敏感词
		Nodeptr  = t     //指向Trie树结点
		Wordptr1 = 0     //以该指向的字符为起始判断是否有敏感词
		Wordptr2 = 1     //指向WordPtr1后未判断的字符
	)
	tmp := []rune(word)
	for ; Wordptr1 < len(tmp); Wordptr1, Wordptr2 = Wordptr1+1, Wordptr2+1 {
		if node, ok := Nodeptr.Child[tmp[Wordptr1]]; ok {
			//有以Wordptr1指向字符开头的敏感词
			if node.isEnd {
				//单个Wordptr1指向的字符就是敏感词
				Has = true
				word = strings.Replace(word, string(tmp[Wordptr1]), ReplaceString, 1)
				continue
			} else {
				//继续判断后面是否有敏感词
				for Wordptr2 < len(tmp) && ok {
					//判断是否为其他字符使敏感词非敏感化
					if isInvalidChar(tmp[Wordptr2]) {
						Wordptr2++
						continue
					}
					node, ok = node.Child[tmp[Wordptr2]]
					if ok {
						Wordptr2++
						if node.isEnd {
							//找到敏感词、替换
							Has = true
							word = strings.Replace(word, string(tmp[Wordptr1:Wordptr2]), ReplaceString, 1)
							break
						}
					}
				}
				Wordptr1 = Wordptr2 - 1
			}
		}
	}
	return word, Has
}

func isInvalidChar(char rune) bool {
	_, ok := InvalidChars[char]
	return ok
}

func HasInvalidChar(str string) bool {
	for _, v := range str {
		if _, ok := InvalidChars[v]; ok {
			return true
		}
	}
	return false
}

func init() {
	//初始化非法字符map
	InvalidChars = make(map[rune]struct{})
	for _, v := range InvalidChar {
		InvalidChars[v] = struct{}{}
	}
	//初始化Trie树
	T = NewTrie()
}
