public class Hamming {

    public static int compute (String stringA, String stringB) throws IllegalArgumentException {
    // static method b/c called upon Hamming class itself
        int numDifferences = 0;

        if( stringA.length() != stringB.length() ) {
            throw new IllegalArgumentException("Please compare only strings of equal length");
        }
        for(int i=0; i < stringA.length(); i++) {
            if( stringA.charAt(i) != stringB.charAt(i)) {
                numDifferences++;
            }
        }

        return numDifferences;
    }
}
