defmodule Tcp.Server do
  require Logger

  def start_link(port, protocol) do
    Logger.info("Listening on port #{port}")
    :ranch.start_listener(make_ref(), :ranch_tcp, [port: port, max_connections: 100], protocol, [])
  end
end
