from common import *
import in_given_out
import out_given_in
import one_for_zero as ofz
import common.sdk_dec as sdk_dec
from common.common import SqrtPriceRange

def estimate_compute_swap_step_out_given_in_reach_ofz():
    """
    go test -timeout 30s -run TestStrategyTestSuite/TestComputeSwapStepOutGivenIn_OneForZero/1 github.com/osmosis-labs/osmosis/v14/x/concentrated-liquidity/internal/swapstrategy -count=1
    """
    liquidity =  sdk_dec.new("1517882343.751510418088349649")
    sqrt_price_current = sdk_dec.new("70.710678118654752440")
    sqrt_price_next = sdk_dec.new("74.161984870956629487")

    token_in, token_out, _ = ofz.calc_test_case_with_next_sqrt_price_out_given_in(liquidity, sqrt_price_current, sqrt_price_next, sdk_dec.zero)
    print("token_in", token_in)
    print("token_out", token_out)

def estimate_compute_swap_step_out_given_in_no_reach_ofz():
    """
    go test -timeout 30s -run TestStrategyTestSuite/TestComputeSwapStepOutGivenIn_OneForZero/1 github.com/osmosis-labs/osmosis/v14/x/concentrated-liquidity/internal/swapstrategy -count=1
    """
    liquidity =  sdk_dec.new("1517882343.751510418088349649")
    sqrt_price_current = sdk_dec.new("70.710678118654752440")
    token_in_remaining = sdk_dec.new("5238677583") - sdk_dec.new("100")

    sqrt_price_next, token_out, _ = ofz.calc_test_case_out_given_in(liquidity, sqrt_price_current, token_in_remaining, sdk_dec.zero)
    print("sqrt_price_next", sqrt_price_next)
    print("token_out", token_out)

def main():
    out_given_in.test()
    
    in_given_out.test()

    estimate_compute_swap_step_out_given_in_reach_ofz()

    estimate_compute_swap_step_out_given_in_no_reach_ofz()

if __name__ == "__main__":
    main()
