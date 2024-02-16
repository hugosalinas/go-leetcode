package add_two_numbers_2

import (
	"slices"
	"strings"

	"github.com/austingebauer/go-leetcode/structures"
	// "math/big"

	"strconv"
)

func addTwoNumbers(l1 *structures.ListNode, l2 *structures.ListNode) *structures.ListNode {
	l1NumString := ""
	l2NumString := ""

	for true {
		l1NumString = strconv.Itoa(l1.Val) + l1NumString
		if l1.Next == nil {
			break
		}
		l1 = l1.Next
	}

	for true {
		l2NumString = strconv.Itoa(l2.Val) + l2NumString
		if l2.Next == nil {
			break
		}
		l2 = l2.Next
	}

	num1, _ := strconv.Atoi(l1NumString)
	num2, _ := strconv.Atoi(l2NumString)
	numSum := strconv.Itoa(num1 + num2)

	numSumList := strings.Split(numSum, "")
	slices.Reverse(numSumList)

	numLinkedList := &structures.ListNode{}
	numLinkedList = addNode(numLinkedList, 0, numSumList)

	return numLinkedList
}

func addNode(nodeList *structures.ListNode, index int, numList []string) *structures.ListNode {
	currentValue, _ := strconv.Atoi(numList[index])
	nodeList = &structures.ListNode{
		Val:  currentValue,
		Next: nil,
	}
	if index < len(numList)-1 {
		index++
		nodeList.Next = addNode(nodeList, index, numList)
	}
	return nodeList
}

// func addTwoNumbers(l1 *structures.ListNode, l2 *structures.ListNode) *structures.ListNode {
// 	num1 := depthFirstNum(l1)
// 	num2 := depthFirstNum(l2)

// 	var sum big.Int
// 	sum.Add(num1, num2)

// 	// single digit sum
// 	if len(sum.String()) == 1 {
// 		return &structures.ListNode{
// 			Val:  int(sum.Int64()),
// 			Next: nil,
// 		}
// 	}

// 	var sumList *structures.ListNode
// 	var current *structures.ListNode

// 	for sum.String() != "0" {
// 		var lastDigit big.Int
// 		lastDigit.Mod(&sum, big.NewInt(10))

// 		node := &structures.ListNode{
// 			Val:  int(lastDigit.Int64()),
// 			Next: nil,
// 		}
// 		if sumList == nil {
// 			sumList = node
// 		}
// 		if current == nil {
// 			current = node
// 		} else {
// 			current.Next = node
// 			current = current.Next
// 		}

// 		sum.Div(&sum, big.NewInt(10))
// 	}

// 	return sumList
// }

// func depthFirstNum(l *structures.ListNode) *big.Int {
// 	if l == nil {
// 		return big.NewInt(0)
// 	}

// 	v := depthFirstNum(l.Next)
// 	return v.Mul(v, big.NewInt(10)).Add(v, big.NewInt(int64(l.Val)))
// }

// func addTwoNumbers2(l1 *structures.ListNode, l2 *structures.ListNode) *structures.ListNode {
// 	carry := 0
// 	sumList := &structures.ListNode{
// 		Val:  0,
// 		Next: nil,
// 	}
// 	frontSumList := sumList

// 	for l1 != nil || l2 != nil {
// 		// calculate sum and carry values
// 		l1Val := 0
// 		l2Val := 0
// 		if l1 != nil {
// 			l1Val = l1.Val
// 			l1 = l1.Next
// 		}
// 		if l2 != nil {
// 			l2Val = l2.Val
// 			l2 = l2.Next
// 		}

// 		sum := l1Val + l2Val + carry
// 		carry = sum / 10
// 		sumList.Val = sum % 10

// 		// no more list to process, but carry is not 0
// 		if l1 == nil && l2 == nil && carry > 0 {
// 			sumList.Next = &structures.ListNode{
// 				Val:  carry,
// 				Next: nil,
// 			}
// 		}

// 		// one or both lists still can be processed
// 		if l1 != nil || l2 != nil {
// 			sumList.Next = &structures.ListNode{
// 				Val:  0,
// 				Next: nil,
// 			}
// 			sumList = sumList.Next
// 		}
// 	}

// 	return frontSumList
// }
