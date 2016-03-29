class Gigasecond
  VERSION = 1
  def self.from(birth)
    giga = 10**9 + birth.to_i
    giga_time = Time.at(giga)

  end

end
