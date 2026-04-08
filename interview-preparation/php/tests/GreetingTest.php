<?php

use PHPUnit\Framework\TestCase;
use App\Greeting;

class GreetingTest extends TestCase
{
    public function testSayHelloReturnsCorrectMessage()
    {
        $greeting = new Greeting();
        $result = $greeting->sayHello('TestUser');
        
        $this->assertEquals('Hello, TestUser! Good luck with your upcoming interview!', $result);
    }
}
