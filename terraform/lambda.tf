resource "aws_lambda_function" "go_lambda" {
  description   = ""
  function_name = "go_lambda"
  architectures = [
    "x86_64"
  ]
  runtime     = "go1.x"
  filename    = "lambda_function_payload.zip"
  memory_size = 1024
  role        = aws_iam_role.lambda-role.arn
  timeout     = 30
  handler     = "main.go"
}

data "archive_file" "lambda_zip" {
  type        = "zip"
  source_file = "../main.go"
  output_path = "lambda_function_payload.zip"
}
