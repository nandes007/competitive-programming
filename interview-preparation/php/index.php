<?php

require_once __DIR__ . '/vendor/autoload.php';

use App\Greeting;
use App\LoanInstallmentCalculator;

// echo "--- PHP Interview Preparation --- \n\n";

// $greeting = new Greeting();
// echo $greeting->sayHello('Nandes') . "\n\n";

// echo "You are ready to start coding!\n";

$loanInstallmentCalculator = new LoanInstallmentCalculator();
$loanInstallmentCalculator->calculateInstallments(1000000, 12, 5);