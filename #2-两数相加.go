/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // 我学到了在用new和&创建结构体时，go语言会直接实例化并返回对应的地址，因此要用指针接收赋值
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var p1, p2, pa *ListNode;
    var ans *ListNode = &ListNode{0,nil}
    pa = ans;
    var upper int = 0;
    p1 = l1;
    p2 = l2;
    for {
        if p1 != nil && p2 != nil {
            pa.Val = p1.Val + p2.Val;
            p1 = p1.Next;
            p2 = p2.Next;
        } else if p2 == nil && p1 != nil {
            pa.Val = p1.Val;
            p1 = p1.Next
        } else if p1 == nil && p2 != nil{
            pa.Val = p2.Val;
            p2 = p2.Next
        }
        if upper == 1 {
            pa.Val += 1
            upper = 0;
        }
        if pa.Val >= 10 {
            pa.Val -= 10;
            upper = 1;
        }
        if p1 == nil && p2 == nil {
            if upper == 1 {
                pa.Next = &ListNode{1,nil};
            }
            break;
        }
        pa.Next = &ListNode{0,nil};
        pa = pa.Next;
    }
    return ans;
}
