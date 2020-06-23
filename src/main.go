package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

var validChar = map[rune]rune{
	'[': '[',
	']': ']',
	'{': '{',
	'}': '}',
	'(': '(',
	')': ')',
}

var leftChar = map[rune]rune{
	'}': '{',
	']': '[',
	')': '(',
}

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	matchParentheses(input)
}

// 利用状态转换匹配括号
func matchParentheses(input string) {
	// 因int默认是0,所以0-6状态用1-7表示,s[0]数组用于计数
	var s = [8][128]int{}

	// 初始化状态转化表
	// 第二列
	for i := 1; i < len(s); i++ {
		s[i]['{'] = 2
	}
	// 第三列
	s[2]['}'] = 3
	s[3]['}'] = 3
	s[5]['}'] = 3
	s[7]['}'] = 3
	// 第四列
	for i := 1; i < len(s); i++ {
		s[i]['['] = 4
	}
	// 第五列
	s[3][']'] = 5
	s[4][']'] = 5
	s[5][']'] = 5
	s[7][']'] = 5
	// 第六列
	for i := 1; i < len(s); i++ {
		s[i]['('] = 6
	}
	// 第七列
	s[3][')'] = 7
	s[5][')'] = 7
	s[6][')'] = 7
	s[7][')'] = 7

	flag := 1            // 默认状态
	var buf bytes.Buffer // 缓冲区
	for _, v := range input[0 : len(input)-1] {
		if _, ok := validChar[v]; !ok { // 忽略其它字符
			continue
		}
		buf.WriteRune(v)

		f := s[flag][v]
		if f == 0 {
			fmt.Println("no match", buf.String())
			return
		}

		flag = f
		s[0][v]++

		if l, ok := leftChar[v]; ok {
			// 右括号不能比左括号多
			if s[0][v] > s[0][l] {
				fmt.Println("no match", buf.String())
				return
			}
		}
	}

	fmt.Printf("\"%s\":%d, \"%s\":%d\n", "{", s[0]['{'], "}", s[0]['}'])
	fmt.Printf("\"%s\":%d, \"%s\":%d\n", "[", s[0]['['], "]", s[0][']'])
	fmt.Printf("\"%s\":%d, \"%s\":%d\n", "(", s[0]['('], ")", s[0][')'])

	// 暂时用计数法判断是否成对
	if s[0]['{'] == s[0]['}'] && s[0]['['] == s[0][']'] && s[0]['('] == s[0][')'] {
		fmt.Println(buf.String())
	} else {
		fmt.Println("no match")
	}
}
