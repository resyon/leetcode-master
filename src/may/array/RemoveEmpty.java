package may.array;


import common.ListNode;
public class RemoveEmpty {
    public ListNode removeElements(ListNode head, int val) {
        while(head != null && head.val == val){
            head = head.next;
        }
        var rtn = head;
        while(head != null && head.next != null){
            if(head.next.val == val ){
                head.next = head.next.next;
            }
            head = head.next;
        }
        return rtn;
    }
    public static void main(String[] args) {
        var ls = ListNode.stringToListNode("[1,2,6,3,4,5,6]");
        var t = new RemoveEmpty();
        var ans = t.removeElements(ls, 6);
        System.out.println(ListNode.listNodeToString(ans));
    }
}
