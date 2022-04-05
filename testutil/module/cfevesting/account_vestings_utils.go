package cfevesting

import (
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var TestEnvTime = time.Now()

func AssertAccountVestings(t *testing.T, expected types.AccountVestings, actual types.AccountVestings) {

	numOfFields := reflect.TypeOf(types.AccountVestings{}).NumField()
	j :=0
	require.EqualValues(t, len(expected.Vestings), len(actual.Vestings)); j++
	require.EqualValues(t, expected.Address, actual.Address); j++
	require.EqualValues(t, expected.DelegableAddress, actual.DelegableAddress); j++
	require.EqualValues(t, numOfFields, j);

	numOfFields = reflect.TypeOf(types.Vesting{}).NumField()
	for i, expectedVesting := range expected.Vestings {
		actualVesting := actual.Vestings[i]
		j :=0
		require.EqualValues(t, expectedVesting.Id, actualVesting.Id); j++
		require.EqualValues(t, expectedVesting.VestingType, actualVesting.VestingType); j++
		require.EqualValues(t, true, expectedVesting.VestingStart.Equal(actualVesting.VestingStart)); j++
		require.EqualValues(t, true, expectedVesting.LockEnd.Equal(actualVesting.LockEnd)); j++
		require.EqualValues(t, true, expectedVesting.VestingEnd.Equal(actualVesting.VestingEnd)); j++
		require.EqualValues(t, expectedVesting.Vested, actualVesting.Vested); j++
		require.EqualValues(t, expectedVesting.ReleasePeriod, actualVesting.ReleasePeriod); j++
		require.EqualValues(t, expectedVesting.DelegationAllowed, actualVesting.DelegationAllowed); j++
		require.EqualValues(t, expectedVesting.Withdrawn, actualVesting.Withdrawn); j++
		require.EqualValues(t, expectedVesting.Sent, actualVesting.Sent); j++
		require.EqualValues(t, true, expectedVesting.LastModification.Equal(actualVesting.LastModification)); j++
		require.EqualValues(t, expectedVesting.LastModificationVested, actualVesting.LastModificationVested); j++
		require.EqualValues(t, expectedVesting.LastModificationWithdrawn, actualVesting.LastModificationWithdrawn); j++
		require.EqualValues(t, expectedVesting.TransferAllowed, actualVesting.TransferAllowed); j++
		require.EqualValues(t, numOfFields, j);

	}

}

func AssertAccountVestingsArrays(t *testing.T, expected []*types.AccountVestings, actual []*types.AccountVestings) {
	for _, accVest := range expected {
		found := false
		for _, accVestExp := range actual {
			if accVest.Address == accVestExp.Address {
				AssertAccountVestings(t, *accVest, *accVestExp)
				found = true
			}
		}
		require.True(t, found, "not found: "+accVest.Address)

	}
}

func GenerateOneAccountVestingsWithAddressWithRandomVestings(numberOfVestingsPerAccount int,
	accountId int, vestingStartId int) types.AccountVestings {
	return *GenerateAccountVestingsWithRandomVestings(1, numberOfVestingsPerAccount, accountId, vestingStartId)[0]
}

func GenerateAccountVestingsWithRandomVestings(numberOfAccounts int, numberOfVestingsPerAccount int,
	accountStartId int, vestingStartId int) []*types.AccountVestings {
	return generateAccountVestings(numberOfAccounts, numberOfVestingsPerAccount,
		accountStartId, vestingStartId, generateRandomVesting)
}

func GenerateOneAccountVestingsWithAddressWith10BasedVestings(numberOfVestingsPerAccount int,
	accountId int, vestingStartId int) types.AccountVestings {
	return *GenerateAccountVestingsWith10BasedVestings(1, numberOfVestingsPerAccount, accountId, vestingStartId)[0]
}

func GenerateAccountVestingsWith10BasedVestings(numberOfAccounts int, numberOfVestingsPerAccount int,
	accountStartId int, vestingStartId int) []*types.AccountVestings {
	return generateAccountVestings(numberOfAccounts, numberOfVestingsPerAccount,
		accountStartId, vestingStartId, generate10BasedVesting)
}

func generateAccountVestings(numberOfAccounts int, numberOfVestingsPerAccount int,
	accountStartId int, vestingStartId int, generateVesting func(accuntId int, vestingId int) types.Vesting) []*types.AccountVestings {
	accountVestingsArr := []*types.AccountVestings{}
	for i := 0; i < numberOfAccounts; i++ {
		accountVestings := types.AccountVestings{}
		accountVestings.Address = "test-vesting-account-addr-" + strconv.Itoa(i+accountStartId)
		accountVestings.DelegableAddress = "test-vesting-account-del-addr-" + strconv.Itoa(i+accountStartId)

		vestings := []*types.Vesting{}
		for j := 0; j < numberOfVestingsPerAccount; j++ {
			vesting := generateVesting(i+accountStartId, j+vestingStartId)
			vestings = append(vestings, &vesting)
		}
		accountVestings.Vestings = vestings

		accountVestingsArr = append(accountVestingsArr, &accountVestings)
	}

	return accountVestingsArr
}

func generateRandomVesting(accuntId int, vestingId int) types.Vesting {
	rgen := rand.New(rand.NewSource(time.Now().UnixNano()))
	return types.Vesting{
		Id:                        int32(vestingId),
		VestingType:               "test-vesting-account-" + strconv.Itoa(accuntId) + "-" + strconv.Itoa(accuntId),
		VestingStart:         CreateTimeFromNumOfHours(int64(rgen.Intn(100000))),
		LockEnd:              CreateTimeFromNumOfHours(int64(rgen.Intn(100000))),
		VestingEnd:           CreateTimeFromNumOfHours(int64(rgen.Intn(100000))),
		Vested:                    sdk.NewInt(int64(rgen.Intn(10000000))),
		ReleasePeriod:      CreateDurationFromNumOfHours(int64(rgen.Intn(1000))),
		DelegationAllowed:         rgen.Intn(2) == 1,
		Withdrawn:                 sdk.NewInt(int64(rgen.Intn(10000000))),
		Sent:                      sdk.NewInt(int64(rgen.Intn(10000000))),
		LastModification:     CreateTimeFromNumOfHours(int64(rgen.Intn(100000))),
		LastModificationVested:    sdk.NewInt(int64(rgen.Intn(10000000))),
		LastModificationWithdrawn: sdk.NewInt(int64(rgen.Intn(10000000))),
	}
}

func generate10BasedVesting(accuntId int, vestingId int) types.Vesting {
	return types.Vesting{
		Id:                        int32(vestingId),
		VestingType:               "test-vesting-account-" + strconv.Itoa(accuntId) + "-" + strconv.Itoa(accuntId),
		VestingStart:         CreateTimeFromNumOfHours(1000),
		LockEnd:              CreateTimeFromNumOfHours(10000),
		VestingEnd:           CreateTimeFromNumOfHours(110000),
		Vested:                    sdk.NewInt(1000000),
		ReleasePeriod:      CreateDurationFromNumOfHours(10),
		DelegationAllowed:         true,
		Withdrawn:                 sdk.ZeroInt(),
		Sent:                      sdk.ZeroInt(),
		LastModification:     CreateTimeFromNumOfHours(1000),
		LastModificationVested:    sdk.NewInt(1000000),
		LastModificationWithdrawn: sdk.ZeroInt(),
	}
}



func ToAccountVestingsPointersArray(src []types.AccountVestings) []*types.AccountVestings {
	result := []*types.AccountVestings{}
	for i := 0; i < len(src); i++ {
		result = append(result, &src[i])
	}
	return result
}

func GetExpectedWithdrawableForVesting(vesting types.Vesting, current time.Time) sdk.Int {
	unlockingStart := vesting.LockEnd
	if vesting.VestingStart.After(unlockingStart) {
		unlockingStart = vesting.VestingStart
	}
	if vesting.LastModification.After(unlockingStart) {
		unlockingStart = vesting.LastModification
	}
	expected := GetExpectedWithdrawable(unlockingStart, vesting.VestingEnd, vesting.ReleasePeriod, current, vesting.LastModificationVested)
	return expected.Sub(vesting.LastModificationWithdrawn)
}

func GetExpectedWithdrawable(unlockingStart time.Time, vestingEnd time.Time, period time.Duration, current time.Time, amount sdk.Int) sdk.Int {
	if current.Equal(vestingEnd) || current.After(vestingEnd) {
		return amount
	}
	if current.Before(unlockingStart) {
		return sdk.ZeroInt()
	}
	
	numOfAllPeriodsF := float64(vestingEnd.Sub(unlockingStart)) / float64(period)
	numOfAllPeriods := int64(math.Ceil(numOfAllPeriodsF))

	numOfPeriodsF := float64(current.Sub(unlockingStart)) / float64(period)
	numOfPeriods := int64(math.Floor(numOfPeriodsF))

	amountDec := sdk.NewDecFromInt(amount)

	resultDec := amountDec.MulInt64(numOfPeriods).QuoInt64(numOfAllPeriods)
	return resultDec.TruncateInt()
}

func CreateTimeFromNumOfHours(numOfHours int64) time.Time {
	return TestEnvTime.Add(time.Hour * time.Duration(numOfHours))
}

func CreateDurationFromNumOfHours(numOfHours int64) time.Duration {
	return time.Hour * time.Duration(numOfHours)
}