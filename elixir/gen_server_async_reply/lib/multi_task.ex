defmodule MultiTask do
  use GenServer

  def get(pid, key) do
    GenServer.call(pid, {:get, key}, 1_500)
  end

  def start_link do
    GenServer.start_link(__MODULE__, %{replies: %{}, tasks: %{}})
  end

  def handle_call({:get, key}, from, state) do
    tasks = case Map.get(state.tasks, key) do
      nil ->
        current = self()
        pid = spawn(fn ->
          # Trapping and monitoring can be added too.
          :timer.sleep(1_000) # long running task
          GenServer.cast(current, {:done, key})
        end)

        Map.put(state.tasks, key, pid)

      _ -> state.tasks
    end

    replies = Map.update(state.replies, key, [from], fn r -> r ++ [from] end)
    {:noreply, %{state | tasks: tasks, replies: replies}}
  end

  def handle_cast({:done, key}, state) do
    state
    |> Map.fetch!(:replies)
    |> Map.fetch!(key)
    |> Enum.each(fn from -> GenServer.reply(from, {:ok, key}) end)

    tasks = Map.delete(state.tasks, key)
    replies = Map.delete(state.replies, key)
    {:noreply, %{state | tasks: tasks, replies: replies}}
  end
end
