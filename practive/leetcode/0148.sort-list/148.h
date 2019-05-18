/*
 * 在 O(nlogn) 时间复杂度和常数级空间复杂度下，对链表进行排序。
 * */

struct ListNode {
    int val;
    ListNode* next;
    ListNode(int x)
        :val(x)
        ,next(nullptr)
    {}
};

/*
 * 参考：Sort List——经典（链表中的归并排序） https://www.cnblogs.com/qiaozhoulin/p/4585401.html
 * 归并排序法：在动手之前一直觉得空间复杂度为常量不太可能，因为原来使用归并时，都是 O(N)的，
 * 需要复制出相等的空间来进行赋值归并。对于链表，实际上是可以实现常数空间占用的（链表的归并
 * 排序不需要额外的空间）。利用归并的思想，递归地将当前链表分为两段，然后merge，分两段的方
 * 法是使用 fast-slow 法，用两个指针，一个每次走两步，一个走一步，知道快的走到了末尾，然后
 * 慢的所在位置就是中间位置，这样就分成了两段。merge时，把两段头部节点值比较，用一个 p 指向
 * 较小的，且记录第一个节点，然后 两段的头一步一步向后走，p也一直向后走，总是指向较小节点，
 * 直至其中一个头为NULL，处理剩下的元素。最后返回记录的头即可。
 * 主要考察3个知识点，
 * 知识点1：归并排序的整体思想
 * 知识点2：找到一个链表的中间节点的方法
 * 知识点3：合并两个已排好序的链表为一个新的有序链表
 * */

class Solution {
public:
    ListNode* sortList(ListNode* head) {
        if (head == nullptr || head->next == nullptr) {
            return head;                    
        }
        ListNode* left = head, *middle = left, *right = head;
        while(right && right->next) {
            right = right->next->next;
            middle = left;
            //跳出循环前，left 为中间节点，跳出后往后移动一位，指向 middle 后面的链表
            left = left->next; 
        }
        //将链表分成两部分，前一部分以 middle 结尾，应该置空，后一阶段以 left 开头
        middle->next = nullptr;
        //分别对 head、left 链表进行排序（合并两个有序链表）
        return merge(sortList(head), sortList(left));
    }
private:
    //merge list
    ListNode* merge(ListNode* l1, ListNode* l2) {
        if(l1 == nullptr) {
            return l2;                         
        }
        if(l2 == nullptr) {
            return l1;                    
        }
        if(l1->val <= l2->val) {
            l1->next = merge(l1->next, l2);
            return l1;                                
        } else {
            l2->next = merge(l1, l2->next);
            return l2;                                
        }
    }
};

/*
 * 执行用时 : 68 ms, 在Sort List的C++提交中击败了96.75% 的用户
 * 内存消耗 : 12.7 MB, 在Sort List的C++提交中击败了55.93% 的用户
 * */
