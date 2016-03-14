import java.nio.channels.FileChannel;
import java.util.HashMap;
import java.util.Map;

public class DNA {
     enum Bases {"G", "A", "T", "C" };
    public String mNuc;

    public DNA(String nuc) {
        mNuc = nuc;
    }

    public int count(char base) {
        int baseCount = 0;
        // find that base in the DNA string
        // for 1 base
        for(int i=0; i < mNuc.length(); i++) {
            if(mNuc.charAt(i) ==  base) {
                // add to count
                baseCount++;
            }
        }

        if(Bases.valueOf(base) == null) {
            throw new IllegalArgumentException("That is not a valid base.");
        }

        }

        return baseCount;
    }

    public Map<Character, Integer> nucleotideCounts() {
        Map<Character, Integer> totalCount = new HashMap<>();

        totalCount.put('A', count('A'));
        totalCount.put('T', count('T'));
        totalCount.put('G', count('G'));
        totalCount.put('C', count('C'));

        return totalCount;
    }

}
