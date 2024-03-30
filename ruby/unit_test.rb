require 'test/unit'

# Function to be tested
def sum(a, b)
    # write code here
    return
end

class TestSum < Test::Unit::TestCase
    def test_sum
      # Test case 1: Addition of two positive numbers
      result = sum(3, 5)
      expected_result = 8
      assert_equal(expected_result, result, "Test case 1 failed")
  
      # Test case 2: Addition of a positive and a negative number
      result = sum(10, -7)
      expected_result = 3
      assert_equal(expected_result, result, "Test case 2 failed")
  
      # Test case 3: Addition of two negative numbers
      result = sum(-5, -3)
      expected_result = -8
      assert_equal(expected_result, result, "Test case 3 failed")
    end
end


# Question:
# Given an array of integers nums and an integer target,
# return indices of the two numbers such that they add up to target.
# You may assume that each input would have exactly one solution,
# and you may not use the same element twice.
    
# Example:
# Input: nums = [2,7,11,15], target = 9
# Output: [0,1]
# Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].


def two_sum(nums, target)
  # write code here
  nil
end

class TestTwoSum < Test::Unit::TestCase
  def test_two_sum
    tests = [
      { name: "test 1", nums: [2, 7, 11, 15], target: 9, expected: [0, 1] },
    ]
    tests.each do |test|
      result = two_sum(test[:nums], test[:target])
      assert_equal(test[:expected], result, "Failed; got: #{result}, expected: #{test[:expected]}")
    end
  end
end
