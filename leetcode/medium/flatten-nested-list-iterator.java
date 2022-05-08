/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * public interface NestedInteger {
 *
 *     // @return true if this NestedInteger holds a single integer, rather than a nested list.
 *     public boolean isInteger();
 *
 *     // @return the single integer that this NestedInteger holds, if it holds a single integer
 *     // Return null if this NestedInteger holds a nested list
 *     public Integer getInteger();
 *
 *     // @return the nested list that this NestedInteger holds, if it holds a nested list
 *     // Return empty list if this NestedInteger holds a single integer
 *     public List<NestedInteger> getList();
 * }
 */

public class NestedIterator implements Iterator<Integer> {
    
    private final Queue<Integer> list = new LinkedList();

    public NestedIterator(List<NestedInteger> nestedList) {
        for (var node : nestedList) {
            flatten(list, node);
        }
    }
    
    private void flatten(Queue<Integer> list, NestedInteger node) {
        if (node.isInteger()) {
            list.add(node.getInteger());
            return;
        }
        
        for (var child : node.getList()) {
            flatten(list, child);
        }
    }

    @Override
    public Integer next() {
        return list.poll();    
    }

    @Override
    public boolean hasNext() {
        return !list.isEmpty();
    }
}

/**
 * Your NestedIterator object will be instantiated and called as such:
 * NestedIterator i = new NestedIterator(nestedList);
 * while (i.hasNext()) v[f()] = i.next();
 */