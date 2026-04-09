<?php

namespace App;

class LoanInstallmentCalculator
{
    public function calculateInstallments($amount, $tenor, $interestRate) {
        $monthlyInstallmentWithoutInterest = $amount / $tenor;
        $monthlyInterest = ($amount * $interestRate) / 100 / $tenor;
        $totalMonthlyInstallment = $monthlyInstallmentWithoutInterest + $monthlyInterest;
        $monthlyInstallments = [];

        for ($i = 1; $i <= $tenor; $i++) {
            $monthlyInstallments[] = [
                'month' => $i,
                'installment' => number_format($totalMonthlyInstallment, 2)
            ];
        }

        foreach ($monthlyInstallments as $installment) {
            echo "Month " . $installment['month'] . ": " . $installment['installment'] . "\n";
        }

        // echo "Monthly Installment without Interest: " . number_format($totalMonthlyInstallment, 2) . "\n";
    }
}