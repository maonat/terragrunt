locals {
  parent_var = run_cmd("echo", "parent_hcl_file")
}

dependency "dependency" {
  config_path = "../dependency"

  mock_outputs = {
    mock_key = "mock_value"
  }

}
