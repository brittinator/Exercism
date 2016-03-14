import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

public class Anagram {
    private String mWord;
    private HashMap<Character, Integer> mCount;

    public Anagram(String word) {
        mWord = word.toLowerCase();
        mCount = countLetters(mWord);
    }

    public ArrayList<String> match(List<String> list) {
        // compare every string to mWord
        // by checking characters
        ArrayList<String> listMatches = new ArrayList<String>();

        for(String word : list) {
            if(isAnagram(word)) {
                listMatches.add(word);
            }
        }

        return listMatches;
    }

    private boolean isAnagram(String word) {
        word = word.toLowerCase();
        HashMap<Character, Integer> wordHash = countLetters(word);
        System.out.println(word);
        System.out.println(mWord);
        if( word.equals(mWord) ) {
            return false;
        }
        return wordHash.equals(mCount);
    }

    private HashMap<Character, Integer> countLetters(String word) {
        HashMap<Character, Integer> letters = new HashMap<Character, Integer>();

        for( char letter : word.toCharArray() ) {
            if( letters.containsKey(letter) ) {
                letters.put(letter, letters.get(letter) + 1);
            } else {
                letters.put(letter, 1);
            }
        }
        return letters;
    }
}
