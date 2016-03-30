class Gigasecond
  VERSION = 1
  GIGASECOND = 10**9
  def self.from(birth)
    giga = GIGASECOND + birth.to_i
    Time.at(giga)
  end
end
