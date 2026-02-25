package nandestech;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

import java.util.Arrays;

import org.junit.jupiter.api.Test;

public class SortOutTheManFromBoysTest {
    @Test
    void testMenFromBoys() {
        System.out.println(Arrays.toString(SortOutTheManFromBoys.menFromBoys(new int[]{-282, -282, 818, 900, 928, 281, 49, -1}))); // [14, 17, 7, 3]
        // assertArrayEquals(new int[]{14, 17, 7, 3}, SortOutTheManFromBoys.menFromBoys(new int[]{7, 3, 14, 17}));
    }
}
