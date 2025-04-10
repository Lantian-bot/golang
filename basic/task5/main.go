package main

//1.借助ai工具，编写go代码实现一个简单的Merkle树，并验证其正确性。

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Node 表示Merkle树中的节点
type Node struct {
	Left  *Node  // 左子节点
	Right *Node  // 右子节点
	Hash  []byte // 当前节点的哈希值
}

// MerkleTree 表示整个Merkle树结构
type MerkleTree struct {
	Root *Node // 根节点
}

// NewMerkleTree 根据输入数据构建Merkle树
func NewMerkleTree(data [][]byte) *MerkleTree {
	if len(data) == 0 {
		return nil
	}

	// 创建叶子节点
	var leaves []*Node
	for _, d := range data {
		hash := sha256.Sum256(d)
		leaves = append(leaves, &Node{Hash: hash[:]})
	}

	// 逐层构建树结构
	for len(leaves) > 1 {
		// 处理奇数节点情况，复制最后一个节点
		if len(leaves)%2 != 0 {
			leaves = append(leaves, leaves[len(leaves)-1])
		}

		var parents []*Node
		for i := 0; i < len(leaves); i += 2 {
			// 创建父节点并计算组合哈希
			parent := &Node{
				Left:  leaves[i],
				Right: leaves[i+1],
				Hash:  computeCombinedHash(leaves[i].Hash, leaves[i+1].Hash),
			}
			parents = append(parents, parent)
		}
		leaves = parents
	}

	return &MerkleTree{Root: leaves[0]}
}

// computeCombinedHash 计算两个子节点哈希的组合哈希
func computeCombinedHash(left, right []byte) []byte {
	combined := append(left, right...)
	hash := sha256.Sum256(combined)
	return hash[:]
}

// GetRootHash 返回根哈希的十六进制字符串表示
func (mt *MerkleTree) GetRootHash() string {
	return hex.EncodeToString(mt.Root.Hash)
}

func main() {
	// 示例数据
	data := [][]byte{
		[]byte("Blockchain"),
		[]byte("Merkle"),
		[]byte("Tree"),
		[]byte("Example"),
	}

	// 构建Merkle树
	mt := NewMerkleTree(data)

	// 输出根哈希
	fmt.Printf("Merkle Root Hash: %s\n", mt.GetRootHash())
}
