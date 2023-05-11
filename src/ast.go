package main

import (
	"math"
)

type AST interface {
	evaluate(c chan float64)
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

func (node *AddNode) evaluate(c chan float64) {
	left := node.GetLeftSubTree()
	right := node.GetRightSubTree()
	controller := make(chan float64, 2)
	go left.evaluate(controller)
	go right.evaluate(controller)
	first := <-controller
	second := <-controller
	c <- (first + second)
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

func (node *SubNode) evaluate(c chan float64) {
	left := node.GetLeftSubTree()
	right := node.GetRightSubTree()
	leftChan := make(chan float64)
	rightChan := make(chan float64)
	go left.evaluate(leftChan)
	go right.evaluate(rightChan)
	first := <-leftChan
	second := <-rightChan
	c <- (first - second)
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

func (node *TimesNode) evaluate(c chan float64) {
	left := node.GetLeftSubTree()
	right := node.GetRightSubTree()
	controller := make(chan float64, 2)
	go left.evaluate(controller)
	go right.evaluate(controller)
	first := <-controller
	second := <-controller
	c <- (first * second)
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

func (node *DivideNode) evaluate(c chan float64) {
	left := node.GetLeftSubTree()
	right := node.GetRightSubTree()
	leftChan := make(chan float64)
	rightChan := make(chan float64)
	go left.evaluate(leftChan)
	go right.evaluate(rightChan)
	first := <-leftChan
	second := <-rightChan
	c <- (first / second)
}

type NumNode struct {
	val float64
}

func (node NumNode) evaluate(c chan float64) {
	c <- node.val
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

func (node *LbyteNode) evaluate(c chan float64) {
	left := node.GetLeftSubTree()
	right := node.GetRightSubTree()
	leftChan := make(chan float64)
	rightChan := make(chan float64)
	go left.evaluate(leftChan)
	go right.evaluate(rightChan)
	first := int(<-leftChan)
	second := int(<-rightChan)
	c <- float64(first << second)
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

func (node *RbyteNode) evaluate(c chan float64) {
	left := node.GetLeftSubTree()
	right := node.GetRightSubTree()
	leftChan := make(chan float64)
	rightChan := make(chan float64)
	go left.evaluate(leftChan)
	go right.evaluate(rightChan)
	first := int(<-leftChan)
	second := int(<-rightChan)
	c <- float64(first >> second)
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

func (node *ModNode) evaluate(c chan float64) {
	left := node.GetLeftSubTree()
	right := node.GetRightSubTree()
	leftChan := make(chan float64)
	rightChan := make(chan float64)
	go left.evaluate(leftChan)
	go right.evaluate(rightChan)
	first := int(<-leftChan)
	second := int(<-rightChan)
	c <- float64(first % second)
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

func (node *PotNode) evaluate(c chan float64) {
	left := node.GetLeftSubTree()
	right := node.GetRightSubTree()
	leftChan := make(chan float64)
	rightChan := make(chan float64)
	go left.evaluate(leftChan)
	go right.evaluate(rightChan)
	first := <-leftChan
	second := <-rightChan
	c <- math.Pow(first, second)
}

type SqrtNode struct {
	subTree AST
}

func (node *SqrtNode) evaluate(c chan float64) {
	chann := make(chan float64)
	/* print("here") */
	go node.subTree.evaluate(chann)
	sub := <-chann
	c <- math.Sqrt(sub)
}
