sm_list_asc = (1..100) |> Enum.into([])
md_list_asc = (1..10_000) |> Enum.into([])
lg_list_asc = (1..10_000) |> Enum.into([])

inputs = %{
  "sm_list_asc" => sm_list_asc,
  "sm_list_dsc" => Enum.reverse(sm_list_asc),
  "sm_list_rnd" => Enum.shuffle(sm_list_asc),
  "md_list_asc" => md_list_asc,
  "md_list_dsc" => Enum.reverse(md_list_asc),
  "md_list_rnd" => Enum.shuffle(md_list_asc),
  "lg_list_asc" => lg_list_asc,
  "lg_list_dsc" => Enum.reverse(lg_list_asc),
  "lg_list_rnd" => Enum.shuffle(lg_list_asc)
}

Benchee.run(%{
  "bubble"  => fn (list) -> Sort.Bubble.sort(list) end,
  "merge"   => fn (list) -> Sort.Merge.sort(list) end,
  "quick"   => fn (list) -> Sort.Quick.sort(list) end
}, inputs: inputs)
