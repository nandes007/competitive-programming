package nandestech;

// TODO:: need explaination of this one more...
public class GetExcelColumnTitle {
    public static String getColumnTitle(int num) {
        if (num <= 0) {
            throw new IllegalArgumentException("Index Error: " + num);
        }

        StringBuilder result = new StringBuilder();

        while (num > 0) {
            num--;
            System.out.println("num = " + num);

            // System.out.println(num % 26);
            System.out.println("Mod result: " + ('A' + (num % 26)));
            char letter = (char) ('A' + (num % 26));
            System.out.println("latter " + letter);
            System.out.println("num after modulo " + num);

            result.append(letter);

            num /= 26;
            System.out.println("num after division " + num);
        }

        return result.reverse().toString();
    }
}
