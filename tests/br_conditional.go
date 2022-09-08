package main

import (
	"fmt"
	"time"
	"math/big"

	"github.com/lol-chain/go-kekchain/common"
	"github.com/lol-chain/go-kekchain/params"
)

// to test build the program 
// go build br_conditional.go
// and run it (on WIN) ./br_conditional.exe || (on UNIX) ./br_conditional 
// or simply run the program in go
// go run br_conditional.go

var (
	mainnetChainConfig = params.ChainConfig{
		ChainID:             big.NewInt(420690),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        nil,
		DAOForkSupport:      false,
		EIP150Block:         big.NewInt(0),
		EIP150Hash:          common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		BerlinBlock:         big.NewInt(0),
		LondonBlock:         big.NewInt(0),
		BRBlock:             big.NewInt(3),
		BRHalving:           big.NewInt(6),
		BRFinalSubsidy:      big.NewInt(9),
	}
    	block = big.NewInt(0)
    	ConstantBlockReward = big.NewInt(2e+18) // Block reward in wei for successfully mining a block upward from BR activator fork
    	ConstantHalfBlockReward = big.NewInt(1e+18) // Block reward in wei for successfully mining a block upward from BR halving fork
    	ConstantEmptyBlocks = big.NewInt(1e+1) // Block reward in wei for successfully mining a block upward from BR activator fork

)

func main() {
    go blockSimulator()
	time.Sleep(time.Second * 15)
}


// isForked returns whether a fork scheduled at block s is active at the given head block.
func isForked(s, head *big.Int) bool {
	if s == nil || head == nil {
		return false
	}
	return s.Cmp(head) <= 0
}

// IsBRonline returns whether num is either 
// equal to the block reward activator fork block or greater.
func IsBRonline(num *big.Int) bool {
	return isForked(mainnetChainConfig.BRBlock, num)
}

// IsBRHalving returns whether num is either 
// equal to the block reward halnving fork block or greater.
func IsBRHalving(num *big.Int) bool {
	return isForked(mainnetChainConfig.BRHalving, num)
}

// IsBRFinalSubsidy returns whether num is either 
// equal to the block reward final subsidy fork block or greater.
func IsBRFinalSubsidy(num *big.Int) bool {
	return isForked(mainnetChainConfig.BRFinalSubsidy, num)
}

// simulate block production, 1 secon per "block"
// iterations occur for up to 15 blocks
func blockSimulator() {
	for range time.Tick(time.Second * 1) {
        // increment block number, simulating block production 
		block.Add(block, big.NewInt(1))
        // begin conditional checks 
        // detect block rebate activation
		if IsBRonline(block) {
            // alias for setting block_rebate
            rebate := ConstantEmptyBlocks
            // ensure final subsidy is handled first,
            // when block > final subsidy activation
            // remainder of code will not compute
            // i.e no rebates post final subsidy
		    if IsBRFinalSubsidy(block) {
                // won't be computed within log @ line 105
                // computation ends here if block height > final subsidy 
                rebate = ConstantEmptyBlocks
			    fmt.Println("Final subsidy entered @ ",block)
		    } else {
                // if block rebate halving event triggered
                // process and computation ends here
                if IsBRHalving(block) {
                    rebate = ConstantHalfBlockReward
                    fmt.Println("Block Halving event entered @ ",block)
                } else {
                    // else set standard BR case 
        		    rebate = ConstantBlockReward
                    fmt.Println("Block Rebates ACTIVE @ ",block)
                }
                // assign the block_rebate only once
                // set from rebate variable above
                block_rebate := rebate
                // this log represents a placeholder used 
                // to add BR to balance of rebate oracle
                fmt.Println("REBATE AMOUNT: ",block_rebate)
		    }
		} else {
            // only runs pre rebate activation
            // in main net this line won't be computed
		    fmt.Println("Disabled rebates",block)
		}
	}
}