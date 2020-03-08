log_level = "TRACE"

addresses {
  http = "unix:///tmp/consul-test.sock"
}

connect {
  enabled = true
}

services {
  name = "doughboy-native-responder"
  port = 5000

  connect {
    native = true
  }
}

services {
  name = "doughboy-native-requester"
  port = 5100

  connect {
    native = true
  }
}
