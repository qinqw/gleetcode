package main

import "fmt"

func main() {

	fmt.Println(isValid("abcabc"))
	fmt.Println(isValid("ababcc"))
	fmt.Println(isValid("aabcbc"))
	fmt.Println(isValid("abc"))
	fmt.Println(isValid("aabcbabcc"))

	fmt.Println(isValid("aaabcbabcc"))

	fmt.Println(isValid("cbbbbcaaaccbaabccabacabcacbabbabbcbacabacabcabbccbccbccbcaccbacaabaaababacbbccabbcbcbabbcbabcacabcbababcaccbcbbbaaabbcbccaaabcbababaaaccbacbbcaabbacbbbabcccacbaacacccbcccbbbbccaaabaaababcbbaacbcccaaabcbacccaacbbbaabababbbbababcbaaaabbcacbacaaabbcacbccbbbbaabcacbccababaaaacbabbbaaccbccaacacccaaabacaaaacacbbaabcabcbcccbcabaaaccbcbaaaccaacbabcaabbccacbbcbbbcaccbaccacaabcbcccacacaccaabbcaaababbbbbbcabcbbacabbbbbaabcbabbbccaabccabccaaaccabcbcbabbcaccabcbaabaaabacccacacabbaccbabbbacaaabacacabacbcabaabaccacbbaacbcaccaccbacbaacbcbabccbcacbbbcabaabaabbabbcbbbcbbaabababcacacacbcabbccbbbacacccbbbbbbcbabaacacbbbaccbcaaccaccabbabbcbaababbcacaababaababcccbbcaaaabccabbbcbbaacacabaabbaabcaabaacaabcbaaabaccbbcbaaababcacaacabcbbbbabaccbabbccacbbbcbacccbbbcabaaaacacaacbbcabbbbaacacababccaaacaccaaaacbabbcacaabaacaaabcacaabcaaaaaaacbbabbababcbbccbabbabaccbbcccacaabcbccacaacbacbcccbabababababbacbabacbcccccbbcbcbabbbbaacbbabacccbcccbcaccabacaacbbbabacccccbcbbbacaccacccbccbbbcbcbaacbbaabaacbacabacacaaaccaaacabbbaabcaacacbaaabcbbbbcaacaaccbbbccbababcbcacbcbcbbaabcacbbcacaacacccababbcacbbcaaccaacbbabccabbccbacacaccccccbcbbcbabcabbcaabbaabaccababcbcbcbabbcabcbabacaababcbcacbaacbcccbbbcbaacabcccbbabccacaaaaccaabccbbaabababcacabbaccbbacbcbcaacabbaccccabbccbbcabcbcccaabcccbabbacabccaacababcaaaacccbabcccabcccabccbacaacabcbbaacbaacbbcbaaaacbbcaccbacaccaaaccabaaabcbabababbcbbcbacbabaacbbbaacabbbaacbaccccaaccccabcabaabaccacaaacabbbbcbccccbcacaacbcbbbcbabcbccaaabbaccbaacccbaacbbbbcccacbcacabccbbcbabbccabbbaabaacbbacacacccbcbcbaabcccacbccbbcacaababbcbaaabaccaaaaacbbbacaaaaabbacaabacbcaccaacbacbaaaabcbaaabcccccbcacbccacbccccaacaaccbbcbabcbaacbbabbbbaabcacbaaacbbcabcbcbbbcabbbbcbcbacccbcaacaccccbcccacbbbbbacbccaabcbabcacacbbaccbbbaccbcaacacabaaccbccabbccbacbaacbbbbacacccabbbbabaaaabcbabcccacccababbabcabbbaabacbcbaaccccacbbbccccacaaaacabaacaccaabccacbbbccaabbccccbccbbbabacbabacbabcbaccbabbaccaccbabcababbbbcbbcabababaacaccacbacccabcacbbbcccacacabcabccaacbacbcaccbaaaacbbccccccaabbabcbaaaacbbbaabbcacabaccabbcbbbbabbccabcccccbbcbbcccbbacacbaaacccccbbcaacaacabbbcacaaaaaacabacbbbbbacbbacbbcbbbcbbaccaacbcbaaaabbbababbbccccaabacbcaacacbbcacccaacccccbbccbbbbbacbabcabbcacccbcbabbaabbacaabaacabcccbbbbcacabbbaacaaacbabbaccaacabbcbbaaabaccaccacccbbbacaccacbabbaacabbcbbcbbccbbbacbbacbbbabcaababaacbaacaabbacbabbccccaccccaabbaaabcccabccbccaaacbbccabbbcbaaaacabbbbbbcbbacbbcccbbabbacbbccbccccccaabacaabaaaccccaabcacbbccbbabccaacccaababbccbabcbabcccbbabaaaababaabbbcbabbccaacbacccaabccababccabcbbcccaacabbbcbaaabbcaacbbbaabaacbbacbbcacbcabccbbccbcbcbbcbaacacbbbbcacbabbaacccaccbacbbcabbbbabcabaabbcaacabbbbbbccbacacaaacabaacbcbacbaaacabcbbcbbbabaccabaaaacaaaaaacaccbccaabaabcaaabaaababcabcccbbcacbcacaabbaaacbbbbcbabaabaabacacaaccbaacabcacbaaabbacccbccaabbaccabbbbacbcbccbcbbcbcbcbcabcccbccaacacccccbbbabbabcbcbabbacaabcbaabbacbcaccacacccacbabbabacababcbaccaabcbacabbacaaaabccbbbbcaabbccbcccccbccacaababcccacbcbcbccabbcbcbcccccccacaaccbbccccbbbcabbacaacacbbabcaacacbcccacbcbcbaaaaccaabaaabcabcaacaacabaacccabaabbaac"))

}

func isValid(S string) bool {
	res := false
	S = cutS(S)
	if len(S) == 0 {
		res = true
	}
	return res
}

func cutS(S string) string {
	//res := S
	//fmt.Println("0 ", S)
	abcIndex := []int{}
	for i := 0; i < len(S)-2; i++ {
		if S[i]+1 == S[i+1] && S[i]+2 == S[i+2] && S[i] == "a"[0] {
			abcIndex = append(abcIndex, i)
		}
	}
	for j := len(abcIndex) - 1; j >= 0; j-- {
		//fmt.Println("1 ", S)
		S = S[:abcIndex[j]] + S[abcIndex[j]+3:]
		//fmt.Println("2 ", S)
	}

	if len(S) < 3 {
		return S
	}

	if len(abcIndex) > 0 {
		return cutS(S)
	}
	return S
}
