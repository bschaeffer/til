defmodule Tcp.EchoProtocol do
  use GenServer
  require Logger

  @behaviour :ranch_protocol

  def start_link(ref, socket, transport, _opts) do
    {:ok, :proc_lib.spawn_link(__MODULE__, :init, [ref, socket, transport])}
  end

  def init(ref, socket, transport) do
    Logger.info("Connection!")
    :ok = :ranch.accept_ack(ref)
    :ok = transport.setopts(socket, [{:active, true}])
    :gen_server.enter_loop(__MODULE__, [], %{socket: socket, transport: transport})
  end

  def handle_info({:tcp, socket, data}, state = %{socket: socket, transport: transport}) do
    Logger.info("Received: #{inspect(data)}")
    transport.send(socket, "OK")
    {:noreply, state}
  end

  def handle_info({:tcp_closed, socket}, state = %{socket: socket, transport: transport}) do
    Logger.info("Closing connection")
    transport.close(socket)
    {:stop, :normal, state}
  end

  def handle_info(msg, state) do
    Logger.warn("Unexpected message: #{inspect(msg)}")
    {:noreply, state}
  end
end
