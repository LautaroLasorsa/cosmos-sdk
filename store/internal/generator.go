package main

import (
	"encoding/hex"
	"fmt"

	"cosmossdk.io/math/unsafe"
	"cosmossdk.io/store/internal/proofs"
)

func main() {
	fmt.Println("Hello, World!")
	size := 1 << 10
	data := proofs.BuildMap(size)
	allkeys := proofs.SortedKeys(data)

	tests := 1
	keys := make([]string, 0, tests)
	values := make([]string, 0, tests)
	roots := make([]string, 0, tests)
	arrProofs := make([]string, 0, tests)

	token := ""
	base := unsafe.NewRand().Int() % (size - tests + 1)

	for i := 0; i < tests; i++ {
		key := allkeys[base+i] //proofs.GetKey(allkeys, proofs.Middle)
		val := data[key]
		root := proofs.CalcRoot(data)
		proof, _ := proofs.CreateMembershipProof(data, []byte(key))

		keys = append(keys, token+"hex\""+hex.EncodeToString([]byte(key))+"\"")
		values = append(values, token+"hex\""+hex.EncodeToString(val)+"\"")
		roots = append(roots, token+"hex\""+hex.EncodeToString(root)+"\"")

		marshalProof, _ := proof.Marshal()
		arrProofs = append(arrProofs, token+"hex\""+hex.EncodeToString(marshalProof)+"\"")
		//token = ","
	}

	fmt.Println("function getKeys() private returns (bytes[] memory) {\nbytes[] memory keys = new bytes[](", tests, ");")
	for i, key := range keys {
		fmt.Printf("keys[%d] = %s; ", i, key)
	}
	fmt.Println("return keys;\n}")

	fmt.Println("function getValues() private returns (bytes[] memory) {\nbytes[] memory values = new bytes[](", tests, ");")
	for i, value := range values {
		fmt.Printf("values[%d] = %s; ", i, value)
	}
	fmt.Println("return values;\n}")

	fmt.Println("function getRoots() private returns (bytes[] memory) {\nbytes[] memory roots = new bytes[](", tests, ");")
	for i, root := range roots {
		fmt.Printf("roots[%d] = %s; ", i, root)
	}
	fmt.Println("return roots;\n}")

	fmt.Println("function getProofs() private returns (bytes[] memory) {\nbytes[] memory proofs = new bytes[](", tests, ");")
	for i, proof := range arrProofs {
		fmt.Printf("proofs[%d] = %s; ", i, proof)
	}
	fmt.Println("return proofs;\n}")

	// fmt.Println("Keys:\n", keys)
	// fmt.Println("Values:\n", values)
	// fmt.Println("Roots:\n", roots)
	// fmt.Println("Proofs:\n", arrProofs)
}
