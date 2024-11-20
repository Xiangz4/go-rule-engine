package handler

import (
	"go-rule-engine/flowengine/config"
	"go-rule-engine/flowengine/context"
	"go-rule-engine/flowengine/enums"
)

type Handler func(ctx *config.FlowConfig) error

type CallbackHandler func(ctx *config.FlowConfig) error

type HandlerNode struct {
	//节点类型
	NodeType enums.HandlerNodeType

	//下一状态
	NextState enums.FlowState
	//主处理方法
	Handler Handler

	//回调处理方法
	CallbackHandler CallbackHandler

	//子节点
	Children []*HandlerNode
}

// NewHandlerNode 创建新的处理器节点
func NewHandlerNode(nodeType enums.HandlerNodeType, nextState enums.FlowState) *HandlerNode {
	return &HandlerNode{
		NodeType:  nodeType,
		NextState: nextState,
		Children:  make([]*HandlerNode, 0),
	}
}

// 设置主处理器
func (n *HandlerNode) WithHandler(handler Handler) *HandlerNode {
	n.Handler = handler
	return n
}

// WithCallBack 设置回调
func (n *HandlerNode) WithCallback(callback CallbackHandler) *HandlerNode {
	n.CallbackHandler = callback
	return n
}

// AddChild 添加子节点
func (n *HandlerNode) AddChild(child *HandlerNode) *HandlerNode {
	n.Children = append(n.Children, child)
	return n
}

// 执行处理器节点以及其所有子节点
func (n *HandlerNode) Execute(ctx *context.FlowContext) error {
	if n.Handler != nil {
		if err := n.Handler(ctx); err != nil {
			return err
		}
	}

	//执行回调处理
	if n.CallbackHandler != nil {
		if err := n.CallbackHandler(ctx); err != nil {
			return err
		}
	}

	//执行所有子节点
	for _, child := range n.Children {
		if err := child.Execute(ctx); err != nil {
			return err
		}
	}
	return nil
}
