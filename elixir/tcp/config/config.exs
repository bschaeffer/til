# This file is responsible for configuring your application
# and its dependencies with the aid of the Mix.Config module.
use Mix.Config

config :tcp,
  port: String.to_integer(System.get_env("TCP_PORT") || "9876")
