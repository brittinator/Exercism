class Raindrops
  VERSION = 1
  
  def self.convert(num)
    drops = ""
    if num%105==0 # 3,5,7
      return "PlingPlangPlong"
    elsif num %15==0 #3,5
      return "PlingPlang"
    elsif num%21==0 #3,7
      return "PlingPlong"
    elsif num%35==0 #5,7
      return "PlangPlong"
    elsif num %3==0
      return "Pling"
    elsif num % 5 == 0
      return 'Plang'
    elsif num % 7 == 0
      return 'Plong'
    else
      return num.to_s
    end
  end
end
