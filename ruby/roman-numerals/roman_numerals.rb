class Fixnum
  VERSION = 1
  LATIN_TO_ROMAN = {
    1000 => "M",
    900 => "CM",
    500 => "D",
    400 => "CD",
    100 => "C",
    90 => "XC",
    50 => "L",
    40 => "XL",
    10 => "X",
    9 => "IX",
    5 => "V",
    4 => "IV",
    1 => "I"
  }

  def to_roman
    latin = self
    roman_str = ""
    LATIN_TO_ROMAN.each do |num, roman|
      if latin/num > 0
        until latin < num
          roman_str += roman
          latin -= num
        end
      end
      break if latin == 0
    end
    return roman_str
  end
end
