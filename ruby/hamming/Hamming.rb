class Hamming
  VERSION = 1
  def self.compute(strandA, strandB)
    return 0 if strandA == strandB
    raise ArgumentError, "Strand lengths must be the same." if strandA.length != strandB.length
    (0..strandA.length).count do |i|
      strandA[i] != strandB[i]
    end
  end
end
