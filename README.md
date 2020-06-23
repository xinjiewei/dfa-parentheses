# dfa-parentheses

#### 确定有限状态自动机求解括号匹配

- 状态图对应的状态表，x轴:前一个状态 y轴:后一个状态，x由[x,y]变化为y。初始状态为0。如:0状态遇到 { 变成1状态；1状态遇到 } 变成2状态。

|      | 0    | 1    | 2    | 3    | 4    | 5    | 6    |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
| 0    |      | {    |      | [    |      | (    |      |
| 1    |      | {    | }    | [    |      | (    |      |
| 2    |      | {    | }    | [    | ]    | (    | )    |
| 3    |      | {    |      | [    | ]    | (    |      |
| 4    |      | {    | }    | [    | ]    | (    | )    |
| 5    |      | {    |      | [    |      | (    | )    |
| 6    |      | {    | }    | [    | ]    | (    | )    |

- 目录说明

> **src/mai.go** 给出了golang的实现

> **src/java** 给出java的实现