package main

import (
	"math"
)

type AST interface {
	evaluate() int
}

type BinaryNode interface {
	GetLeftSubTree() AST
	GetRightSubTree() AST
}

type UnaryNode interface {
	getSubTree() AST
}

type AddNode struct {
	leftTree  AST
	rightTree AST
}

func (node *AddNode) GetLeftSubTree() AST {
	return node.leftTree
}

func (node *AddNode) GetRightSubTree() AST {
	return node.rightTree
}

func (node *AddNode) evaluate() int {
	left := node.GetLeftSubTree()
	right := node.GetRightSubTree()
	return left.evaluate() + right.evaluate()
}

type SubNode struct {
	leftTree  AST
	rightTree AST
}

func (node *SubNode) GetLeftSubTree() AST {
	return node.leftTree
}

func (node *SubNode) GetRightSubTree() AST {
	return node.rightTree
}

func (node *SubNode) evaluate() int {
	left := node.GetLeftSubTree()
	right := node.GetRightSubTree()
	return left.evaluate() - right.evaluate()
}

type TimesNode struct {
	leftTree  AST
	rightTree AST
}

func (node *TimesNode) GetLeftSubTree() AST {
	return node.leftTree
}

func (node *TimesNode) GetRightSubTree() AST {
	return node.rightTree
}

func (node *TimesNode) evaluate() int {
	left := node.GetLeftSubTree()
	right := node.GetRightSubTree()
	return left.evaluate() * right.evaluate()
}

type DivideNode struct {
	leftTree  AST
	rightTree AST
}

func (node *DivideNode) GetLeftSubTree() AST {
	return node.leftTree
}

func (node *DivideNode) GetRightSubTree() AST {
	return node.rightTree
}

func (node *DivideNode) evaluate() int {
	left := node.GetLeftSubTree()
	right := node.GetRightSubTree()
	return left.evaluate() / right.evaluate()
}

type NumNode struct {
	val int
}

func (node NumNode) evaluate() int {
	return node.val
}

type LbyteNode struct {
	leftTree  AST
	rightTree AST
}

func (node *LbyteNode) GetLeftSubTree() AST {
	return node.leftTree
}

func (node *LbyteNode) GetRightSubTree() AST {
	return node.rightTree
}

func (node *LbyteNode) evaluate() int {
	left := node.GetLeftSubTree()
	right := node.GetRightSubTree()
	return left.evaluate() << right.evaluate()
}

type RbyteNode struct {
	leftTree  AST
	rightTree AST
}

func (node *RbyteNode) GetLeftSubTree() AST {
	return node.leftTree
}

func (node *RbyteNode) GetRightSubTree() AST {
	return node.rightTree
}

func (node *RbyteNode) evaluate() int {
	left := node.GetLeftSubTree()
	right := node.GetRightSubTree()
	return left.evaluate() >> right.evaluate()
}

type ModNode struct {
	leftTree  AST
	rightTree AST
}

func (node *ModNode) GetLeftSubTree() AST {
	return node.leftTree
}

func (node *ModNode) GetRightSubTree() AST {
	return node.rightTree
}

func (node *ModNode) evaluate() int {
	left := node.GetLeftSubTree()
	right := node.GetRightSubTree()
	return left.evaluate() % right.evaluate()
}

type PotNode struct {
	leftTree  AST
	rightTree AST
}

func (node *PotNode) GetLeftSubTree() AST {
	return node.leftTree
}

func (node *PotNode) GetRightSubTree() AST {
	return node.rightTree
}

func (node *PotNode) evaluate() int {
	left := node.GetLeftSubTree()
	right := node.GetRightSubTree()
	return int(math.Pow(float64(left.evaluate()), float64(right.evaluate())))
}

type SqrtNode struct {
	subTree AST
}

func (node *SqrtNode) evaluate() int {
	return int(math.Sqrt(float64(node.subTree.evaluate())))
}
