package nandestech;

import java.util.ArrayList;
import java.util.List;

public class SortOutTheManFromBoys {
    public static int[] menFromBoys(final int[] values) {
        List<Integer> evens = new ArrayList<>();
        List<Integer> odds = new ArrayList<>();
        List<Integer> results = new ArrayList<>();
        for (int i = 0; i < values.length; i++) {
            if (values[i] % 2 == 0 && evens.indexOf(values[i]) == -1) {
                evens.add(values[i]);
            }

            if (values[i] % 2 != 0 && odds.indexOf(values[i]) == -1) {
                odds.add(values[i]);
            }
        }

        evens.sort((a, b) -> a - b);
        odds.sort((a, b) -> b - a);

        results.addAll(evens);
        results.addAll(odds);

        return results.stream().mapToInt(i -> i).toArray();
    }
}