// Write a function solution that, given two integers A and B, returns a string containing exactly A letters 'a' and exactly B letters 'b' with no three consecutive letters being the same (in other words, neither "aaa" nor "bbb" may occur in the returned string).

// Examples:

// 1. Given A = 5 and B = 3, your function may return "aabaabab". Note that "abaabbaa" would also be a correct answer. Your function may return any correct answer.

// 2. Given A = 3 and B = 3, your function should return "ababab", "aababb", "abaabb" or any of several other strings.

// 3. Given A = 1 and B = 4, your function should return "bbabb", which is the only correct answer in this case.

// Assume that:

// A and B are integers within the range [0..100];
// at least one solution exists for the given A and B.
// In your solution, focus on correctness. The performance of your solution will not be the focus of the assessment.

// you can also use imports, for example:
// import java.util.*;

// you can write to stdout for debugging purposes, e.g.
// System.out.println("this is a debug message");

class Solution {
    public String solution(int A, int B) {
        // Implement your solution here
        String s = "";
        while (A > 0 || B > 0) {
            if (A > B) {
                if (isLast2CharsAreSame(s, 'a')) {
                    B--;
                    s += "b";
                } else {
                    A--;
                    s += "a";
                }
            } else {
                if (isLast2CharsAreSame(s, 'b')) {
                    A--;
                    s += "a";
                } else {
                    B--;
                    s += "b";
                }
            }
        }

        return s;
    }

    private boolean isLast2CharsAreSame(String s, char c) {
        int len = s.length();
        if (len < 2) {
            return false;
        }

        return s.charAt(len-1) == s.charAt(len-2) && s.charAt(len-1) == c;
    }
}
