// const { UserServiceClient } = require('./proto/user/user_grpc_web_pb');
// const { User } = require('./proto/user/user_pb.js');

import { UserServiceClient } from './proto/user/user_grpc_web_pb'
import { User } from './proto/user/user_pb.js'

function initClient() {
  const client = new UserServiceClient('localhost:8080');

  const request = new User();
  // request.setMessage('Hello World!');

  // const metadata = { 'custom-header-1': 'value1' };

  client.register(request/*, metadata*/, (err, resp) => {
    console.log(err, resp)
  })
}

export default initClient
