数组的下标是一个隐含的很有用的数组，特别是在统计一些数字，或者判断一些整型数是否出现过的时候。例如，给你一串字母，让你判断这些字母出现的次数时，我们就可以把这些字母作为下标，在遍历的时候，如果字母a遍历到，则arr[a]就可以加1了，即  arr[a]++;

对于双指针，在做关于单链表的题是特别有用，比如“判断单链表是否有环”、“如何一次遍历就找到链表中间位置节点”、“单链表中倒数第 k 个节点”等问题。对于这种问题，我们就可以使用双指针了，会方便很多。我顺便说下这三个问题怎么用双指针解决吧。

作者：Markorg
链接：https://www.zhihu.com/question/24964987/answer/660815955
来源：知乎
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

我面试经验很丰富，就随意说几个1.支付宝面试题（经典题目）// Single Number:
     Given 2*n + 1 numbers, every numbers occurs twice except one, find it.

Example:
   Given [1,2,2,1,3,4,3], return 4

Code:

public class Solution {
    /**
     * @param A: An integer array
     * @return: An integer
     */
    public int singleNumber(int[] A) {
        // write your code here
      
    }
}
这种题目使用位运算    public int singleNumber(int[] nums) {
        int result = 0;
        for (int i = 0; i < nums.length; i++) {
            result ^= nums[i];
        }
        return result;
    }2.蚂蚁金服，合并有序链表//评测题目: /**已知两个链表head1 和head2 各自有序，请把它们归并成一个链表依然有序**/
class Array{
   int data;
   Node next;
}
int[] Merge(Node head1 , Node head2) {
//TODO

  
}
代码如下    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        if(l1==null) return l2;
        if(l2==null) return l1;

        ListNode temp;
        if(l1.val<l2.val){
            temp = l1;
            temp.next= mergeTwoLists(l1.next,l2);
        }else{
            temp = l2;
            temp.next= mergeTwoLists(l1,l2.next);
        }
        return temp;
    }

相同数字做&运算，会得到相同的数字。

https://blog.csdn.net/biezhihua/article/details/79571917

^=符号

二叉查找树（BST：Binary Search Tree）是一种特殊的二叉树，它改善了二叉树节点查找的效率。二叉查找树有以下性质：

对于任意一个节点 n，

其左子树（left subtree）下的每个后代节点（descendant node）的值都小于节点 n 的值；
其右子树（right subtree）下的每个后代节点的值都大于节点 n 的值。
