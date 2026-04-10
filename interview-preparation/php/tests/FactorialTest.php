<?php

use App\Factorial;
use PHPUnit\Framework\TestCase;

class FactorialTest extends TestCase
{
    public function testCalculation()
    {
        $factorial = new Factorial();
        $this->assertEquals(120, $factorial->calculate(5));
    }

    public function testCalculationRecursive()
    {
        $factorial = new Factorial();
        $this->assertEquals(120, $factorial->calculateRecursive(5));
    }
}
