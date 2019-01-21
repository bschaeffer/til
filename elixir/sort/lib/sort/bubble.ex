defmodule Sort.Bubble do
  @moduledoc """
  Bubble Sort
  """

  @doc "Sorts the list using the bubble sort algorithim"
  @spec sort(list()) :: list()
  def sort(list), do: do_sort(list, [], :sorted)

  defp do_sort([h1 | [h2 | tail]], acc, _status) when h1 > h2 do
    do_sort([h1 | tail], [h2 | acc], :unsorted)
  end
  defp do_sort([head | tail], acc, status) do
    do_sort(tail, [head | acc], status)
  end
  defp do_sort([], acc, :unsorted) do
    do_sort(Enum.reverse(acc), [], :sorted)
  end
  defp do_sort([], acc, _status) do
    Enum.reverse(acc)
  end
end
