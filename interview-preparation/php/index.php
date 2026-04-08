<?php

require_once __DIR__ . '/vendor/autoload.php';

use App\Greeting;

echo "--- PHP Interview Preparation --- \n\n";

$greeting = new Greeting();
echo $greeting->sayHello('Nandes') . "\n\n";

echo "You are ready to start coding!\n";
