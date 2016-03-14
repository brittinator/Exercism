public class Hello { //creating public class
    private int skip;
    // constructor
    private Hello(int skip) {
        // System.out.println(skip);
        this.skip = skip; // this.skip referring to private skip
    }

    public void loop(int number) {
        // static and non-static is the same as class and instance
        for(int i = 0; i <= number; i += skip) {
            System.out.println(i);
        }
    }


    public static void main(String[] args) { //creates method, static is class method
        // main is the meth that's run when you init Hello class
        // should go at the end
        // void returns nothing
        // String[] args, bracket is an array of strings
        // method only takes 1 type, no overloading
        // overloading is making 2 versions of method that takes diff types or num or args
        System.out.println("Hello, World!");
        new Hello(2).loop(10);
        new Hello(8).loop(10);

    }
}

