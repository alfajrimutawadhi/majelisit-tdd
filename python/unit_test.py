import unittest

# Function to be tested
def sum(a, b):
    # write code here
    return


# Unit test class
class TestSumFunction(unittest.TestCase):
    
    # Test case 1: Addition of two positive numbers
    def test_sum_positive_numbers(self):
        result = sum(3, 5)
        self.assertEqual(result, 8)  # Assert that the result is as expected
    
    # Test case 2: Addition of a positive and a negative number
    def test_sum_positive_and_negative_numbers(self):
        result = sum(10, -7)
        self.assertEqual(result, 3)  # Assert that the result is as expected

# Execute the tests if the script is run directly
if __name__ == '__main__':
    unittest.main()




# Question:
# Given an array of integers nums and an integer target,
# return indices of the two numbers such that they add up to target.
# You may assume that each input would have exactly one solution,
# and you may not use the same element twice.
    
# Example:
# Input: nums = [2,7,11,15], target = 9
# Output: [0,1]
# Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
    
    
def two_sum(nums, target):
    # write code here
    return None

class TestTwoSumFunction(unittest.TestCase):
    def test_two_sum(self):
        tests = [
            {"name": "test 1", "nums": [2, 7, 11, 15], "target": 9, "expected": [0, 1]},
        ]
        for test in tests:
            with self.subTest(name=test["name"]):
                result = two_sum(test["nums"], test["target"])
                self.assertEqual(result, test["expected"], f"Failed; got: {result}, expected: {test['expected']}")

if __name__ == '__main__':
    unittest.main()