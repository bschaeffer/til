defmodule Sort.MergeTest do
  use ExUnit.Case
  use ExUnitProperties

  doctest Sort.Bubble

  describe ".sort/1" do
    test "sorts an empty list" do
      assert [] = Sort.Merge.sort([])
    end

    test "sorts a single item list" do
      assert [1] = Sort.Merge.sort([1])
    end

    test "sorts a two item list" do
      assert [1, 2] = Sort.Merge.sort([2, 1])
    end

    test "sorts a list" do
      check all list <- list_of(integer(), min_length: 5, max_length: 1_000) do
        sorted_list = Enum.sort(list)
        assert ^sorted_list = Sort.Merge.sort(list)
      end
    end
  end
end
