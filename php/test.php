<?php


function sum($a, $b) {
    // write code here
    return;
}

class SumTest extends \PHPUnit\Framework\TestCase {
    public function testSum() {
        $tests = [
            [
                'name' => 'test 1',
                'num1' => 1,
                'num2' => 2,
                'want' => 3,
            ],
            [
                'name' => 'test 2',
                'num1' => 1,
                'num2' => 5,
                'want' => 6,
            ],
        ];
        foreach ($tests as $test) {
            $name = $test['name'];
            $num1 = $test['num1'];
            $num2 = $test['num2'];
            $want = $test['want'];
            $this->assertEquals($want, sum($num1, $num2), "Test case $name failed");
        }
    }
}


// Question:
// Given an array of integers nums and an integer target,
// return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution,
// and you may not use the same element twice.

// Example:
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].


function twoSum($nums, $target) {
    // write code here
    return null;
}

// class TwoSumTest extends \PHPUnit\Framework\TestCase {
//     public function testTwoSum() {
//         $tests = [
//             [
//                 'name' => 'test 1',
//                 'in' => [2, 7, 11, 15],
//                 'target' => 9,
//                 'out' => [0, 1],
//             ],
//         ];
//         foreach ($tests as $test) {
//             $name = $test['name'];
//             $in = $test['in'];
//             $target = $test['target'];
//             $out = $test['out'];
//             $this->assertEquals($out, twoSum($in, $target), "Test case $name failed");
//         }
//     }
// }
