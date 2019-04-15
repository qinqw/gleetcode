
class Solution:
    def twoSum(self, nums: [int], target: int) -> [int]:
        print(nums)
        print(target)
        for i in range(len(nums)):
            for j in range(len(nums)):
                if (target == nums[i]+nums[j]) & (nums[i]!=target/2):
                    return [i,j]

s = Solution()
print(s.twoSum([3,2,4],6))