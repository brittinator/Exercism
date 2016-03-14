import java.util.List;
import java.util.ArrayList;

public class CustomerSupportRep {
    private String mName;
    private List<Customer> mAssistedCustomers;

    public CustomerSupportRep(String name) {
        mName = name;
        mAssistedCustomers = new ArrayList<Customer>();
    }

    public void assist(Customer customer) {
        System.out.printf("Hello %s, my name is %s, how can I assist you.%n",
                customer.getName(),
                mName);
        System.out.println("...");
        System.out.println("Is there anything else I can help you with?");
        mAssistedCustomers.add(customer);
    }

    public List<Customer> getAssistedCustomers() {
        return mAssistedCustomers;
    }

}