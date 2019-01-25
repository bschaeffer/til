defmodule Sort.Quick do
  @moduledoc """
  Quick Sort

  Not particularly useful here because of the linked list situation. Really great blog post going down the rabbit hole
  explaining why this is difficult with linked lists, and some different approaches you could use in the situation:

    https://harfangk.github.io/2017/01/19/various-quicksort-implementations-in-elixir.html
  """

  @doc "Sorts the list using the quick sort algorithim"
  @spec sort(list()) :: list()
  def sort([]), do: []
  def sort([pivot | tail]) do
    {left, right} = Enum.partition(tail, &(&1 < pivot))
    sort(left) ++ [pivot | sort(right)]
  end
end
