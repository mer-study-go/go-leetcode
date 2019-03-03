# [Valid Palindrome](https://leetcode.com/problems/valid-palindrome/)

## Description

Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

**Note**: For the purpose of this problem, we define empty string as valid palindrome.

**Example 1:**

```
Input: "A man, a plan, a canal: Panama"
Output: true
```

**Example 2:**

```
Input: "race a car"
Output: false
```

## Hint
Two Pointers

## Example Questions Candidate Might Ask
Q: What about an empty string? Is it a valid palindrome? 
A: For the purpose of this problem, we define string as valid palindrome. 

## Solution 
**O(n) runtime, O(1) space**

The idea is simple, have two pointers -- one at the head while the other one at the tail. 
Move them towards each other until they meet while skipping non-alphanumeric characters. 

Consider the case where given string contains only non-alphanumeric characters. This is a valid palindrome bacause the empty string is also a valid palindrome. 

**Go Solution**

```go
func isPalindrome(s string) bool {
    if len(s) == 0 {
        return true
    }
    s = strings.ToLower(s)
    head := 0
    tail := len(s) - 1
    for head < tail {
        for head < tail && !isChar(s[head]) {
            head++
        } 
        for head < tail && !isChar(s[tail]) {
            tail--
        }
        if (s[head] != s[tail]) {
            return false;
        }
        head++
        tail--
    }
    return true
}

func isChar(c byte) bool {
    if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
        return true
    }
    return false
}
```

**Java Solution**

```java
class Solution {
    public boolean isPalindrome(String s) {
        if (s.length() == 0 || s == null) {
            return true;
        }
        int head = 0, tail = s.length() - 1;
        while (head < tail) {
            while (head < tail && !Character.isLetterOrDigit(s.charAt(head))) {
                head++;
            }
            while (head < tail && !Character.isLetterOrDigit(s.charAt(tail))) {
                tail--;
            }
            if (Character.toLowerCase(s.charAt(head)) != Character.toLowerCase(s.charAt(tail))) {
                return false;
            }
            head++;
            tail--;
        }
        return true;
    }
}

```