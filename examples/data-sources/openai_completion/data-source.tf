data "openai_completion" "my_completion" {
  best_of           = 6
  echo              = false
  frequency_penalty = 71.52
  logit_bias = {
    "quibusdam" = 6
    "nulla"     = 5
  }
  logprobs         = 9
  max_tokens       = 16
  model            = "text-curie-001"
  n                = 1
  presence_penalty = 38.44
  prompt           = "...my_prompt..."
  stream           = false
  suffix           = "test."
  temperature      = 1
  top_p            = 1
  user             = "user-1234"
}