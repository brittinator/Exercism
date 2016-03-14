/**
 * Created by brittlwalentin on 11/6/15.
 */
public class Bob {

    public String hey(String statement) {
        String response;

        if( statement.equals(statement.toUpperCase()) && hasLetters(statement) ) {
            response = "Whoa, chill out!";
        } else if(statement.endsWith("?")) {
            response = "Sure.";
        } else if(statement.trim().length() == 0) {
            response = "Fine. Be that way!";
        } else {response = "Whatever."; }

        return response;
    }

    private boolean hasLetters(String statement) {
        boolean answer = false;

        for(Character letter : statement.toCharArray()) {
            if(Character.isAlphabetic(letter)) {
                answer = true;
            }
        }
        return answer;
    }
}
