package transaction

import "testing"

func TestGetVersion(t *testing.T) {
	const tx = "01000000014c6ec863cf3e0284b407a1a1b8138c76f98280812cb9653231f385a0305fc76f010000006b483045022100f01c1a1679c9437398d691c8497f278fa2d615efc05115688bf2c3335b45c88602201b54437e54fb53bc50545de44ea8c64e9e583952771fcc663c8687dc2638f7854121037e87bbd3b680748a74372640628a8f32d3a841ceeef6f75626ab030c1a04824fffffffff021d784500000000001976a914e9b62e25d4c6f97287dfe62f8063b79a9638c84688ac60d64f00000000001976a914bb4bca2306df66d72c6e44a470873484d8808b8888ac00000000"
	bt, err := NewFromString(tx)
	if err != nil {
		t.Error(err)
		return
	}

	res := bt.Version()
	if res != 1 {
		t.Errorf("Expecting 1, got %d", res)
	}
}

func TestIsCoinbase(t *testing.T) {
	const coinbase = "01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff4303bfea07322f53696d6f6e204f726469736820616e642053747561727420467265656d616e206d61646520746869732068617070656e2f9a46434790f7dbdea3430000ffffffff018a08ac4a000000001976a9148bf10d323ac757268eb715e613cb8e8e1d1793aa88ac00000000"
	bt1, err := NewFromString(coinbase)
	if err != nil {
		t.Error(err)
		return
	}

	cb1 := bt1.IsCoinbase()
	if cb1 == false {
		t.Errorf("Expecting true, got %t", cb1)
	}

	const tx = "01000000014c6ec863cf3e0284b407a1a1b8138c76f98280812cb9653231f385a0305fc76f010000006b483045022100f01c1a1679c9437398d691c8497f278fa2d615efc05115688bf2c3335b45c88602201b54437e54fb53bc50545de44ea8c64e9e583952771fcc663c8687dc2638f7854121037e87bbd3b680748a74372640628a8f32d3a841ceeef6f75626ab030c1a04824fffffffff021d784500000000001976a914e9b62e25d4c6f97287dfe62f8063b79a9638c84688ac60d64f00000000001976a914bb4bca2306df66d72c6e44a470873484d8808b8888ac00000000"
	bt2, err := NewFromString(tx)
	if err != nil {
		t.Error(err)
		return
	}

	cb2 := bt2.IsCoinbase()
	if cb2 == true {
		t.Errorf("Expecting false, got %t", cb2)
	}
}

/*
48
30
45
02
21
00f4de422896e461da647b21d800a4ca9ace98dbd08c2dc9b8e049c93197c314f5
02
20
68836c3dfa6650ebeff73b1e3caa8761cd107ed13d6cc713856ebde3f874dd41
41

21
02aea77c449eeeef2746562e56ad053202755f9844276e3f0c684f9d59cdb9458d
ac OP_CHECKSIG

*/
// func TestMyTransaction(t *testing.T) {
// 	fromTx, err := NewFromString("02000000010000000000000000000000000000000000000000000000000000000000000000ffffffff0d510101092f45423132382e302fffffffff0100f2052a01000000232102aea77c449eeeef2746562e56ad053202755f9844276e3f0c684f9d59cdb9458dac00000000")
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}

// 	toTx, err := NewFromString("02000000019bb2dea27bcff46bca60e46ba2fdce706a8eb9d22c9b05e54166b8f9ac57d6de0000000049483045022100f4de422896e461da647b21d800a4ca9ace98dbd08c2dc9b8e049c93197c314f5022068836c3dfa6650ebeff73b1e3caa8761cd107ed13d6cc713856ebde3f874dd4141feffffff0200ca9a3b000000001976a9143c134f3ccd097be40242efd6fb370fc62501afe788ac00196bee000000001976a914c3d737cb0d93ded96a35d240aa3f01b34edc4e5d88ac65000000")
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}

// 	t.Errorf("%x\n%x\n", fromTx.GetOutputs()[0].GetOutputScript(), toTx.GetInputs()[0].GetInputScript())
// }
