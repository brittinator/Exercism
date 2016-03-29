class Complement
  VERSION = 3
  COMPLEMENTS = {"C" => "G", "A" => "U", "G" => "C", "T" => "A"}
  def self.of_dna(nucleotide)
    complement_strand = ""
    nucleotide.each_char do |base|
      if COMPLEMENTS.include?(base) == false
        raise ArgumentError, "Must be a valid strand with valid bases."
        exit
      else
        complement_strand += COMPLEMENTS[base]
      end
    end
    return complement_strand
  end
end
