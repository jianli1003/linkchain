package types

import "math/big"

/*
	if (x % (1e+18) != 0) {
		fee = (x/(1e+18)+1) * rate *(1e+13)
	} else {
		fee = (x/(1e+18)) * rate *(1e+13)
	}
	if fee < min {
		fee = min
	}
	if max != 0 && fee > max {
		fee = max
	}
*/

const (
	MaxGasLimit  int64 = 5e9     // max gas limit
	MinGasLimit  int64 = 5e5     // min gas limit
	MaxFeeCounts int64 = 1024000

	everLiankeFee         int64 = 5e4  // ever poundage fee unit(gas)
	everContractLiankeFee int64 = 25e3 // ever poundate contract fee uint(gas)
	gasToLiankeRate       int64 = 1e7  // lianke = 1e+7 gas
	GasPrice              int64 = 1e11

)

func CalNewAmountGas(value *big.Int) uint64 {
	var liankeCount *big.Int
	lianke := new(big.Int).Mul(big.NewInt(GasPrice), big.NewInt(gasToLiankeRate))
	if new(big.Int).Mod(value, lianke).Uint64() != 0 {
		liankeCount = new(big.Int).Div(value, lianke)
		liankeCount.Add(liankeCount, big.NewInt(1))
	} else {
		liankeCount = new(big.Int).Div(value, lianke)
	}
	calFeeGas := new(big.Int).Mul(big.NewInt(everLiankeFee), liankeCount)

	if calFeeGas.Cmp(big.NewInt(MinGasLimit)) < 0 {
		calFeeGas.Set(big.NewInt(MinGasLimit))
	}
	if MaxGasLimit != 0 && calFeeGas.Cmp(big.NewInt(MaxGasLimit)) > 0 {
		calFeeGas.Set(big.NewInt(MaxGasLimit))
	}
	return calFeeGas.Uint64()
}

func CalNewContractAmountGas(value *big.Int) uint64 {
	var liankeCount *big.Int
	lianke := new(big.Int).Mul(big.NewInt(GasPrice), big.NewInt(gasToLiankeRate))
	if new(big.Int).Mod(value, lianke).Uint64() != 0 {
		liankeCount = new(big.Int).Div(value, lianke)
		liankeCount.Add(liankeCount, big.NewInt(1))
	} else {
		liankeCount = new(big.Int).Div(value, lianke)
	}
	calFeeGas := new(big.Int).Mul(big.NewInt(everContractLiankeFee), liankeCount)

	if calFeeGas.Cmp(big.NewInt(MinGasLimit)) < 0 {
		calFeeGas.Set(big.NewInt(MinGasLimit))
	}
	if MaxGasLimit != 0 && calFeeGas.Cmp(big.NewInt(MaxGasLimit)) > 0 {
		calFeeGas.Set(big.NewInt(MaxGasLimit))
	}
	return calFeeGas.Uint64()
}
