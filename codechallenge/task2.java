// You parked your car in a parking lot and want to compute the total cost of the ticket. The billing rules are as follows:

// The entrance fee of the car parking lot is 2;
// The first full or partial hour costs 3;
// Each successive full or partial hour (after the first) costs 4.
// You entered the car parking lot at time E and left at time L. In this task, times are represented as strings in the format "HH:MM" (where "HH" is a two-digit number between 0 and 23, which stands for hours, and "MM" is a two-digit number between 0 and 59, which stands for minutes).

// Write a function:

// class Solution { public int solution(String E, String L); }

// that, given strings E and L specifying points in time in the format "HH:MM", returns the total cost of the parking bill from your entry at time E to your exit at time L. You can assume that E describes a time before L on the same day.

// For example, given "10:00" and "13:21" your function should return 17, because the entrance fee equals 2, the first hour costs 3 and there are two more full hours and part of a further hour, so the total cost is 2 + 3 + (3 * 4) = 17. Given "09:42" and "11:42" your function should return 9, because the entrance fee equals 2, the first hour costs 3 and the second hour costs 4, so the total cost is 2 + 3 + 4 = 9.

// Assume that:

// strings E and L follow the format "HH:MM" strictly;
// string E describes a time before L on the same day.
// In your solution, focus on correctness. The performance of your solution will not be the focus of the assessment.

//==================================================
// you can also use imports, for example:
// import java.util.*;

// you can write to stdout for debugging purposes, e.g.
// System.out.println("this is a debug message");

class Solution {
    public int solution(String E, String L) {
        // Implement your solution here
        int diff = roundUp(diffHours(E, L));
        int cost = 2;
        for (int i = 1; i <= diff; i++) {
            if (i == 1) {
                cost += 3;
            } else {
                cost += 4;
            }
        }
        return cost;
    }

    private int roundUp(double num) {
       int n = (int)(num*100);
       int mod = n%100; 
       return mod > 0 ? (n/100+1):n/100;
    }
    private double diffHours(String E, String L) {
        int hh1 = parseHour(E);
        int hh2 = parseHour(L);
        int mm1 = parseMinute(E);
        int mm2 = parseMinute(L);

        double d1 = hh1 + mm1*1.0/100;
        double d2 = hh2 + mm2*1.0/100;
        return d2-d1;
    }

    private int parseHour(String timeStr) {
        String hh = timeStr.split(":")[0];
        return Integer.parseInt(hh);
    }

    private int parseMinute(String timeStr) {
        String mm = timeStr.split(":")[1];
        return Integer.parseInt(mm);
    }
}
