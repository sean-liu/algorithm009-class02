import java.util.Deque;
import java.util.LinkedList;

public class deque_old_codes_converting {
    public static void main(final String[] args) {
        final Deque<String> deque = new LinkedList<String>();
        deque.addFirst("a");
        deque.addFirst("b");
        deque.addFirst("c");
        System.out.println(deque);

        final String str = deque.removeFirst();
        System.out.println(str);
        System.out.println(deque);
    
        while (deque.size() > 0) {
            System.out.println(deque.removeFirst());
        }
    
        System.out.println(deque); 
    }
}