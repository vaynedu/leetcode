package interview

// 127. 单词阶梯
// BFS 最短路：从 beginWord 到 endWord，每次变换一个字符
func ladderLength(beginWord string, endWord string, wordList []string) int {
    wordSet := map[string]bool{}
    for _, w := range wordList {
        wordSet[w] = true
    }
    if !wordSet[endWord] {
        return 0
    }
    delete(wordSet, beginWord)

    queue := []string{beginWord}
    step := 1
    for len(queue) > 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            word := queue[0]
            queue = queue[1:]
            // 尝试改变每个字符
            chars := []byte(word)
            for j := 0; j < len(chars); j++ {
                original := chars[j]
                for c := 'a'; c <= 'z'; c++ {
                    if byte(c) == original {
                        continue
                    }
                    chars[j] = byte(c)
                    newWord := string(chars)
                    if newWord == endWord {
                        return step + 1
                    }
                    if wordSet[newWord] {
                        delete(wordSet, newWord)
                        queue = append(queue, newWord)
                    }
                }
                chars[j] = original
            }
        }
        step++
    }
    return 0
}
