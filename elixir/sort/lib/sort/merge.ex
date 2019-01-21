defmodule Sort.Merge do
  @moduledoc """
  Merge Sort

  This is the sort implementation used by erlang under the hood, as it has the best performance on the average case for
  sorting linked lists. Implemented the **merge** step instead of using `:lists.merge/2`) so I can actually see how that
  process works.
  """
  @doc "Sorts a list using the merge sort algorithim"
  @spec sort(list()) :: list()
  def sort([]), do: []
  def sort([i]), do: [i]
  def sort(list) do
    {left, right} = Enum.split(list, div(length(list), 2))
    merge(sort(left), sort(right), [])
  end

  defp merge([], [], acc) do
    :lists.reverse(acc)
  end

  defp merge([], [h | t], acc) do
    merge([], t, [h | acc])
  end

  defp merge([h | t], [], acc) do
    merge([], t, [h | acc])
  end

  defp merge(l1 = [h1 | t1], l2 = [h2 | t2], acc) do
    if h1 <= h2 do
      merge(t1, l2, [h1 | acc])
    else
      merge(t2, l1, [h2 | acc])
    end
  end
end
