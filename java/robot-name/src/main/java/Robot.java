import java.util.ArrayList;
import java.util.List;
import java.util.Random;

public class Robot {
    String mName;
    List<String> names = new ArrayList<String>();

    public Robot() {
        generateName();
    }

    public void generateName() {
        // prereqs: 2 letters followed by 3 numbers
        // check against other names
        Random random = new Random();
        char a = (char)(random.nextInt(26) + 'a');
        char b = (char)(random.nextInt(26) + 'a');

        Integer number = random.nextInt(999);
        mName = Character.toString(a).toUpperCase() + Character.toString(b).toUpperCase() + Integer.toString(number);
        addNameToList(mName);
    }

    private void addNameToList(String name) {
        names.add(name);
    }

    public String getName() {
        return mName;
    }

    public void reset() {
        generateName();
    }

    public static void main(String[] args) {
        Robot robot = new Robot();

    }

}
