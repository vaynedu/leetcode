#!/usr/bin/env python3
"""
快速生成算法面试题相关文件

用法：
  python3 gen_interview.py --list          # 列出所有题目
  python3 gen_interview.py --new 23       # 生成新题目模板
  python3 gen_interview.py --test 23       # 运行题目测试
  python3 gen_interview.py --all           # 运行所有测试
"""

import os
import sys

QUESTIONS = {
    "23": {
        "name": "合并K个升序链表",
        "file": "23.merge-k-sorted-lists",
        "data_struct": "链表",
        "technique": "分治归并",
        "complexity": "O(N log K)"
    },
    "146": {
        "name": "LRU 缓存",
        "file": "146.lru-cache",
        "data_struct": "HashMap + 双向链表",
        "technique": "表头表尾操作",
        "complexity": "O(1)"
    },
    "56": {
        "name": "合并区间",
        "file": "56.merge-intervals",
        "data_struct": "数组/排序",
        "technique": "排序 + 贪心",
        "complexity": "O(N log N)"
    },
    "200": {
        "name": "岛屿数量",
        "file": "200.number-of-islands",
        "data_struct": "网格/DFS",
        "technique": "感染标记",
        "complexity": "O(MN)"
    },
    "102": {
        "name": "二叉树层序遍历",
        "file": "102.binary-tree-level-order-traversal",
        "data_struct": "队列",
        "technique": "size分层",
        "complexity": "O(N)"
    },
}

TEMPLATE_GO = '''package interview

// {id}. {name}
func {func_name}() {{
    // TODO: 实现代码
}}
'''

TEMPLATE_MD = '''# {id}. {name}

## 题目
TODO: 题目描述

## 难度
**频率**：🔥🔥🔥🔥🔥（面试高频）

## 一句话思路
**{technique}**

## 代码实现
```go
// TODO: 实现代码
```

## 核心考点
```
数据结构：{data_struct}
时间复杂度：{complexity}
空间复杂度：O(?)
```

## 记忆口诀
TODO: 编口诀

## 测试
```go
func Test{func_name}() {{
    // TODO: 测试用例
}}
```
'''

def list_questions():
    print("\\n" + "="*60)
    print("        算法面试高频题速查")
    print("="*60)
    print(f"{"题号":<6} {"名称":<20} {"数据结构":<20} {"技巧":<15} {"复杂度":<10}")
    print("-"*60)
    for id_, q in QUESTIONS.items():
        print(f"{id_:<6} {q['name']:<20} {q['data_struct']:<20} {q['technique']:<15} {q['complexity']:<10}")
    print("="*60 + "\\n")

def run_tests(qid=None):
    """运行测试"""
    cmd = "cd /Users/nikki/go/src/leetcode/interview && go test -v"
    if qid:
        cmd += f" -run Test{QUESTIONS[qid]['name'].split()[0]}"
    os.system(cmd)

def main():
    if len(sys.argv) < 2:
        list_questions()
        sys.exit(0)

    arg = sys.argv[1]

    if arg == "--list" or arg == "-l":
        list_questions()

    elif arg == "--test" or arg == "-t":
        if len(sys.argv) < 3:
            print("请指定题目编号，如: --test 23")
            sys.exit(1)
        run_tests(sys.argv[2])

    elif arg == "--all" or arg == "-a":
        run_tests()

    elif arg == "--new" or arg == "-n":
        if len(sys.argv) < 3:
            print("请指定题目编号，如: --new 23")
            sys.exit(1)
        qid = sys.argv[2]
        if qid not in QUESTIONS:
            print(f"未知题目: {qid}")
            sys.exit(1)
        # TODO: 生成新题目文件
        print(f"生成题目 {qid}: {QUESTIONS[qid]['name']}")

    else:
        print(f"未知参数: {arg}")
        print(__doc__)

if __name__ == "__main__":
    main()
