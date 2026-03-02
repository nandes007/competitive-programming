package nandestech;

// Case 1: num = 28, return "AB"
//  -> num = 28 (before decrement),
//  -> num = 27 (after decrement)
//  -> num % 26 = 27 % 26 = 1
//  -> letter = 'B' (the character 'A' represents numeric or have numeric value 65, so: 65 + 1 = 66, whis is is alphabet ascii character it is 'B')
//  -> and then the last step is revers it because the code accepts that it builds the string backwards.
public class GetExcelColumnTitle {
    public static String getColumnTitle(int num) {
        if (num <= 0) {
            throw new IllegalArgumentException("Index Error: " + num);
        }

        StringBuilder result = new StringBuilder();

        while (num > 0) {
            num--;
            char letter = (char) ('A' + (num % 26));
            System.err.println("letter: " + letter + ", calculation: " + ('A' + (num % 26)));

            result.append(letter);

            num /= 26;
        }

        return result.reverse().toString();
    }
}
