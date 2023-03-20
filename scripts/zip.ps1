$compress = @{
  Path = "."
  DestinationPath = "go.zip"
}
Compress-Archive -Update @compress;