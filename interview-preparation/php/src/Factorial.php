<?php

namespace App;

class Factorial
{
    public function calculate(int $number): int
    {
        if ($number <= 0) {
            return 1;
        }

        $result = 1;
        for ($i = $number; $i >= 1; $i--) {
            $result *= $i;
        }

        return $result;
    }

    public function calculateRecursive(int $value): int
    {
        if ($value <= 0) return 1;

        return $value * $this->calculateRecursive($value - 1);
    }
}