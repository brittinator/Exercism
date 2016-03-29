class HelloWorld
  def self.hello(name=nil)
    if name.nil?
      return 'Hello, World!'
    else
      return "Hello, #{name}!"
    end
  end


end
