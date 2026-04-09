<?php

namespace App;

class LoanInstallmentCalculator
{
    public function calculateInstallments($amount, $tenor, $interestRate) {
        $monthlyInstallmentWithoutInterest = $amount / $tenor;
        $monthlyInterest = ($amount * $interestRate) / 100 / $tenor;
        $totalMonthlyInstallment = $monthlyInstallmentWithoutInterest + $monthlyInterest;
        $monthlyInstallments = [];

        $total = 0;
        for ($i = 1; $i <= $tenor; $i++) {
            $monthlyInstallments[] = [
                'month' => $i,
                'installment' => number_format($totalMonthlyInstallment, 2)
            ];

            $total += $totalMonthlyInstallment;
        }

        echo $total;

        // $total = 0.0;
        // foreach ($monthlyInstallments as $installment) {
        //     $total += $installment['installment'];
        // }

        return $monthlyInstallments;
    }
}