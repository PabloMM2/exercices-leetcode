class Solution:
    def fizzBuzz(self, n: int) -> List[str]:
        response : List[str] = []

        for i in range(1, n +1):
          if i % 5 == 0 and i % 3 == 0:
            response.append("FizzBuzz")
          elif i % 5 == 0:
            response.append("Buzz")
          elif i%3 == 0:
            response.append("Fizz")
          else:
            response.append(str(i))

        return response
        