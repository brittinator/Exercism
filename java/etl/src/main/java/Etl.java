import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Set;

public class Etl {

   public Map<String, Integer> transform(Map<Integer, List<String>> old) {
      Map<String, Integer> newVals = new HashMap<>();

      Set<Map.Entry<Integer, List<String>>> entries = old.entrySet();

      // all the alpha and values in the map, of which old is an instance of
      for(Map.Entry<Integer, List<String>> entry : entries ) {
         // for each value of an entry 'a', 'b', etc
         // assign the key to the letter, the value to the points
         for(String letter : entry.getValue()) {
            newVals.put(letter.toLowerCase(), entry.getKey());
         }
      }

      return newVals;
   }
}
