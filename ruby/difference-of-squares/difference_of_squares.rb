class Squares
  VERSION = 2
  def initialize(n)
    @n = n
  end

  def square_of_sum
    return 0 if @n == 0
    sum = (1..@n).reduce(:+)
    sum**2
  end

  def difference
    square_sum = square_of_sum
    sum_square = sum_of_squares

    square_sum - sum_square
  end
  def sum_of_squares
    return 0 if @n == 0
    count = 0
    @n.times do |num|
      count += (num+1)*(num+1)
    end
    return count
  end

end
