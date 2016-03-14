import java.util.HashMap;

public class WordCount {

    public HashMap<String, Integer> Phrase(String phrase) {
        HashMap<String, Integer> counts = new HashMap<String, Integer>();
        phrase = filter(phrase);

        // separate string by spaces
        String[] wordsAsArray = phrase.split(" ");
        // use for loop to iterate over words to count
        for(String word : wordsAsArray) {
            //System.out.println(word);
            if(word.equals("") ) { continue; } // use .equals on strings
            // push word into HashMap if not already
            if (counts.containsKey(word)) {
                // increment that word's count
                //int oldCount = counts.get(word); get(word) gets the value at word
                //oldCount++;
                //counts.put(word, oldCount);
                // one liner below:
                counts.put(word, counts.get(word) + 1);
            } else {
                counts.put(word, 1);
                // increment count of word
            }

        }
        return counts;
    }

    private String filter(String unfilteredString) {
        // to lowercase everything
        unfilteredString = unfilteredString.toLowerCase();
        String filtered = "";
        // filter special characters
        // loop through string to get chars
        for( Character c : unfilteredString.toCharArray() ) {
            // .isLetterOrDigit static method called on class (Character)
            if( Character.isLetterOrDigit(c) || c == ' ' ) {
                filtered += c; // concatenate c onto filtered
            }
        }
        //System.out.println(filtered);
        return filtered;
    }
}