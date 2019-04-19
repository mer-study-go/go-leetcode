# [Unique Email Addresses](https://leetcode.com/problems/unique-email-addresses/)

## Description

Every email consists of a local name and a domain name, separated by the @ sign.

For example, in `alice@leetcode.com`, `alice` is the local name, and `leetcode.com` is the domain name.

Besides lowercase letters, these emails may contain `.`s or `+`s.

If you add periods(`.`) between some characters in the **local name** part of an email address, mail sent there will be forwarded to the same address without dots in the local name. For example, `alice.z@leetcode.com` and `alicez@leetcode.com` forward to the same email address. (Note that this rule does not apply for domain names.)

If you add a plus (`+`) in the **local name**, everything after the first plus sign will be **ignored**. This allows certain emails to be filtered, for example `m.y.+name@email.com` will be forwarded to `my@email.com`. (Again, this rule does not apply for domain names.)

It is possible to use both of these rules at the same time. 

Given a list of `emails`, we send one email to each address in the list. How many different addresses actually addresses actually recieve mails? 

**Example 1:**

```
Input: ["test.mail+ale@leetcode.com", "test.email+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"]
Output: 2
Explanation: "testemail@leetcode.com" and "testemail@lee.tcode.com" actually receive mails
```

**Note**:

* `1 <= email[i].length <= 100`
* `1 <= emails.length <= 100`
* Each `email[i]` contains exactly one `@` character. 
* All local and domain names are non-empty.
* Local names do not start with a `+` character. 

## Hint

Hash Set

## Solution In Java

```java
class Solution {
    public int numUniqueEmails(String[] emails) {
        HashSet set = new HashSet();
        for (int i = 0; i < emails.length; i++) {
            System.out.println(transform(emails[i]));
            if (!set.contains(transform(emails[i]))) {

                set.add(transform(emails[i]));
            }
        }
        return set.size();
    }

    public String transform(String email) {

        String local = email.substring(0, email.indexOf('@'));
        if (local.contains("+")) {
            local = local.substring(0, email.indexOf('+'));
        }

        char[] input = local.toCharArray();
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < input.length; i++) {
            if (input[i] != '.') {
                sb.append(input[i]);
            }
        }

        sb.append(email.substring(email.indexOf('@'), email.length()));
        return sb.toString();
    }
}
```

## Solution In Go

```go
func numUniqueEmails(emails []string) int {
	dict := make(map[string]int)
	for index := range emails {
		if _, ok := dict[transform(emails[index])]; !ok {
			dict[transform(emails[index])] = 1
		}
	}
	return len(dict)
}

func transform(email string) string {
	local := email[:strings.Index(email, "@")]
	if strings.Contains(email, "+") {
		local = email[:strings.Index(local, "+")]
	}
	local = strings.Replace(local, ".", "", -1)
	result := local + email[strings.Index(email, "@"):len(email)]
	return result
}
```

## Complexity Analysis

* Time Complexity:  O(C), where C is the total length of `emails`.
* Space Complexity: O(C), new set is introduced. 