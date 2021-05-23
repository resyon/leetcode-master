package main

type Node struct{
    Val int 
    Next *Node
    Pre *Node
}
//循环双链表
type MyLinkedList struct {
	dummy *Node
}

//仅保存哑节点，pre-> rear, next-> head 
/** Initialize your data structure here. */
func Constructor() MyLinkedList {
    rear := &Node{
        Val: -1,
        Next:nil, 
        Pre:nil,
    }
    rear.Next = rear
    rear.Pre = rear
    return MyLinkedList{rear}
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
    head := this.dummy.Next
    //head == this, 遍历完全
    for head != this.dummy && index > 0{
        index--
        head = head.Next
    }
    //否则, head == this, 索引无效
    if 0 != index {
        return -1
    }
    return head.Val
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	dummy := this.dummy
    node := &Node{
        Val: val,
        //head.Next指向原头节点
        Next: dummy.Next,
        //head.Pre 指向哑节点
        Pre: dummy,
    }

    //更新原头节点
    dummy.Next.Pre = node
    //更新哑节点
    dummy.Next = node
    //以上两步不能反
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	dummy := this.dummy
    rear := &Node{
        Val: val,
        //rear.Next = dummy(哑节点)
        Next: dummy,
        //rear.Pre = ori_rear
        Pre: dummy.Pre,
    }
    
    //ori_rear.Next = rear
    dummy.Pre.Next = rear
    //update dummy
    dummy.Pre = rear
    //以上两步不能反
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    head := this.dummy.Next
    //head = MyLinkedList[index]
    for head != this.dummy && index > 0 {
        head = head.Next
        index--
    }
    node := &Node{
        Val: val,
        //node.Next = MyLinkedList[index]
        Next: head,
        //node.Pre = MyLinkedList[index-1]
        Pre: head.Pre,
    }
    //MyLinkedList[index-1].Next = node
    head.Pre.Next = node
    //MyLinkedList[index].Pre = node
    head.Pre = node
    //以上两步不能反
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	//链表为空
    if this.dummy.Next == this.dummy{
        return
    }
    head := this.dummy.Next
    //head = MyLinkedList[index]
    for head.Next != this.dummy && index > 0 {
        head = head.Next
        index--
    }
    //验证index有效
    if(index == 0){
        //MyLinkedList[index].Pre = index[index-2]
        head.Next.Pre = head.Pre
        //MyLinedList[index-2].Next = index[index]
        head.Pre.Next = head.Next
        //以上两步顺序无所谓
    }
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

//  func main(){
// 	obj := Constructor()
// 	obj.AddAtHead(1)
// 	obj.AddAtTail(3)
// 	obj.Get(1)
// 	obj.AddAtIndex(1, 2)
// 	obj.DeleteAtIndex(1)
// 	obj.Get(1)
//  }
func main(){
	/*
	["MyLinkedList","addAtHead","addAtIndex","addAtTail","addAtTail","addAtTail","addAtIndex",
	"addAtTail","addAtHead","deleteAtIndex","deleteAtIndex","deleteAtIndex","addAtIndex","addAtTail",
	"get","get","addAtHead","addAtTail","addAtTail","get",
	"addAtTail","addAtTail","deleteAtIndex","deleteAtIndex","addAtHead","addAtTail","addAtIndex",
	"get","addAtTail","addAtIndex","addAtHead","addAtTail","addAtIndex",
	"get","addAtHead","addAtTail","addAtIndex","addAtHead","addAtIndex","addAtTail","addAtHead","addAtIndex","addAtTail","addAtHead","deleteAtIndex",
	"get","addAtIndex","get","addAtIndex","addAtTail","addAtTail",
	"get","deleteAtIndex","get","addAtHead","addAtTail","addAtIndex","addAtIndex","addAtIndex","addAtHead","addAtTail","addAtIndex","deleteAtIndex","addAtHead","addAtHead","addAtTail",
	"get","addAtTail","addAtIndex","addAtHead","deleteAtIndex","addAtHead","deleteAtIndex",
	"get","get","addAtTail","addAtIndex","get","deleteAtIndex","deleteAtIndex","addAtHead","addAtHead","addAtIndex",
	"get","addAt...
	 [[],[55],[1,90],[51],[91],[12],[2,72],[17],[82],[4],[7],[7],[5,75],[54],[6],[2],[8],[35],[36],[10],[40],[43],[12],[3],[78],[89],[3,41],[10],[96],[5,37],[51],[26],[16,91],[18],[11],[66],[22,20],[44],[17,16],[95],[2],[14,2],[99],[51],[1],[11],[22,99],[20],[25,42],[72],[45],[2],[4],[32],[55],[84],[32,64],[26,14],[30,80],[88],[51],[27,71],[15],[8],[60],[37],[25],[96],[25,53],[36],[8],[85],[42],[20],[34],[78],[42,76],[26],[30],[39],[27],[93],[19,75],[8],[24],[32],[25,98],[21],[95],[18],[45],[24],[38],[8],[20],[83],[71],[78],[55],[29],[11],[84]]
*/
	obj := Constructor()
	obj.AddAtHead(55)
	obj.AddAtIndex(1, 90)
	obj.AddAtTail(51)
	obj.AddAtTail(91)
	obj.AddAtTail(12)
	obj.AddAtIndex(2, 72)
	obj.AddAtTail(17)
	obj.AddAtHead(82)
	obj.AddAtHead(4)
	obj.DeleteAtIndex(7)
	obj.DeleteAtIndex(7)
	obj.DeleteAtIndex(7)
	obj.AddAtIndex(5, 75)
	obj.AddAtTail(54)
	obj.Get(6)
	obj.Get(2)
	obj.AddAtHead(8)
	obj.AddAtTail(35)
	obj.AddAtTail(36)
	obj.Get(10)
	// "addAtTail","addAtTail","deleteAtIndex","deleteAtIndex","addAtHead","addAtTail","addAtIndex",
	// "get","addAtTail","addAtIndex","addAtHead","addAtTail","addAtIndex",
	// [40],[43],[12],[3],[78],[89],[3,41],[10],[96],[5,37],[51],[26],[16,91],[18],[11],[66],[22,20],[44],[17,16],[95],[2],[14,2],[99],[51],[1],[11],[22,99],[20],[25,42],[72],[45],[2],[4],[32],[55],[84],[32,64],[26,14],[30,80],[88],[51],[27,71],[15],[8],[60],[37],[25],[96],[25,53],[36],[8],[85],[42],[20],[34],[78],[42,76],[26],[30],[39],[27],[93],[19,75],[8],[24],[32],[25,98],[21],[95],[18],[45],[24],[38],[8],[20],[83],[71],[78],[55],[29],[11],[84]]
	obj.AddAtTail(40)
	obj.AddAtTail(43)
	obj.DeleteAtIndex(12)
	obj.DeleteAtIndex(3)
	obj.AddAtHead(78)
	obj.AddAtTail(89)
	obj.AddAtIndex(3, 41)
	obj.Get(10)
}