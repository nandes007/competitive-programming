<?php

namespace App;

class Greeting
{
    /**
     * Returns a personalized greeting message.
     *
     * @param string $name
     * @return string
     */
    public function sayHello(string $name): string
    {
        return "Hello, {$name}! Good luck with your upcoming interview!";
    }
}
