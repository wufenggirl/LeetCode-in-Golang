# [841. Keys and Rooms](https://leetcode.com/problems/keys-and-rooms/)

## 题目

There are N rooms and you start in room 0. Each room has a distinct number in 0, 1, 2, ..., N-1, and each room may havesome keys to access the next room.

Formally, each room ihas a list of keys rooms[i], and each key rooms[i][j] is an integer in [0, 1, ..., N-1] where N = rooms.length. A key rooms[i][j] = vopens the room with number v.

Initially, all the rooms start locked (except for room 0).

You can walk back and forth between rooms freely.

Return trueif and only if you can enterevery room.

Example 1:

```text
Input: [[1],[2],[3],[]]
Output: true
Explanation:
We start in room 0, and pick up key 1.
We then go to room 1, and pick up key 2.
We then go to room 2, and pick up key 3.
We then go to room 3.  Since we were able to go to every room, we return true.
```

Example 2:

```text
Input: [[1,3],[3,0,1],[2],[0]]
Output: false
Explanation: We can't enter the room with number 2.
```

Note:

1. 1 <= rooms.length <=1000
1. 0 <= rooms[i].length <= 1000
1. The number of keys in all rooms combined is at most3000.

## 解题思路

见程序注释
