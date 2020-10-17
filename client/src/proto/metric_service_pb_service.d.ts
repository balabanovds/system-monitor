// package: metric
// file: metric_service.proto

import * as metric_service_pb from "./metric_service_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import {grpc} from "@improbable-eng/grpc-web";

type MetricsGetStream = {
  readonly methodName: string;
  readonly service: typeof Metrics;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof metric_service_pb.Request;
  readonly responseType: typeof metric_service_pb.Metric;
};

type MetricsParsersInfo = {
  readonly methodName: string;
  readonly service: typeof Metrics;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof google_protobuf_empty_pb.Empty;
  readonly responseType: typeof metric_service_pb.ParsersInfoResponse;
};

export class Metrics {
  static readonly serviceName: string;
  static readonly GetStream: MetricsGetStream;
  static readonly ParsersInfo: MetricsParsersInfo;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class MetricsClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  getStream(requestMessage: metric_service_pb.Request, metadata?: grpc.Metadata): ResponseStream<metric_service_pb.Metric>;
  parsersInfo(
    requestMessage: google_protobuf_empty_pb.Empty,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: metric_service_pb.ParsersInfoResponse|null) => void
  ): UnaryResponse;
  parsersInfo(
    requestMessage: google_protobuf_empty_pb.Empty,
    callback: (error: ServiceError|null, responseMessage: metric_service_pb.ParsersInfoResponse|null) => void
  ): UnaryResponse;
}

