

/**
 * Created by brittlwalentin on 3/21/16.
 */
public class Pangrams {


    public static boolean isPangram(String str) {
        boolean isPangram = true; // until proven false

        if (str.length() < 26)
        { //26 letters in alphabet
            return false;
        }

        str = str.toLowerCase();

        for(char c='a'; c <='z'; c++) {
            if(str.indexOf(c) == -1) {
                isPangram = false;
            }
        }
        return isPangram;
    }

}
