// --- gRPC official client imports ---
import { HelloRequest } from "./pb/hello_pb.js";
import { HelloServiceClient } from "./pb/hello_grpc_web_pb.js";

import { PingServiceClient } from "./pb/ping_grpc_web_pb.js";

// --- google Empty ---
import { Empty } from "google-protobuf/google/protobuf/empty_pb.js";

// --- ENV ---
const GRPC_WEB_PORT = process.env.GRPC_WEB_PORT || "2222";
const PROTOCOL = window.location.protocol;
const HOSTNAME = window.location.hostname;
const GRPC_HOST = `${PROTOCOL}//${HOSTNAME}:${GRPC_WEB_PORT}`;

// --- gRPC clients ---
const helloClient = new HelloServiceClient(GRPC_HOST, null, null);
const pingClient = new PingServiceClient(GRPC_HOST, null, null);

/**
 * SayHello
 */
export function sayHello(name, token) {
    return new Promise((resolve, reject) => {
        console.log("Say hello.")
        const req = new HelloRequest();
        req.setName(name);

        helloClient.sayHello(
            req,
            { "Authorization": "Bearer " + token}, (err, resp) => {
            if (err) {
                reject(err);
            } else {
                resolve(resp);
            }
        });
    });
}

/**
 * Ping
 */
export function ping() {
    return new Promise((resolve, reject) => {
        console.log("Ping.")
        const req = new Empty();

        pingClient.ping(req, {}, (err, resp) => {
            if (err) {
                reject(err);
            } else {
                resolve(resp);
            }
        });
    });
}