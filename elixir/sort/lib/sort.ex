defmodule Sort do
  @moduledoc """
  Sorting algorithims.
  """

  @doc "Sorts a list using `Sort.Bubble.sort/1"
  @spec bubble(list()) :: list()
  def bubble(list), do: Sort.Bubble.sort(list)

  @doc "Sorts a list using `Sort.Merge.sort/1"
  @spec merge(list()) :: list()
  def merge(list), do: Sort.Merge.sort(list)
end
